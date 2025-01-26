package natsnodev2

import (
	typev2pb "libs/protobuf/go/protobuf/gen/platform/type/v2"

	"github.com/nats-io/nats.go"
)

// Stream represents an interface for defining operations related to event plane stream types.
// EventPlaneStreamType provides the stream type as defined by typev2pb.Stream.
// StreamPrefix returns the prefix associated with the stream.
type Stream interface {
	EventPlaneStreamType() typev2pb.Stream
	StreamPrefix() string
}

// InboundStream represents a configuration for an inbound NATS JetStream, including topic wildcard and stream information.
type InboundStream struct {
	TopicWildcard string
	Stream        *nats.StreamInfo
	StreamName    string
}

// InternalStream represents an internal stream configuration used for managing NATS stream details and topic wildcard.
type InternalStream struct {
	TopicWildcard string
	Stream        *nats.StreamInfo
	StreamName    string
}

// OutboundStream represents an outbound stream configuration and metadata.
// TopicWildcard specifies the topic pattern for the stream.
// Stream holds information about the nats stream.
// StreamName defines the name of the stream.
type OutboundStream struct {
	TopicWildcard string
	Stream        *nats.StreamInfo
	StreamName    string
}

// NewInboundStream creates and returns a new instance of InboundStream, representing an inbound event stream.
func NewInboundStream() *InboundStream {
	s := new(InboundStream)
	return s
}

// EventPlaneStreamType returns the type of the stream as typev2pb.Stream_STREAM_INBOUND for the InboundStream.
func (s *InboundStream) EventPlaneStreamType() typev2pb.Stream {
	return typev2pb.Stream_STREAM_INBOUND
}

// StreamPrefix returns the prefix string "inbound" associated with the InboundStream.
func (s *InboundStream) StreamPrefix() string {
	return "inbound"
}

// NewInternalStream creates a new instance of InternalStream and returns its pointer.
func NewInternalStream() *InternalStream {
	s := new(InternalStream)
	return s
}

// EventPlaneStreamType returns the stream type for the internal event plane stream, which is STREAM_INTERNAL.
func (s *InternalStream) EventPlaneStreamType() typev2pb.Stream {
	return typev2pb.Stream_STREAM_INTERNAL
}

// StreamPrefix returns the prefix string for identifying the internal stream type.
func (s *InternalStream) StreamPrefix() string {
	return "internal"
}

// NewOutboundStream initializes and returns a new OutboundStream instance.
func NewOutboundStream() *OutboundStream {
	s := new(OutboundStream)
	return s
}

// EventPlaneStreamType returns the default event plane stream type for outbound streams as `typev2pb.Stream_STREAM_OUTBOUND`.
func (s *OutboundStream) EventPlaneStreamType() typev2pb.Stream {
	return typev2pb.Stream_STREAM_OUTBOUND
}

// StreamPrefix returns the prefix string used for identifying outbound streams.
func (s *OutboundStream) StreamPrefix() string {
	return "outbound"
}
