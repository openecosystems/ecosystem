use fastly::http::Version;
use fastly::Request;

// JSON Content Type
pub(crate) const  PROTOCOL_APPLICATION_JSON: &str = "application/json";

// GRPC Protocols
pub(crate) const  PROTOCOL_APPLICATION_GRPC: &str = "application/grpc";
pub(crate) const  PROTOCOL_APPLICATION_GRPC_PROTO: &str = "application/grpc+proto";
pub(crate) const  PROTOCOL_APPLICATION_GRPC_JSON: &str = "application/grpc+json";

// GRPC Web Protocols
pub(crate) const  PROTOCOL_APPLICATION_GRPC_WEB: &str = "application/grpc-web";
pub(crate) const  PROTOCOL_APPLICATION_GRPC_WEB_PROTO: &str = "application/grpc-web+proto";
pub(crate) const  PROTOCOL_APPLICATION_GRPC_WEB_JSON: &str = "application/grpc-web+json";

// Connect Protocols
pub(crate) const  PROTOCOL_APPLICATION_PROTO: &str = "application/proto";


// Connect Streaming Protocols
pub(crate) const  PROTOCOL_APPLICATION_CONNECT_PROTO: &str = "application/connect+proto";
pub(crate) const  PROTOCOL_APPLICATION_CONNECT_JSON: &str = "application/connect+json";

// GraphQL
pub(crate) const  PROTOCOL_APPLICATION_GRAPHQL: &str = "application/graphql";
pub(crate) const  PROTOCOL_APPLICATION_GRAPHQL_JSON: &str = "application/graphql+json";

// Soap
pub(crate) const  PROTOCOL_APPLICATION_SOAP: &str = "application/soap";
pub(crate) const  PROTOCOL_APPLICATION_SOAP_XML: &str = "application/soap+xml";


#[derive(serde::Serialize, Debug)]
pub(crate) enum SpecProtocol {
    HTTP,
    Rest,
    Graphql,
    Grpc,
    GrpcWeb,
    Connect,
    Soap
}

#[derive(serde::Serialize, Debug)]
pub(crate) struct Protocol {
    pub content_type: String,
    //pub http_protocol: Version,
    pub spec_protocol: SpecProtocol,
}

impl Protocol {
    pub(crate) fn get_spec_protocol(&self) -> SpecProtocol {
        todo!()

    }
}


//
pub(crate) fn get_protocol(req: &mut Request) -> Protocol {

    match req.get_version() {

        Version::HTTP_09 => {
            return determine_known_mime_types(req, Version::HTTP_09);
        }

        Version::HTTP_10 => {
            return determine_known_mime_types(req, Version::HTTP_10);
        }

        Version::HTTP_11 => {
            return determine_known_mime_types(req, Version::HTTP_11);
        }

        // https://chromium.googlesource.com/external/github.com/grpc/grpc/+/HEAD/doc/PROTOCOL-HTTP2.md
        Version::HTTP_2 => {
            return determine_known_mime_types(req, Version::HTTP_2);
        },

        Version::HTTP_3 => {
            return determine_known_mime_types(req, Version::HTTP_3);
        },

        _ => {
            return Protocol{
                content_type: "".to_string(),
                //http_protocol: Version::HTTP_2,
                spec_protocol: SpecProtocol::Rest,
            }
        },
    };
}

//
fn determine_known_mime_types(req: &mut Request, _: Version) -> Protocol {

    let mime_type = req.get_content_type();
    let m = mime_type.clone();
    match mime_type {
        None => {}
        Some(mime) => {

            if mime.type_() == fastly::mime::APPLICATION
            {

                match mime.subtype().as_str() {
                    "grpc" => {
                        return Protocol{
                            content_type: mime.to_string(),
                            //http_protocol: version,
                            spec_protocol: SpecProtocol::Grpc,
                        }
                    },
                    "grpc-web" => {
                        return Protocol{
                            content_type: mime.to_string(),
                            //http_protocol: version,
                            spec_protocol: SpecProtocol::GrpcWeb,
                        }
                    },
                    "connect" => {
                        return Protocol{
                            content_type: mime.to_string(),
                            //http_protocol: version,
                            spec_protocol: SpecProtocol::Connect,
                        }
                    },
                    "graphql" => {
                        return Protocol{
                            content_type: mime.to_string(),
                            //http_protocol: version,
                            spec_protocol: SpecProtocol::Graphql,
                        }
                    },
                    "soap" => {
                        return Protocol{
                            content_type: mime.to_string(),
                            //http_protocol: version,
                            spec_protocol: SpecProtocol::Soap,
                        }
                    },
                    "json" => {
                        match req.get_path() {
                            "/graphql" => {
                                return Protocol{
                                    content_type: mime.to_string(),
                                    //http_protocol: version,
                                    spec_protocol: SpecProtocol::Graphql,
                                }
                            },
                            _ => {
                                // TODO: Refine this. Find a better way to check for the Connect case without using periods
                                // Handle Connect Case which can either be HTTP/1.1 or HTTP/2 with either application/proto or application/json
                                if req.get_path().contains(".") {
                                    return Protocol{
                                        content_type: mime.to_string(),
                                        //http_protocol: version,
                                        spec_protocol: SpecProtocol::Connect,
                                    }
                                } else {
                                    return Protocol{
                                        content_type: mime.to_string(),
                                        //http_protocol: version,
                                        spec_protocol: SpecProtocol::Rest,
                                    }
                                }
                            }
                        }
                    },
                    _ => {}
                }
            } else {
                return Protocol{
                    content_type: mime.to_string(),
                    //http_protocol: version,
                    spec_protocol: SpecProtocol::Rest,
                }
            }
        }
    };

    Protocol{
        content_type: "".to_string(),
        //http_protocol: version,
        spec_protocol: SpecProtocol::HTTP,
    }
}
