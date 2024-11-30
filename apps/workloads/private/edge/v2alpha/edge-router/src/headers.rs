use fastly::http::header::{CONTENT_LENGTH, CONTENT_TYPE, ACCEPT, ACCEPT_CHARSET, ACCEPT_ENCODING, ACCEPT_LANGUAGE, ACCEPT_RANGES, ACCESS_CONTROL_ALLOW_CREDENTIALS, ACCESS_CONTROL_ALLOW_HEADERS, ACCESS_CONTROL_ALLOW_METHODS, ACCESS_CONTROL_ALLOW_ORIGIN, ACCESS_CONTROL_EXPOSE_HEADERS, ACCESS_CONTROL_MAX_AGE, ACCESS_CONTROL_REQUEST_HEADERS, ACCESS_CONTROL_REQUEST_METHOD, AUTHORIZATION, HeaderName, HOST, ORIGIN, REFERER, SEC_WEBSOCKET_ACCEPT, SEC_WEBSOCKET_EXTENSIONS, SEC_WEBSOCKET_KEY, SEC_WEBSOCKET_PROTOCOL, SEC_WEBSOCKET_VERSION, USER_AGENT, STRICT_TRANSPORT_SECURITY, REFERRER_POLICY, DATE, CONNECTION, X_FRAME_OPTIONS, X_XSS_PROTECTION, X_CONTENT_TYPE_OPTIONS, CONTENT_SECURITY_POLICY_REPORT_ONLY, CONTENT_SECURITY_POLICY};
use fastly::{Request, Response};
use fastly::http::{header};
use crate::routing::RoutingRule;
use fastly::geo::geo_lookup;

/// Define a Content Security Policy for content that can load on your site.
pub(crate) const CONTENT_SECURITY_POLICY_VALUE: &str =
    "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com; frame-src api.platform.health";

pub(crate) const X_CORRELATION_ID: HeaderName = HeaderName::from_static("x-correlation-id");
pub(crate) const X_CLIENT_TRACE_ID: HeaderName = HeaderName::from_static("x-client-trace-id");
pub(crate) const X_B3_TRACEID: HeaderName = HeaderName::from_static("x-b3-traceid");
pub(crate) const X_B3_SPANID: HeaderName = HeaderName::from_static("x-b3-spanid");
pub(crate) const X_B3_PARENTSPANID: HeaderName = HeaderName::from_static("x-b3-parentspanid");
pub(crate) const X_B3_SAMPLED: HeaderName = HeaderName::from_static("x-b3-sampled");
pub(crate) const X_B3_FLAGS: HeaderName = HeaderName::from_static("x-b3-flags");
pub(crate) const X_SPEC_ROUTER_KEY: HeaderName = HeaderName::from_static("x-spec-router-key");
pub(crate) const X_SPEC_ENVIRONMENT: HeaderName = HeaderName::from_static("x-spec-environment");
pub(crate) const X_SPEC_PLATFORM_HOST: HeaderName = HeaderName::from_static("x-spec-platform-host");
pub(crate) const X_SPEC_SENT_AT: HeaderName = HeaderName::from_static("x-spec-sent-at");
pub(crate) const X_SPEC_LOCALE: HeaderName = HeaderName::from_static("x-spec-locale");
pub(crate) const X_SPEC_TIMEZONE: HeaderName = HeaderName::from_static("x-spec-timezone");
pub(crate) const X_SPEC_API_KEY: HeaderName = HeaderName::from_static("x-spec-apikey");
pub(crate) const X_SPEC_ORGANIZATION: HeaderName = HeaderName::from_static("x-spec-organization-slug");
pub(crate) const X_SPEC_WORKSPACE: HeaderName = HeaderName::from_static("x-spec-workspace-slug");
pub(crate) const X_SPEC_WORKSPACE_JAN: HeaderName = HeaderName::from_static("x-spec-workspace-jan");
pub(crate) const X_SPEC_VALIDATE_ONLY: HeaderName = HeaderName::from_static("x-spec-validate-only");
pub(crate) const X_SPEC_FIELDMASK: HeaderName = HeaderName::from_static("x-spec-fieldmask");
pub(crate) const X_SPEC_DEVICE_ID: HeaderName = HeaderName::from_static("x-spec-device-id");
pub(crate) const X_SPEC_DEVICE_ADV_ID: HeaderName = HeaderName::from_static("x-spec-device-adv-id");
pub(crate) const X_SPEC_DEVICE_MANUFACTURER: HeaderName = HeaderName::from_static("x-spec-device-manufacturer");
pub(crate) const X_SPEC_DEVICE_MODEL: HeaderName = HeaderName::from_static("x-spec-device-model");
pub(crate) const X_SPEC_DEVICE_NAME: HeaderName = HeaderName::from_static("x-spec-device-name");
pub(crate) const X_SPEC_DEVICE_TYPE: HeaderName = HeaderName::from_static("x-spec-device-type");
pub(crate) const X_SPEC_DEVICE_TOKEN: HeaderName = HeaderName::from_static("x-spec-device-token");
pub(crate) const X_SPEC_DEVICE_BLUETOOTH: HeaderName = HeaderName::from_static("x-spec-bluetooth");
pub(crate) const X_SPEC_DEVICE_CELLULAR: HeaderName = HeaderName::from_static("x-spec-cellular");
pub(crate) const X_SPEC_DEVICE_WIFI: HeaderName = HeaderName::from_static("x-spec-wifi");
pub(crate) const X_SPEC_DEVICE_CARRIER: HeaderName = HeaderName::from_static("x-spec-carrier");
pub(crate) const X_SPEC_OS_NAME: HeaderName = HeaderName::from_static("x-spec-os-name");
pub(crate) const X_SPEC_OS_VERSION: HeaderName = HeaderName::from_static("x-spec-os-version");

pub(crate) const GRPC_ACCEPT_ENCODING: HeaderName = HeaderName::from_static("grpc-accept-encoding");
pub(crate) const GRPC_ENCODING: HeaderName = HeaderName::from_static("grpc-encoding");
pub(crate) const GRPC_STATUS: HeaderName = HeaderName::from_static("grpc-status");
pub(crate) const GRPC_MESSAGE: HeaderName = HeaderName::from_static("grpc-message");
pub(crate) const GRPC_TIMEOUT: HeaderName = HeaderName::from_static("grpc-timeout");
pub(crate) const TE: HeaderName = HeaderName::from_static("te");

pub(crate) const SEC_CH_PREFERS_COLOR_SCHEME: HeaderName = HeaderName::from_static("sec-ch-prefers-color-scheme");
pub(crate) const SEC_CH_PREFERS_REDUCED_MOTION: HeaderName = HeaderName::from_static("sec-ch-prefers-reduced-motion");
pub(crate) const SEC_CH_PREFERS_REDUCED_TRANSPARENCY: HeaderName = HeaderName::from_static("sec-ch-prefers-reduced-transparency");
pub(crate) const SEC_CH_UA: HeaderName = HeaderName::from_static("sec-ch-ua");
pub(crate) const SEC_CH_UA_ARCH: HeaderName = HeaderName::from_static("sec-ch-ua-arch");
pub(crate) const SEC_CH_UA_BITNESS: HeaderName = HeaderName::from_static("sec-ch-ua-bitness");
pub(crate) const SEC_CH_UA_FULL_VERSION_LIST: HeaderName = HeaderName::from_static("sec-ch-ua-full-version-list");
pub(crate) const SEC_CH_UA_MOBILE: HeaderName = HeaderName::from_static("sec-ch-ua-mobile");
pub(crate) const SEC_CH_UA_MODEL: HeaderName = HeaderName::from_static("sec-ch-ua-model");
pub(crate) const SEC_CH_UA_PLATFORM: HeaderName = HeaderName::from_static("sec-ch-ua-platform");
pub(crate) const SEC_CH_UA_PLATFORM_VERSION: HeaderName = HeaderName::from_static("sec-ch-ua-platform-version");
pub(crate) const SEC_FETCH_DEST: HeaderName = HeaderName::from_static("sec-fetch-dest");
pub(crate) const SEC_FETCH_MODE: HeaderName = HeaderName::from_static("sec-fetch-mode");
pub(crate) const SEC_FETCH_SITE: HeaderName = HeaderName::from_static("sec-fetch-site");
pub(crate) const SEC_GPC: HeaderName = HeaderName::from_static("sec-gpc");
pub(crate) const SEC_PURPOSE: HeaderName = HeaderName::from_static("sec-purpose");

pub(crate) static ALLOWED_REQUEST_HEADERS: [HeaderName; 70] =
[
    HOST,
    ORIGIN,
    ACCEPT,
    ACCEPT_ENCODING,
    ACCEPT_CHARSET,
    ACCEPT_LANGUAGE,
    ACCEPT_RANGES,
    AUTHORIZATION,
    CONTENT_TYPE,
    CONTENT_LENGTH,
    REFERER,
    USER_AGENT,

    GRPC_ACCEPT_ENCODING,
    GRPC_TIMEOUT,
    TE,

    SEC_CH_PREFERS_COLOR_SCHEME,
    SEC_CH_PREFERS_REDUCED_MOTION,
    SEC_CH_PREFERS_REDUCED_TRANSPARENCY,
    SEC_CH_UA,
    SEC_CH_UA_ARCH,
    SEC_CH_UA_BITNESS,
    SEC_CH_UA_FULL_VERSION_LIST,
    SEC_CH_UA_MOBILE,
    SEC_CH_UA_MODEL,
    SEC_CH_UA_PLATFORM,
    SEC_CH_UA_PLATFORM_VERSION,
    SEC_FETCH_DEST,
    SEC_FETCH_MODE,
    SEC_FETCH_SITE,
    SEC_GPC,
    SEC_PURPOSE,
    SEC_WEBSOCKET_ACCEPT,
    SEC_WEBSOCKET_ACCEPT,
    SEC_WEBSOCKET_EXTENSIONS,
    SEC_WEBSOCKET_KEY,
    SEC_WEBSOCKET_PROTOCOL,
    SEC_WEBSOCKET_VERSION,

    //ACCESS_CONTROL_ALLOW_HEADERS,
    //ACCESS_CONTROL_ALLOW_METHODS,
    //ACCESS_CONTROL_ALLOW_CREDENTIALS,
    //ACCESS_CONTROL_ALLOW_ORIGIN,
    ACCESS_CONTROL_EXPOSE_HEADERS,
    ACCESS_CONTROL_MAX_AGE,
    ACCESS_CONTROL_REQUEST_HEADERS,
    ACCESS_CONTROL_REQUEST_METHOD,

    X_CORRELATION_ID,
    X_CLIENT_TRACE_ID,
    X_B3_TRACEID,
    X_B3_SPANID,
    X_B3_PARENTSPANID,
    X_B3_SAMPLED,
    X_B3_FLAGS,

    X_SPEC_API_KEY,
    X_SPEC_ROUTER_KEY,
    X_SPEC_ENVIRONMENT,
    X_SPEC_PLATFORM_HOST,
    X_SPEC_FIELDMASK,
    X_SPEC_SENT_AT,
    X_SPEC_LOCALE,
    X_SPEC_TIMEZONE,
    //X_SPEC_WORKSPACE,
    X_SPEC_VALIDATE_ONLY,
    X_SPEC_DEVICE_ID,
    X_SPEC_DEVICE_ADV_ID,
    X_SPEC_DEVICE_MANUFACTURER,
    X_SPEC_DEVICE_MODEL,
    X_SPEC_DEVICE_NAME,
    X_SPEC_DEVICE_TYPE,
    X_SPEC_DEVICE_TOKEN,
    X_SPEC_DEVICE_BLUETOOTH,
    X_SPEC_DEVICE_CELLULAR,
    X_SPEC_DEVICE_WIFI,
    X_SPEC_DEVICE_CARRIER,
    X_SPEC_OS_NAME,
    X_SPEC_OS_VERSION,
];

pub(crate) static ALLOWED_RESPONSE_HEADERS: [HeaderName; 34] =
[
    //HOST,
    //ORIGIN,
    ACCEPT,
    ACCEPT_ENCODING,
    ACCEPT_CHARSET,
    ACCEPT_LANGUAGE,
    ACCEPT_RANGES,
    CONTENT_TYPE,
    CONTENT_LENGTH,
    REFERER,
    USER_AGENT,
    DATE,
    CONNECTION,

    CONTENT_SECURITY_POLICY,
    CONTENT_SECURITY_POLICY_REPORT_ONLY,
    STRICT_TRANSPORT_SECURITY,
    REFERRER_POLICY,

    GRPC_ACCEPT_ENCODING,
    GRPC_ENCODING,
    GRPC_STATUS,
    GRPC_MESSAGE,
    TE,

    ACCESS_CONTROL_ALLOW_HEADERS,
    ACCESS_CONTROL_ALLOW_METHODS,
    ACCESS_CONTROL_ALLOW_CREDENTIALS,
    ACCESS_CONTROL_ALLOW_ORIGIN,
    ACCESS_CONTROL_MAX_AGE,

    X_FRAME_OPTIONS,
    X_XSS_PROTECTION,
    X_CONTENT_TYPE_OPTIONS,

    X_SPEC_SENT_AT,
    X_SPEC_LOCALE,
    X_SPEC_ORGANIZATION,
    X_SPEC_WORKSPACE,
    X_SPEC_WORKSPACE_JAN,

    X_CORRELATION_ID,
];


pub(crate) fn sanitize_request_headers(req: &mut Request) {
    let to_remove: Vec<_> = req
        .get_header_names()
        .filter(|header| !ALLOWED_REQUEST_HEADERS.contains(header))
        .cloned()
        .collect();

    for header in to_remove {
        req.remove_header(header);
    }
}

pub(crate) fn decorate_request_headers(req: &mut Request, routing_rule: &RoutingRule) {
    req.set_header("x-spec-workspace-slug", &routing_rule.workspace_slug);
    req.set_header("x-spec-organization-slug", &routing_rule.organization_slug);
    req.set_header("x-spec-workspace-jan", &routing_rule.jan);
    if let Some(geo) = req.get_client_ip_addr().and_then(geo_lookup) {
        req.set_header("x-spec-city", geo.city());
        req.set_header("x-spec-continent", format!("{:?}", geo.continent()));
        req.set_header("x-spec-country", geo.country_code());
        req.set_header("x-spec-lat", geo.latitude().to_string());
        req.set_header("x-spec-long", geo.longitude().to_string());
        req.set_header("x-spec-speed", format!("{:?}", geo.conn_speed()));
    }

    if let Some(ip) = req.get_client_ip_addr() {
        req.set_header("x-spec-ip", &ip.to_string());
        req.set_header("x-forwarded-for", &ip.to_string());
    }

    if let Some(request_id) = req.get_client_request_id() {
        req.set_header("x-spec-request-id", request_id.to_string());
    }
}

pub(crate) fn sanitize_response_headers(resp: &mut Response) {
    let to_remove: Vec<_> = resp
        .get_header_names()
        .filter(|header| !ALLOWED_RESPONSE_HEADERS.contains(header))
        .cloned()
        .collect();

    for header in to_remove {
        resp.remove_header(header);
    }
}

pub(crate) fn secure_response_headers(beresp: &mut Response) {

    beresp.set_header(header::CONTENT_SECURITY_POLICY, "default-src 'self'");
    beresp.set_header(header::X_FRAME_OPTIONS, "SAMEORIGIN");
    beresp.set_header(header::X_XSS_PROTECTION, "1");
    beresp.set_header(header::X_CONTENT_TYPE_OPTIONS, "nosniff");
    beresp.set_header(header::REFERRER_POLICY, "origin-when-cross-origin");
    beresp.set_header(
        header::STRICT_TRANSPORT_SECURITY,
        "max-age=31536000; includeSubDomains",
    );

    let has_max_age = beresp
        .get_header_str(header::CACHE_CONTROL)
        .unwrap_or("")
        .split(',')
        .any(|directive| {
            directive
                .split('=')
                .next()
                .unwrap_or("")
                .trim()
                .eq_ignore_ascii_case("max-age")
        });
    if has_max_age {
        beresp.remove_header(header::EXPIRES);
    }
}

//pub(crate) fn list_headers(req: &mut Request) {
pub(crate) fn list_headers(req: &mut Request) {
    for header in req.get_header_names() {

        let value = match req.get_header_str(header) {
            None => "".to_string(),
            Some(val) => val.to_string()
        };

        println!(
            "    Request Header: {}: {:?}",
            header,
            value
        );
    }
}

pub(crate) fn list_response_headers(req: &mut Response) {
    for header in req.get_header_names() {

        let value = match req.get_header_str(header) {
            None => "".to_string(),
            Some(val) => val.to_string()
        };

        println!(
            "    Response Header: {}: {:?}",
            header,
            value
        );
    }
}
