package natsnodev2

import (
	"google.golang.org/protobuf/proto"
)

type Command interface {
	CommandName() string

	CommandTopic() string

	CommandTopicWildcard() string
}

type Event interface {
	EventName() string

	EventTopic() string

	EventTopicWildcard() string
}

type Topic interface {
	String() string
}

type Type interface {
	TypeName() string

	CommandTopic() string

	EventTopic() string
}

type SpecCommand struct {
	Request        proto.Message
	Stream         Stream
	CommandName    string
	CommandTopic   string
	EntityTypeName string
}

type SpecEvent struct {
	Request        proto.Message
	Stream         Stream
	EventName      string
	EventTopic     string
	EntityTypeName string
}

func GetListenerGroup(source Type, sink Type) string {
	return "listener-" + source.TypeName() + "+" + sink.TypeName()
}
