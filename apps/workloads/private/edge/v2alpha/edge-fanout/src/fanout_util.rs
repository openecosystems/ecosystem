use fastly::http::StatusCode;
use fastly::Response;

/// Returns a GRIP response to initialize a stream
///
/// When Compute receives a non-WebSocket request (i.e. normal HTTP) and wants
/// to make it long lived (longpoll or SSE), we call handoff_fanout on it, and
/// Fanout will then forward that request to the nominated backend.  In this app,
/// that backend is this same Compute service, where we then need to respond
/// with some Grip headers to tell Fanout to hold the connection for streaming.
/// This function constructs such a response.
pub fn grip_response(ctype: &str, ghold: &str, chan: &str) -> Response {
    Response::from_status(StatusCode::OK)
        .with_header("Content-Type", ctype)
        .with_header("Grip-Hold", ghold)
        .with_header("Grip-Channel", chan)
        .with_body("")
}

/// Returns a WebSocket-over-HTTP formatted TEXT message
pub fn ws_text(msg: &str) -> Vec<u8> {
    format!("TEXT {:02x}\r\n{}\r\n", msg.len(), msg)
        .as_bytes()
        .to_vec()
}

// Returns a channel-subscription command in a WebSocket-over-HTTP format
pub fn ws_sub(ch: &str) -> Vec<u8> {
    ws_text(format!("c:{{\"type\":\"subscribe\",\"channel\":\"{}\"}}", ch).as_str())
}
