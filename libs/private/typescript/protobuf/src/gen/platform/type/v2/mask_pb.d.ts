// @generated by protoc-gen-es v2.2.3
// @generated from file platform/type/v2/mask.proto (package platform.type.v2, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";
import type { FieldMask } from "@bufbuild/protobuf/wkt";

/**
 * Describes the file platform/type/v2/mask.proto.
 */
export declare const file_platform_type_v2_mask: GenFile;

/**
 * @generated from message platform.type.v2.ResponseMask
 */
export declare type ResponseMask = Message<"platform.type.v2.ResponseMask"> & {
  /**
   * @generated from field: google.protobuf.FieldMask field_mask = 1;
   */
  fieldMask?: FieldMask;

  /**
   * @generated from field: google.protobuf.FieldMask policy_mask = 2;
   */
  policyMask?: FieldMask;
};

/**
 * Describes the message platform.type.v2.ResponseMask.
 * Use `create(ResponseMaskSchema)` to create a new message.
 */
export declare const ResponseMaskSchema: GenMessage<ResponseMask>;
