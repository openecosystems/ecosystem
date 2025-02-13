// Code generated by protoc-gen-platform typescript/spec. DO NOT EDIT.
// source: kevel/advertisement/v1/decision.proto

export type DecisionCommand = number;
export type DecisionEvent = number;

// Constants for Decision Type Names
export const DecisionTypeName = "decision";
export const DecisionTypeNamePlural = "decisions";
export const DecisionTypeNameScreamingCamelCase = "DECISION";
export const DecisionTypeNamePluralScreamingCamelCase = "DECISIONS";
export const DecisionTypeNameEventPrefix = "decision.";

// Enums for DecisionCommands
export enum DecisionCommands {
  
  DecisionCommandsUnspecified = 0,
  DecisionCommandsOptOut = 1,
  UnrecognizedDecisionCommand = -1,
}

// Enums for Decision Events
export enum DecisionEvents {

  DecisionEventsUnspecified = 0,
  DecisionEventsOptedOut = 1,
  UnrecognizedDecisionEvent  = -1,
}

// Topics
export const CommandDataDecisionTopic = "decision.data.command";
export const EventDataDecisionTopic = "decision.data.event";
export const RoutineDataDecisionTopic = "decision.data.routine";
export const UnrecognizedDecisionTopic = "unrecognized";

// Command Methods
export class DecisionCommandHelper {
  static commandName(command: DecisionCommands): string {
    switch (command) {
      case DecisionCommands.DecisionCommandsUnspecified:
        return "DecisionCommandsUnspecified"
      case DecisionCommands.DecisionCommandsOptOut:
        return "DecisionCommandsOptOut"
      default:
        return "UnrecognizedDecisionCommand"
    }
  }

  static commandTopic(command: DecisionCommands): string {
    switch (command) {
      case DecisionCommands.DecisionCommandsUnspecified:
      case DecisionCommands.DecisionCommandsOptOut:
        return CommandDataDecisionTopic;
      default:
		    return UnrecognizedDecisionTopic;
    }
  }

  static commandTopicWildcard(): string {
    return DecisionTypeNameEventPrefix + ">";
  }

  static getDecisionCommand(command: string): DecisionCommands {
    switch (command) {
      case "DecisionCommandsUnspecified":
        return DecisionCommands.DecisionCommandsUnspecified;
      case "DecisionCommandsOptOut":
        return DecisionCommands.DecisionCommandsOptOut;
      default:
        return DecisionCommands.UnrecognizedDecisionCommand;
    }
  }
}

// Event Methods
export class DecisionEventHelper {
  static eventName(event: DecisionEvents): string {
    switch (event) {
      case DecisionEvents.DecisionEventsUnspecified:
        return "DecisionEventsUnspecified";
      case DecisionEvents.DecisionEventsOptedOut:
        return "DecisionEventsOptedOut";
      default:
        return "UnrecognizedDecisionEvent";
    }
  }

  static eventTopic(event: DecisionEvents): string {
    switch (event) {
      case DecisionEvents.DecisionEventsUnspecified:
      case DecisionEvents.DecisionEventsOptedOut:
        return EventDataDecisionTopic;
      default:
        return UnrecognizedDecisionTopic;
    }
  }

  static eventTopicWildcard(): string {
    return DecisionTypeNameEventPrefix + ">";
  }

  static getDecisionEvent(event: string): DecisionEvents {
    switch (event) {
      case "DecisionEventsUnspecified":
        return DecisionEvents.DecisionEventsUnspecified;
      case "DecisionEventsOptedOut":
        return DecisionEvents.DecisionEventsOptedOut;
      default:
        return DecisionEvents.UnrecognizedDecisionEvent;
    }
  }
}

