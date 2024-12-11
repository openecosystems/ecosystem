package sdkv2alphalib

import (
	"encoding/json"
	"errors"

	specv2pb "libs/protobuf/go/protobuf/gen/platform/spec/v2"

	"google.golang.org/protobuf/types/dynamicpb"

	"google.golang.org/protobuf/types/known/anypb"
)

type (
	DynamicCommand int
	DynamicEvent   int
)

const (
	DynamicTypeName                         string = "dynamic"
	DynamicTypeNamePlural                   string = "dynamics"
	DynamicTypeNameScreamingCamelCase       string = "DYNAMIC"
	DynamicTypeNamePluralScreamingCamelCase string = "DYNAMICS"
	DynamicTypeNameEventPrefix              string = "dynamic."
)

const (
	DynamicCommandsUnspecified DynamicCommand = iota
	DynamicCommandsCreate      DynamicCommand = iota
	DynamicCommandsUpdate      DynamicCommand = iota
	DynamicCommandsDelete      DynamicCommand = iota
	UnrecognizedDynamicCommand DynamicCommand = -1
)

const (
	DynamicEventsUnspecified DynamicEvent = iota
	DynamicEventsCreated     DynamicEvent = iota
	DynamicEventsUpdated     DynamicEvent = iota
	DynamicEventsDeleted     DynamicEvent = iota
	UnrecognizedDynamicEvent DynamicEvent = -1
)

const (
	CommandDataDynamicTopic  string = "dynamic.data.command"
	EventDataDynamicTopic    string = "dynamic.data.event"
	RoutineDataDynamicTopic  string = "dynamic.data.routine"
	UnrecognizedDynamicTopic string = "unrecognized"
)

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

func (c DynamicCommand) CommandTopicWildcard() string {
	return DynamicTypeNameEventPrefix + ">"
}

func (e DynamicEvent) EventTopicWildcard() string {
	return DynamicTypeNameEventPrefix + ">"
}

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

type DynamicSpecEntity struct{}

func NewDynamicSpecEntity(_ *specv2pb.SpecContext) (*DynamicSpecEntity, error) {
	return &DynamicSpecEntity{}, nil
}

func (entity *DynamicSpecEntity) ToProto() (*dynamicpb.Message, error) {
	return &dynamicpb.Message{}, nil
}

func (entity *DynamicSpecEntity) ToEvent() (*string, error) {
	bytes, err := json.Marshal(entity)
	if err != nil {
		return nil, err
	}

	event := string(bytes)

	return &event, nil
}

func (entity *DynamicSpecEntity) FromEvent(event *string) (*DynamicSpecEntity, error) {
	bytes := []byte(*event)
	err := json.Unmarshal(bytes, entity)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (entity *DynamicSpecEntity) MarshalEntity() (*anypb.Any, error) {
	d, err := anypb.New(dynamicpb.NewMessage(nil))
	if err != nil {
		return nil, ErrServerInternal.WithInternalErrorDetail(errors.New("failed to marshall entity"), err)
	}

	return d, nil
}

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

func (entity *DynamicSpecEntity) TypeName() string {
	return "dynamic"
}

func (entity *DynamicSpecEntity) CommandTopic() string {
	return CommandDataDynamicTopic
}

func (entity *DynamicSpecEntity) EventTopic() string {
	return EventDataDynamicTopic
}

func (entity *DynamicSpecEntity) RoutineTopic() string {
	return RoutineDataDynamicTopic
}

func (entity *DynamicSpecEntity) TopicWildcard() string {
	return DynamicTypeNameEventPrefix + ">"
}

func (entity *DynamicSpecEntity) SystemName() string {
	return "dynamic"
}
