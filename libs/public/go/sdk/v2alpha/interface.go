package sdkv2alphalib

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
)

// Entity represents an interface for defining domain-specific entities with serialization and messaging capabilities.
// ToEvent converts the entity into an event representation as a string pointer.
// MarshalEntity serializes the entity into a protocol buffer Any type.
// MarshalProto serializes the Proto object representation of the entity into a protocol buffer Any type.
// TypeName returns the name of the entity type as a string.
// CommandTopic retrieves the messaging topic for commands associated with the entity.
// EventTopic retrieves the messaging topic for events associated with the entity.
// RoutineTopic retrieves the messaging topic for routines related to the entity.
// TopicWildcard retrieves the wildcard string for topic-based messaging for the entity type.
// SystemName returns the name of the system to which the entity belongs.
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

// Connector defines an interface for managing and resolving methods by their paths.
type Connector interface {
	MethodsByPath() map[string]*Method
}

// Service is an interface that defines a contract for implementing service-related functionalities.
type Service interface{}

// Method defines an interface for describing a method in a protocol with its name, input, output, and schema details.
type Method interface {
	ProcedureName() string
	Input() protoreflect.MessageDescriptor
	Output() protoreflect.MessageDescriptor
	Schema() protoreflect.MethodDescriptor
}

// Client represents an abstract interface for defining client-side functionality or behavior.
type Client interface{}
