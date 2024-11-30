use fastly::{Request, Response};
use fastly::http::{header, Method, StatusCode};
use crate::errors;
use fastly::http::Version;

pub(crate) fn sanitize_methods(req: &mut Request) -> Option<Response> {

    // TODO: Filter according to protocol instead
    // Filter request methods...
    match req.get_version() {

        Version::HTTP_11 | Version::HTTP_10 | Version::HTTP_09 => {

            match req.get_method() {

                // Let these methods requests through
                &Method::GET | &Method::POST | &Method::PUT | &Method::PATCH | &Method::DELETE => {}

                // Block requests with unexpected methods
                _ => {
                    let e = "Only the following methods are allowed: OPTIONS, HEAD, GET, POST, PUT, PATCH, DELETE".to_string();
                    return Some(Response::from_status(StatusCode::METHOD_NOT_ALLOWED)
                        .with_header(header::ALLOW, "OPTIONS, HEAD, GET, POST, PUT, PATCH, DELETE")
                        .with_body_json(&errors::create_error(&e)).unwrap())
                },
            };

        }

        // https://chromium.googlesource.com/external/github.com/grpc/grpc/+/HEAD/doc/PROTOCOL-HTTP2.md
        Version::HTTP_2 | Version::HTTP_3 => {

            match req.get_method() {

                // Let these methods requests through
                &Method::POST => {}

                // Block requests with unexpected methods
                _ => {
                    let e = "When using HTTP2 or HTTP3, the following methods are allowed: POST".to_string();
                    return Some(Response::from_status(StatusCode::METHOD_NOT_ALLOWED)
                        .with_header(header::ALLOW, "POST")
                        .with_body_json(&errors::create_error(&e)).unwrap())
                },
            };
        },
        _ => {},
    };

    None
}