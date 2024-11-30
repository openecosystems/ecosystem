use fastly::{Request, Response};
use fastly::http::{header, HeaderValue, Method, StatusCode};
use crate::{errors, headers};

pub(crate) fn respond_to_cors_preflight(req: &mut Request) -> Option<Response> {

    // Used later to generate CORS headers.
    // Usually you would want an allowlist of domains here, but this example allows any origin to make requests.
    let allowed_origins = match req.get_header(header::ORIGIN) {
        Some(val) => val.clone(),
        _ => HeaderValue::from_static("*"),
    };

    // Respond to CORS preflight requests.
    if req.get_method() == Method::OPTIONS
        && req.contains_header(header::ORIGIN)
        && (req.contains_header(header::ACCESS_CONTROL_REQUEST_HEADERS)
        || req.contains_header(header::ACCESS_CONTROL_REQUEST_METHOD))
    {
        log::info!("    =============RESPONDING TO CORS:");

        return Some(Response::from_status(StatusCode::NO_CONTENT)
            .with_header(header::ACCESS_CONTROL_ALLOW_ORIGIN, allowed_origins)
            .with_header(
                header::ACCESS_CONTROL_ALLOW_METHODS,
                "HEAD,OPTIONS,GET,PUT,PATCH,POST,DELETE",
            )
            .with_header(header::ACCESS_CONTROL_ALLOW_HEADERS, "*")
            .with_header(header::ACCESS_CONTROL_MAX_AGE, "86400")
            .with_header(header::CACHE_CONTROL, "public, max-age=86400"))
    }

    None
}

pub(crate) fn is_fetch_site_mode_cors(req: &mut Request) -> bool {

    let mode = req.get_header(headers::SEC_FETCH_MODE);
    if mode.is_some() {
        let mode = mode.unwrap().to_str();
        if mode.is_ok() {
            return true
        }
    }

    return false
}


pub(crate) fn resource_isolation_policy(beresp: &mut Response) {

    beresp.set_header(header::ACCESS_CONTROL_ALLOW_ORIGIN, "*");
    beresp.set_header(header::ACCESS_CONTROL_ALLOW_HEADERS, "*");
    beresp.set_header(header::ACCESS_CONTROL_ALLOW_METHODS,
                      "HEAD,OPTIONS,GET,PUT,PATCH,POST,DELETE",
    );
}