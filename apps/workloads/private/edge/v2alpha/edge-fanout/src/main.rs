use fastly::http::StatusCode;
use fastly::{Error, Request, Response};

mod fanout_util;

fn handle_fanout_ws(mut req: Request, chan: &str) -> Response {
    if req.get_header_str("Content-Type") != Some("application/websocket-events") {
        return Response::from_status(StatusCode::BAD_REQUEST)
            .with_body("Not a WebSocket-over-HTTP request.\n");
    }

    let req_body = req.take_body().into_bytes();
    let mut resp_body: Vec<u8> = [].to_vec();

    let mut resp = Response::from_status(StatusCode::OK)
        .with_header("Content-Type", "application/websocket-events");

    if req_body.starts_with(b"OPEN\r\n") {
        resp.set_header("Sec-WebSocket-Extensions", "grip; message-prefix=\"\"");
        resp_body.extend("OPEN\r\n".as_bytes());
        resp_body.extend(fanout_util::ws_sub(chan));
    } else if req_body.starts_with(b"TEXT ") {
        resp_body.extend(fanout_util::ws_text(
            format!("You said: {}", std::str::from_utf8(&req_body).unwrap_or("")).as_str(),
        ));
    }

    resp.set_body(resp_body);
    resp
}

fn handle_test(req: Request, chan: &str) -> Response {
    match req.get_url().path() {
        "/test/long-poll" => fanout_util::grip_response("text/plain", "response", chan),
        "/test/stream" => fanout_util::grip_response("text/plain", "stream", chan),
        "/test/sse" => fanout_util::grip_response("text/event-stream", "stream", chan),
        "/test/websocket" => handle_fanout_ws(req, chan),
        _ => Response::from_status(StatusCode::NOT_FOUND).with_body("No such test endpoint\n"),
    }
}

fn is_tls(req: &Request) -> bool {
    req.get_url().scheme().eq_ignore_ascii_case("https")
}

fn main() -> Result<(), Error> {
    // Log service version
    println!(
        "FASTLY_SERVICE_VERSION: {}",
        std::env::var("FASTLY_SERVICE_VERSION").unwrap_or_else(|_| String::new())
    );

    let mut req = Request::from_client();

    let host = match req.get_url().host_str() {
        Some(s) => s.to_string(),
        None => {
            return Ok(Response::from_status(StatusCode::NOT_FOUND)
                .with_body("Unknown host\n")
                .send_to_client());
        }
    };

    let path = req.get_path().to_string();

    if let Some(addr) = req.get_client_ip_addr() {
        req.set_header("X-Forwarded-For", addr.to_string());
    }

    if is_tls(&req) {
        req.set_header("X-Forwarded-Proto", "https");
    }

    // Request is a test request - from client, or from Fanout
    if host.ends_with(".edgecompute.app") && path.starts_with("/test/") {
        if req.get_header_str("Grip-Sig").is_some() {
            // Request is from Fanout, handle it here
            return Ok(handle_test(req, "test").send_to_client());
        } else {
            // Not from Fanout, route it through Fanout first
            return Ok(req.handoff_fanout("self")?);
        }
    }

    // Forward all non-test requests to the origin
    Ok(req.handoff_fanout("origin")?)
}
