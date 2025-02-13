// Code generated by protoc-gen-platform typescript/spec. DO NOT EDIT.
// source: platform/ecosystem/v2alpha/ecosystem.proto

export type EcosystemCommand = number;
export type EcosystemEvent = number;

// Constants for Ecosystem Type Names
export const EcosystemTypeName = "ecosystem";
export const EcosystemTypeNamePlural = "ecosystems";
export const EcosystemTypeNameScreamingCamelCase = "ECOSYSTEM";
export const EcosystemTypeNamePluralScreamingCamelCase = "ECOSYSTEMS";
export const EcosystemTypeNameEventPrefix = "ecosystem.";

// Enums for EcosystemCommands
export enum EcosystemCommands {
  
  EcosystemCommandsUnspecified = 0,
  EcosystemCommandsCreate = 1,
  EcosystemCommandsUpdate = 2,
  EcosystemCommandsDelete = 3,
  UnrecognizedEcosystemCommand = -1,
}

// Enums for Ecosystem Events
export enum EcosystemEvents {

  EcosystemEventsUnspecified = 0,
  EcosystemEventsCreated = 1,
  EcosystemEventsUpdated = 2,
  EcosystemEventsDeleted = 3,
  EcosystemEventsErrored = 4,
  UnrecognizedEcosystemEvent  = -1,
}

// Topics
export const CommandDataEcosystemTopic = "ecosystem.data.command";
export const EventDataEcosystemTopic = "ecosystem.data.event";
export const RoutineDataEcosystemTopic = "ecosystem.data.routine";
export const UnrecognizedEcosystemTopic = "unrecognized";

// Command Methods
export class EcosystemCommandHelper {
  static commandName(command: EcosystemCommands): string {
    switch (command) {
      case EcosystemCommands.EcosystemCommandsUnspecified:
        return "EcosystemCommandsUnspecified"
      case EcosystemCommands.EcosystemCommandsCreate:
        return "EcosystemCommandsCreate"
      case EcosystemCommands.EcosystemCommandsUpdate:
        return "EcosystemCommandsUpdate"
      case EcosystemCommands.EcosystemCommandsDelete:
        return "EcosystemCommandsDelete"
      default:
        return "UnrecognizedEcosystemCommand"
    }
  }

  static commandTopic(command: EcosystemCommands): string {
    switch (command) {
      case EcosystemCommands.EcosystemCommandsUnspecified:
      case EcosystemCommands.EcosystemCommandsCreate:
      case EcosystemCommands.EcosystemCommandsUpdate:
      case EcosystemCommands.EcosystemCommandsDelete:
        return CommandDataEcosystemTopic;
      default:
		    return UnrecognizedEcosystemTopic;
    }
  }

  static commandTopicWildcard(): string {
    return EcosystemTypeNameEventPrefix + ">";
  }

  static getEcosystemCommand(command: string): EcosystemCommands {
    switch (command) {
      case "EcosystemCommandsUnspecified":
        return EcosystemCommands.EcosystemCommandsUnspecified;
      case "EcosystemCommandsCreate":
        return EcosystemCommands.EcosystemCommandsCreate;
      case "EcosystemCommandsUpdate":
        return EcosystemCommands.EcosystemCommandsUpdate;
      case "EcosystemCommandsDelete":
        return EcosystemCommands.EcosystemCommandsDelete;
      default:
        return EcosystemCommands.UnrecognizedEcosystemCommand;
    }
  }
}

// Event Methods
export class EcosystemEventHelper {
  static eventName(event: EcosystemEvents): string {
    switch (event) {
      case EcosystemEvents.EcosystemEventsUnspecified:
        return "EcosystemEventsUnspecified";
      case EcosystemEvents.EcosystemEventsCreated:
        return "EcosystemEventsCreated";
      case EcosystemEvents.EcosystemEventsUpdated:
        return "EcosystemEventsUpdated";
      case EcosystemEvents.EcosystemEventsDeleted:
        return "EcosystemEventsDeleted";
      case EcosystemEvents.EcosystemEventsErrored:
        return "EcosystemEventsErrored";
      default:
        return "UnrecognizedEcosystemEvent";
    }
  }

  static eventTopic(event: EcosystemEvents): string {
    switch (event) {
      case EcosystemEvents.EcosystemEventsUnspecified:
      case EcosystemEvents.EcosystemEventsCreated:
      case EcosystemEvents.EcosystemEventsUpdated:
      case EcosystemEvents.EcosystemEventsDeleted:
      case EcosystemEvents.EcosystemEventsErrored:
        return EventDataEcosystemTopic;
      default:
        return UnrecognizedEcosystemTopic;
    }
  }

  static eventTopicWildcard(): string {
    return EcosystemTypeNameEventPrefix + ">";
  }

  static getEcosystemEvent(event: string): EcosystemEvents {
    switch (event) {
      case "EcosystemEventsUnspecified":
        return EcosystemEvents.EcosystemEventsUnspecified;
      case "EcosystemEventsCreated":
        return EcosystemEvents.EcosystemEventsCreated;
      case "EcosystemEventsUpdated":
        return EcosystemEvents.EcosystemEventsUpdated;
      case "EcosystemEventsDeleted":
        return EcosystemEvents.EcosystemEventsDeleted;
      case "EcosystemEventsErrored":
        return EcosystemEvents.EcosystemEventsErrored;
      default:
        return EcosystemEvents.UnrecognizedEcosystemEvent;
    }
  }
}

