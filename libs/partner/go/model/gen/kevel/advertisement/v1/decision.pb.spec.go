// Code generated by protoc-gen-platform go/spec. DO NOT EDIT.
// source: kevel/advertisement/v1/decision.proto

package advertisementv1pbmodel

type DecisionCommand int
type DecisionEvent int

const (
	DecisionTypeName                         string = "decision"
	DecisionTypeNamePlural                   string = "decisions"
	DecisionTypeNameScreamingCamelCase       string = "DECISION"
	DecisionTypeNamePluralScreamingCamelCase string = "DECISIONS"
	DecisionTypeNameEventPrefix              string = "decision."
)

const (
	UnrecognizedDecisionCommand DecisionCommand = -1
)

const (
	UnrecognizedDecisionEvent DecisionEvent = -1
)

const (
	CommandDataDecisionTopic  string = "decision.data.command"
	EventDataDecisionTopic    string = "decision.data.event"
	RoutineDataDecisionTopic  string = "decision.data.routine"
	UnrecognizedDecisionTopic string = "unrecognized"
)

func (c DecisionCommand) CommandName() string {

	switch c {

	default:
		return "UnrecognizedDecisionCommand"
	}

}

func (e DecisionEvent) EventName() string {

	switch e {

	default:
		return "UnrecognizedDecisionEvent"
	}

}

func (c DecisionCommand) CommandTopic() string {

	switch c {

	default:
		return UnrecognizedDecisionTopic
	}

}

func (e DecisionEvent) EventTopic() string {

	switch e {

	default:
		return UnrecognizedDecisionTopic
	}

}

func (c DecisionCommand) CommandTopicWildcard() string {
	return DecisionTypeNameEventPrefix + ">"
}

func (e DecisionEvent) EventTopicWildcard() string {
	return DecisionTypeNameEventPrefix + ">"
}

func GetDecisionCommand(command string) DecisionCommand {

	switch command {

	default:
		return UnrecognizedDecisionCommand
	}
}

func GetDecisionEvent(event string) DecisionEvent {

	switch event {

	default:
		return UnrecognizedDecisionEvent
	}
}