// Code generated by protoc-gen-platform typescript/spec. DO NOT EDIT.
// source: platform/configuration/v2alpha/configuration.proto

export type ConfigurationCommand = number;
export type ConfigurationEvent = number;

// Constants for Configuration Type Names
export const ConfigurationTypeName = "configuration";
export const ConfigurationTypeNamePlural = "configurations";
export const ConfigurationTypeNameScreamingCamelCase = "CONFIGURATION";
export const ConfigurationTypeNamePluralScreamingCamelCase = "CONFIGURATIONS";
export const ConfigurationTypeNameEventPrefix = "configuration.";

// Enums for ConfigurationCommands
export enum ConfigurationCommands {
  
  ConfigurationCommandsUnspecified = 0,
  ConfigurationCommandsCreate = 1,
  ConfigurationCommandsUpdate = 2,
  ConfigurationCommandsDelete = 3,
  UnrecognizedConfigurationCommand = -1,
}

// Enums for Configuration Events
export enum ConfigurationEvents {

  ConfigurationEventsUnspecified = 0,
  ConfigurationEventsCreated = 1,
  ConfigurationEventsUpdated = 2,
  ConfigurationEventsDeleted = 3,
  UnrecognizedConfigurationEvent  = -1,
}

// Topics
export const CommandDataConfigurationTopic = "configuration.data.command";
export const EventDataConfigurationTopic = "configuration.data.event";
export const RoutineDataConfigurationTopic = "configuration.data.routine";
export const UnrecognizedConfigurationTopic = "unrecognized";

// Command Methods
export class ConfigurationCommandHelper {
  static commandName(command: ConfigurationCommands): string {
    switch (command) {
      case ConfigurationCommands.ConfigurationCommandsUnspecified:
        return "ConfigurationCommandsUnspecified"
      case ConfigurationCommands.ConfigurationCommandsCreate:
        return "ConfigurationCommandsCreate"
      case ConfigurationCommands.ConfigurationCommandsUpdate:
        return "ConfigurationCommandsUpdate"
      case ConfigurationCommands.ConfigurationCommandsDelete:
        return "ConfigurationCommandsDelete"
      default:
        return "UnrecognizedConfigurationCommand"
    }
  }

  static commandTopic(command: ConfigurationCommands): string {
    switch (command) {
      case ConfigurationCommands.ConfigurationCommandsUnspecified:
      case ConfigurationCommands.ConfigurationCommandsCreate:
      case ConfigurationCommands.ConfigurationCommandsUpdate:
      case ConfigurationCommands.ConfigurationCommandsDelete:
        return CommandDataConfigurationTopic;
      default:
		    return UnrecognizedConfigurationTopic;
    }
  }

  static commandTopicWildcard(): string {
    return ConfigurationTypeNameEventPrefix + ">";
  }

  static getConfigurationCommand(command: string): ConfigurationCommands {
    switch (command) {
      case "ConfigurationCommandsUnspecified":
        return ConfigurationCommands.ConfigurationCommandsUnspecified;
      case "ConfigurationCommandsCreate":
        return ConfigurationCommands.ConfigurationCommandsCreate;
      case "ConfigurationCommandsUpdate":
        return ConfigurationCommands.ConfigurationCommandsUpdate;
      case "ConfigurationCommandsDelete":
        return ConfigurationCommands.ConfigurationCommandsDelete;
      default:
        return ConfigurationCommands.UnrecognizedConfigurationCommand;
    }
  }
}

// Event Methods
export class ConfigurationEventHelper {
  static eventName(event: ConfigurationEvents): string {
    switch (event) {
      case ConfigurationEvents.ConfigurationEventsUnspecified:
        return "ConfigurationEventsUnspecified";
      case ConfigurationEvents.ConfigurationEventsCreated:
        return "ConfigurationEventsCreated";
      case ConfigurationEvents.ConfigurationEventsUpdated:
        return "ConfigurationEventsUpdated";
      case ConfigurationEvents.ConfigurationEventsDeleted:
        return "ConfigurationEventsDeleted";
      default:
        return "UnrecognizedConfigurationEvent";
    }
  }

  static eventTopic(event: ConfigurationEvents): string {
    switch (event) {
      case ConfigurationEvents.ConfigurationEventsUnspecified:
      case ConfigurationEvents.ConfigurationEventsCreated:
      case ConfigurationEvents.ConfigurationEventsUpdated:
      case ConfigurationEvents.ConfigurationEventsDeleted:
        return EventDataConfigurationTopic;
      default:
        return UnrecognizedConfigurationTopic;
    }
  }

  static eventTopicWildcard(): string {
    return ConfigurationTypeNameEventPrefix + ">";
  }

  static getConfigurationEvent(event: string): ConfigurationEvents {
    switch (event) {
      case "ConfigurationEventsUnspecified":
        return ConfigurationEvents.ConfigurationEventsUnspecified;
      case "ConfigurationEventsCreated":
        return ConfigurationEvents.ConfigurationEventsCreated;
      case "ConfigurationEventsUpdated":
        return ConfigurationEvents.ConfigurationEventsUpdated;
      case "ConfigurationEventsDeleted":
        return ConfigurationEvents.ConfigurationEventsDeleted;
      default:
        return ConfigurationEvents.UnrecognizedConfigurationEvent;
    }
  }
}

