// @generated by protoc-gen-es v2.2.3
// @generated from file platform/options/v2/annotations.proto (package platform.options.v2, syntax proto3)
/* eslint-disable */

import { enumDesc, extDesc, fileDesc, messageDesc, tsEnum } from "@bufbuild/protobuf/codegenv1";
import { file_google_protobuf_descriptor } from "@bufbuild/protobuf/wkt";

/**
 * Describes the file platform/options/v2/annotations.proto.
 */
export const file_platform_options_v2_annotations = /*@__PURE__*/
  fileDesc("CiVwbGF0Zm9ybS9vcHRpb25zL3YyL2Fubm90YXRpb25zLnByb3RvEhNwbGF0Zm9ybS5vcHRpb25zLnYyIkAKDk5ldHdvcmtPcHRpb25zEi4KBHR5cGUYASABKA4yIC5wbGF0Zm9ybS5vcHRpb25zLnYyLk5ldHdvcmtUeXBlIqQBCgpBcGlPcHRpb25zEioKBHR5cGUYASABKA4yHC5wbGF0Zm9ybS5vcHRpb25zLnYyLkFwaVR5cGUSMAoFY3ljbGUYAiABKA4yIS5wbGF0Zm9ybS5vcHRpb25zLnYyLkFwaUxpZmVjeWNsZRI4CglpbnRlcmZhY2UYAyABKA4yJS5wbGF0Zm9ybS5vcHRpb25zLnYyLkFwaUludGVyZmFjZVR5cGUigwIKDUVudGl0eU9wdGlvbnMSDgoGZW50aXR5GAEgASgJEhUKDWVudGl0eV9wbHVyYWwYAiABKAkSLQoEdHlwZRgDIAEoDjIfLnBsYXRmb3JtLm9wdGlvbnMudjIuRW50aXR5VHlwZRIRCgluYW1lc3BhY2UYBCABKAkSOwoLY29uc2lzdGVuY3kYBSABKA4yJi5wbGF0Zm9ybS5vcHRpb25zLnYyLkVudGl0eUNvbnNpc3RlbmN5EjcKCWhpZXJhcmNoeRgGIAEoDjIkLnBsYXRmb3JtLm9wdGlvbnMudjIuRW50aXR5SGllcmFyY2h5EhMKC3ZlcnNpb25hYmxlGAcgASgIIkcKD0xhbmd1YWdlT3B0aW9ucxI0CglsYW5ndWFnZXMYASADKA4yIS5wbGF0Zm9ybS5vcHRpb25zLnYyLkxhbmd1YWdlVHlwZSIhCg5HcmFwaHFsT3B0aW9ucxIPCgdlbmFibGVkGAEgASgIIjoKElNwZWNTZXJ2aWNlT3B0aW9ucxIRCglncnBjX3BvcnQYASABKAUSEQoJaHR0cF9wb3J0GAIgASgFIjIKD1JhdGVMaW1pdEZpbHRlchIPCgdlbmFibGVkGAEgASgIEg4KBm1ldHJpYxgCIAEoCSImChNBdXRob3JpemF0aW9uRmlsdGVyEg8KB2VuYWJsZWQYASABKAgiIAoNQ29uc2VudEZpbHRlchIPCgdlbmFibGVkGAEgASgIIr4BCgxQcm94eU9wdGlvbnMSPwoNYXV0aG9yaXphdGlvbhgBIAEoCzIoLnBsYXRmb3JtLm9wdGlvbnMudjIuQXV0aG9yaXphdGlvbkZpbHRlchIzCgdjb25zZW50GAIgASgLMiIucGxhdGZvcm0ub3B0aW9ucy52Mi5Db25zZW50RmlsdGVyEjgKCnJhdGVfbGltaXQYAyABKAsyJC5wbGF0Zm9ybS5vcHRpb25zLnYyLlJhdGVMaW1pdEZpbHRlciJEChBDb25uZWN0b3JPcHRpb25zEjAKBHR5cGUYASABKA4yIi5wbGF0Zm9ybS5vcHRpb25zLnYyLkNvbm5lY3RvclR5cGUiOgoLQ1FSU09wdGlvbnMSKwoEdHlwZRgBIAEoDjIdLnBsYXRmb3JtLm9wdGlvbnMudjIuQ1FSU1R5cGUiIQoQUmF0ZUxpbWl0T3B0aW9ucxINCgVsaW1pdBgBIAEoCCJVChFQZXJtaXNzaW9uT3B0aW9ucxISCgpwZXJtaXNzaW9uGAEgASgJEiwKBXJvbGVzGAIgAygOMh0ucGxhdGZvcm0ub3B0aW9ucy52Mi5BdXRoUm9sZSI8CgxHcmFwaE9wdGlvbnMSLAoEdHlwZRgBIAEoDjIeLnBsYXRmb3JtLm9wdGlvbnMudjIuR3JhcGhUeXBlImwKEkVudGl0eUZpZWxkT3B0aW9ucxILCgNrZXkYASABKAgSNAoIYmVoYXZpb3IYAiADKA4yIi5wbGF0Zm9ybS5vcHRpb25zLnYyLkZpZWxkQmVoYXZpb3ISEwoLdmVyc2lvbl9rZXkYAyABKAgilQMKEVNwZWNDb25maWd1cmF0aW9uEgsKA2tleRgBIAEoCRJKCgR0eXBlGAIgASgOMjwucGxhdGZvcm0ub3B0aW9ucy52Mi5TcGVjQ29uZmlndXJhdGlvbi5TcGVjQ29uZmlndXJhdGlvblR5cGUSEwoLZGVzY3JpcHRpb24YAyABKAkSFAoMb3ZlcnJpZGVhYmxlGAQgASgIEg8KB2VuYWJsZWQYBSABKAgi6gEKFVNwZWNDb25maWd1cmF0aW9uVHlwZRInCiNTUEVDX0NPTkZJR1VSQVRJT05fVFlQRV9VTlNQRUNJRklFRBAAEiIKHlNQRUNfQ09ORklHVVJBVElPTl9UWVBFX1NUUklORxABEiAKHFNQRUNfQ09ORklHVVJBVElPTl9UWVBFX0JPT0wQAhIfChtTUEVDX0NPTkZJR1VSQVRJT05fVFlQRV9JTlQQAxIgChxTUEVDX0NPTkZJR1VSQVRJT05fVFlQRV9MSVNUEAQSHwobU1BFQ19DT05GSUdVUkFUSU9OX1RZUEVfTUFQEAUiPQoUQ29uZmlndXJhdGlvbk9wdGlvbnMSDwoHZW5hYmxlZBgBIAEoCBIUCgxmaWVsZF9udW1iZXIYAiABKAUiWgoZQ29uZmlndXJhdGlvbkZpZWxkT3B0aW9ucxI9Cg1jb25maWd1cmF0aW9uGAEgASgLMiYucGxhdGZvcm0ub3B0aW9ucy52Mi5TcGVjQ29uZmlndXJhdGlvbiJcChBTeW50aGV0aWNPcHRpb25zEhYKDmRpY3Rpb25hcnlfa2V5GAEgASgJEjAKBHR5cGUYAiABKA4yIi5wbGF0Zm9ybS5vcHRpb25zLnYyLlN5bnRoZXRpY1R5cGUiRwoOQ2xhc3NpZmljYXRpb24SNQoEdHlwZRgBIAEoDjInLnBsYXRmb3JtLm9wdGlvbnMudjIuQ2xhc3NpZmljYXRpb25UeXBlIj4KC1NwZWNPcHRpb25zEi8KBHR5cGUYASABKA4yIS5wbGF0Zm9ybS5vcHRpb25zLnYyLlNwZWNFbnVtVHlwZSJNCg5CaWxsaW5nT3B0aW9ucxIQCghiaWxsYWJsZRgBIAEoCBIYChBwYXJ0bmVyX2JpbGxhYmxlGAIgASgIEg8KB21ldGVyZWQYAyABKAgiRAoRRXZlbnRTY29wZU9wdGlvbnMSLwoGc2NvcGVzGAEgAygOMh8ucGxhdGZvcm0ub3B0aW9ucy52Mi5FdmVudFNjb3BlIiMKDEV2ZW50T3B0aW9ucxITCgt2ZXJzaW9uYWJsZRgBIAEoCCJHCg9BdXRoUm9sZU9wdGlvbnMSNAoJcm9sZV90eXBlGAEgASgOMiEucGxhdGZvcm0ub3B0aW9ucy52Mi5BdXRoUm9sZVR5cGUiIgoOUm91dGluZU9wdGlvbnMSEAoIbGlzdGVuZXIYASABKAkqewoLTmV0d29ya1R5cGUSHAoYTkVUV09SS19UWVBFX1VOU1BFQ0lGSUVEEAASGQoVTkVUV09SS19UWVBFX1VOREVSTEFZEAESGQoVTkVUV09SS19UWVBFX0lOVEVSTkVUEAISGAoUTkVUV09SS19UWVBFX09WRVJMQVkQAyp2CgdBcGlUeXBlEhgKFEFQSV9UWVBFX1VOU1BFQ0lGSUVEEAASFAoQQVBJX1RZUEVfUFJJVkFURRABEhQKEEFQSV9UWVBFX1BBUlRORVIQAhITCg9BUElfVFlQRV9QVUJMSUMQAxIQCgxBUElfVFlQRV9QT0MQBCqHAgoQQXBpSW50ZXJmYWNlVHlwZRIiCh5BUElfSU5URVJGQUNFX1RZUEVfVU5TUEVDSUZJRUQQABIbChdBUElfSU5URVJGQUNFX1RZUEVfTUVUQRABEiIKHkFQSV9JTlRFUkZBQ0VfVFlQRV9PUEVSQVRJT05BTBACEiMKH0FQSV9JTlRFUkZBQ0VfVFlQRV9DT05UUklCVVRJT04QAxIhCh1BUElfSU5URVJGQUNFX1RZUEVfQU5BTFlUSUNBTBAEEiEKHUFQSV9JTlRFUkZBQ0VfVFlQRV9PQlNFUlZBQkxFEAUSIwofQVBJX0lOVEVSRkFDRV9UWVBFX0RJU0NPVkVSQUJMRRAGKucBCgxBcGlMaWZlY3ljbGUSHQoZQVBJX0xJRkVDWUNMRV9VTlNQRUNJRklFRBAAEhcKE0FQSV9MSUZFQ1lDTEVfQUxQSEEQARIWChJBUElfTElGRUNZQ0xFX0JFVEEQAhImCiJBUElfTElGRUNZQ0xFX0xJTUlURURfQVZBSUxBQklMSVRZEAMSJgoiQVBJX0xJRkVDWUNMRV9HRU5FUkFMX0FWQUlMQUJJTElUWRAEEhwKGEFQSV9MSUZFQ1lDTEVfREVQUkVDQVRFRBAFEhkKFUFQSV9MSUZFQ1lDTEVfUkVUSVJFRBAGKtoBCgpFbnRpdHlUeXBlEhsKF0VOVElUWV9UWVBFX1VOU1BFQ0lGSUVEEAASGQoVRU5USVRZX1RZUEVfQUVST1NQSUtFEAESFgoSRU5USVRZX1RZUEVfREdSQVBIEAISFwoTRU5USVRZX1RZUEVfTU9OR09EQhADEhgKFEVOVElUWV9UWVBFX0JJR1FVRVJZEAQSFQoRRU5USVRZX1RZUEVfUkVESVMQBRIXChNFTlRJVFlfVFlQRV9ST0NLU0RCEAYSGQoVRU5USVRZX1RZUEVfQ09VQ0hCQVNFEAcqdwoRRW50aXR5Q29uc2lzdGVuY3kSIgoeRU5USVRZX0NPTlNJU1RFTkNZX1VOU1BFQ0lGSUVEEAASHQoZRU5USVRZX0NPTlNJU1RFTkNZX1NUUk9ORxABEh8KG0VOVElUWV9DT05TSVNURU5DWV9FVkVOVFVBTBACKpUBCg9FbnRpdHlIaWVyYXJjaHkSIAocRU5USVRZX0hJRVJBUkNIWV9VTlNQRUNJRklFRBAAEh0KGUVOVElUWV9ISUVSQVJDSFlfUExBVEZPUk0QARIhCh1FTlRJVFlfSElFUkFSQ0hZX09SR0FOSVpBVElPThACEh4KGkVOVElUWV9ISUVSQVJDSFlfV09SS1NQQUNFEAMqzAIKDExhbmd1YWdlVHlwZRIdChlMQU5HVUFHRV9UWVBFX1VOU1BFQ0lGSUVEEAASGwoXTEFOR1VBR0VfVFlQRV9DUExVU1BMVVMQARIWChJMQU5HVUFHRV9UWVBFX1JVU1QQAhIYChRMQU5HVUFHRV9UWVBFX0dPTEFORxADEhYKEkxBTkdVQUdFX1RZUEVfSkFWQRAEEhgKFExBTkdVQUdFX1RZUEVfUFlUSE9OEAUSHAoYTEFOR1VBR0VfVFlQRV9UWVBFU0NSSVBUEAYSGAoUTEFOR1VBR0VfVFlQRV9DU0hBUlAQBxIXChNMQU5HVUFHRV9UWVBFX1NXSUZUEAgSGQoVTEFOR1VBR0VfVFlQRV9BTkRST0lEEAkSGQoVTEFOR1VBR0VfVFlQRV9HUkFQSFFMEAoSFQoRTEFOR1VBR0VfVFlQRV9MVUEQCypNCg1Db25uZWN0b3JUeXBlEh4KGkNPTk5FQ1RPUl9UWVBFX1VOU1BFQ0lGSUVEEAASHAoYQ09OTkVDVE9SX1RZUEVfUkVGRVJFTkNFEAEqrQEKDEF1dGhSb2xlVHlwZRIeChpBVVRIX1JPTEVfVFlQRV9VTlNQRUNJRklFRBAAEhsKF0FVVEhfUk9MRV9UWVBFX1BMQVRGT1JNEAESHwobQVVUSF9ST0xFX1RZUEVfT1JHQU5JWkFUSU9OEAISHAoYQVVUSF9ST0xFX1RZUEVfV09SS1NQQUNFEAMSIQodQVVUSF9ST0xFX1RZUEVfQ09OTkVDVEVEX1RFU1QQBCr6AwoIQ1FSU1R5cGUSGQoVQ1FSU19UWVBFX1VOU1BFQ0lGSUVEEAASEgoOQ1FSU19UWVBFX05PTkUQARIdChlDUVJTX1RZUEVfTVVUQVRJT05fQ1JFQVRFEAISHQoZQ1FSU19UWVBFX01VVEFUSU9OX1VQREFURRADEh0KGUNRUlNfVFlQRV9NVVRBVElPTl9ERUxFVEUQBBIkCiBDUVJTX1RZUEVfTVVUQVRJT05fQ0xJRU5UX1NUUkVBTRAFEiQKIENRUlNfVFlQRV9NVVRBVElPTl9TRVJWRVJfU1RSRUFNEAYSIgoeQ1FSU19UWVBFX01VVEFUSU9OX0JJRElfU1RSRUFNEAcSGAoUQ1FSU19UWVBFX1FVRVJZX0xJU1QQCBIaChZDUVJTX1RZUEVfUVVFUllfU1RSRUFNEAkSFwoTQ1FSU19UWVBFX1FVRVJZX0dFVBAKEiAKHENRUlNfVFlQRV9RVUVSWV9FVkVOVF9TVFJFQU0QCxIhCh1DUVJTX1RZUEVfUVVFUllfQ0xJRU5UX1NUUkVBTRAMEiEKHUNRUlNfVFlQRV9RVUVSWV9TRVJWRVJfU1RSRUFNEA0SHwobQ1FSU19UWVBFX1FVRVJZX0JJRElfU1RSRUFNEA4SGgoWQ1FSU19UWVBFX1FVRVJZX0VYSVNUUxAPKtUICghBdXRoUm9sZRIZChVBVVRIX1JPTEVfVU5TUEVDSUZJRUQQABIqCh5BVVRIX1JPTEVfUExBVEZPUk1fU1VQRVJfQURNSU4QZBoG8rgYAggBEi0KIUFVVEhfUk9MRV9QTEFURk9STV9DTElOSUNBTF9BRE1JThBlGgbyuBgCCAESLAogQVVUSF9ST0xFX1BMQVRGT1JNX0JJTExJTkdfQURNSU4QZhoG8rgYAggBEiQKGEFVVEhfUk9MRV9QTEFURk9STV9BRE1JThBnGgbyuBgCCAESJgoaQVVUSF9ST0xFX1BMQVRGT1JNX01BTkFHRVIQaBoG8rgYAggBEiMKF0FVVEhfUk9MRV9QTEFURk9STV9VU0VSEGkaBvK4GAIIARIlChlBVVRIX1JPTEVfUExBVEZPUk1fVklFV0VSEGoaBvK4GAIIARIvCiJBVVRIX1JPTEVfT1JHQU5JWkFUSU9OX1NVUEVSX0FETUlOEMgBGgbyuBgCCAISMgolQVVUSF9ST0xFX09SR0FOSVpBVElPTl9DTElOSUNBTF9BRE1JThDJARoG8rgYAggCEjEKJEFVVEhfUk9MRV9PUkdBTklaQVRJT05fQklMTElOR19BRE1JThDKARoG8rgYAggCEikKHEFVVEhfUk9MRV9PUkdBTklaQVRJT05fQURNSU4QywEaBvK4GAIIAhIrCh5BVVRIX1JPTEVfT1JHQU5JWkFUSU9OX01BTkFHRVIQzAEaBvK4GAIIAhIoChtBVVRIX1JPTEVfT1JHQU5JWkFUSU9OX1VTRVIQzQEaBvK4GAIIAhIqCh1BVVRIX1JPTEVfT1JHQU5JWkFUSU9OX1ZJRVdFUhDOARoG8rgYAggCEiwKH0FVVEhfUk9MRV9XT1JLU1BBQ0VfU1VQRVJfQURNSU4QrAIaBvK4GAIIAxIvCiJBVVRIX1JPTEVfV09SS1NQQUNFX0NMSU5JQ0FMX0FETUlOEK0CGgbyuBgCCAMSLgohQVVUSF9ST0xFX1dPUktTUEFDRV9CSUxMSU5HX0FETUlOEK4CGgbyuBgCCAMSJgoZQVVUSF9ST0xFX1dPUktTUEFDRV9BRE1JThCvAhoG8rgYAggDEigKG0FVVEhfUk9MRV9XT1JLU1BBQ0VfTUFOQUdFUhCwAhoG8rgYAggDEiUKGEFVVEhfUk9MRV9XT1JLU1BBQ0VfVVNFUhCxAhoG8rgYAggDEicKGkFVVEhfUk9MRV9XT1JLU1BBQ0VfVklFV0VSELICGgbyuBgCCAMSJAogQVVUSF9ST0xFX0NPTk5FQ1RFRF9URVNUX1BBVElFTlQQDxIlCiFBVVRIX1JPTEVfQ09OTkVDVEVEX1RFU1RfUFJPVklERVIQEBIiCh5BVVRIX1JPTEVfQ09OTkVDVEVEX1RFU1RfUFJPWFkQERIjCh9BVVRIX1JPTEVfQ09OTkVDVEVEX1RFU1RfVklFV0VSEBIqVAoJR3JhcGhUeXBlEhoKFkdSQVBIX1RZUEVfVU5TUEVDSUZJRUQQABIUChBHUkFQSF9UWVBFX0lOUFVUEAESFQoRR1JBUEhfVFlQRV9PVVRQVVQQAiqPAgoNRmllbGRCZWhhdmlvchIeChpGSUVMRF9CRUhBVklPUl9VTlNQRUNJRklFRBAAEhsKF0ZJRUxEX0JFSEFWSU9SX09QVElPTkFMEAESGwoXRklFTERfQkVIQVZJT1JfUkVRVUlSRUQQAhIeChpGSUVMRF9CRUhBVklPUl9PVVRQVVRfT05MWRADEh0KGUZJRUxEX0JFSEFWSU9SX0lOUFVUX09OTFkQBBIcChhGSUVMRF9CRUhBVklPUl9JTU1VVEFCTEUQBRIhCh1GSUVMRF9CRUhBVklPUl9VTk9SREVSRURfTElTVBAGEiQKIEZJRUxEX0JFSEFWSU9SX05PTl9FTVBUWV9ERUZBVUxUEAcqzAEKDVN5bnRoZXRpY1R5cGUSHgoaU1lOVEhFVElDX1RZUEVfVU5TUEVDSUZJRUQQABIpCiVTWU5USEVUSUNfVFlQRV9ESVJFQ1RfRlJPTV9ESUNUSU9OQVJZEAESKgomU1lOVEhFVElDX1RZUEVfU0VMRUNUX1JBTkRPTV9GUk9NX0xJU1QQAhIhCh1TWU5USEVUSUNfVFlQRV9MSVNUX0ZST01fTElTVBADEiEKHVNZTlRIRVRJQ19UWVBFX0dFTkVSQVRFRF9MT0dPEAYqxAIKEkNsYXNzaWZpY2F0aW9uVHlwZRIjCh9DTEFTU0lGSUNBVElPTl9UWVBFX1VOU1BFQ0lGSUVEEAASJwojQ0xBU1NJRklDQVRJT05fVFlQRV9ERVJJVkFUSVZFX0RBVEEQARIlCiFDTEFTU0lGSUNBVElPTl9UWVBFX0RFX0lERU5USUZJRUQQAhIeChpDTEFTU0lGSUNBVElPTl9UWVBFX1BVQkxJQxADEiQKIENMQVNTSUZJQ0FUSU9OX1RZUEVfSU5URVJOQUxfVVNFEAQSJAogQ0xBU1NJRklDQVRJT05fVFlQRV9DT05GSURFTlRJQUwQBRIiCh5DTEFTU0lGSUNBVElPTl9UWVBFX1JFU1RSSUNURUQQBhIpCiVDTEFTU0lGSUNBVElPTl9UWVBFX0hJR0hMWV9SRVNUUklDVEVEEAcqwAEKDFNwZWNFbnVtVHlwZRIeChpTUEVDX0VOVU1fVFlQRV9VTlNQRUNJRklFRBAAEhcKE1NQRUNfRU5VTV9UWVBFX05PTkUQARIZChVTUEVDX0VOVU1fVFlQRV9UT1BJQ1MQAhIbChdTUEVDX0VOVU1fVFlQRV9DT01NQU5EUxADEhkKFVNQRUNfRU5VTV9UWVBFX0VWRU5UUxAEEiQKIFNQRUNfRU5VTV9UWVBFX1JPVVRJTkVfTElTVEVORVJTEAUqeAoKRXZlbnRTY29wZRIbChdFVkVOVF9TQ09QRV9VTlNQRUNJRklFRBAAEhQKEEVWRU5UX1NDT1BFX1VTRVIQARIZChVFVkVOVF9TQ09QRV9XT1JLU1BBQ0UQAhIcChhFVkVOVF9TQ09QRV9PUkdBTklaQVRJT04QAzpmCgxuZXR3b3JrX2ZpbGUSHC5nb29nbGUucHJvdG9idWYuRmlsZU9wdGlvbnMYwLgCIAEoCzIjLnBsYXRmb3JtLm9wdGlvbnMudjIuTmV0d29ya09wdGlvbnNSC25ldHdvcmtGaWxlOloKCGFwaV9maWxlEhwuZ29vZ2xlLnByb3RvYnVmLkZpbGVPcHRpb25zGNCGAyABKAsyHy5wbGF0Zm9ybS5vcHRpb25zLnYyLkFwaU9wdGlvbnNSB2FwaUZpbGU6WgoGZW50aXR5EhwuZ29vZ2xlLnByb3RvYnVmLkZpbGVPcHRpb25zGNGGAyABKAsyIi5wbGF0Zm9ybS5vcHRpb25zLnYyLkVudGl0eU9wdGlvbnNSBmVudGl0eTpgCghsYW5ndWFnZRIcLmdvb2dsZS5wcm90b2J1Zi5GaWxlT3B0aW9ucxjShgMgASgLMiQucGxhdGZvcm0ub3B0aW9ucy52Mi5MYW5ndWFnZU9wdGlvbnNSCGxhbmd1YWdlOl0KB2dyYXBocWwSHC5nb29nbGUucHJvdG9idWYuRmlsZU9wdGlvbnMY04YDIAEoCzIjLnBsYXRmb3JtLm9wdGlvbnMudjIuR3JhcGhxbE9wdGlvbnNSB2dyYXBocWw6bwoNY29uZmlndXJhdGlvbhIcLmdvb2dsZS5wcm90b2J1Zi5GaWxlT3B0aW9ucxjUhgMgASgLMikucGxhdGZvcm0ub3B0aW9ucy52Mi5Db25maWd1cmF0aW9uT3B0aW9uc1INY29uZmlndXJhdGlvbjpgChxoYXNfbXVsdGlwbGVfaW1wbGVtZW50YXRpb25zEhwuZ29vZ2xlLnByb3RvYnVmLkZpbGVPcHRpb25zGNWGAyABKAhSGmhhc011bHRpcGxlSW1wbGVtZW50YXRpb25zOmMKC2FwaV9zZXJ2aWNlEh8uZ29vZ2xlLnByb3RvYnVmLlNlcnZpY2VPcHRpb25zGNqGAyABKAsyHy5wbGF0Zm9ybS5vcHRpb25zLnYyLkFwaU9wdGlvbnNSCmFwaVNlcnZpY2U6ZAoHc2VydmljZRIfLmdvb2dsZS5wcm90b2J1Zi5TZXJ2aWNlT3B0aW9ucxjbhgMgASgLMicucGxhdGZvcm0ub3B0aW9ucy52Mi5TcGVjU2VydmljZU9wdGlvbnNSB3NlcnZpY2U6WgoFcHJveHkSHy5nb29nbGUucHJvdG9idWYuU2VydmljZU9wdGlvbnMY3IYDIAEoCzIhLnBsYXRmb3JtLm9wdGlvbnMudjIuUHJveHlPcHRpb25zUgVwcm94eTpmCgljb25uZWN0b3ISHy5nb29nbGUucHJvdG9idWYuU2VydmljZU9wdGlvbnMY3YYDIAEoCzIlLnBsYXRmb3JtLm9wdGlvbnMudjIuQ29ubmVjdG9yT3B0aW9uc1IJY29ubmVjdG9yOmAKCmFwaV9tZXRob2QSHi5nb29nbGUucHJvdG9idWYuTWV0aG9kT3B0aW9ucxjkhgMgASgLMh8ucGxhdGZvcm0ub3B0aW9ucy52Mi5BcGlPcHRpb25zUglhcGlNZXRob2Q6VgoEY3FycxIeLmdvb2dsZS5wcm90b2J1Zi5NZXRob2RPcHRpb25zGOWGAyABKAsyIC5wbGF0Zm9ybS5vcHRpb25zLnYyLkNRUlNPcHRpb25zUgRjcXJzOmgKCnBlcm1pc3Npb24SHi5nb29nbGUucHJvdG9idWYuTWV0aG9kT3B0aW9ucxjmhgMgASgLMiYucGxhdGZvcm0ub3B0aW9ucy52Mi5QZXJtaXNzaW9uT3B0aW9uc1IKcGVybWlzc2lvbjpbCgRyYXRlEh4uZ29vZ2xlLnByb3RvYnVmLk1ldGhvZE9wdGlvbnMY54YDIAEoCzIlLnBsYXRmb3JtLm9wdGlvbnMudjIuUmF0ZUxpbWl0T3B0aW9uc1IEcmF0ZTpjCgthcGlfbWVzc2FnZRIfLmdvb2dsZS5wcm90b2J1Zi5NZXNzYWdlT3B0aW9ucxjuhgMgASgLMh8ucGxhdGZvcm0ub3B0aW9ucy52Mi5BcGlPcHRpb25zUgphcGlNZXNzYWdlOloKBWdyYXBoEh8uZ29vZ2xlLnByb3RvYnVmLk1lc3NhZ2VPcHRpb25zGO+GAyABKAsyIS5wbGF0Zm9ybS5vcHRpb25zLnYyLkdyYXBoT3B0aW9uc1IFZ3JhcGg6YAoHcm91dGluZRIfLmdvb2dsZS5wcm90b2J1Zi5NZXNzYWdlT3B0aW9ucxjwhgMgASgLMiMucGxhdGZvcm0ub3B0aW9ucy52Mi5Sb3V0aW5lT3B0aW9uc1IHcm91dGluZTpdCglhcGlfZmllbGQSHS5nb29nbGUucHJvdG9idWYuRmllbGRPcHRpb25zGPiGAyABKAsyHy5wbGF0Zm9ybS5vcHRpb25zLnYyLkFwaU9wdGlvbnNSCGFwaUZpZWxkOmsKDGVudGl0eV9maWVsZBIdLmdvb2dsZS5wcm90b2J1Zi5GaWVsZE9wdGlvbnMY+YYDIAEoCzInLnBsYXRmb3JtLm9wdGlvbnMudjIuRW50aXR5RmllbGRPcHRpb25zUgtlbnRpdHlGaWVsZDqAAQoTY29uZmlndXJhdGlvbl9maWVsZBIdLmdvb2dsZS5wcm90b2J1Zi5GaWVsZE9wdGlvbnMY+oYDIAEoCzIuLnBsYXRmb3JtLm9wdGlvbnMudjIuQ29uZmlndXJhdGlvbkZpZWxkT3B0aW9uc1ISY29uZmlndXJhdGlvbkZpZWxkOmQKCXN5bnRoZXRpYxIdLmdvb2dsZS5wcm90b2J1Zi5GaWVsZE9wdGlvbnMY+4YDIAEoCzIlLnBsYXRmb3JtLm9wdGlvbnMudjIuU3ludGhldGljT3B0aW9uc1IJc3ludGhldGljOlQKBHNwZWMSHC5nb29nbGUucHJvdG9idWYuRW51bU9wdGlvbnMYgocDIAEoCzIgLnBsYXRmb3JtLm9wdGlvbnMudjIuU3BlY09wdGlvbnNSBHNwZWM6YgoHYmlsbGluZxIhLmdvb2dsZS5wcm90b2J1Zi5FbnVtVmFsdWVPcHRpb25zGIyHAyABKAsyIy5wbGF0Zm9ybS5vcHRpb25zLnYyLkJpbGxpbmdPcHRpb25zUgdiaWxsaW5nOmwKC2V2ZW50X3Njb3BlEiEuZ29vZ2xlLnByb3RvYnVmLkVudW1WYWx1ZU9wdGlvbnMYjYcDIAEoCzImLnBsYXRmb3JtLm9wdGlvbnMudjIuRXZlbnRTY29wZU9wdGlvbnNSCmV2ZW50U2NvcGU6ZgoJYXV0aF9yb2xlEiEuZ29vZ2xlLnByb3RvYnVmLkVudW1WYWx1ZU9wdGlvbnMYjocDIAEoCzIkLnBsYXRmb3JtLm9wdGlvbnMudjIuQXV0aFJvbGVPcHRpb25zUghhdXRoUm9sZTpcCgVldmVudBIhLmdvb2dsZS5wcm90b2J1Zi5FbnVtVmFsdWVPcHRpb25zGI+HAyABKAsyIS5wbGF0Zm9ybS5vcHRpb25zLnYyLkV2ZW50T3B0aW9uc1IFZXZlbnRCPlo8bGlicy9wcm90b2J1Zi9nby9wcm90b2J1Zi9nZW4vcGxhdGZvcm0vb3B0aW9ucy92MjtvcHRpb252MnBiYgZwcm90bzM", [file_google_protobuf_descriptor]);

/**
 * Describes the message platform.options.v2.NetworkOptions.
 * Use `create(NetworkOptionsSchema)` to create a new message.
 */
export const NetworkOptionsSchema = /*@__PURE__*/
  messageDesc(file_platform_options_v2_annotations, 0);

/**
 * Describes the message platform.options.v2.ApiOptions.
 * Use `create(ApiOptionsSchema)` to create a new message.
 */
export const ApiOptionsSchema = /*@__PURE__*/
  messageDesc(file_platform_options_v2_annotations, 1);

/**
 * Describes the message platform.options.v2.EntityOptions.
 * Use `create(EntityOptionsSchema)` to create a new message.
 */
export const EntityOptionsSchema = /*@__PURE__*/
  messageDesc(file_platform_options_v2_annotations, 2);

/**
 * Describes the message platform.options.v2.LanguageOptions.
 * Use `create(LanguageOptionsSchema)` to create a new message.
 */
export const LanguageOptionsSchema = /*@__PURE__*/
  messageDesc(file_platform_options_v2_annotations, 3);

/**
 * Describes the message platform.options.v2.GraphqlOptions.
 * Use `create(GraphqlOptionsSchema)` to create a new message.
 */
export const GraphqlOptionsSchema = /*@__PURE__*/
  messageDesc(file_platform_options_v2_annotations, 4);

/**
 * Describes the message platform.options.v2.SpecServiceOptions.
 * Use `create(SpecServiceOptionsSchema)` to create a new message.
 */
export const SpecServiceOptionsSchema = /*@__PURE__*/
  messageDesc(file_platform_options_v2_annotations, 5);

/**
 * Describes the message platform.options.v2.RateLimitFilter.
 * Use `create(RateLimitFilterSchema)` to create a new message.
 */
export const RateLimitFilterSchema = /*@__PURE__*/
  messageDesc(file_platform_options_v2_annotations, 6);

/**
 * Describes the message platform.options.v2.AuthorizationFilter.
 * Use `create(AuthorizationFilterSchema)` to create a new message.
 */
export const AuthorizationFilterSchema = /*@__PURE__*/
  messageDesc(file_platform_options_v2_annotations, 7);

/**
 * Describes the message platform.options.v2.ConsentFilter.
 * Use `create(ConsentFilterSchema)` to create a new message.
 */
export const ConsentFilterSchema = /*@__PURE__*/
  messageDesc(file_platform_options_v2_annotations, 8);

/**
 * Describes the message platform.options.v2.ProxyOptions.
 * Use `create(ProxyOptionsSchema)` to create a new message.
 */
export const ProxyOptionsSchema = /*@__PURE__*/
  messageDesc(file_platform_options_v2_annotations, 9);

/**
 * Describes the message platform.options.v2.ConnectorOptions.
 * Use `create(ConnectorOptionsSchema)` to create a new message.
 */
export const ConnectorOptionsSchema = /*@__PURE__*/
  messageDesc(file_platform_options_v2_annotations, 10);

/**
 * Describes the message platform.options.v2.CQRSOptions.
 * Use `create(CQRSOptionsSchema)` to create a new message.
 */
export const CQRSOptionsSchema = /*@__PURE__*/
  messageDesc(file_platform_options_v2_annotations, 11);

/**
 * Describes the message platform.options.v2.RateLimitOptions.
 * Use `create(RateLimitOptionsSchema)` to create a new message.
 */
export const RateLimitOptionsSchema = /*@__PURE__*/
  messageDesc(file_platform_options_v2_annotations, 12);

/**
 * Describes the message platform.options.v2.PermissionOptions.
 * Use `create(PermissionOptionsSchema)` to create a new message.
 */
export const PermissionOptionsSchema = /*@__PURE__*/
  messageDesc(file_platform_options_v2_annotations, 13);

/**
 * Describes the message platform.options.v2.GraphOptions.
 * Use `create(GraphOptionsSchema)` to create a new message.
 */
export const GraphOptionsSchema = /*@__PURE__*/
  messageDesc(file_platform_options_v2_annotations, 14);

/**
 * Describes the message platform.options.v2.EntityFieldOptions.
 * Use `create(EntityFieldOptionsSchema)` to create a new message.
 */
export const EntityFieldOptionsSchema = /*@__PURE__*/
  messageDesc(file_platform_options_v2_annotations, 15);

/**
 * Describes the message platform.options.v2.SpecConfiguration.
 * Use `create(SpecConfigurationSchema)` to create a new message.
 */
export const SpecConfigurationSchema = /*@__PURE__*/
  messageDesc(file_platform_options_v2_annotations, 16);

/**
 * Describes the enum platform.options.v2.SpecConfiguration.SpecConfigurationType.
 */
export const SpecConfiguration_SpecConfigurationTypeSchema = /*@__PURE__*/
  enumDesc(file_platform_options_v2_annotations, 16, 0);

/**
 * Value types that can be used as label values.
 *
 * @generated from enum platform.options.v2.SpecConfiguration.SpecConfigurationType
 */
export const SpecConfiguration_SpecConfigurationType = /*@__PURE__*/
  tsEnum(SpecConfiguration_SpecConfigurationTypeSchema);

/**
 * Describes the message platform.options.v2.ConfigurationOptions.
 * Use `create(ConfigurationOptionsSchema)` to create a new message.
 */
export const ConfigurationOptionsSchema = /*@__PURE__*/
  messageDesc(file_platform_options_v2_annotations, 17);

/**
 * Describes the message platform.options.v2.ConfigurationFieldOptions.
 * Use `create(ConfigurationFieldOptionsSchema)` to create a new message.
 */
export const ConfigurationFieldOptionsSchema = /*@__PURE__*/
  messageDesc(file_platform_options_v2_annotations, 18);

/**
 * Describes the message platform.options.v2.SyntheticOptions.
 * Use `create(SyntheticOptionsSchema)` to create a new message.
 */
export const SyntheticOptionsSchema = /*@__PURE__*/
  messageDesc(file_platform_options_v2_annotations, 19);

/**
 * Describes the message platform.options.v2.Classification.
 * Use `create(ClassificationSchema)` to create a new message.
 */
export const ClassificationSchema = /*@__PURE__*/
  messageDesc(file_platform_options_v2_annotations, 20);

/**
 * Describes the message platform.options.v2.SpecOptions.
 * Use `create(SpecOptionsSchema)` to create a new message.
 */
export const SpecOptionsSchema = /*@__PURE__*/
  messageDesc(file_platform_options_v2_annotations, 21);

/**
 * Describes the message platform.options.v2.BillingOptions.
 * Use `create(BillingOptionsSchema)` to create a new message.
 */
export const BillingOptionsSchema = /*@__PURE__*/
  messageDesc(file_platform_options_v2_annotations, 22);

/**
 * Describes the message platform.options.v2.EventScopeOptions.
 * Use `create(EventScopeOptionsSchema)` to create a new message.
 */
export const EventScopeOptionsSchema = /*@__PURE__*/
  messageDesc(file_platform_options_v2_annotations, 23);

/**
 * Describes the message platform.options.v2.EventOptions.
 * Use `create(EventOptionsSchema)` to create a new message.
 */
export const EventOptionsSchema = /*@__PURE__*/
  messageDesc(file_platform_options_v2_annotations, 24);

/**
 * Describes the message platform.options.v2.AuthRoleOptions.
 * Use `create(AuthRoleOptionsSchema)` to create a new message.
 */
export const AuthRoleOptionsSchema = /*@__PURE__*/
  messageDesc(file_platform_options_v2_annotations, 25);

/**
 * Describes the message platform.options.v2.RoutineOptions.
 * Use `create(RoutineOptionsSchema)` to create a new message.
 */
export const RoutineOptionsSchema = /*@__PURE__*/
  messageDesc(file_platform_options_v2_annotations, 26);

/**
 * Describes the enum platform.options.v2.NetworkType.
 */
export const NetworkTypeSchema = /*@__PURE__*/
  enumDesc(file_platform_options_v2_annotations, 0);

/**
 * @generated from enum platform.options.v2.NetworkType
 */
export const NetworkType = /*@__PURE__*/
  tsEnum(NetworkTypeSchema);

/**
 * Describes the enum platform.options.v2.ApiType.
 */
export const ApiTypeSchema = /*@__PURE__*/
  enumDesc(file_platform_options_v2_annotations, 1);

/**
 * @generated from enum platform.options.v2.ApiType
 */
export const ApiType = /*@__PURE__*/
  tsEnum(ApiTypeSchema);

/**
 * Describes the enum platform.options.v2.ApiInterfaceType.
 */
export const ApiInterfaceTypeSchema = /*@__PURE__*/
  enumDesc(file_platform_options_v2_annotations, 2);

/**
 * @generated from enum platform.options.v2.ApiInterfaceType
 */
export const ApiInterfaceType = /*@__PURE__*/
  tsEnum(ApiInterfaceTypeSchema);

/**
 * Describes the enum platform.options.v2.ApiLifecycle.
 */
export const ApiLifecycleSchema = /*@__PURE__*/
  enumDesc(file_platform_options_v2_annotations, 3);

/**
 * @generated from enum platform.options.v2.ApiLifecycle
 */
export const ApiLifecycle = /*@__PURE__*/
  tsEnum(ApiLifecycleSchema);

/**
 * Describes the enum platform.options.v2.EntityType.
 */
export const EntityTypeSchema = /*@__PURE__*/
  enumDesc(file_platform_options_v2_annotations, 4);

/**
 * @generated from enum platform.options.v2.EntityType
 */
export const EntityType = /*@__PURE__*/
  tsEnum(EntityTypeSchema);

/**
 * Describes the enum platform.options.v2.EntityConsistency.
 */
export const EntityConsistencySchema = /*@__PURE__*/
  enumDesc(file_platform_options_v2_annotations, 5);

/**
 * @generated from enum platform.options.v2.EntityConsistency
 */
export const EntityConsistency = /*@__PURE__*/
  tsEnum(EntityConsistencySchema);

/**
 * Describes the enum platform.options.v2.EntityHierarchy.
 */
export const EntityHierarchySchema = /*@__PURE__*/
  enumDesc(file_platform_options_v2_annotations, 6);

/**
 * @generated from enum platform.options.v2.EntityHierarchy
 */
export const EntityHierarchy = /*@__PURE__*/
  tsEnum(EntityHierarchySchema);

/**
 * Describes the enum platform.options.v2.LanguageType.
 */
export const LanguageTypeSchema = /*@__PURE__*/
  enumDesc(file_platform_options_v2_annotations, 7);

/**
 * @generated from enum platform.options.v2.LanguageType
 */
export const LanguageType = /*@__PURE__*/
  tsEnum(LanguageTypeSchema);

/**
 * Describes the enum platform.options.v2.ConnectorType.
 */
export const ConnectorTypeSchema = /*@__PURE__*/
  enumDesc(file_platform_options_v2_annotations, 8);

/**
 * @generated from enum platform.options.v2.ConnectorType
 */
export const ConnectorType = /*@__PURE__*/
  tsEnum(ConnectorTypeSchema);

/**
 * Describes the enum platform.options.v2.AuthRoleType.
 */
export const AuthRoleTypeSchema = /*@__PURE__*/
  enumDesc(file_platform_options_v2_annotations, 9);

/**
 * @generated from enum platform.options.v2.AuthRoleType
 */
export const AuthRoleType = /*@__PURE__*/
  tsEnum(AuthRoleTypeSchema);

/**
 * Describes the enum platform.options.v2.CQRSType.
 */
export const CQRSTypeSchema = /*@__PURE__*/
  enumDesc(file_platform_options_v2_annotations, 10);

/**
 * @generated from enum platform.options.v2.CQRSType
 */
export const CQRSType = /*@__PURE__*/
  tsEnum(CQRSTypeSchema);

/**
 * Describes the enum platform.options.v2.AuthRole.
 */
export const AuthRoleSchema = /*@__PURE__*/
  enumDesc(file_platform_options_v2_annotations, 11);

/**
 * @generated from enum platform.options.v2.AuthRole
 */
export const AuthRole = /*@__PURE__*/
  tsEnum(AuthRoleSchema);

/**
 * Describes the enum platform.options.v2.GraphType.
 */
export const GraphTypeSchema = /*@__PURE__*/
  enumDesc(file_platform_options_v2_annotations, 12);

/**
 * @generated from enum platform.options.v2.GraphType
 */
export const GraphType = /*@__PURE__*/
  tsEnum(GraphTypeSchema);

/**
 * Describes the enum platform.options.v2.FieldBehavior.
 */
export const FieldBehaviorSchema = /*@__PURE__*/
  enumDesc(file_platform_options_v2_annotations, 13);

/**
 * An indicator of the behavior of a given field (for example, that a field
 * is required in requests, or given as output but ignored as input).
 * This **does not** change the behavior in protocol buffers itself; it only
 * denotes the behavior and may affect how API tooling handles the field.
 *
 * @generated from enum platform.options.v2.FieldBehavior
 */
export const FieldBehavior = /*@__PURE__*/
  tsEnum(FieldBehaviorSchema);

/**
 * Describes the enum platform.options.v2.SyntheticType.
 */
export const SyntheticTypeSchema = /*@__PURE__*/
  enumDesc(file_platform_options_v2_annotations, 14);

/**
 * @generated from enum platform.options.v2.SyntheticType
 */
export const SyntheticType = /*@__PURE__*/
  tsEnum(SyntheticTypeSchema);

/**
 * Describes the enum platform.options.v2.ClassificationType.
 */
export const ClassificationTypeSchema = /*@__PURE__*/
  enumDesc(file_platform_options_v2_annotations, 15);

/**
 * Supported workspace type
 *
 * @generated from enum platform.options.v2.ClassificationType
 */
export const ClassificationType = /*@__PURE__*/
  tsEnum(ClassificationTypeSchema);

/**
 * Describes the enum platform.options.v2.SpecEnumType.
 */
export const SpecEnumTypeSchema = /*@__PURE__*/
  enumDesc(file_platform_options_v2_annotations, 16);

/**
 * @generated from enum platform.options.v2.SpecEnumType
 */
export const SpecEnumType = /*@__PURE__*/
  tsEnum(SpecEnumTypeSchema);

/**
 * Describes the enum platform.options.v2.EventScope.
 */
export const EventScopeSchema = /*@__PURE__*/
  enumDesc(file_platform_options_v2_annotations, 17);

/**
 * @generated from enum platform.options.v2.EventScope
 */
export const EventScope = /*@__PURE__*/
  tsEnum(EventScopeSchema);

/**
 * @generated from extension: platform.options.v2.NetworkOptions network_file = 40000;
 */
export const network_file = /*@__PURE__*/
  extDesc(file_platform_options_v2_annotations, 0);

/**
 * @generated from extension: platform.options.v2.ApiOptions api_file = 50000;
 */
export const api_file = /*@__PURE__*/
  extDesc(file_platform_options_v2_annotations, 1);

/**
 * @generated from extension: platform.options.v2.EntityOptions entity = 50001;
 */
export const entity = /*@__PURE__*/
  extDesc(file_platform_options_v2_annotations, 2);

/**
 * @generated from extension: platform.options.v2.LanguageOptions language = 50002;
 */
export const language = /*@__PURE__*/
  extDesc(file_platform_options_v2_annotations, 3);

/**
 * @generated from extension: platform.options.v2.GraphqlOptions graphql = 50003;
 */
export const graphql = /*@__PURE__*/
  extDesc(file_platform_options_v2_annotations, 4);

/**
 * @generated from extension: platform.options.v2.ConfigurationOptions configuration = 50004;
 */
export const configuration = /*@__PURE__*/
  extDesc(file_platform_options_v2_annotations, 5);

/**
 * @generated from extension: bool has_multiple_implementations = 50005;
 */
export const has_multiple_implementations = /*@__PURE__*/
  extDesc(file_platform_options_v2_annotations, 6);

/**
 * @generated from extension: platform.options.v2.ApiOptions api_service = 50010;
 */
export const api_service = /*@__PURE__*/
  extDesc(file_platform_options_v2_annotations, 7);

/**
 * @generated from extension: platform.options.v2.SpecServiceOptions service = 50011;
 */
export const service = /*@__PURE__*/
  extDesc(file_platform_options_v2_annotations, 8);

/**
 * @generated from extension: platform.options.v2.ProxyOptions proxy = 50012;
 */
export const proxy = /*@__PURE__*/
  extDesc(file_platform_options_v2_annotations, 9);

/**
 * @generated from extension: platform.options.v2.ConnectorOptions connector = 50013;
 */
export const connector = /*@__PURE__*/
  extDesc(file_platform_options_v2_annotations, 10);

/**
 * @generated from extension: platform.options.v2.ApiOptions api_method = 50020;
 */
export const api_method = /*@__PURE__*/
  extDesc(file_platform_options_v2_annotations, 11);

/**
 * @generated from extension: platform.options.v2.CQRSOptions cqrs = 50021;
 */
export const cqrs = /*@__PURE__*/
  extDesc(file_platform_options_v2_annotations, 12);

/**
 * @generated from extension: platform.options.v2.PermissionOptions permission = 50022;
 */
export const permission = /*@__PURE__*/
  extDesc(file_platform_options_v2_annotations, 13);

/**
 * @generated from extension: platform.options.v2.RateLimitOptions rate = 50023;
 */
export const rate = /*@__PURE__*/
  extDesc(file_platform_options_v2_annotations, 14);

/**
 * @generated from extension: platform.options.v2.ApiOptions api_message = 50030;
 */
export const api_message = /*@__PURE__*/
  extDesc(file_platform_options_v2_annotations, 15);

/**
 * @generated from extension: platform.options.v2.GraphOptions graph = 50031;
 */
export const graph = /*@__PURE__*/
  extDesc(file_platform_options_v2_annotations, 16);

/**
 * @generated from extension: platform.options.v2.RoutineOptions routine = 50032;
 */
export const routine = /*@__PURE__*/
  extDesc(file_platform_options_v2_annotations, 17);

/**
 * @generated from extension: platform.options.v2.ApiOptions api_field = 50040;
 */
export const api_field = /*@__PURE__*/
  extDesc(file_platform_options_v2_annotations, 18);

/**
 * @generated from extension: platform.options.v2.EntityFieldOptions entity_field = 50041;
 */
export const entity_field = /*@__PURE__*/
  extDesc(file_platform_options_v2_annotations, 19);

/**
 * @generated from extension: platform.options.v2.ConfigurationFieldOptions configuration_field = 50042;
 */
export const configuration_field = /*@__PURE__*/
  extDesc(file_platform_options_v2_annotations, 20);

/**
 * @generated from extension: platform.options.v2.SyntheticOptions synthetic = 50043;
 */
export const synthetic = /*@__PURE__*/
  extDesc(file_platform_options_v2_annotations, 21);

/**
 * @generated from extension: platform.options.v2.SpecOptions spec = 50050;
 */
export const spec = /*@__PURE__*/
  extDesc(file_platform_options_v2_annotations, 22);

/**
 * @generated from extension: platform.options.v2.BillingOptions billing = 50060;
 */
export const billing = /*@__PURE__*/
  extDesc(file_platform_options_v2_annotations, 23);

/**
 * @generated from extension: platform.options.v2.EventScopeOptions event_scope = 50061;
 */
export const event_scope = /*@__PURE__*/
  extDesc(file_platform_options_v2_annotations, 24);

/**
 * @generated from extension: platform.options.v2.AuthRoleOptions auth_role = 50062;
 */
export const auth_role = /*@__PURE__*/
  extDesc(file_platform_options_v2_annotations, 25);

/**
 * @generated from extension: platform.options.v2.EventOptions event = 50063;
 */
export const event = /*@__PURE__*/
  extDesc(file_platform_options_v2_annotations, 26);
