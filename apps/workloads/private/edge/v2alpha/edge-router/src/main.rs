//! Platform Edge Router
mod debug;
mod robots;
mod headers;
mod cors;
mod methods;
mod errors;
mod routing;
mod query;
mod context;
mod observability;
mod otlp;
mod ingress;
mod protocol;
mod normalize;

use fastly::{Error, Request, Response};
use std::string::ToString;
use std::{time::{SystemTime, UNIX_EPOCH}};
use std::io::{ErrorKind, Read};
use fastly::http::StatusCode;
use log_fastly::Logger;
/// If `main` returns an error, a 500 error response will be delivered to the client.

#[fastly::main]
fn main(mut req: Request) -> Result<Response, Error> {

    let start_time = SystemTime::now().duration_since(UNIX_EPOCH)
        .unwrap_or_default()
        .as_nanos();

    Logger::builder()
        .max_level(log::LevelFilter::Info)
        .default_endpoint("frontsight")
        .echo_stdout(true)
        .init();

    // Log service version
    log::info!("FASTLY_SERVICE_VERSION CHECK: {}", std::env::var("FASTLY_SERVICE_VERSION").unwrap_or_else(|_| String::new()));

    let debug = debug::is_debug(&mut req);

    if debug {
        log::info!("    =============Printing raw headers:");
        headers::list_headers(&mut req);
    }

    let protocol = protocol::get_protocol(&mut req);

    let ctx = match context::extract_context(&mut req, debug) {
        Some(val) => val,
        _ => {
            let _ = observability::observe_failed_distributed_transaction(start_time, &mut req);
            return errors::misdirected("This is a misdirected request and should not be retried".to_string(), &protocol)
        },
    };

    match robots::respond_to_robots(&mut req) {
        Some(val) => return Ok(val),
        _ => (),
    };

    match cors::respond_to_cors_preflight(&mut req) {
        Some(val) => return Ok(val),
        _ => (),
    };

    let fetch_site_mode_cors = cors::is_fetch_site_mode_cors(&mut req);


    let api_key = match context::extract_api_key(&mut req) {
        Some(val) => val.clone(),
        _ => {
            return errors::precondition("Please provide the x-spec-apikey header".to_string(), &protocol)
        },
    };

    match methods::sanitize_methods(&mut req) {
        Some(val) => return Ok(val),
        _ => (),
    };
    headers::sanitize_request_headers(&mut req);
    let sanitized_query_strings = query::sanitize_query_strings(&mut req);
    normalize::normalize_request(&mut req);

    let routing_rule = match routing::extract_workspace_routing_rules(&api_key, &ctx.service_id) {
        Some(val) => val,
        _ => return errors::precondition("Please check if API Key is correct".to_string(), &protocol),
    };

    // Set Caching Rules
    //req.set_ttl(60);

    // HIPAA Compliant Caching
    req.set_pci(true);

    headers::decorate_request_headers(&mut req, &routing_rule);
    if debug {
        log::info!("    =============Printing decorated headers:");
        headers::list_headers(&mut req);
    }

    let ingress_gateway = match ingress::determine_ingress_gateway(&mut req, &ctx, &routing_rule, &sanitized_query_strings, &protocol) {
        Some(val) => val,
        _ => return errors::fail("Could not determine ingress gateway".to_string(), &protocol),
    };

    //let _ = observability::observe_successful_distributed_transaction(start_time, &mut req, &ctx);

    let mut _beresp = match ingress::route_to_ingress_gateway(req, ingress_gateway, &protocol) {
        Some(val) => val,
        _ => return errors::fail("Could not route to ingress gateway".to_string(), &protocol),
    };

    if debug {
        let mut beresp = _beresp.clone_with_body();

        println!("Ingress Gateway Backend Request: : {:?}",  beresp.get_backend_request());
        println!("Ingress Gateway Backend Response Status: : {:?}",  beresp.get_status());
        println!("Ingress Gateway Backend Response Version: : {:?}",  beresp.get_version());

        headers::list_response_headers(&mut beresp);
        // Used to debug proto binary issues
        let mut s = String::new();
        let _ = beresp.get_body_mut().read_to_string(&mut s);
        println!("Ingress Gateway Backend Response Body Length: : {:}",  s.len());
        println!("Ingress Gateway Backend Response Body: : {:?}",  s);
    }

    headers::sanitize_response_headers(&mut _beresp);
    headers::secure_response_headers(&mut _beresp);
    if fetch_site_mode_cors {
        cors::resource_isolation_policy(&mut _beresp);
    }

    Ok(_beresp)

}
