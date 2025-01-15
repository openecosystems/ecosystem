use fastly::backend::{Backend};
use fastly::{Request, Response};
use fastly::experimental::GrpcBackend;
use crate::context::Context;
use crate::routing::RoutingRule;
use std::{time::{Duration}};
use std::fmt::Write;
use crate::{context, errors};
use crate::protocol::{Protocol, SpecProtocol};

pub(crate) fn determine_ingress_gateway(req: &mut Request, ctx: &Context, routing_rule: &RoutingRule, sanitized_query_string: &str, protocol: &Protocol) -> Option<Backend> {

    let local = context::is_local();

    let mut _name = String::new();
    write!(&mut _name, "api_{}_{}_{}", &ctx.system, &ctx.env, &routing_rule.jan).unwrap();
    let name = _name.as_str();

    let mut _target = String::new();
    write!(&mut _target, "api.{}.{}.{}.{}:{}", &ctx.system, &ctx.env, &routing_rule.jan, &ctx.host, &ctx.port).unwrap();
    let mut target = _target.as_str();

    if let "nosystem" = &*ctx.system {
        _target = String::new();
        write!(&mut _target, "api.{}.{}.{}:{}", &ctx.env, &routing_rule.jan, &ctx.host, &ctx.port).unwrap();
        target = _target.as_str();
    }

    let mut _path = String::new();
    write!(&mut _path, "{}{}{}", target, req.get_path(),  sanitized_query_string).unwrap();
    let path = _path.as_str();

    let mut _url = String::new();
    // TODO: Revert this back to https once settled on server approach
    write!(&mut _url, "http://{}", path).unwrap();
    let mut url = _url.as_str();

    if local {
        _target = String::new();
        write!(&mut _target, "{}:{}", &ctx.host, &ctx.port).unwrap();
        target = _target.as_str();

        _url = String::new();
        write!(&mut _url, "http://{}{}{}", target, req.get_path(),  sanitized_query_string).unwrap();
        url = _url.as_str();
    }

    if ctx.debug {
        println!("Traffic Type: {:?}", protocol.spec_protocol);
        println!("Ingress Name: {}", name);
        println!("Ingress Target: {}", target);
        println!("Ingress URL: {}",  url);

        if local { println!("Running locally"); }
    }

    req.set_url(url);

    let mut backend_builder = Backend::builder(name, target);
    backend_builder = backend_builder.connect_timeout(Duration::from_secs(10));
    backend_builder = backend_builder.first_byte_timeout(Duration::from_secs(15));
    backend_builder = backend_builder.between_bytes_timeout(Duration::from_secs(10));
    backend_builder = backend_builder.enable_pooling(false);


    match protocol.spec_protocol {
        SpecProtocol::HTTP | SpecProtocol::Rest | SpecProtocol::Graphql | SpecProtocol::GrpcWeb | SpecProtocol::Connect | SpecProtocol::Soap => {}
        SpecProtocol::Grpc => {
            backend_builder = backend_builder.for_grpc(true);
        }
    }

    if local {
        backend_builder = backend_builder.override_host(&ctx.host);
    } else {
        // TODO Revert this once SSL is set
        //backend_builder = backend_builder.enable_ssl();
        backend_builder = backend_builder.override_host(target);
        //backend_builder = backend_builder.sni_hostname(target);
        //backend_builder = backend_builder.check_certificate(target);
    }

    //let ingress = backend_builder.finish().ok()?;
    let ingress = backend_builder.finish();

    if ingress.is_err() {
        println!("Ingress Error: {}", ingress.err().unwrap());
        return None
    }

    Some(ingress.unwrap())
}


pub(crate) fn route_to_ingress_gateway(req: Request, ingress_gateway: Backend, protocol: &Protocol) -> Option<Response> {

    match protocol.spec_protocol {
        SpecProtocol::HTTP | SpecProtocol::Rest | SpecProtocol::Grpc | SpecProtocol::GrpcWeb | SpecProtocol::Connect | SpecProtocol::Soap => {

            match req.send(ingress_gateway) {
                Ok(val) => {return Some(val)}
                Err(_) => {return None}
            };
        }
        SpecProtocol::Graphql => {

            return errors::misdirected("Please use the graph endpoint".to_string(), protocol).ok();
        }
    }
}
