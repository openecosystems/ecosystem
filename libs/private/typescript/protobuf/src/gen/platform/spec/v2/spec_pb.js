// @generated by protoc-gen-es v2.2.3
// @generated from file platform/spec/v2/spec.proto (package platform.spec.v2, syntax proto3)
/* eslint-disable */

import { enumDesc, fileDesc, messageDesc, tsEnum } from "@bufbuild/protobuf/codegenv1";
import { file_google_protobuf_any, file_google_protobuf_field_mask, file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";
import { file_platform_options_v2_annotations } from "../../options/v2/annotations_pb";
import { file_platform_type_v2_jurisdiction } from "../../type/v2/jurisdiction_pb";
import { file_platform_type_v2_validation } from "../../type/v2/validation_pb";
import { file_platform_type_v2_mask } from "../../type/v2/mask_pb";

/**
 * Describes the file platform/spec/v2/spec.proto.
 */
export const file_platform_spec_v2_spec = /*@__PURE__*/
  fileDesc("ChtwbGF0Zm9ybS9zcGVjL3YyL3NwZWMucHJvdG8SEHBsYXRmb3JtLnNwZWMudjIimgEKB1NwZWNLZXkSGQoRb3JnYW5pemF0aW9uX3NsdWcYAiABKAkSFgoOd29ya3NwYWNlX3NsdWcYAyABKAkSNQoNd29ya3NwYWNlX2phbhgEIAEoDjIeLnBsYXRmb3JtLnR5cGUudjIuSnVyaXNkaWN0aW9uEhEKCXNwZWNfdHlwZRgFIAEoCRIKCgJpZBgGIAEoCToG+rYYAggBItQECgRTcGVjEhQKDHNwZWNfdmVyc2lvbhgBIAEoCRISCgptZXNzYWdlX2lkGAIgASgJEisKB3NlbnRfYXQYAyABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wEi8KC3JlY2VpdmVkX2F0GAQgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcBIwCgxjb21wbGV0ZWRfYXQYBSABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wEhEKCXNwZWNfdHlwZRgGIAEoCRI4Cg9zcGVjX2V2ZW50X3R5cGUYByABKA4yHy5wbGF0Zm9ybS5zcGVjLnYyLlNwZWNFdmVudFR5cGUSEgoKc3BlY19ldmVudBgIIAEoCRIyCglwcmluY2lwYWwYCSABKAsyHy5wbGF0Zm9ybS5zcGVjLnYyLlNwZWNQcmluY2lwYWwSMwoMc3Bhbl9jb250ZXh0GAogASgLMh0ucGxhdGZvcm0uc3BlYy52Mi5TcGFuQ29udGV4dBIuCgdjb250ZXh0GAsgASgLMh0ucGxhdGZvcm0uc3BlYy52Mi5TcGVjQ29udGV4dBI9Cg9yb3V0aW5lX2NvbnRleHQYDCABKAsyJC5wbGF0Zm9ybS5zcGVjLnYyLlNwZWNSb3V0aW5lQ29udGV4dBIiCgRkYXRhGA0gASgLMhQuZ29vZ2xlLnByb3RvYnVmLkFueRItCglzcGVjX2RhdGEYDiABKAsyGi5wbGF0Zm9ybS5zcGVjLnYyLlNwZWNEYXRhOgb6thgCCAEi0wIKClNwZWNQdWJsaWMSFAoMc3BlY192ZXJzaW9uGAEgASgJEhIKCm1lc3NhZ2VfaWQYAiABKAkSKwoHc2VudF9hdBgDIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXASLwoLcmVjZWl2ZWRfYXQYBCABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wEjAKDGNvbXBsZXRlZF9hdBgFIAEoCzIaLmdvb2dsZS5wcm90b2J1Zi5UaW1lc3RhbXASEQoJc3BlY190eXBlGAYgASgJEjgKD3NwZWNfZXZlbnRfdHlwZRgHIAEoDjIfLnBsYXRmb3JtLnNwZWMudjIuU3BlY0V2ZW50VHlwZRISCgpzcGVjX2V2ZW50GAggASgJEiIKBGRhdGEYCSABKAsyFC5nb29nbGUucHJvdG9idWYuQW55Ogb6thgCCAEi3QQKC1NwZWNDb250ZXh0EhQKDGVjb3N5c3RlbV9pZBgBIAEoCRIWCg5lY29zeXN0ZW1fc2x1ZxgCIAEoCRI1Cg1lY29zeXN0ZW1famFuGAMgASgOMh4ucGxhdGZvcm0udHlwZS52Mi5KdXJpc2RpY3Rpb24SFwoPb3JnYW5pemF0aW9uX2lkGDEgASgJEhkKEW9yZ2FuaXphdGlvbl9zbHVnGDIgASgJEhYKDndvcmtzcGFjZV9zbHVnGDMgASgJEjUKDXdvcmtzcGFjZV9qYW4YNCABKA4yHi5wbGF0Zm9ybS50eXBlLnYyLkp1cmlzZGljdGlvbhIKCgJpcBgFIAEoCRIOCgZsb2NhbGUYBiABKAkSEAoIdGltZXpvbmUYByABKAkSEgoKdXNlcl9hZ2VudBgIIAEoCRI0Cgp2YWxpZGF0aW9uGAkgASgLMiAucGxhdGZvcm0uc3BlYy52Mi5TcGVjVmFsaWRhdGlvbhIwCghwcm9kdWNlchgKIAEoCzIeLnBsYXRmb3JtLnNwZWMudjIuU3BlY1Byb2R1Y2VyEiwKBmRldmljZRgLIAEoCzIcLnBsYXRmb3JtLnNwZWMudjIuU3BlY0RldmljZRIwCghsb2NhdGlvbhgMIAEoCzIeLnBsYXRmb3JtLnNwZWMudjIuU3BlY0xvY2F0aW9uEi4KB25ldHdvcmsYDSABKAsyHS5wbGF0Zm9ybS5zcGVjLnYyLlNwZWNOZXR3b3JrEiQKAm9zGA4gASgLMhgucGxhdGZvcm0uc3BlYy52Mi5TcGVjT1M6Bvq2GAIIASJlCgtTcGFuQ29udGV4dBIQCgh0cmFjZV9pZBgBIAEoCRIPCgdzcGFuX2lkGAIgASgJEhYKDnBhcmVudF9zcGFuX2lkGAMgASgJEhMKC3RyYWNlX2ZsYWdzGAQgASgJOgb6thgCCAEixwEKElNwZWNSb3V0aW5lQ29udGV4dBISCgpyb3V0aW5lX2lkGAEgASgJEksKDHJvdXRpbmVfZGF0YRgDIAMoCzI1LnBsYXRmb3JtLnNwZWMudjIuU3BlY1JvdXRpbmVDb250ZXh0LlJvdXRpbmVEYXRhRW50cnkaSAoQUm91dGluZURhdGFFbnRyeRILCgNrZXkYASABKAkSIwoFdmFsdWUYAiABKAsyFC5nb29nbGUucHJvdG9idWYuQW55OgI4AToG+rYYAggBItkBCg1TcGVjUHJpbmNpcGFsEjEKBHR5cGUYASABKA4yIy5wbGF0Zm9ybS5zcGVjLnYyLlNwZWNQcmluY2lwYWxUeXBlEhQKDGFub255bW91c19pZBgCIAEoCRIUCgxwcmluY2lwYWxfaWQYAyABKAkSFwoPcHJpbmNpcGFsX2VtYWlsGAQgASgJEhUKDWNvbm5lY3Rpb25faWQYBSABKAkSMQoKYXV0aF9yb2xlcxgGIAMoDjIdLnBsYXRmb3JtLm9wdGlvbnMudjIuQXV0aFJvbGU6Bvq2GAIIASIvCg5TcGVjVmFsaWRhdGlvbhIVCg12YWxpZGF0ZV9vbmx5GAEgASgIOgb6thgCCAEiVwoMU3BlY1Byb2R1Y2VyEgwKBG5hbWUYASABKAkSDwoHdmVyc2lvbhgCIAEoCRINCgVidWlsZBgDIAEoCRIRCgluYW1lc3BhY2UYBCABKAk6Bvq2GAIIASKIAQoKU3BlY0RldmljZRIKCgJpZBgBIAEoCRIMCgR0eXBlGAIgASgJEhYKDmFkdmVydGlzaW5nX2lkGAMgASgJEhQKDG1hbnVmYWN0dXJlchgEIAEoCRINCgVtb2RlbBgFIAEoCRIMCgRuYW1lGAYgASgJEg0KBXRva2VuGAcgASgJOgb6thgCCAEiaQoMU3BlY0xvY2F0aW9uEgwKBGNpdHkYASABKAkSDwoHY291bnRyeRgCIAEoCRIQCghsYXRpdHVkZRgDIAEoARIRCglsb25naXR1ZGUYBCABKAESDQoFc3BlZWQYBSABKAk6Bvq2GAIIASJZCgtTcGVjTmV0d29yaxIRCglibHVldG9vdGgYASABKAgSEAoIY2VsbHVsYXIYAiABKAgSDAoEd2lmaRgDIAEoCBIPCgdjYXJyaWVyGAQgASgJOgb6thgCCAEiLwoGU3BlY09TEgwKBG5hbWUYASABKAkSDwoHdmVyc2lvbhgCIAEoCToG+rYYAggBIpMBCghTcGVjRGF0YRIrCg1jb25maWd1cmF0aW9uGAEgASgLMhQuZ29vZ2xlLnByb3RvYnVmLkFueRIiCgRkYXRhGAIgASgLMhQuZ29vZ2xlLnByb3RvYnVmLkFueRIuCgpmaWVsZF9tYXNrGAMgASgLMhouZ29vZ2xlLnByb3RvYnVmLkZpZWxkTWFzazoG+rYYAggBIqQBChJTcGVjUmVxdWVzdENvbnRleHQSPwoScmVxdWVzdF92YWxpZGF0aW9uGAEgASgLMiMucGxhdGZvcm0udHlwZS52Mi5SZXF1ZXN0VmFsaWRhdGlvbhIZChFvcmdhbml6YXRpb25fc2x1ZxgCIAEoCRIWCg53b3Jrc3BhY2Vfc2x1ZxgDIAEoCRISCgpyb3V0aW5lX2lkGAQgASgJOgb6thgCCAEirQIKE1NwZWNSZXNwb25zZUNvbnRleHQSQQoTcmVzcG9uc2VfdmFsaWRhdGlvbhgBIAEoCzIkLnBsYXRmb3JtLnR5cGUudjIuUmVzcG9uc2VWYWxpZGF0aW9uEjUKDXJlc3BvbnNlX21hc2sYAiABKAsyHi5wbGF0Zm9ybS50eXBlLnYyLlJlc3BvbnNlTWFzaxIWCg5lY29zeXN0ZW1fc2x1ZxgDIAEoCRIZChFvcmdhbml6YXRpb25fc2x1ZxgyIAEoCRIWCg53b3Jrc3BhY2Vfc2x1ZxgzIAEoCRI1Cg13b3Jrc3BhY2VfamFuGDQgASgOMh4ucGxhdGZvcm0udHlwZS52Mi5KdXJpc2RpY3Rpb24SEgoKcm91dGluZV9pZBg1IAEoCToG+rYYAggCKqACCg1TcGVjRXZlbnRUeXBlEh8KG1NQRUNfRVZFTlRfVFlQRV9VTlNQRUNJRklFRBAAEhsKF1NQRUNfRVZFTlRfVFlQRV9DT01NQU5EEAESGQoVU1BFQ19FVkVOVF9UWVBFX0VWRU5UEAISGwoXU1BFQ19FVkVOVF9UWVBFX1JPVVRJTkUQAxIWChJTUEVDX0VWRU5UX1RZUEVfTUwQBBIaChZTUEVDX0VWRU5UX1RZUEVfU1RSRUFNEAUSFgoSU1BFQ19FVkVOVF9UWVBFX0RCEAYSGwoXU1BFQ19FVkVOVF9UWVBFX1BST0ZJTEUQBxIXChNTUEVDX0VWRU5UX1RZUEVfRVRMEAgSFwoTU1BFQ19FVkVOVF9UWVBFX0xPRxAJKr4BChFTcGVjUHJpbmNpcGFsVHlwZRIjCh9TUEVDX1BSSU5DSVBBTF9UWVBFX1VOU1BFQ0lGSUVEEAASHAoYU1BFQ19QUklOQ0lQQUxfVFlQRV9VU0VSEAESJwojU1BFQ19QUklOQ0lQQUxfVFlQRV9TRVJWSUNFX0FDQ09VTlQQAhIdChlTUEVDX1BSSU5DSVBBTF9UWVBFX0dST1VQEAMSHgoaU1BFQ19QUklOQ0lQQUxfVFlQRV9ET01BSU4QBEJjWltnaXRodWIuY29tL29wZW5lY29zeXN0ZW1zL2Vjb3N5c3RlbS9saWJzL3Byb3RvYnVmL2dvL3Byb3RvYnVmL2dlbi9wbGF0Zm9ybS9zcGVjL3YyO3NwZWN2MnBimrUYAggBYgZwcm90bzM", [file_google_protobuf_any, file_google_protobuf_timestamp, file_platform_options_v2_annotations, file_platform_type_v2_jurisdiction, file_platform_type_v2_validation, file_platform_type_v2_mask, file_google_protobuf_field_mask]);

/**
 * Describes the message platform.spec.v2.SpecKey.
 * Use `create(SpecKeySchema)` to create a new message.
 */
export const SpecKeySchema = /*@__PURE__*/
  messageDesc(file_platform_spec_v2_spec, 0);

/**
 * Describes the message platform.spec.v2.Spec.
 * Use `create(SpecSchema)` to create a new message.
 */
export const SpecSchema = /*@__PURE__*/
  messageDesc(file_platform_spec_v2_spec, 1);

/**
 * Describes the message platform.spec.v2.SpecPublic.
 * Use `create(SpecPublicSchema)` to create a new message.
 */
export const SpecPublicSchema = /*@__PURE__*/
  messageDesc(file_platform_spec_v2_spec, 2);

/**
 * Describes the message platform.spec.v2.SpecContext.
 * Use `create(SpecContextSchema)` to create a new message.
 */
export const SpecContextSchema = /*@__PURE__*/
  messageDesc(file_platform_spec_v2_spec, 3);

/**
 * Describes the message platform.spec.v2.SpanContext.
 * Use `create(SpanContextSchema)` to create a new message.
 */
export const SpanContextSchema = /*@__PURE__*/
  messageDesc(file_platform_spec_v2_spec, 4);

/**
 * Describes the message platform.spec.v2.SpecRoutineContext.
 * Use `create(SpecRoutineContextSchema)` to create a new message.
 */
export const SpecRoutineContextSchema = /*@__PURE__*/
  messageDesc(file_platform_spec_v2_spec, 5);

/**
 * Describes the message platform.spec.v2.SpecPrincipal.
 * Use `create(SpecPrincipalSchema)` to create a new message.
 */
export const SpecPrincipalSchema = /*@__PURE__*/
  messageDesc(file_platform_spec_v2_spec, 6);

/**
 * Describes the message platform.spec.v2.SpecValidation.
 * Use `create(SpecValidationSchema)` to create a new message.
 */
export const SpecValidationSchema = /*@__PURE__*/
  messageDesc(file_platform_spec_v2_spec, 7);

/**
 * Describes the message platform.spec.v2.SpecProducer.
 * Use `create(SpecProducerSchema)` to create a new message.
 */
export const SpecProducerSchema = /*@__PURE__*/
  messageDesc(file_platform_spec_v2_spec, 8);

/**
 * Describes the message platform.spec.v2.SpecDevice.
 * Use `create(SpecDeviceSchema)` to create a new message.
 */
export const SpecDeviceSchema = /*@__PURE__*/
  messageDesc(file_platform_spec_v2_spec, 9);

/**
 * Describes the message platform.spec.v2.SpecLocation.
 * Use `create(SpecLocationSchema)` to create a new message.
 */
export const SpecLocationSchema = /*@__PURE__*/
  messageDesc(file_platform_spec_v2_spec, 10);

/**
 * Describes the message platform.spec.v2.SpecNetwork.
 * Use `create(SpecNetworkSchema)` to create a new message.
 */
export const SpecNetworkSchema = /*@__PURE__*/
  messageDesc(file_platform_spec_v2_spec, 11);

/**
 * Describes the message platform.spec.v2.SpecOS.
 * Use `create(SpecOSSchema)` to create a new message.
 */
export const SpecOSSchema = /*@__PURE__*/
  messageDesc(file_platform_spec_v2_spec, 12);

/**
 * Describes the message platform.spec.v2.SpecData.
 * Use `create(SpecDataSchema)` to create a new message.
 */
export const SpecDataSchema = /*@__PURE__*/
  messageDesc(file_platform_spec_v2_spec, 13);

/**
 * Describes the message platform.spec.v2.SpecRequestContext.
 * Use `create(SpecRequestContextSchema)` to create a new message.
 */
export const SpecRequestContextSchema = /*@__PURE__*/
  messageDesc(file_platform_spec_v2_spec, 14);

/**
 * Describes the message platform.spec.v2.SpecResponseContext.
 * Use `create(SpecResponseContextSchema)` to create a new message.
 */
export const SpecResponseContextSchema = /*@__PURE__*/
  messageDesc(file_platform_spec_v2_spec, 15);

/**
 * Describes the enum platform.spec.v2.SpecEventType.
 */
export const SpecEventTypeSchema = /*@__PURE__*/
  enumDesc(file_platform_spec_v2_spec, 0);

/**
 * @generated from enum platform.spec.v2.SpecEventType
 */
export const SpecEventType = /*@__PURE__*/
  tsEnum(SpecEventTypeSchema);

/**
 * Describes the enum platform.spec.v2.SpecPrincipalType.
 */
export const SpecPrincipalTypeSchema = /*@__PURE__*/
  enumDesc(file_platform_spec_v2_spec, 1);

/**
 * Spec principal types
 *
 * @generated from enum platform.spec.v2.SpecPrincipalType
 */
export const SpecPrincipalType = /*@__PURE__*/
  tsEnum(SpecPrincipalTypeSchema);

