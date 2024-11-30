use std::vec::Vec;
use fastly::{Error, Request, Response};

pub(crate) static ALLOWED_REQUEST_QUERY_PARAMETERS: [&str; 5] =
[
    "query",
    "page",
    "foo",
    "key",
    "hello",
];

pub(crate) fn normalize_request(req: &mut Request) {

    //let _qs: Vec<(String, String)> = req.get_query().unwrap();
    //let mut qs = filter_except(_qs);

    // Ensure all sanitized query params are lowercase
    // for item in &mut qs {
    //     println!("Sanitized Query String: {:?}", item);
    //     if &item.0 == "query" {
    //         item.1 = item.1.to_lowercase();
    //     }
    //     println!("Sanitized Query String Lowercase: {:?}", item);
    // }

    // Sort query strings in alphabetical order
    //qs.sort_by(|a, b| a.0.cmp(&b.0));
}

// norm_accept() function checks the input against a list of accepted values
// and returns a value in the list or the default value if no match is found.
fn norm_accept(input: &str, accept: &[&str], def: &str) -> String {
    for val in accept {
        if input.contains(val) {
            return val.to_string();
        }
    }
    def.to_string()
}

fn filter_except(qs: Vec<(String, String)>) -> Vec<(String, String)> {
    qs.into_iter()
        .filter(|(k, _)| ALLOWED_REQUEST_QUERY_PARAMETERS.contains(&k.as_str()))
        .collect()
}




// fn main22(mut req: Request) -> Result<Response, Error>  {
//     let qs: Vec<(String, String)> = req.get_query().unwrap();
//     // Query strings parameters to keep.
//     let keep = vec!["query", "page", "foo"];
//     // Filter query strings to only include query
//     // parameters that are valid for your site
//     let mut qs = filter_except(qs, &ALLOWED_REQUEST_QUERY_PARAMETERS);
//
//     // Lowercase specific query string values.
//     for item in &mut qs {
//         if &item.0 == "query" {
//             item.1 = item.1.to_lowercase();
//         }
//     }
//
//     //Sort the querystring params in alphabetical order
//     qs.sort_by(|a, b| a.0.cmp(&b.0));
//
//     // Encode() encodes the query string parameters into URL encoded form.
//     let encoded_query = qs.iter()
//         .map(|(name, value)| format!("{}={}", name, value))
//         .collect::<Vec<String>>()
//         .join("&");
//
//     req.set_query_str(&encoded_query);
//
//     // Remove headers that you want to avoid using to vary responses.
//     req.remove_header("user-agent");
//     req.remove_header("cookie");
//
//     // Normalize headers that you may vary on.
//     // Normalize the Accept-Language header.
//     let lang_accept = vec!["en", "de", "fr", "nl"];
//     let lang_default = "de";
//     let lang_norm = norm_accept(req.get_header_str("accept-language").unwrap_or(""), &lang_accept, lang_default);
//     req.set_header("accept-language", lang_norm);
//
//     // Normalize the accept-encoding.
//     let encoding_accept = vec!["br", "compress", "deflate", "gzip"];
//     let encoding_default = "identity";
//     let encoding_norm = norm_accept(req.get_header_str("accept-encoding").unwrap_or(""), &encoding_accept, encoding_default);
//     req.set_header("accept-encoding", encoding_norm);
//
//     // Normalize the accept-charset.
//     let charset_accept = vec!["iso-8859-5", "iso-8859-2", "utf-8"];
//     let charset_default = "utf-8";
//     let charset_norm = norm_accept(req.get_header_str("accept-charset").unwrap_or(""), &charset_accept, charset_default);
//     req.set_header("accept-charset", charset_norm);
//
//     //send to backend and return the response to client
//     Ok(req.send("origin_0")?)
// }