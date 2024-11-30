use fastly::{Request, Response};

pub(crate) fn respond_to_robots(req: &mut Request) -> Option<Response> {
    // Respond to requests for robots.txt.
    if req.get_path() == "/robots.txt" {
        return Some(Response::from_body("User-agent: *\nAllow: /\n")
            .with_content_type(fastly::mime::TEXT_PLAIN));
    }

    None
}