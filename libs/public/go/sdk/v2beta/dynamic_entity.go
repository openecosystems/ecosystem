package sdkv2betalib

import (
	"encoding/json"
	"errors"

	specv2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/spec/v2"

	"google.golang.org/protobuf/types/dynamicpb"

	"google.golang.org/protobuf/types/known/anypb"
)

// DynamicCommand represents a dynamic command identifier as an integer type.
// DynamicEvent represents a dynamic event identifier as an integer type.
type (
	DynamicCommand int
	DynamicEvent   int
)

// DynamicTypeName represents the singular name for the dynamic type.
// DynamicTypeNamePlural represents the plural name for the dynamic type.
// DynamicTypeNameScreamingCamelCase represents the singular dynamic type in screaming camel case.
// DynamicTypeNamePluralScreamingCamelCase represents the plural dynamic type in screaming camel case.
// DynamicTypeNameEventPrefix represents the event prefix for dynamic type events.
const (
	DynamicTypeName                         string = "dynamic"
	DynamicTypeNamePlural                   string = "dynamics"
	DynamicTypeNameScreamingCamelCase       string = "DYNAMIC"
	DynamicTypeNamePluralScreamingCamelCase string = "DYNAMICS"
	DynamicTypeNameEventPrefix              string = "dynamic."
)

// DynamicCommandsUnspecified represents an unspecified dynamic command.
// DynamicCommandsCreate represents the creation of a dynamic command.
// DynamicCommandsUpdate represents the update of a dynamic command.
// DynamicCommandsDelete represents the deletion of a dynamic command.
// UnrecognizedDynamicCommand represents an unrecognized dynamic command.
const (
	DynamicCommandsUnspecified DynamicCommand = iota
	DynamicCommandsCreate      DynamicCommand = iota
	DynamicCommandsUpdate      DynamicCommand = iota
	DynamicCommandsDelete      DynamicCommand = iota
	UnrecognizedDynamicCommand DynamicCommand = -1
)

// DynamicEventsUnspecified represents an unspecified dynamic event.
// DynamicEventsCreated represents a created dynamic event.
// DynamicEventsUpdated represents an updated dynamic event.
// DynamicEventsDeleted represents a deleted dynamic event.
// UnrecognizedDynamicEvent represents an unrecognized dynamic event.
const (
	DynamicEventsUnspecified DynamicEvent = iota
	DynamicEventsCreated     DynamicEvent = iota
	DynamicEventsUpdated     DynamicEvent = iota
	DynamicEventsDeleted     DynamicEvent = iota
	UnrecognizedDynamicEvent DynamicEvent = -1
)

// CommandDataDynamicTopic represents the topic for dynamic data commands.
// EventDataDynamicTopic represents the topic for dynamic data events.
// RoutineDataDynamicTopic represents the topic for dynamic data routines.
// UnrecognizedDynamicTopic represents the topic for unrecognized dynamic data.
const (
	CommandDataDynamicTopic  string = "dynamic.data.command"
	EventDataDynamicTopic    string = "dynamic.data.event"
	RoutineDataDynamicTopic  string = "dynamic.data.routine"
	UnrecognizedDynamicTopic string = "unrecognized"
)

// CommandName returns the string representation of the DynamicCommand. It maps each command to a specific string value.
func (c DynamicCommand) CommandName() string {
	switch c {
	case DynamicCommandsUnspecified:
		return "DynamicCommandsUnspecified"
	case DynamicCommandsCreate:
		return "DynamicCommandsCreate"
	case DynamicCommandsUpdate:
		return "DynamicCommandsUpdate"
	case DynamicCommandsDelete:
		return "DynamicCommandsDelete"
	default:
		return "UnrecognizedDynamicCommand"
	}
}

// EventName returns a string representation of the DynamicEvent type.
func (e DynamicEvent) EventName() string {
	switch e {
	case DynamicEventsUnspecified:
		return "DynamicEventsUnspecified"
	case DynamicEventsCreated:
		return "DynamicEventsCreated"
	case DynamicEventsUpdated:
		return "DynamicEventsUpdated"
	case DynamicEventsDeleted:
		return "DynamicEventsDeleted"
	default:
		return "UnrecognizedDynamicEvent"
	}
}

// CommandTopic returns the topic string associated with the DynamicCommand based on its value.
func (c DynamicCommand) CommandTopic() string {
	switch c {
	case DynamicCommandsUnspecified:
		return CommandDataDynamicTopic
	case DynamicCommandsCreate:
		return CommandDataDynamicTopic
	case DynamicCommandsUpdate:
		return CommandDataDynamicTopic
	case DynamicCommandsDelete:
		return CommandDataDynamicTopic
	default:
		return UnrecognizedDynamicTopic
	}
}

// EventTopic returns the topic string associated with a DynamicEvent, or "unrecognized" for unknown events.
func (e DynamicEvent) EventTopic() string {
	switch e {
	case DynamicEventsUnspecified:
		return EventDataDynamicTopic
	case DynamicEventsCreated:
		return EventDataDynamicTopic
	case DynamicEventsUpdated:
		return EventDataDynamicTopic
	case DynamicEventsDeleted:
		return EventDataDynamicTopic
	default:
		return UnrecognizedDynamicTopic
	}
}

// CommandTopicWildcard returns a wildcard topic string for the command, using the DynamicTypeNameEventPrefix.
func (c DynamicCommand) CommandTopicWildcard() string {
	return DynamicTypeNameEventPrefix + ">"
}

// EventTopicWildcard returns the wildcard topic string for dynamic events, using the DynamicTypeNameEventPrefix constant.
func (e DynamicEvent) EventTopicWildcard() string {
	return DynamicTypeNameEventPrefix + ">"
}

// GetDynamicCommand returns a DynamicCommand based on the given command string.
// Maps known command strings to their respective enumerated DynamicCommand values.
// Returns UnrecognizedDynamicCommand for unknown or unhandled command strings.
func GetDynamicCommand(command string) DynamicCommand {
	switch command {
	case "DynamicCommandsUnspecified":
		return DynamicCommandsUnspecified
	case "DynamicCommandsCreate":
		return DynamicCommandsCreate
	case "DynamicCommandsUpdate":
		return DynamicCommandsUpdate
	case "DynamicCommandsDelete":
		return DynamicCommandsDelete
	default:
		return UnrecognizedDynamicCommand
	}
}

// GetDynamicEvent maps a string identifier to its corresponding DynamicEvent constant.
func GetDynamicEvent(event string) DynamicEvent {
	switch event {
	case "DynamicEventsUnspecified":
		return DynamicEventsUnspecified
	case "DynamicEventsCreated":
		return DynamicEventsCreated
	case "DynamicEventsUpdated":
		return DynamicEventsUpdated
	case "DynamicEventsDeleted":
		return DynamicEventsDeleted
	default:
		return UnrecognizedDynamicEvent
	}
}

// DynamicSpecEntity represents a dynamic specification entity with methods for serialization and deserialization.
type DynamicSpecEntity struct{}

// NewDynamicSpecEntity creates a new instance of DynamicSpecEntity using the provided SpecContext and returns an error if any.
func NewDynamicSpecEntity(_ *specv2pb.SpecContext) (*DynamicSpecEntity, error) {
	return &DynamicSpecEntity{}, nil
}

// ToProto converts the DynamicSpecEntity instance into a protobuf dynamicpb.Message representation.
func (entity *DynamicSpecEntity) ToProto() (*dynamicpb.Message, error) {
	return &dynamicpb.Message{}, nil
}

// ToEvent serializes the DynamicSpecEntity to a JSON string and returns it as a pointer, or an error if serialization fails.
func (entity *DynamicSpecEntity) ToEvent() (*string, error) {
	bytes, err := json.Marshal(entity)
	if err != nil {
		return nil, err
	}

	event := string(bytes)

	return &event, nil
}

// FromEvent populates the DynamicSpecEntity by unmarshaling a JSON string from the provided event pointer.
func (entity *DynamicSpecEntity) FromEvent(event *string) (*DynamicSpecEntity, error) {
	bytes := []byte(*event)
	err := json.Unmarshal(bytes, entity)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

// MarshalEntity serializes the DynamicSpecEntity into an Any protobuf message, returning an error if marshalling fails.
func (entity *DynamicSpecEntity) MarshalEntity() (*anypb.Any, error) {
	d, err := anypb.New(dynamicpb.NewMessage(nil))
	if err != nil {
		return nil, ErrServerInternal.WithInternalErrorDetail(errors.New("failed to marshall entity"), err)
	}

	return d, nil
}

// MarshalProto serializes the DynamicSpecEntity into a protobuf Any message and returns it or an error if the process fails.
func (entity *DynamicSpecEntity) MarshalProto() (*anypb.Any, error) {
	proto, err := entity.ToProto()
	if err != nil {
		return nil, ErrServerInternal.WithInternalErrorDetail(errors.New("failed to convert entity to proto"), err)
	}

	d, err := anypb.New(proto)
	if err != nil {
		return nil, ErrServerInternal.WithInternalErrorDetail(errors.New("failed to marshall proto"), err)
	}

	return d, nil
}

// TypeName returns the type name of the entity as a string.
func (entity *DynamicSpecEntity) TypeName() string {
	return "dynamic"
}

// CommandTopic returns the topic for command-related dynamic data processing.
func (entity *DynamicSpecEntity) CommandTopic() string {
	return CommandDataDynamicTopic
}

// EventTopic returns the event topic associated with the DynamicSpecEntity.
func (entity *DynamicSpecEntity) EventTopic() string {
	return EventDataDynamicTopic
}

// RoutineTopic returns the routine topic string associated with dynamic data.
func (entity *DynamicSpecEntity) RoutineTopic() string {
	return RoutineDataDynamicTopic
}

// TopicWildcard generates a topic wildcard string for dynamic type event subscription.
func (entity *DynamicSpecEntity) TopicWildcard() string {
	return DynamicTypeNameEventPrefix + ">"
}

// SystemName returns the system name associated with the DynamicSpecEntity, represented as a string.
func (entity *DynamicSpecEntity) SystemName() string {
	return "dynamic"
}
