use serde::{Deserialize, Serialize};


// https://github.com/open-telemetry/opentelemetry-proto/blob/v1.0.0/opentelemetry/proto/trace/v1/trace.proto




/// AnyValue is used to represent any type of attribute value. AnyValue may contain a
/// primitive value such as a string or integer or it may contain an arbitrary nested
/// object containing arrays, key-value lists and primitives.
#[derive(Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub(crate) struct AnyValue {
    /// The value is one of the listed fields. It is valid for all values to be unspecified
    /// in which case this AnyValue is considered to be "empty".
    pub(crate) value: Option<any_value::Value>,
}

/// Nested message and enum types in `AnyValue`.
pub(crate) mod any_value {
    use serde::{Deserialize, Serialize};

    /// The value is one of the listed fields. It is valid for all values to be unspecified
    /// in which case this AnyValue is considered to be "empty".
    #[derive(Serialize, Deserialize)]
    #[serde(rename_all = "camelCase")]
    pub(crate) enum Value {
        StringValue(String),
        BoolValue(bool),
        IntValue(i64),
        DoubleValue(f64),
        ArrayValue(super::ArrayValue),
        KvlistValue(super::KeyValueList),
        BytesValue(Vec<u8>),
    }
}
/// ArrayValue is a list of AnyValue messages. We need ArrayValue as a message
/// since oneof in AnyValue does not allow repeated fields.
#[derive(Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub(crate) struct ArrayValue {
    /// Array of values. The array may be empty (contain 0 elements).
    pub(crate) values: Vec<AnyValue>,
}

/// KeyValueList is a list of KeyValue messages. We need KeyValueList as a message
/// since `oneof` in AnyValue does not allow repeated fields. Everywhere else where we need
/// a list of KeyValue messages (e.g. in Span) we use `repeated KeyValue` directly to
/// avoid unnecessary extra wrapping (which slows down the protocol). The 2 approaches
/// are semantically equivalent.
#[derive(Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub(crate) struct KeyValueList {
    /// A collection of key/value pairs of key-value pairs. The list may be empty (may
    /// contain 0 elements).
    /// The keys MUST be unique (it is not allowed to have more than one
    /// value with the same key).
    pub(crate) values: Vec<KeyValue>,
}
/// KeyValue is a key-value pair that is used to store Span attributes, Link
/// attributes, etc.
#[derive(Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub(crate) struct KeyValue {
    pub(crate) key: String,
    pub(crate) value: Option<any_value::Value>,
}
/// InstrumentationScope is a message representing the instrumentation scope information
/// such as the fully qualified name and version.
#[derive(Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub(crate) struct InstrumentationScope {
    /// An empty instrumentation scope name means the name is unknown.
    pub(crate) name: String,
    pub(crate) version: String,
    /// Additional attributes that describe the scope. \[Optional\].
    /// Attribute keys MUST be unique (it is not allowed to have more than one
    /// attribute with the same key).
    pub(crate) attributes: Vec<KeyValue>,
    //pub(crate) dropped_attributes_count: u32,
}

#[derive(Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub(crate) struct Resource {
    /// Set of attributes that describe the resource.
    /// Attribute keys MUST be unique (it is not allowed to have more than one
    /// attribute with the same key).
    pub(crate) attributes: Vec<KeyValue>,
    // /// dropped_attributes_count is the number of dropped attributes. If the value is 0, then
    // /// no attributes were dropped.
    // pub(crate) dropped_attributes_count: u32,
}

/// TracesData represents the traces data that can be stored in a persistent storage,
/// OR can be embedded by other protocols that transfer OTLP traces data but do
/// not implement the OTLP protocol.
///
/// The main difference between this message and collector protocol is that
/// in this message there will not be any "control" or "metadata" specific to
/// OTLP protocol.
///
/// When new fields are added into this message, the OTLP request MUST be updated
/// as well.
#[derive(Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub(crate) struct TracesData {
    /// An array of ResourceSpans.
    /// For data coming from a single resource this array will typically contain
    /// one element. Intermediary nodes that receive data from multiple origins
    /// typically batch the data before forwarding further and in that case this
    /// array will contain multiple elements.
    pub(crate) resource_spans: Vec<ResourceSpans>,
}
/// A collection of ScopeSpans from a Resource.
#[derive(Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub(crate) struct ResourceSpans {
    /// The resource for the spans in this message.
    /// If this field is not set then no resource info is known.
    pub(crate) resource: Option<Resource>,
    /// A list of ScopeSpans that originate from a resource.
    pub(crate) scope_spans: Vec<ScopeSpans>,
}
/// A collection of Spans produced by an InstrumentationScope.
#[derive(Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub(crate) struct ScopeSpans {
    /// The instrumentation scope information for the spans in this message.
    /// Semantically when InstrumentationScope isn't set, it is equivalent with
    /// an empty instrumentation scope name (unknown).
    pub(crate) scope: Option<InstrumentationScope>,
    /// A list of Spans that originate from an instrumentation scope.
    pub(crate) spans: Vec<Span>,
}

/// A Span represents a single operation performed by a single component of the system.
///
/// The next available field id is 17.
#[derive(Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub(crate) struct Span {
    /// A unique identifier for a trace. All spans from the same trace share
    /// the same `trace_id`. The ID is a 16-byte array. An ID with all zeroes OR
    /// of length other than 16 bytes is considered invalid (empty string in OTLP/JSON
    /// is zero-length and thus is also invalid).
    ///
    /// This field is required.
    pub(crate) trace_id: String,
    /// A unique identifier for a span within a trace, assigned when the span
    /// is created. The ID is an 8-byte array. An ID with all zeroes OR of length
    /// other than 8 bytes is considered invalid (empty string in OTLP/JSON
    /// is zero-length and thus is also invalid).
    ///
    /// This field is required.
    pub(crate) span_id: String,
    // /// trace_state conveys information about request position in multiple distributed tracing graphs.
    // /// It is a trace_state in w3c-trace-context format: <https://www.w3.org/TR/trace-context/#tracestate-header>
    // /// See also <https://github.com/w3c/distributed-tracing> for more details about this field.
    // pub(crate) trace_state: String,
    /// The `span_id` of this span's parent span. If this is a root span, then this
    /// field must be empty. The ID is an 8-byte array.
    pub(crate) parent_span_id: String,
    /// A description of the span's operation.
    ///
    /// For example, the name can be a qualified method name or a file name
    /// and a line number where the operation is called. A best practice is to use
    /// the same display name at the same call point in an application.
    /// This makes it easier to correlate spans in different traces.
    ///
    /// This field is semantically required to be set to non-empty string.
    /// Empty value is equivalent to an unknown span name.
    ///
    /// This field is required.
    pub(crate) name: String,
    /// Distinguishes between spans generated in a particular context. For example,
    /// two spans with the same name may be distinguished using `CLIENT` (caller)
    /// and `SERVER` (callee) to identify queueing latency associated with the span.
    pub(crate) kind: i32,
    /// start_time_unix_nano is the start time of the span. On the client side, this is the time
    /// kept by the local machine where the span execution starts. On the server side, this
    /// is the time when the server's application handler starts running.
    /// Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
    ///
    /// This field is semantically required and it is expected that end_time >= start_time.
    pub(crate) start_time_unix_nano: u64,
    /// end_time_unix_nano is the end time of the span. On the client side, this is the time
    /// kept by the local machine where the span execution ends. On the server side, this
    /// is the time when the server application handler stops running.
    /// Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
    ///
    /// This field is semantically required and it is expected that end_time >= start_time.
    pub(crate) end_time_unix_nano: u64,
    /// attributes is a collection of key/value pairs. Note, global attributes
    /// like server name can be set using the resource API. Examples of attributes:
    ///
    ///      "/http/user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"
    ///      "/http/server_latency": 300
    ///      "example.com/myattribute": true
    ///      "example.com/score": 10.239
    ///
    /// The OpenTelemetry API specification further restricts the allowed value types:
    /// <https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/common/README.md#attribute>
    /// Attribute keys MUST be unique (it is not allowed to have more than one
    /// attribute with the same key).
    pub(crate) attributes: Vec<KeyValue>,
    /// dropped_attributes_count is the number of attributes that were discarded. Attributes
    /// can be discarded because their keys are too long or because there are too many
    /// attributes. If this value is 0, then no attributes were dropped.
    pub(crate) dropped_attributes_count: u32,
    /// events is a collection of Event items.
    pub(crate) events: Vec<span::Event>,
    /// dropped_events_count is the number of dropped events. If the value is 0, then no
    /// events were dropped.
    pub(crate) dropped_events_count: u32,
    // /// links is a collection of Links, which are references from this span to a span
    // /// in the same or different trace.
    // pub(crate) links: Vec<span::Link>,
    // /// dropped_links_count is the number of dropped links after the maximum size was
    // /// enforced. If this value is 0, then no links were dropped.
    // pub(crate) dropped_links_count: u32,
    /// An optional final status for this span. Semantically when Status isn't set, it means
    /// span's status code is unset, i.e. assume STATUS_CODE_UNSET (code = 0).
    pub(crate) status: Option<Status>,
}
/// Nested message and enum types in `Span`.
pub(crate) mod span {
    use serde::{Deserialize, Serialize};

    /// Event is a time-stamped annotation of the span, consisting of user-supplied
    /// text description and key-value pairs.
    #[derive(Serialize, Deserialize)]
    #[serde(rename_all = "camelCase")]
    pub(crate) struct Event {
        /// time_unix_nano is the time the event occurred.
        pub(crate) time_unix_nano: u64,
        /// name of the event.
        /// This field is semantically required to be set to non-empty string.
        pub(crate) name: String,
        /// attributes is a collection of attribute key/value pairs on the event.
        /// Attribute keys MUST be unique (it is not allowed to have more than one
        /// attribute with the same key).
        pub(crate) attributes: Vec<
            super::KeyValue,
        >,
        /// dropped_attributes_count is the number of dropped attributes. If the value is 0,
        /// then no attributes were dropped.
        pub(crate) dropped_attributes_count: u32,
    }
    /// A pointer from the current span to another span in the same trace or in a
    /// different trace. For example, this can be used in batching operations,
    /// where a single batch handler processes multiple requests from different
    /// traces or when the handler receives a request from a different project.
    #[derive(Serialize, Deserialize)]
    #[serde(rename_all = "camelCase")]
    pub(crate) struct Link {
        /// A unique identifier of a trace that this linked span is part of. The ID is a
        /// 16-byte array.
        pub(crate) trace_id: Vec<u8>,
        /// A unique identifier for the linked span. The ID is an 8-byte array.
        pub(crate) span_id: Vec<u8>,
        /// The trace_state associated with the link.
        pub(crate) trace_state: String,
        /// attributes is a collection of attribute key/value pairs on the link.
        /// Attribute keys MUST be unique (it is not allowed to have more than one
        /// attribute with the same key).
        pub(crate) attributes: Vec<
            super::KeyValue,
        >,
        /// dropped_attributes_count is the number of dropped attributes. If the value is 0,
        /// then no attributes were dropped.
        pub(crate) dropped_attributes_count: u32,
    }
    /// SpanKind is the type of span. Can be used to specify additional relationships between spans
    /// in addition to a parent/child relationship.
    #[derive(
    Clone,
    Copy,
    Debug,
    PartialEq,
    Eq,
    Hash,
    PartialOrd,
    Ord,
    )]
    #[repr(i32)]
    #[derive(Serialize, Deserialize)]
    #[serde(rename_all = "camelCase")]
    pub(crate) enum SpanKind {
        /// Unspecified. Do NOT use as default.
        /// Implementations MAY assume SpanKind to be INTERNAL when receiving UNSPECIFIED.
        Unspecified = 0,
        /// Indicates that the span represents an internal operation within an application,
        /// as opposed to an operation happening at the boundaries. Default value.
        Internal = 1,
        /// Indicates that the span covers server-side handling of an RPC or other
        /// remote network request.
        Server = 2,
        /// Indicates that the span describes a request to some remote service.
        Client = 3,
        /// Indicates that the span describes a producer sending a message to a broker.
        /// Unlike CLIENT and SERVER, there is often no direct critical path latency relationship
        /// between producer and consumer spans. A PRODUCER span ends when the message was accepted
        /// by the broker while the logical processing of the message might span a much longer time.
        Producer = 4,
        /// Indicates that the span describes consumer receiving a message from a broker.
        /// Like the PRODUCER kind, there is often no direct critical path latency relationship
        /// between producer and consumer spans.
        Consumer = 5,
    }
    impl SpanKind {
        /// String value of the enum field names used in the ProtoBuf definition.
        ///
        /// The values are not transformed in any way and thus are considered stable
        /// (if the ProtoBuf definition does not change) and safe for programmatic use.
        pub(crate) fn as_str_name(&self) -> &'static str {
            match self {
                SpanKind::Unspecified => "SPAN_KIND_UNSPECIFIED",
                SpanKind::Internal => "SPAN_KIND_INTERNAL",
                SpanKind::Server => "SPAN_KIND_SERVER",
                SpanKind::Client => "SPAN_KIND_CLIENT",
                SpanKind::Producer => "SPAN_KIND_PRODUCER",
                SpanKind::Consumer => "SPAN_KIND_CONSUMER",
            }
        }
        /// Creates an enum from field names used in the ProtoBuf definition.
        pub(crate) fn from_str_name(value: &str) -> Option<Self> {
            match value {
                "SPAN_KIND_UNSPECIFIED" => Some(Self::Unspecified),
                "SPAN_KIND_INTERNAL" => Some(Self::Internal),
                "SPAN_KIND_SERVER" => Some(Self::Server),
                "SPAN_KIND_CLIENT" => Some(Self::Client),
                "SPAN_KIND_PRODUCER" => Some(Self::Producer),
                "SPAN_KIND_CONSUMER" => Some(Self::Consumer),
                _ => None,
            }
        }
    }
}
/// The Status type defines a logical error model that is suitable for different
/// programming environments, including REST APIs and RPC APIs.
#[derive(Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub(crate) struct Status {
    /// A developer-facing human readable error message.
    pub(crate) message: String,
    /// The status code.
    pub(crate) code: i32,
}
/// Nested message and enum types in `Status`.
pub mod status {
    use serde::{Deserialize, Serialize};

    /// For the semantics of status codes see
    /// <https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/api.md#set-status>
    #[derive(
    Clone,
    Copy,
    Debug,
    PartialEq,
    Eq,
    Hash,
    PartialOrd,
    Ord,
    )]
    #[repr(i32)]
    #[derive(Serialize, Deserialize)]
    #[serde(rename_all = "camelCase")]
    pub(crate) enum StatusCode {
        /// The default status.
        Unset = 0,
        /// The Span has been validated by an Application developer or Operator to
        /// have completed successfully.
        Ok = 1,
        /// The Span contains an error.
        SpecError = 2,
    }
    impl StatusCode {
        /// String value of the enum field names used in the ProtoBuf definition.
        ///
        /// The values are not transformed in any way and thus are considered stable
        /// (if the ProtoBuf definition does not change) and safe for programmatic use.
        pub(crate) fn as_str_name(&self) -> &'static str {
            match self {
                StatusCode::Unset => "STATUS_CODE_UNSET",
                StatusCode::Ok => "STATUS_CODE_OK",
                StatusCode::SpecError => "STATUS_CODE_ERROR",
            }
        }
        /// Creates an enum from field names used in the ProtoBuf definition.
        pub(crate) fn from_str_name(value: &str) -> Option<Self> {
            match value {
                "STATUS_CODE_UNSET" => Some(Self::Unset),
                "STATUS_CODE_OK" => Some(Self::Ok),
                "STATUS_CODE_ERROR" => Some(Self::SpecError),
                _ => None,
            }
        }
    }
}
