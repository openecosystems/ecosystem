// Code generated by protoc-gen-platform typescript/spec. DO NOT EDIT.
// source: platform/audit/v2alpha/audit.proto

export type AuditCommand = number;
export type AuditEvent = number;

// Constants for Audit Type Names
export const AuditTypeName = "audit";
export const AuditTypeNamePlural = "audits";
export const AuditTypeNameScreamingCamelCase = "AUDIT";
export const AuditTypeNamePluralScreamingCamelCase = "AUDITS";
export const AuditTypeNameEventPrefix = "audit.";

// Enums for AuditCommands
export enum AuditCommands {
  
  AuditCommandsUnspecified = 0,
  UnrecognizedAuditCommand = -1,
}

// Enums for Audit Events
export enum AuditEvents {

  AuditEventsUnspecified = 0,
  AuditEventsCreated = 1,
  UnrecognizedAuditEvent  = -1,
}

// Topics
export const CommandDataAuditTopic = "audit.data.command";
export const EventDataAuditTopic = "audit.data.event";
export const RoutineDataAuditTopic = "audit.data.routine";
export const UnrecognizedAuditTopic = "unrecognized";

// Command Methods
export class AuditCommandHelper {
  static commandName(command: AuditCommands): string {
    switch (command) {
      case AuditCommands.AuditCommandsUnspecified:
        return "AuditCommandsUnspecified"
      default:
        return "UnrecognizedAuditCommand"
    }
  }

  static commandTopic(command: AuditCommands): string {
    switch (command) {
      case AuditCommands.AuditCommandsUnspecified:
        return CommandDataAuditTopic;
      default:
		    return UnrecognizedAuditTopic;
    }
  }

  static commandTopicWildcard(): string {
    return AuditTypeNameEventPrefix + ">";
  }

  static getAuditCommand(command: string): AuditCommands {
    switch (command) {
      case "AuditCommandsUnspecified":
        return AuditCommands.AuditCommandsUnspecified;
      default:
        return AuditCommands.UnrecognizedAuditCommand;
    }
  }
}

// Event Methods
export class AuditEventHelper {
  static eventName(event: AuditEvents): string {
    switch (event) {
      case AuditEvents.AuditEventsUnspecified:
        return "AuditEventsUnspecified";
      case AuditEvents.AuditEventsCreated:
        return "AuditEventsCreated";
      default:
        return "UnrecognizedAuditEvent";
    }
  }

  static eventTopic(event: AuditEvents): string {
    switch (event) {
      case AuditEvents.AuditEventsUnspecified:
      case AuditEvents.AuditEventsCreated:
        return EventDataAuditTopic;
      default:
        return UnrecognizedAuditTopic;
    }
  }

  static eventTopicWildcard(): string {
    return AuditTypeNameEventPrefix + ">";
  }

  static getAuditEvent(event: string): AuditEvents {
    switch (event) {
      case "AuditEventsUnspecified":
        return AuditEvents.AuditEventsUnspecified;
      case "AuditEventsCreated":
        return AuditEvents.AuditEventsCreated;
      default:
        return AuditEvents.UnrecognizedAuditEvent;
    }
  }
}
