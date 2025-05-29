package natsnodev1

import (
	"google.golang.org/protobuf/proto"
)

// Command defines an interface for commands with methods to retrieve name, topic, and topic wildcard.
type Command interface {
	CommandName() string

	CommandTopic() string

	CommandTopicWildcard() string
}

// Event represents a generic event with methods to retrieve its name, topic, and topic wildcard.
type Event interface {
	EventName() string

	EventTopic() string

	EventTopicWildcard() string
}

// Topic defines an interface that requires implementing a String method to represent the topic as a string.
type Topic interface {
	String() string
}

// Type represents an interface for types with associated metadata and topics for commands and events.
type Type interface {
	TypeName() string

	CommandTopic() string

	EventTopic() string
}

// SpecCommand represents a command that encapsulates a request, stream details, and metadata for execution.
// It includes protocol buffers message, NATS stream, command name, topic, and entity type information.
type SpecCommand struct {
	Request        proto.Message
	Stream         Stream
	Procedure      string
	CommandName    string
	CommandTopic   string
	EntityTypeName string
}

// SpecEvent represents an event specification used for multiplexing streams and event-based communication.
// It encapsulates details like the request payload, associated stream, event name, topic, and entity type.
type SpecEvent struct {
	Request        proto.Message
	Stream         Stream
	Procedure      string
	EventName      string
	EventTopic     string
	EntityTypeName string
}

// GetListenerGroup generates a unique listener group identifier by combining the type names of the source and sink.
func GetListenerGroup(source Type, sink Type) string {
	return "listener-" + source.TypeName() + "+" + sink.TypeName()
}
