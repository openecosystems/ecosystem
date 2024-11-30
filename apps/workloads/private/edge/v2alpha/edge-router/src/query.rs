use fastly::{Request};
use std::fmt::{Write};

pub(crate) fn sanitize_query_strings(req: &mut Request) -> String {
    let mut q = String::new();
    let query = req.get_query_str();
    if !query.is_none() {
        write!(&mut q, "?{}", query.unwrap()).unwrap();
    }

    return q;
}