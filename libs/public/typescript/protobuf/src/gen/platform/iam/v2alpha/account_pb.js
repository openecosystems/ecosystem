// @generated by protoc-gen-es v2.2.3
// @generated from file platform/iam/v2alpha/account.proto (package platform.iam.v2alpha, syntax proto3)
/* eslint-disable */

import { enumDesc, fileDesc, messageDesc, serviceDesc, tsEnum } from "@bufbuild/protobuf/codegenv1";
import { file_google_api_annotations } from "../../../google/api/annotations_pb";
import { file_platform_options_v2_annotations } from "../../options/v2/annotations_pb";
import { file_platform_spec_v2_spec } from "../../spec/v2/spec_pb";
import { file_google_protobuf_duration, file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";
import { file_platform_type_v2_file } from "../../type/v2/file_pb";
import { file_platform_type_v2_credential } from "../../type/v2/credential_pb";
import { file_platform_iam_v2alpha_account_authority } from "./account_authority_pb";
import { file_platform_type_v2_peer } from "../../type/v2/peer_pb";

/**
 * Describes the file platform/iam/v2alpha/account.proto.
 */
export const file_platform_iam_v2alpha_account = /*@__PURE__*/
  fileDesc("CiJwbGF0Zm9ybS9pYW0vdjJhbHBoYS9hY2NvdW50LnByb3RvEhRwbGF0Zm9ybS5pYW0udjJhbHBoYSIWChRBY2NvdW50Q29uZmlndXJhdGlvbiJ6ChRDcmVhdGVBY2NvdW50UmVxdWVzdBIMCgRuYW1lGAEgASgJEiYKBWN1cnZlGAIgASgOMhcucGxhdGZvcm0udHlwZS52Mi5DdXJ2ZRIkCgRjZXJ0GAMgASgLMhYucGxhdGZvcm0udHlwZS52Mi5GaWxlOgb6thgCCAEijAEKFUNyZWF0ZUFjY291bnRSZXNwb25zZRI7CgxzcGVjX2NvbnRleHQYASABKAsyJS5wbGF0Zm9ybS5zcGVjLnYyLlNwZWNSZXNwb25zZUNvbnRleHQSLgoHYWNjb3VudBgCIAEoCzIdLnBsYXRmb3JtLmlhbS52MmFscGhhLkFjY291bnQ6Bvq2GAIIAiJLChRWZXJpZnlBY2NvdW50UmVxdWVzdBIrCgtwdWJsaWNfY2VydBgBIAEoCzIWLnBsYXRmb3JtLnR5cGUudjIuRmlsZToG+rYYAggBIowBChVWZXJpZnlBY2NvdW50UmVzcG9uc2USOwoMc3BlY19jb250ZXh0GAEgASgLMiUucGxhdGZvcm0uc3BlYy52Mi5TcGVjUmVzcG9uc2VDb250ZXh0Ei4KB2FjY291bnQYAiABKAsyHS5wbGF0Zm9ybS5pYW0udjJhbHBoYS5BY2NvdW50Ogb6thgCCAIihgEKElNpZ25BY2NvdW50UmVxdWVzdBIMCgRuYW1lGAEgASgJEi0KCXBlZXJfdHlwZRgCIAEoDjIaLnBsYXRmb3JtLnR5cGUudjIuUGVlclR5cGUSKwoLcHVibGljX2NlcnQYAyABKAsyFi5wbGF0Zm9ybS50eXBlLnYyLkZpbGU6Bvq2GAIIASKKAQoTU2lnbkFjY291bnRSZXNwb25zZRI7CgxzcGVjX2NvbnRleHQYASABKAsyJS5wbGF0Zm9ybS5zcGVjLnYyLlNwZWNSZXNwb25zZUNvbnRleHQSLgoHYWNjb3VudBgCIAEoCzIdLnBsYXRmb3JtLmlhbS52MmFscGhhLkFjY291bnQ6Bvq2GAIIAiK3AQoHQWNjb3VudBISCgJpZBgBIAEoCUIGyrcYAggBEi4KCmNyZWF0ZWRfYXQYAiABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wEi4KCnVwZGF0ZWRfYXQYAyABKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wEjAKCmNyZWRlbnRpYWwYBCABKAsyHC5wbGF0Zm9ybS50eXBlLnYyLkNyZWRlbnRpYWw6Bvq2GAIIAiphCgtBY2NvdW50VHlwZRIcChhBQ0NPVU5UX1RZUEVfVU5TUEVDSUZJRUQQABIdChlBQ0NPVU5UX1RZUEVfT1JHQU5JWkFUSU9OEAESFQoRQUNDT1VOVF9UWVBFX1VTRVIQAirDAQoSRXZlbnRBY2NvdW50U3RhdHVzEiQKIEVWRU5UX0FDQ09VTlRfU1RBVFVTX1VOU1BFQ0lGSUVEEAASIQodRVZFTlRfQUNDT1VOVF9TVEFUVVNfQ1JFQVRJTkcQARIiCh5FVkVOVF9BQ0NPVU5UX1NUQVRVU19WRVJJRllJTkcQAhIgChxFVkVOVF9BQ0NPVU5UX1NUQVRVU19TSUdOSU5HEAMSHgoaRVZFTlRfQUNDT1VOVF9TVEFUVVNfRVJST1IQBCqoAQoPQWNjb3VudENvbW1hbmRzEiAKHEFDQ09VTlRfQ09NTUFORFNfVU5TUEVDSUZJRUQQABIjCh9BQ0NPVU5UX0NPTU1BTkRTX0NSRUFURV9BQ0NPVU5UEAESIwofQUNDT1VOVF9DT01NQU5EU19WRVJJRllfQUNDT1VOVBACEiEKHUFDQ09VTlRfQ09NTUFORFNfU0lHTl9BQ0NPVU5UEAMaBpK4GAIIAyqtAQoNQWNjb3VudEV2ZW50cxIeChpBQ0NPVU5UX0VWRU5UU19VTlNQRUNJRklFRBAAEiwKHkFDQ09VTlRfRVZFTlRTX0NSRUFURURfQUNDT1VOVBABGgjiuBgECAEYARIjCh9BQ0NPVU5UX0VWRU5UU19WRVJJRklFRF9BQ0NPVU5UEAISIQodQUNDT1VOVF9FVkVOVFNfU0lHTkVEX0FDQ09VTlQQAxoGkrgYAggEMuADCg5BY2NvdW50U2VydmljZRKXAQoNQ3JlYXRlQWNjb3VudBIqLnBsYXRmb3JtLmlhbS52MmFscGhhLkNyZWF0ZUFjY291bnRSZXF1ZXN0GisucGxhdGZvcm0uaWFtLnYyYWxwaGEuQ3JlYXRlQWNjb3VudFJlc3BvbnNlIi2ithgKIAIyBmNyZWF0Zaq2GAIIAoLT5JMCEzoBKiIOL3YyYWxwaGEvaWFtL2ESngEKDVZlcmlmeUFjY291bnQSKi5wbGF0Zm9ybS5pYW0udjJhbHBoYS5WZXJpZnlBY2NvdW50UmVxdWVzdBorLnBsYXRmb3JtLmlhbS52MmFscGhhLlZlcmlmeUFjY291bnRSZXNwb25zZSI0orYYCiACMgZ2ZXJpZnmqthgCCAOC0+STAho6ASoiFS92MmFscGhhL2lhbS9hL3ZlcmlmeRKSAQoLU2lnbkFjY291bnQSKC5wbGF0Zm9ybS5pYW0udjJhbHBoYS5TaWduQWNjb3VudFJlcXVlc3QaKS5wbGF0Zm9ybS5pYW0udjJhbHBoYS5TaWduQWNjb3VudFJlc3BvbnNlIi6ithgGMgRzaWduqrYYAggDgtPkkwIYOgEqIhMvdjJhbHBoYS9pYW0vYS9zaWduQqEBWlxnaXRodWIuY29tL29wZW5lY29zeXN0ZW1zL2Vjb3N5c3RlbS9saWJzL3B1YmxpYy9nby9zZGsvZ2VuL3BsYXRmb3JtL2lhbS92MmFscGhhO2lhbXYyYWxwaGFwYoLEEwIIAYK1GAYIAxABGAKKtRgcCgdhY2NvdW50EghhY2NvdW50cyIDamFuKAI4AZK1GAMKAQOatRgCCAGitRgCCAFiBnByb3RvMw", [file_google_api_annotations, file_platform_options_v2_annotations, file_platform_spec_v2_spec, file_google_protobuf_duration, file_google_protobuf_timestamp, file_platform_type_v2_file, file_platform_type_v2_credential, file_platform_iam_v2alpha_account_authority, file_platform_type_v2_peer]);

/**
 * Describes the message platform.iam.v2alpha.AccountConfiguration.
 * Use `create(AccountConfigurationSchema)` to create a new message.
 */
export const AccountConfigurationSchema = /*@__PURE__*/
  messageDesc(file_platform_iam_v2alpha_account, 0);

/**
 * Describes the message platform.iam.v2alpha.CreateAccountRequest.
 * Use `create(CreateAccountRequestSchema)` to create a new message.
 */
export const CreateAccountRequestSchema = /*@__PURE__*/
  messageDesc(file_platform_iam_v2alpha_account, 1);

/**
 * Describes the message platform.iam.v2alpha.CreateAccountResponse.
 * Use `create(CreateAccountResponseSchema)` to create a new message.
 */
export const CreateAccountResponseSchema = /*@__PURE__*/
  messageDesc(file_platform_iam_v2alpha_account, 2);

/**
 * Describes the message platform.iam.v2alpha.VerifyAccountRequest.
 * Use `create(VerifyAccountRequestSchema)` to create a new message.
 */
export const VerifyAccountRequestSchema = /*@__PURE__*/
  messageDesc(file_platform_iam_v2alpha_account, 3);

/**
 * Describes the message platform.iam.v2alpha.VerifyAccountResponse.
 * Use `create(VerifyAccountResponseSchema)` to create a new message.
 */
export const VerifyAccountResponseSchema = /*@__PURE__*/
  messageDesc(file_platform_iam_v2alpha_account, 4);

/**
 * Describes the message platform.iam.v2alpha.SignAccountRequest.
 * Use `create(SignAccountRequestSchema)` to create a new message.
 */
export const SignAccountRequestSchema = /*@__PURE__*/
  messageDesc(file_platform_iam_v2alpha_account, 5);

/**
 * Describes the message platform.iam.v2alpha.SignAccountResponse.
 * Use `create(SignAccountResponseSchema)` to create a new message.
 */
export const SignAccountResponseSchema = /*@__PURE__*/
  messageDesc(file_platform_iam_v2alpha_account, 6);

/**
 * Describes the message platform.iam.v2alpha.Account.
 * Use `create(AccountSchema)` to create a new message.
 */
export const AccountSchema = /*@__PURE__*/
  messageDesc(file_platform_iam_v2alpha_account, 7);

/**
 * Describes the enum platform.iam.v2alpha.AccountType.
 */
export const AccountTypeSchema = /*@__PURE__*/
  enumDesc(file_platform_iam_v2alpha_account, 0);

/**
 * Supported account type for subscription.
 *
 * @generated from enum platform.iam.v2alpha.AccountType
 */
export const AccountType = /*@__PURE__*/
  tsEnum(AccountTypeSchema);

/**
 * Describes the enum platform.iam.v2alpha.EventAccountStatus.
 */
export const EventAccountStatusSchema = /*@__PURE__*/
  enumDesc(file_platform_iam_v2alpha_account, 1);

/**
 * The current status of a account
 *
 * @generated from enum platform.iam.v2alpha.EventAccountStatus
 */
export const EventAccountStatus = /*@__PURE__*/
  tsEnum(EventAccountStatusSchema);

/**
 * Describes the enum platform.iam.v2alpha.AccountCommands.
 */
export const AccountCommandsSchema = /*@__PURE__*/
  enumDesc(file_platform_iam_v2alpha_account, 2);

/**
 * @generated from enum platform.iam.v2alpha.AccountCommands
 */
export const AccountCommands = /*@__PURE__*/
  tsEnum(AccountCommandsSchema);

/**
 * Describes the enum platform.iam.v2alpha.AccountEvents.
 */
export const AccountEventsSchema = /*@__PURE__*/
  enumDesc(file_platform_iam_v2alpha_account, 3);

/**
 * @generated from enum platform.iam.v2alpha.AccountEvents
 */
export const AccountEvents = /*@__PURE__*/
  tsEnum(AccountEventsSchema);

/**
 * Account Service exposes capabilities to connect to an Ecosystem
 *
 * @generated from service platform.iam.v2alpha.AccountService
 */
export const AccountService = /*@__PURE__*/
  serviceDesc(file_platform_iam_v2alpha_account, 0);

