use fastly::{SpecError, Response};
use fastly::http::{StatusCode, Version};
use fastly::http::header::CONTENT_TYPE;
use crate::headers;
use crate::headers::{GRPC_ACCEPT_ENCODING, GRPC_ENCODING};
use crate::protocol::{SpecProtocol, Protocol};

#[derive(serde::Serialize)]
pub(crate) struct ResultCode {
    message: String,
    status: String,
    r#type: String,
    api_status: u32,
    user_message: String,
}

#[derive(serde::Serialize)]
pub(crate) struct ErrorResponse {
    correlation_id: String,
    result_code: ResultCode,
}

pub(crate) fn create_error(err: &String) -> ErrorResponse {
    ErrorResponse {
        correlation_id: "".to_string(),
        result_code: ResultCode {
            message: err.to_string(),
            status: "".to_string(),
            r#type: "".to_string(),
            api_status: 0,
            user_message: "".to_string(),
        }
    }
}

pub(crate) fn fail(err: String, protocol: &Protocol) -> Result<Response, SpecError> {
    let error = create_error(&err);

    let response = match protocol.spec_protocol {
        SpecProtocol::HTTP => {
            println!("HTTP internal error: {}", &err);
            Response::from_status(StatusCode::INTERNAL_SERVER_ERROR)
                .with_body_json(&error).unwrap()
                .with_version(Version::HTTP_11)
        }
        SpecProtocol::Rest => {
            println!("REST internal error: {}", &err);
            Response::from_status(StatusCode::INTERNAL_SERVER_ERROR)
                .with_body_json(&error).unwrap()
                .with_version(Version::HTTP_11)
        }
        SpecProtocol::Graphql => {
            println!("GraphQL internal error: {}", &err);
            Response::from_status(StatusCode::INTERNAL_SERVER_ERROR)
                .with_body_json(&error).unwrap()
                .with_version(Version::HTTP_11)
        }
        SpecProtocol::Grpc => {
            println!("GRPC internal error: {}", &err);
            Response::from_status(StatusCode::OK)
                .with_body_json(&error).unwrap()
                .with_version(Version::HTTP_2)
                .with_header(CONTENT_TYPE, "application/grpc")
                .with_header(GRPC_ENCODING, "identity")
                .with_header(GRPC_ACCEPT_ENCODING, "gzip")
                .with_header(headers::GRPC_STATUS, "9")
                .with_header(headers::GRPC_MESSAGE, error.result_code.message)
        }
        SpecProtocol::GrpcWeb => {
            println!("GRPCWEB internal error: {}", &err);
            Response::from_status(StatusCode::INTERNAL_SERVER_ERROR)
                .with_body_json(&error).unwrap()
        }
        SpecProtocol::Connect => {
            println!("Connect internal error: {}", &err);
            Response::from_status(StatusCode::INTERNAL_SERVER_ERROR)
                .with_body_json(&error).unwrap()
        }
        SpecProtocol::Soap => {
            println!("SOAP internal error: {}", &err);
            Response::from_status(StatusCode::INTERNAL_SERVER_ERROR)
                .with_body_json(&error).unwrap()
        }
    };

    Ok(response)

}

pub(crate) fn misdirected(err: String, protocol: &Protocol) -> Result<Response, SpecError> {
    let error = create_error(&err);

    let response = match protocol.spec_protocol {

        SpecProtocol::HTTP=> {
            println!("HTTP misdirected error: {}", &err);
            Response::from_status(StatusCode::MISDIRECTED_REQUEST)
                .with_body_json(&error).unwrap()
                .with_version(Version::HTTP_11)
        }
        SpecProtocol::Rest => {
            println!("REST misdirected error: {}", &err);
            Response::from_status(StatusCode::MISDIRECTED_REQUEST)
                .with_body_json(&error).unwrap()
                .with_version(Version::HTTP_11)
        }
        SpecProtocol::Graphql => {
            println!("GraphQL misdirected error: {}", &err);
            Response::from_status(StatusCode::MISDIRECTED_REQUEST)
                .with_body_json(&error).unwrap()
                .with_version(Version::HTTP_11)
        }
        SpecProtocol::Grpc => {
            println!("GRPC misdirected error: {}", &err);
            Response::from_status(StatusCode::OK)
                .with_body_json(&error).unwrap()
                .with_version(Version::HTTP_2)
                .with_header(CONTENT_TYPE, "application/grpc")
                .with_header(GRPC_ENCODING, "identity")
                .with_header(GRPC_ACCEPT_ENCODING, "gzip")
                .with_header(headers::GRPC_STATUS, "9")
                .with_header(headers::GRPC_MESSAGE, error.result_code.message)
        }
        SpecProtocol::GrpcWeb => {
            println!("GRPCWEB misdirected error: {}", &err);
            Response::from_status(StatusCode::MISDIRECTED_REQUEST)
                .with_body_json(&error).unwrap()
        }
        SpecProtocol::Connect => {
            println!("Connect misdirected error: {}", &err);
            Response::from_status(StatusCode::MISDIRECTED_REQUEST)
                .with_body_json(&error).unwrap()
        }
        SpecProtocol::Soap => {
            println!("SOAP misdirected error: {}", err);
            Response::from_status(StatusCode::MISDIRECTED_REQUEST)
                .with_body_json(&error).unwrap()
        }
    };

    Ok(response)

}

pub(crate) fn precondition(err: String, protocol: &Protocol) -> Result<Response, SpecError> {
    let error = create_error(&err);

    let response = match protocol.spec_protocol {

        SpecProtocol::HTTP => {
            println!("HTTP precondition failed error: {}", &err);
            Response::from_status(StatusCode::PRECONDITION_FAILED)
                .with_body_json(&error).unwrap()
                .with_version(Version::HTTP_11)
        }
        SpecProtocol::Rest => {
            println!("REST precondition failed error: {}", &err);
            Response::from_status(StatusCode::PRECONDITION_FAILED)
                .with_body_json(&error).unwrap()
                .with_version(Version::HTTP_11)
        }
        SpecProtocol::Graphql => {
            println!("GraphQL precondition failed error: {}", &err);
            Response::from_status(StatusCode::PRECONDITION_FAILED)
                .with_body_json(&error).unwrap()
                .with_version(Version::HTTP_11)
        }
        SpecProtocol::Grpc => {
            println!("GRPC precondition failed error: {}", &err);
            Response::from_status(StatusCode::OK)
                .with_body_json(&error).unwrap()
                .with_version(Version::HTTP_2)
                .with_header(CONTENT_TYPE, "application/grpc")
                .with_header(GRPC_ENCODING, "identity")
                .with_header(GRPC_ACCEPT_ENCODING, "gzip")
                .with_header(headers::GRPC_STATUS, "9")
                .with_header(headers::GRPC_MESSAGE, error.result_code.message)
        }
        SpecProtocol::GrpcWeb => {
            println!("GRPCWEB precondition failed error: {}", &err);
            Response::from_status(StatusCode::PRECONDITION_FAILED)
                .with_body_json(&error).unwrap()
        }
        SpecProtocol::Connect => {
            println!("Connect precondition failed error: {}", &err);
            Response::from_status(StatusCode::PRECONDITION_FAILED)
                .with_body_json(&error).unwrap()
        }
        SpecProtocol::Soap => {
            println!("SOAP precondition failed error: {}", &err);
            Response::from_status(StatusCode::PRECONDITION_FAILED)
                .with_body_json(&error).unwrap()
        }
    };

    Ok(response)

}