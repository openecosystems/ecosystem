// Code generated by protoc-gen-platform typescript/spec. DO NOT EDIT.
// source: platform/cryptography/v2alpha/certificate_authority.proto

export type CertificateAuthorityCommand = number;
export type CertificateAuthorityEvent = number;

// Constants for CertificateAuthority Type Names
export const CertificateAuthorityTypeName = "certificateAuthority";
export const CertificateAuthorityTypeNamePlural = "certificateAuthorities";
export const CertificateAuthorityTypeNameScreamingCamelCase = "CERTIFICATE_AUTHORITY";
export const CertificateAuthorityTypeNamePluralScreamingCamelCase = "CERTIFICATE_AUTHORITIES";
export const CertificateAuthorityTypeNameEventPrefix = "certificateAuthority.";

// Enums for CertificateAuthorityCommands
export enum CertificateAuthorityCommands {
  
  CertificateAuthorityCommandsUnspecified = 0,
  CertificateAuthorityCommandsCreateCertificateAuthority = 1,
  UnrecognizedCertificateAuthorityCommand = -1,
}

// Enums for CertificateAuthority Events
export enum CertificateAuthorityEvents {

  CertificateAuthorityEventsUnspecified = 0,
  CertificateAuthorityEventsCreatedCertificateAuthority = 1,
  UnrecognizedCertificateAuthorityEvent  = -1,
}

// Topics
export const CommandDataCertificateAuthorityTopic = "certificateAuthority.data.command";
export const EventDataCertificateAuthorityTopic = "certificateAuthority.data.event";
export const RoutineDataCertificateAuthorityTopic = "certificateAuthority.data.routine";
export const UnrecognizedCertificateAuthorityTopic = "unrecognized";

// Command Methods
export class CertificateAuthorityCommandHelper {
  static commandName(command: CertificateAuthorityCommands): string {
    switch (command) {
      case CertificateAuthorityCommands.CertificateAuthorityCommandsUnspecified:
        return "CertificateAuthorityCommandsUnspecified"
      case CertificateAuthorityCommands.CertificateAuthorityCommandsCreateCertificateAuthority:
        return "CertificateAuthorityCommandsCreateCertificateAuthority"
      default:
        return "UnrecognizedCertificateAuthorityCommand"
    }
  }

  static commandTopic(command: CertificateAuthorityCommands): string {
    switch (command) {
      case CertificateAuthorityCommands.CertificateAuthorityCommandsUnspecified:
      case CertificateAuthorityCommands.CertificateAuthorityCommandsCreateCertificateAuthority:
        return CommandDataCertificateAuthorityTopic;
      default:
		    return UnrecognizedCertificateAuthorityTopic;
    }
  }

  static commandTopicWildcard(): string {
    return CertificateAuthorityTypeNameEventPrefix + ">";
  }

  static getCertificateAuthorityCommand(command: string): CertificateAuthorityCommands {
    switch (command) {
      case "CertificateAuthorityCommandsUnspecified":
        return CertificateAuthorityCommands.CertificateAuthorityCommandsUnspecified;
      case "CertificateAuthorityCommandsCreateCertificateAuthority":
        return CertificateAuthorityCommands.CertificateAuthorityCommandsCreateCertificateAuthority;
      default:
        return CertificateAuthorityCommands.UnrecognizedCertificateAuthorityCommand;
    }
  }
}

// Event Methods
export class CertificateAuthorityEventHelper {
  static eventName(event: CertificateAuthorityEvents): string {
    switch (event) {
      case CertificateAuthorityEvents.CertificateAuthorityEventsUnspecified:
        return "CertificateAuthorityEventsUnspecified";
      case CertificateAuthorityEvents.CertificateAuthorityEventsCreatedCertificateAuthority:
        return "CertificateAuthorityEventsCreatedCertificateAuthority";
      default:
        return "UnrecognizedCertificateAuthorityEvent";
    }
  }

  static eventTopic(event: CertificateAuthorityEvents): string {
    switch (event) {
      case CertificateAuthorityEvents.CertificateAuthorityEventsUnspecified:
      case CertificateAuthorityEvents.CertificateAuthorityEventsCreatedCertificateAuthority:
        return EventDataCertificateAuthorityTopic;
      default:
        return UnrecognizedCertificateAuthorityTopic;
    }
  }

  static eventTopicWildcard(): string {
    return CertificateAuthorityTypeNameEventPrefix + ">";
  }

  static getCertificateAuthorityEvent(event: string): CertificateAuthorityEvents {
    switch (event) {
      case "CertificateAuthorityEventsUnspecified":
        return CertificateAuthorityEvents.CertificateAuthorityEventsUnspecified;
      case "CertificateAuthorityEventsCreatedCertificateAuthority":
        return CertificateAuthorityEvents.CertificateAuthorityEventsCreatedCertificateAuthority;
      default:
        return CertificateAuthorityEvents.UnrecognizedCertificateAuthorityEvent;
    }
  }
}
