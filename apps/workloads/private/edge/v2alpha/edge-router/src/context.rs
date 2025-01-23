use fastly::{ConfigStore, Request};
use std::fmt::Write;
use std::string::ToString;

#[derive(serde::Serialize)]
pub(crate) struct Context {
    pub service_id: String,
    pub env: String,
    pub host: String,
    pub port: String,
    pub system: String,
    pub debug: bool,
}

const DEFAULT_FASTLY_SERVICE_ID: &str = "0000000000000000000000";

pub(crate) fn is_local() -> bool {
    return std::env::var("FASTLY_HOSTNAME").unwrap() == "localhost";
}

pub(crate) fn get_service_id() -> String {
    let service_id = std::env::var("FASTLY_SERVICE_ID");

    if !service_id.is_err() {
        let sid = service_id.unwrap();
        if sid == DEFAULT_FASTLY_SERVICE_ID {
            "localhost".to_string()
        } else {
            sid
        }
    } else if service_id.is_err() {
        "localhost".to_string()
    } else {
        std::env::var("FASTLY_SERVICE_ID").unwrap()
    }
}

pub(crate) fn extract_context(req: &mut Request, debug: bool) -> Option<Context> {

    let service_id = get_service_id();

    let mut store = String::new();
    write!(&mut store, "context_{}", service_id).unwrap();

    let _config_store = ConfigStore::try_open(&store);
    if _config_store.is_err() {
        println!("Error {}", _config_store.err().unwrap().to_string());
        println!("Context Store could not be opened: {}", store);
        return None
    }

    let config_store = _config_store.unwrap();

    let _context = config_store.try_get("context");
    if _context.is_err() {
        println!("Error {}", _context.err().unwrap().to_string());
        println!("Could not find context value in the store");
        return None
    }

    let context = _context.unwrap();
    if context.is_none() {
        println!("Context store opened and found context, however value is empty");
        return None
    }

    let _c: Vec<&str> = context.as_ref().unwrap()
        .split(";")
        .collect();

    if _c.len() != 4 {
        return None
    }

    let host = _c[0];
    let env = _c[1];
    let port = _c[2];
    let system = _c[3];

    if debug {
        println!("Service ID: {}", service_id);
        println!("Environment: {}", env);
        println!("Host: {}", host);
        println!("System: {}", system);
        println!("Port: {}", port);
        println!("Method: {}", req.get_method_str());
        println!("Path: {}", req.get_path());
        println!("Version: {:?}", req.get_version());
        let mime_type = req.get_content_type();
        if mime_type.is_some() {
            let mime = mime_type.unwrap();
            println!("Mime: {}", mime.to_string());
            println!("Mime Subtype: {}", mime.subtype().to_string());
            if mime.suffix().is_some() {
                println!("Mime Suffix: {}", mime.suffix().unwrap().to_string());
            }
        }
    }

    Some(Context {
        service_id: service_id,
        env: env.to_string(),
        host: host.to_string(),
        port: port.to_string(),
        system: system.to_string(),
        debug,
    })
}

pub(crate) fn extract_api_key(req: &mut Request) -> Option<String> {

    // TODO: Sanitize header length no weird characters
    // TODO: Limit length
    // TODO: data.retain(|x| {!['(', ')', ',', '\"', '.', ';', ':', '\''].contains(&x)});
    // TODO: https://users.rust-lang.org/t/fast-removing-chars-from-string/24554/11
    let api_key = req.get_header_str("x-spec-apikey");

    if api_key.is_none() {
        return None;
    }

    Some(api_key.unwrap().to_string())
}