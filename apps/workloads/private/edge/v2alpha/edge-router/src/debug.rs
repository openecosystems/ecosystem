use fastly::Request;

pub(crate) fn is_debug(req: &mut Request) -> bool {

    return req.contains_header("Fastly-Debug") || req.contains_header("x-spec-debug");
}
