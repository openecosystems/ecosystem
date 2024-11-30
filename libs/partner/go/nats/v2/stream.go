package natsnodev2

import (
	"github.com/nats-io/nats.go"
	"libs/protobuf/go/protobuf/gen/platform/type/v2"
)

type Stream interface {
	EventPlaneStreamType() typev2pb.Stream
	StreamPrefix() string
}

type InboundStream struct {
	TopicWildcard string
	Stream        *nats.StreamInfo
	StreamName    string
}

type InternalStream struct {
	TopicWildcard string
	Stream        *nats.StreamInfo
	StreamName    string
}

type OutboundStream struct {
	TopicWildcard string
	Stream        *nats.StreamInfo
	StreamName    string
}

func NewInboundStream() *InboundStream {
	s := new(InboundStream)
	return s
}

func (s *InboundStream) EventPlaneStreamType() typev2pb.Stream {

	return typev2pb.Stream_STREAM_INBOUND
}

func (s *InboundStream) StreamPrefix() string {
	return "inbound"
}

func NewInternalStream() *InternalStream {
	s := new(InternalStream)
	return s
}

func (s *InternalStream) EventPlaneStreamType() typev2pb.Stream {
	return typev2pb.Stream_STREAM_INTERNAL
}

func (s *InternalStream) StreamPrefix() string {
	return "internal"
}

func NewOutboundStream() *OutboundStream {
	s := new(OutboundStream)
	return s
}

func (s *OutboundStream) EventPlaneStreamType() typev2pb.Stream {
	return typev2pb.Stream_STREAM_OUTBOUND
}

func (s *OutboundStream) StreamPrefix() string {
	return "outbound"
}
