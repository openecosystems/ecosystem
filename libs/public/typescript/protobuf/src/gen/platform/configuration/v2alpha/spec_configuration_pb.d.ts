// Code generated by protoc-gen-platform protobuf/configuration. DO NOT EDIT.

// @generated by protoc-gen-es v2.2.3
// @generated from file platform/configuration/v2alpha/spec_configuration.proto (package platform.configuration.v2alpha, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";
import type { AuditConfiguration } from "../../audit/v2alpha/audit_pb";
import type { OecoConfiguration } from "../../cli/v2alpha/oeco_pb";
import type { CertificateConfiguration } from "../../cryptography/v2alpha/certificate_pb";
import type { CertificateAuthorityConfiguration } from "../../cryptography/v2alpha/certificate_authority_pb";
import type { EncryptionConfiguration } from "../../cryptography/v2alpha/encryption_pb";
import type { DynamicDnsConfiguration } from "../../dns/v2alpha/dynamic_dns_pb";
import type { EdgeRouterConfiguration } from "../../edge/v2alpha/edge_router_pb";
import type { EventMultiplexerConfiguration } from "../../event/v2alpha/event_multiplexer_pb";
import type { EventSubscriptionConfiguration } from "../../event/v2alpha/event_subscription_pb";
import type { IamApiKeyConfiguration } from "../../iam/v2alpha/iam_api_key_pb";
import type { IamAuthenticationConfiguration } from "../../iam/v2alpha/iam_authentication_pb";
import type { CryptographicMeshConfiguration } from "../../mesh/v2alpha/cryptographic_mesh_pb";
import type { ReferenceConfiguration } from "../../reference/v2alpha/reference_pb";
import type { SystemConfiguration } from "../../system/v2alpha/system_pb";
import type { PreferenceCenterConfiguration } from "../../communication/v1alpha/preference_center_pb";
import type { PreferenceCenterConfiguration as PreferenceCenterConfiguration$1 } from "../../communication/v1beta/preference_center_pb";

/**
 * Describes the file platform/configuration/v2alpha/spec_configuration.proto.
 */
export declare const file_platform_configuration_v2alpha_spec_configuration: GenFile;

/**
 * @generated from message platform.configuration.v2alpha.SpecPlatformConfiguration
 */
export declare type SpecPlatformConfiguration = Message<"platform.configuration.v2alpha.SpecPlatformConfiguration"> & {
  /**
   * @generated from field: platform.audit.v2alpha.AuditConfiguration audit_configuration_v2alpha = 2;
   */
  auditConfigurationV2alpha?: AuditConfiguration;

  /**
   * @generated from field: platform.cli.v2alpha.OecoConfiguration oeco_configuration_v2alpha = 3;
   */
  oecoConfigurationV2alpha?: OecoConfiguration;

  /**
   * @generated from field: platform.cryptography.v2alpha.CertificateConfiguration certificate_configuration_v2alpha = 4;
   */
  certificateConfigurationV2alpha?: CertificateConfiguration;

  /**
   * @generated from field: platform.cryptography.v2alpha.CertificateAuthorityConfiguration certificate_authority_configuration_v2alpha = 5;
   */
  certificateAuthorityConfigurationV2alpha?: CertificateAuthorityConfiguration;

  /**
   * @generated from field: platform.cryptography.v2alpha.EncryptionConfiguration encryption_configuration_v2alpha = 6;
   */
  encryptionConfigurationV2alpha?: EncryptionConfiguration;

  /**
   * @generated from field: platform.dns.v2alpha.DynamicDnsConfiguration dynamic_dns_configuration_v2alpha = 7;
   */
  dynamicDnsConfigurationV2alpha?: DynamicDnsConfiguration;

  /**
   * @generated from field: platform.edge.v2alpha.EdgeRouterConfiguration edge_router_configuration_v2alpha = 8;
   */
  edgeRouterConfigurationV2alpha?: EdgeRouterConfiguration;

  /**
   * @generated from field: platform.event.v2alpha.EventMultiplexerConfiguration event_multiplexer_configuration_v2alpha = 9;
   */
  eventMultiplexerConfigurationV2alpha?: EventMultiplexerConfiguration;

  /**
   * @generated from field: platform.event.v2alpha.EventSubscriptionConfiguration event_subscription_configuration_v2alpha = 10;
   */
  eventSubscriptionConfigurationV2alpha?: EventSubscriptionConfiguration;

  /**
   * @generated from field: platform.iam.v2alpha.IamApiKeyConfiguration iam_api_key_configuration_v2alpha = 11;
   */
  iamApiKeyConfigurationV2alpha?: IamApiKeyConfiguration;

  /**
   * @generated from field: platform.iam.v2alpha.IamAuthenticationConfiguration iam_authentication_configuration_v2alpha = 12;
   */
  iamAuthenticationConfigurationV2alpha?: IamAuthenticationConfiguration;

  /**
   * @generated from field: platform.mesh.v2alpha.CryptographicMeshConfiguration cryptographic_mesh_configuration_v2alpha = 13;
   */
  cryptographicMeshConfigurationV2alpha?: CryptographicMeshConfiguration;

  /**
   * @generated from field: platform.reference.v2alpha.ReferenceConfiguration reference_configuration_v2alpha = 14;
   */
  referenceConfigurationV2alpha?: ReferenceConfiguration;

  /**
   * @generated from field: platform.system.v2alpha.SystemConfiguration system_configuration_v2alpha = 15;
   */
  systemConfigurationV2alpha?: SystemConfiguration;

  /**
   * @generated from field: platform.communication.v1alpha.PreferenceCenterConfiguration preference_center_configuration_v1alpha = 16;
   */
  preferenceCenterConfigurationV1alpha?: PreferenceCenterConfiguration;

  /**
   * @generated from field: platform.communication.v1beta.PreferenceCenterConfiguration preference_center_configuration_v1beta = 17;
   */
  preferenceCenterConfigurationV1beta?: PreferenceCenterConfiguration$1;
};

/**
 * Describes the message platform.configuration.v2alpha.SpecPlatformConfiguration.
 * Use `create(SpecPlatformConfigurationSchema)` to create a new message.
 */
export declare const SpecPlatformConfigurationSchema: GenMessage<SpecPlatformConfiguration>;
