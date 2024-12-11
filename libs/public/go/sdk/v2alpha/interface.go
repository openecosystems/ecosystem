package sdkv2alphalib

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
)

type Entity interface {
	// ToProto() (*interface{}, error)

	ToEvent() (*string, error)

	MarshalEntity() (*anypb.Any, error)

	MarshalProto() (*anypb.Any, error)

	TypeName() string

	CommandTopic() string

	EventTopic() string

	RoutineTopic() string

	TopicWildcard() string

	SystemName() string
}

type Server interface{}

type Connector interface {
	MethodsByPath() map[string]*Method
}

type Service interface{}

type Method interface {
	ProcedureName() string
	Input() protoreflect.MessageDescriptor
	Output() protoreflect.MessageDescriptor
	Schema() protoreflect.MethodDescriptor
}

type Client interface{}
