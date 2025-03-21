// Code generated by protoc-gen-platform go/spec. DO NOT EDIT.
// source: platform/communication/v1beta/preference_center.proto

package communicationv1betapb

type PreferenceCenterCommand int
type PreferenceCenterEvent int

const (
	PreferenceCenterTypeName                         string = "preferenceCenter"
	PreferenceCenterTypeNamePlural                   string = "preferenceCenters"
	PreferenceCenterTypeNameScreamingCamelCase       string = "PREFERENCE_CENTER"
	PreferenceCenterTypeNamePluralScreamingCamelCase string = "PREFERENCE_CENTERS"
	PreferenceCenterTypeNameEventPrefix              string = "preferenceCenter."
)

const (
	PreferenceCenterCommandsUnspecified        PreferenceCenterCommand = iota
	PreferenceCenterCommandsSubscribe          PreferenceCenterCommand = iota
	PreferenceCenterCommandsUpdateSubscription PreferenceCenterCommand = iota
	PreferenceCenterCommandsUnsubscribe        PreferenceCenterCommand = iota
	UnrecognizedPreferenceCenterCommand        PreferenceCenterCommand = -1
)

const (
	PreferenceCenterEventsUnspecified  PreferenceCenterEvent = iota
	PreferenceCenterEventsSubscribed   PreferenceCenterEvent = iota
	PreferenceCenterEventsUpdated      PreferenceCenterEvent = iota
	PreferenceCenterEventsUnsubscribed PreferenceCenterEvent = iota
	UnrecognizedPreferenceCenterEvent  PreferenceCenterEvent = -1
)

const (
	CommandDataPreferenceCenterTopic  string = "preferenceCenter.data.command"
	EventDataPreferenceCenterTopic    string = "preferenceCenter.data.event"
	RoutineDataPreferenceCenterTopic  string = "preferenceCenter.data.routine"
	UnrecognizedPreferenceCenterTopic string = "unrecognized"
)

func (c PreferenceCenterCommand) CommandName() string {

	switch c {

	case PreferenceCenterCommandsUnspecified:
		return "PreferenceCenterCommandsUnspecified"
	case PreferenceCenterCommandsSubscribe:
		return "PreferenceCenterCommandsSubscribe"
	case PreferenceCenterCommandsUpdateSubscription:
		return "PreferenceCenterCommandsUpdateSubscription"
	case PreferenceCenterCommandsUnsubscribe:
		return "PreferenceCenterCommandsUnsubscribe"
	default:
		return "UnrecognizedPreferenceCenterCommand"
	}

}

func (e PreferenceCenterEvent) EventName() string {

	switch e {

	case PreferenceCenterEventsUnspecified:
		return "PreferenceCenterEventsUnspecified"
	case PreferenceCenterEventsSubscribed:
		return "PreferenceCenterEventsSubscribed"
	case PreferenceCenterEventsUpdated:
		return "PreferenceCenterEventsUpdated"
	case PreferenceCenterEventsUnsubscribed:
		return "PreferenceCenterEventsUnsubscribed"
	default:
		return "UnrecognizedPreferenceCenterEvent"
	}

}

func (c PreferenceCenterCommand) CommandTopic() string {

	switch c {

	case PreferenceCenterCommandsUnspecified:
		return CommandDataPreferenceCenterTopic
	case PreferenceCenterCommandsSubscribe:
		return CommandDataPreferenceCenterTopic
	case PreferenceCenterCommandsUpdateSubscription:
		return CommandDataPreferenceCenterTopic
	case PreferenceCenterCommandsUnsubscribe:
		return CommandDataPreferenceCenterTopic
	default:
		return UnrecognizedPreferenceCenterTopic
	}

}

func (e PreferenceCenterEvent) EventTopic() string {

	switch e {

	case PreferenceCenterEventsUnspecified:
		return EventDataPreferenceCenterTopic
	case PreferenceCenterEventsSubscribed:
		return EventDataPreferenceCenterTopic
	case PreferenceCenterEventsUpdated:
		return EventDataPreferenceCenterTopic
	case PreferenceCenterEventsUnsubscribed:
		return EventDataPreferenceCenterTopic
	default:
		return UnrecognizedPreferenceCenterTopic
	}

}

func (c PreferenceCenterCommand) CommandTopicWildcard() string {
	return PreferenceCenterTypeNameEventPrefix + ">"
}

func (e PreferenceCenterEvent) EventTopicWildcard() string {
	return PreferenceCenterTypeNameEventPrefix + ">"
}

func GetPreferenceCenterCommand(command string) PreferenceCenterCommand {

	switch command {

	case "PreferenceCenterCommandsUnspecified":
		return PreferenceCenterCommandsUnspecified
	case "PreferenceCenterCommandsSubscribe":
		return PreferenceCenterCommandsSubscribe
	case "PreferenceCenterCommandsUpdateSubscription":
		return PreferenceCenterCommandsUpdateSubscription
	case "PreferenceCenterCommandsUnsubscribe":
		return PreferenceCenterCommandsUnsubscribe
	default:
		return UnrecognizedPreferenceCenterCommand
	}
}

func GetPreferenceCenterEvent(event string) PreferenceCenterEvent {

	switch event {

	case "PreferenceCenterEventsUnspecified":
		return PreferenceCenterEventsUnspecified
	case "PreferenceCenterEventsSubscribed":
		return PreferenceCenterEventsSubscribed
	case "PreferenceCenterEventsUpdated":
		return PreferenceCenterEventsUpdated
	case "PreferenceCenterEventsUnsubscribed":
		return PreferenceCenterEventsUnsubscribed
	default:
		return UnrecognizedPreferenceCenterEvent
	}
}
