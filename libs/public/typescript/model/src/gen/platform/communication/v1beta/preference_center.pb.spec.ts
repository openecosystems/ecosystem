// Code generated by protoc-gen-platform typescript/spec. DO NOT EDIT.
// source: platform/communication/v1beta/preference_center.proto

export type PreferenceCenterCommand = number;
export type PreferenceCenterEvent = number;

// Constants for PreferenceCenter Type Names
export const PreferenceCenterTypeName = "preferenceCenter";
export const PreferenceCenterTypeNamePlural = "preferenceCenters";
export const PreferenceCenterTypeNameScreamingCamelCase = "PREFERENCE_CENTER";
export const PreferenceCenterTypeNamePluralScreamingCamelCase = "PREFERENCE_CENTERS";
export const PreferenceCenterTypeNameEventPrefix = "preferenceCenter.";

// Enums for PreferenceCenterCommands
export enum PreferenceCenterCommands {
  
  PreferenceCenterCommandsUnspecified = 0,
  PreferenceCenterCommandsSubscribe = 1,
  PreferenceCenterCommandsUpdateSubscription = 2,
  PreferenceCenterCommandsUnsubscribe = 3,
  UnrecognizedPreferenceCenterCommand = -1,
}

// Enums for PreferenceCenter Events
export enum PreferenceCenterEvents {

  PreferenceCenterEventsUnspecified = 0,
  PreferenceCenterEventsSubscribed = 1,
  PreferenceCenterEventsUpdated = 2,
  PreferenceCenterEventsUnsubscribed = 3,
  UnrecognizedPreferenceCenterEvent  = -1,
}

// Topics
export const CommandDataPreferenceCenterTopic = "preferenceCenter.data.command";
export const EventDataPreferenceCenterTopic = "preferenceCenter.data.event";
export const RoutineDataPreferenceCenterTopic = "preferenceCenter.data.routine";
export const UnrecognizedPreferenceCenterTopic = "unrecognized";

// Command Methods
export class PreferenceCenterCommandHelper {
  static commandName(command: PreferenceCenterCommands): string {
    switch (command) {
      case PreferenceCenterCommands.PreferenceCenterCommandsUnspecified:
        return "PreferenceCenterCommandsUnspecified"
      case PreferenceCenterCommands.PreferenceCenterCommandsSubscribe:
        return "PreferenceCenterCommandsSubscribe"
      case PreferenceCenterCommands.PreferenceCenterCommandsUpdateSubscription:
        return "PreferenceCenterCommandsUpdateSubscription"
      case PreferenceCenterCommands.PreferenceCenterCommandsUnsubscribe:
        return "PreferenceCenterCommandsUnsubscribe"
      default:
        return "UnrecognizedPreferenceCenterCommand"
    }
  }

  static commandTopic(command: PreferenceCenterCommands): string {
    switch (command) {
      case PreferenceCenterCommands.PreferenceCenterCommandsUnspecified:
      case PreferenceCenterCommands.PreferenceCenterCommandsSubscribe:
      case PreferenceCenterCommands.PreferenceCenterCommandsUpdateSubscription:
      case PreferenceCenterCommands.PreferenceCenterCommandsUnsubscribe:
        return CommandDataPreferenceCenterTopic;
      default:
		    return UnrecognizedPreferenceCenterTopic;
    }
  }

  static commandTopicWildcard(): string {
    return PreferenceCenterTypeNameEventPrefix + ">";
  }

  static getPreferenceCenterCommand(command: string): PreferenceCenterCommands {
    switch (command) {
      case "PreferenceCenterCommandsUnspecified":
        return PreferenceCenterCommands.PreferenceCenterCommandsUnspecified;
      case "PreferenceCenterCommandsSubscribe":
        return PreferenceCenterCommands.PreferenceCenterCommandsSubscribe;
      case "PreferenceCenterCommandsUpdateSubscription":
        return PreferenceCenterCommands.PreferenceCenterCommandsUpdateSubscription;
      case "PreferenceCenterCommandsUnsubscribe":
        return PreferenceCenterCommands.PreferenceCenterCommandsUnsubscribe;
      default:
        return PreferenceCenterCommands.UnrecognizedPreferenceCenterCommand;
    }
  }
}

// Event Methods
export class PreferenceCenterEventHelper {
  static eventName(event: PreferenceCenterEvents): string {
    switch (event) {
      case PreferenceCenterEvents.PreferenceCenterEventsUnspecified:
        return "PreferenceCenterEventsUnspecified";
      case PreferenceCenterEvents.PreferenceCenterEventsSubscribed:
        return "PreferenceCenterEventsSubscribed";
      case PreferenceCenterEvents.PreferenceCenterEventsUpdated:
        return "PreferenceCenterEventsUpdated";
      case PreferenceCenterEvents.PreferenceCenterEventsUnsubscribed:
        return "PreferenceCenterEventsUnsubscribed";
      default:
        return "UnrecognizedPreferenceCenterEvent";
    }
  }

  static eventTopic(event: PreferenceCenterEvents): string {
    switch (event) {
      case PreferenceCenterEvents.PreferenceCenterEventsUnspecified:
      case PreferenceCenterEvents.PreferenceCenterEventsSubscribed:
      case PreferenceCenterEvents.PreferenceCenterEventsUpdated:
      case PreferenceCenterEvents.PreferenceCenterEventsUnsubscribed:
        return EventDataPreferenceCenterTopic;
      default:
        return UnrecognizedPreferenceCenterTopic;
    }
  }

  static eventTopicWildcard(): string {
    return PreferenceCenterTypeNameEventPrefix + ">";
  }

  static getPreferenceCenterEvent(event: string): PreferenceCenterEvents {
    switch (event) {
      case "PreferenceCenterEventsUnspecified":
        return PreferenceCenterEvents.PreferenceCenterEventsUnspecified;
      case "PreferenceCenterEventsSubscribed":
        return PreferenceCenterEvents.PreferenceCenterEventsSubscribed;
      case "PreferenceCenterEventsUpdated":
        return PreferenceCenterEvents.PreferenceCenterEventsUpdated;
      case "PreferenceCenterEventsUnsubscribed":
        return PreferenceCenterEvents.PreferenceCenterEventsUnsubscribed;
      default:
        return PreferenceCenterEvents.UnrecognizedPreferenceCenterEvent;
    }
  }
}

