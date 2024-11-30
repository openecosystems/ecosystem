use crate::otlp::{ResourceSpans, Resource, KeyValue, InstrumentationScope, ScopeSpans, Span, TracesData, Status};
use serde_json::Result;
use fastly::{Request};
use crate::context::Context;
use opentelemetry_sdk::{trace::{RandomIdGenerator}};
use opentelemetry_sdk::trace::IdGenerator;
use crate::otlp::any_value::Value;
use std::{time::{SystemTime, UNIX_EPOCH}};

pub(crate) fn observe_successful_distributed_transaction(start_time: u128, req: &mut Request, ctx: &Context) -> Result<()> {

    log_fastly::init_simple("frontsight", log::LevelFilter::Info);

    let resource = Resource{
        attributes: get_resource_attributes(req, ctx),
        //dropped_attributes_count: 0
    };

    let random_trace_id = RandomIdGenerator::default().new_trace_id().to_string();
    let trace_id = match req.get_header_str("x-correlation-id") {
        Some(val) => val.clone(),
        _ => &random_trace_id,
    };

    let parent_span_id = match req.get_header_str("x-b3-parentspanid") {
        Some(val) => val.clone(),
        _ => &"",
    };

    let end_time = SystemTime::now().duration_since(UNIX_EPOCH)
        .unwrap_or_default()
        .as_nanos();

    let mut spans = Vec::new();
    spans.push(Span{
        trace_id: trace_id.to_string(),
        span_id: RandomIdGenerator::default().new_span_id().to_string(),
        //trace_state: "".to_string(),
        parent_span_id: parent_span_id.to_string(),
        name: req.get_path().to_string(),
        kind: 1, //INTERNAL
        start_time_unix_nano: start_time as u64,
        end_time_unix_nano: end_time as u64,
        attributes: get_span_attributes(req, ctx),
        dropped_attributes_count: 0,
        events: vec![],
        dropped_events_count: 0,
        //links: vec![],
        //dropped_links_count: 0,
        status: Some(Status{ message: "OK".to_string(), code: 1 }),
    });

    let scope = InstrumentationScope{
        name: "edge-router".to_string(),
        //version: std::env::var("CARGO_PKG_VERSION").unwrap_or_else(|_| String::new()),
        version: "1.0.0".to_string(),
        attributes: get_service_attributes(req, ctx),
        //dropped_attributes_count: 0,
    };

    let mut scope_spans = Vec::new();
    scope_spans.push(ScopeSpans{
        scope: Some(scope),
        spans: spans,
    });

    let mut resource_spans = Vec::new();
    resource_spans.push(ResourceSpans{
        resource: Some(resource),
        scope_spans: scope_spans,
    });

    let traces_data = TracesData{ resource_spans: resource_spans };
    let traces_data_json = serde_json::to_string(&traces_data)?;

    log::info!("{}", traces_data_json);

    Ok(())
}

pub(crate) fn observe_failed_distributed_transaction(_start_time: u128, _req: &mut Request) {
    //OK(())
}

// fn observe_distributed_transaction(req: &mut Request, ctx: &Context) -> Result<()> {
//     OK(())
// }

fn get_service_attributes(_req: &mut Request, _ctx: &Context) -> Vec<KeyValue> {

    let mut service_attributes = Vec::new();
    service_attributes.push(KeyValue{ key: "service.name".to_string(), value: Option::from(Value::StringValue("edge-router".to_string()))});
    //service_attributes.push(KeyValue{ key: "service.version".to_string(), value: Option::from(Value::StringValue(std::env::var("CARGO_PKG_VERSION").unwrap_or_else(|_| String::new())))});
    service_attributes.push(KeyValue{ key: "service.version".to_string(), value: Option::from(Value::StringValue("0.1".to_string()))});

    return service_attributes

}

fn get_resource_attributes(_req: &mut Request, ctx: &Context) -> Vec<KeyValue> {

    let mut resource_attributes = get_service_attributes(_req, ctx);

    resource_attributes.push(KeyValue{ key: "telemetry.sdk.language".to_string(), value: Option::from(Value::StringValue("rust".to_string()))});
    resource_attributes.push(KeyValue{ key: "telemetry.sdk.name".to_string(), value: Option::from(Value::StringValue("opentelemetry".to_string()))});
    resource_attributes.push(KeyValue{ key: "telemetry.sdk.version".to_string(), value: Option::from(Value::StringValue("1.0.1".to_string()))});
    resource_attributes.push(KeyValue{ key: "host.name".to_string(), value: Option::from(Value::StringValue(ctx.host.to_string()))});

    return resource_attributes

}

fn get_span_attributes(req: &mut Request, _ctx: &Context) -> Vec<KeyValue> {

    let mut span_attributes = get_service_attributes(req, _ctx);
    span_attributes.push(KeyValue{ key: "http.method".to_string(), value: Option::from(Value::StringValue(req.get_method_str().to_string()))});
    span_attributes.push(KeyValue{ key: "http.target".to_string(), value: Option::from(Value::StringValue(req.get_url_str().to_string()))});

    // span_attributes.push(KeyValue::new("http.method", req.get_method_str()));
    // span_attributes.push(KeyValue::new("http.target", req.get_url_str()));
    //span_attributes.push(KeyValue::new("http.host", req.get));
    //span_attributes.push(KeyValue::new("http.protocol", req.));
    //span_attributes.push(KeyValue::new("http.client_ip", req.));

    return span_attributes

}