// @generated by protoc-gen-es v2.2.3
// @generated from file platform/options/v2/annotations.proto (package platform.options.v2, syntax proto3)
/* eslint-disable */

import { enumDesc, extDesc, fileDesc, messageDesc, tsEnum } from "@bufbuild/protobuf/codegenv1";
import { file_google_protobuf_descriptor } from "@bufbuild/protobuf/wkt";

/**
 * Describes the file platform/options/v2/annotations.proto.
 */
export const file_platform_options_v2_annotations = /*@__PURE__*/
  fileDesc("CiVwbGF0Zm9ybS9vcHRpb25zL3YyL2Fubm90YXRpb25zLnByb3RvEhNwbGF0Zm9ybS5vcHRpb25zLnYyIkAKDk5ldHdvcmtPcHRpb25zEi4KBHR5cGUYASABKA4yIC5wbGF0Zm9ybS5vcHRpb25zLnYyLk5ldHdvcmtUeXBlIuoBCgpBcGlPcHRpb25zEioKBHR5cGUYASABKA4yHC5wbGF0Zm9ybS5vcHRpb25zLnYyLkFwaVR5cGUSMAoFY3ljbGUYAiABKA4yIS5wbGF0Zm9ybS5vcHRpb25zLnYyLkFwaUxpZmVjeWNsZRI4CglpbnRlcmZhY2UYAyABKA4yJS5wbGF0Zm9ybS5vcHRpb25zLnYyLkFwaUludGVyZmFjZVR5cGUSMQoHbmV0d29yaxgEIAEoDjIgLnBsYXRmb3JtLm9wdGlvbnMudjIuTmV0d29ya1R5cGUSEQoJc2hvcnRuYW1lGAUgASgJIoMCCg1FbnRpdHlPcHRpb25zEg4KBmVudGl0eRgBIAEoCRIVCg1lbnRpdHlfcGx1cmFsGAIgASgJEi0KBHR5cGUYAyABKA4yHy5wbGF0Zm9ybS5vcHRpb25zLnYyLkVudGl0eVR5cGUSEQoJbmFtZXNwYWNlGAQgASgJEjsKC2NvbnNpc3RlbmN5GAUgASgOMiYucGxhdGZvcm0ub3B0aW9ucy52Mi5FbnRpdHlDb25zaXN0ZW5jeRI3CgloaWVyYXJjaHkYBiABKA4yJC5wbGF0Zm9ybS5vcHRpb25zLnYyLkVudGl0eUhpZXJhcmNoeRITCgt2ZXJzaW9uYWJsZRgHIAEoCCJHCg9MYW5ndWFnZU9wdGlvbnMSNAoJbGFuZ3VhZ2VzGAEgAygOMiEucGxhdGZvcm0ub3B0aW9ucy52Mi5MYW5ndWFnZVR5cGUiIQoOR3JhcGhxbE9wdGlvbnMSDwoHZW5hYmxlZBgBIAEoCCI6ChJTcGVjU2VydmljZU9wdGlvbnMSEQoJZ3JwY19wb3J0GAEgASgFEhEKCWh0dHBfcG9ydBgCIAEoBSIyCg9SYXRlTGltaXRGaWx0ZXISDwoHZW5hYmxlZBgBIAEoCBIOCgZtZXRyaWMYAiABKAkiJgoTQXV0aG9yaXphdGlvbkZpbHRlchIPCgdlbmFibGVkGAEgASgIIiAKDUNvbnNlbnRGaWx0ZXISDwoHZW5hYmxlZBgBIAEoCCK+AQoMUHJveHlPcHRpb25zEj8KDWF1dGhvcml6YXRpb24YASABKAsyKC5wbGF0Zm9ybS5vcHRpb25zLnYyLkF1dGhvcml6YXRpb25GaWx0ZXISMwoHY29uc2VudBgCIAEoCzIiLnBsYXRmb3JtLm9wdGlvbnMudjIuQ29uc2VudEZpbHRlchI4CgpyYXRlX2xpbWl0GAMgASgLMiQucGxhdGZvcm0ub3B0aW9ucy52Mi5SYXRlTGltaXRGaWx0ZXIiRAoQQ29ubmVjdG9yT3B0aW9ucxIwCgR0eXBlGAEgASgOMiIucGxhdGZvcm0ub3B0aW9ucy52Mi5Db25uZWN0b3JUeXBlIjoKC0NRUlNPcHRpb25zEisKBHR5cGUYASABKA4yHS5wbGF0Zm9ybS5vcHRpb25zLnYyLkNRUlNUeXBlIiEKEFJhdGVMaW1pdE9wdGlvbnMSDQoFbGltaXQYASABKAgiVQoRUGVybWlzc2lvbk9wdGlvbnMSEgoKcGVybWlzc2lvbhgBIAEoCRIsCgVyb2xlcxgCIAMoDjIdLnBsYXRmb3JtLm9wdGlvbnMudjIuQXV0aFJvbGUiPAoMR3JhcGhPcHRpb25zEiwKBHR5cGUYASABKA4yHi5wbGF0Zm9ybS5vcHRpb25zLnYyLkdyYXBoVHlwZSJsChJFbnRpdHlGaWVsZE9wdGlvbnMSCwoDa2V5GAEgASgIEjQKCGJlaGF2aW9yGAIgAygOMiIucGxhdGZvcm0ub3B0aW9ucy52Mi5GaWVsZEJlaGF2aW9yEhMKC3ZlcnNpb25fa2V5GAMgASgIIpUDChFTcGVjQ29uZmlndXJhdGlvbhILCgNrZXkYASABKAkSSgoEdHlwZRgCIAEoDjI8LnBsYXRmb3JtLm9wdGlvbnMudjIuU3BlY0NvbmZpZ3VyYXRpb24uU3BlY0NvbmZpZ3VyYXRpb25UeXBlEhMKC2Rlc2NyaXB0aW9uGAMgASgJEhQKDG92ZXJyaWRlYWJsZRgEIAEoCBIPCgdlbmFibGVkGAUgASgIIuoBChVTcGVjQ29uZmlndXJhdGlvblR5cGUSJwojU1BFQ19DT05GSUdVUkFUSU9OX1RZUEVfVU5TUEVDSUZJRUQQABIiCh5TUEVDX0NPTkZJR1VSQVRJT05fVFlQRV9TVFJJTkcQARIgChxTUEVDX0NPTkZJR1VSQVRJT05fVFlQRV9CT09MEAISHwobU1BFQ19DT05GSUdVUkFUSU9OX1RZUEVfSU5UEAMSIAocU1BFQ19DT05GSUdVUkFUSU9OX1RZUEVfTElTVBAEEh8KG1NQRUNfQ09ORklHVVJBVElPTl9UWVBFX01BUBAFIj0KFENvbmZpZ3VyYXRpb25PcHRpb25zEg8KB2VuYWJsZWQYASABKAgSFAoMZmllbGRfbnVtYmVyGAIgASgFIloKGUNvbmZpZ3VyYXRpb25GaWVsZE9wdGlvbnMSPQoNY29uZmlndXJhdGlvbhgBIAEoCzImLnBsYXRmb3JtLm9wdGlvbnMudjIuU3BlY0NvbmZpZ3VyYXRpb24iXAoQU3ludGhldGljT3B0aW9ucxIWCg5kaWN0aW9uYXJ5X2tleRgBIAEoCRIwCgR0eXBlGAIgASgOMiIucGxhdGZvcm0ub3B0aW9ucy52Mi5TeW50aGV0aWNUeXBlIkcKDkNsYXNzaWZpY2F0aW9uEjUKBHR5cGUYASABKA4yJy5wbGF0Zm9ybS5vcHRpb25zLnYyLkNsYXNzaWZpY2F0aW9uVHlwZSI+CgtTcGVjT3B0aW9ucxIvCgR0eXBlGAEgASgOMiEucGxhdGZvcm0ub3B0aW9ucy52Mi5TcGVjRW51bVR5cGUiTQoOQmlsbGluZ09wdGlvbnMSEAoIYmlsbGFibGUYASABKAgSGAoQcGFydG5lcl9iaWxsYWJsZRgCIAEoCBIPCgdtZXRlcmVkGAMgASgIIkQKEUV2ZW50U2NvcGVPcHRpb25zEi8KBnNjb3BlcxgBIAMoDjIfLnBsYXRmb3JtLm9wdGlvbnMudjIuRXZlbnRTY29wZSIjCgxFdmVudE9wdGlvbnMSEwoLdmVyc2lvbmFibGUYASABKAgiRwoPQXV0aFJvbGVPcHRpb25zEjQKCXJvbGVfdHlwZRgBIAEoDjIhLnBsYXRmb3JtLm9wdGlvbnMudjIuQXV0aFJvbGVUeXBlIiIKDlJvdXRpbmVPcHRpb25zEhAKCGxpc3RlbmVyGAEgASgJKnsKC05ldHdvcmtUeXBlEhwKGE5FVFdPUktfVFlQRV9VTlNQRUNJRklFRBAAEhkKFU5FVFdPUktfVFlQRV9VTkRFUkxBWRABEhkKFU5FVFdPUktfVFlQRV9JTlRFUk5FVBACEhgKFE5FVFdPUktfVFlQRV9PVkVSTEFZEAMqdgoHQXBpVHlwZRIYChRBUElfVFlQRV9VTlNQRUNJRklFRBAAEhQKEEFQSV9UWVBFX1BSSVZBVEUQARIUChBBUElfVFlQRV9QQVJUTkVSEAISEwoPQVBJX1RZUEVfUFVCTElDEAMSEAoMQVBJX1RZUEVfUE9DEAQqhwIKEEFwaUludGVyZmFjZVR5cGUSIgoeQVBJX0lOVEVSRkFDRV9UWVBFX1VOU1BFQ0lGSUVEEAASGwoXQVBJX0lOVEVSRkFDRV9UWVBFX01FVEEQARIiCh5BUElfSU5URVJGQUNFX1RZUEVfT1BFUkFUSU9OQUwQAhIjCh9BUElfSU5URVJGQUNFX1RZUEVfQ09OVFJJQlVUSU9OEAMSIQodQVBJX0lOVEVSRkFDRV9UWVBFX0FOQUxZVElDQUwQBBIhCh1BUElfSU5URVJGQUNFX1RZUEVfT0JTRVJWQUJMRRAFEiMKH0FQSV9JTlRFUkZBQ0VfVFlQRV9ESVNDT1ZFUkFCTEUQBirnAQoMQXBpTGlmZWN5Y2xlEh0KGUFQSV9MSUZFQ1lDTEVfVU5TUEVDSUZJRUQQABIXChNBUElfTElGRUNZQ0xFX0FMUEhBEAESFgoSQVBJX0xJRkVDWUNMRV9CRVRBEAISJgoiQVBJX0xJRkVDWUNMRV9MSU1JVEVEX0FWQUlMQUJJTElUWRADEiYKIkFQSV9MSUZFQ1lDTEVfR0VORVJBTF9BVkFJTEFCSUxJVFkQBBIcChhBUElfTElGRUNZQ0xFX0RFUFJFQ0FURUQQBRIZChVBUElfTElGRUNZQ0xFX1JFVElSRUQQBiraAQoKRW50aXR5VHlwZRIbChdFTlRJVFlfVFlQRV9VTlNQRUNJRklFRBAAEhkKFUVOVElUWV9UWVBFX0FFUk9TUElLRRABEhYKEkVOVElUWV9UWVBFX0RHUkFQSBACEhcKE0VOVElUWV9UWVBFX01PTkdPREIQAxIYChRFTlRJVFlfVFlQRV9CSUdRVUVSWRAEEhUKEUVOVElUWV9UWVBFX1JFRElTEAUSFwoTRU5USVRZX1RZUEVfUk9DS1NEQhAGEhkKFUVOVElUWV9UWVBFX0NPVUNIQkFTRRAHKncKEUVudGl0eUNvbnNpc3RlbmN5EiIKHkVOVElUWV9DT05TSVNURU5DWV9VTlNQRUNJRklFRBAAEh0KGUVOVElUWV9DT05TSVNURU5DWV9TVFJPTkcQARIfChtFTlRJVFlfQ09OU0lTVEVOQ1lfRVZFTlRVQUwQAiqVAQoPRW50aXR5SGllcmFyY2h5EiAKHEVOVElUWV9ISUVSQVJDSFlfVU5TUEVDSUZJRUQQABIdChlFTlRJVFlfSElFUkFSQ0hZX1BMQVRGT1JNEAESIQodRU5USVRZX0hJRVJBUkNIWV9PUkdBTklaQVRJT04QAhIeChpFTlRJVFlfSElFUkFSQ0hZX1dPUktTUEFDRRADKswCCgxMYW5ndWFnZVR5cGUSHQoZTEFOR1VBR0VfVFlQRV9VTlNQRUNJRklFRBAAEhsKF0xBTkdVQUdFX1RZUEVfQ1BMVVNQTFVTEAESFgoSTEFOR1VBR0VfVFlQRV9SVVNUEAISGAoUTEFOR1VBR0VfVFlQRV9HT0xBTkcQAxIWChJMQU5HVUFHRV9UWVBFX0pBVkEQBBIYChRMQU5HVUFHRV9UWVBFX1BZVEhPThAFEhwKGExBTkdVQUdFX1RZUEVfVFlQRVNDUklQVBAGEhgKFExBTkdVQUdFX1RZUEVfQ1NIQVJQEAcSFwoTTEFOR1VBR0VfVFlQRV9TV0lGVBAIEhkKFUxBTkdVQUdFX1RZUEVfQU5EUk9JRBAJEhkKFUxBTkdVQUdFX1RZUEVfR1JBUEhRTBAKEhUKEUxBTkdVQUdFX1RZUEVfTFVBEAsqTQoNQ29ubmVjdG9yVHlwZRIeChpDT05ORUNUT1JfVFlQRV9VTlNQRUNJRklFRBAAEhwKGENPTk5FQ1RPUl9UWVBFX1JFRkVSRU5DRRABKq0BCgxBdXRoUm9sZVR5cGUSHgoaQVVUSF9ST0xFX1RZUEVfVU5TUEVDSUZJRUQQABIbChdBVVRIX1JPTEVfVFlQRV9QTEFURk9STRABEh8KG0FVVEhfUk9MRV9UWVBFX09SR0FOSVpBVElPThACEhwKGEFVVEhfUk9MRV9UWVBFX1dPUktTUEFDRRADEiEKHUFVVEhfUk9MRV9UWVBFX0NPTk5FQ1RFRF9URVNUEAQq+gMKCENRUlNUeXBlEhkKFUNRUlNfVFlQRV9VTlNQRUNJRklFRBAAEhIKDkNRUlNfVFlQRV9OT05FEAESHQoZQ1FSU19UWVBFX01VVEFUSU9OX0NSRUFURRACEh0KGUNRUlNfVFlQRV9NVVRBVElPTl9VUERBVEUQAxIdChlDUVJTX1RZUEVfTVVUQVRJT05fREVMRVRFEAQSJAogQ1FSU19UWVBFX01VVEFUSU9OX0NMSUVOVF9TVFJFQU0QBRIkCiBDUVJTX1RZUEVfTVVUQVRJT05fU0VSVkVSX1NUUkVBTRAGEiIKHkNRUlNfVFlQRV9NVVRBVElPTl9CSURJX1NUUkVBTRAHEhgKFENRUlNfVFlQRV9RVUVSWV9MSVNUEAgSGgoWQ1FSU19UWVBFX1FVRVJZX1NUUkVBTRAJEhcKE0NRUlNfVFlQRV9RVUVSWV9HRVQQChIgChxDUVJTX1RZUEVfUVVFUllfRVZFTlRfU1RSRUFNEAsSIQodQ1FSU19UWVBFX1FVRVJZX0NMSUVOVF9TVFJFQU0QDBIhCh1DUVJTX1RZUEVfUVVFUllfU0VSVkVSX1NUUkVBTRANEh8KG0NRUlNfVFlQRV9RVUVSWV9CSURJX1NUUkVBTRAOEhoKFkNRUlNfVFlQRV9RVUVSWV9FWElTVFMQDyrVCAoIQXV0aFJvbGUSGQoVQVVUSF9ST0xFX1VOU1BFQ0lGSUVEEAASKgoeQVVUSF9ST0xFX1BMQVRGT1JNX1NVUEVSX0FETUlOEGQaBvK4GAIIARItCiFBVVRIX1JPTEVfUExBVEZPUk1fQ0xJTklDQUxfQURNSU4QZRoG8rgYAggBEiwKIEFVVEhfUk9MRV9QTEFURk9STV9CSUxMSU5HX0FETUlOEGYaBvK4GAIIARIkChhBVVRIX1JPTEVfUExBVEZPUk1fQURNSU4QZxoG8rgYAggBEiYKGkFVVEhfUk9MRV9QTEFURk9STV9NQU5BR0VSEGgaBvK4GAIIARIjChdBVVRIX1JPTEVfUExBVEZPUk1fVVNFUhBpGgbyuBgCCAESJQoZQVVUSF9ST0xFX1BMQVRGT1JNX1ZJRVdFUhBqGgbyuBgCCAESLwoiQVVUSF9ST0xFX09SR0FOSVpBVElPTl9TVVBFUl9BRE1JThDIARoG8rgYAggCEjIKJUFVVEhfUk9MRV9PUkdBTklaQVRJT05fQ0xJTklDQUxfQURNSU4QyQEaBvK4GAIIAhIxCiRBVVRIX1JPTEVfT1JHQU5JWkFUSU9OX0JJTExJTkdfQURNSU4QygEaBvK4GAIIAhIpChxBVVRIX1JPTEVfT1JHQU5JWkFUSU9OX0FETUlOEMsBGgbyuBgCCAISKwoeQVVUSF9ST0xFX09SR0FOSVpBVElPTl9NQU5BR0VSEMwBGgbyuBgCCAISKAobQVVUSF9ST0xFX09SR0FOSVpBVElPTl9VU0VSEM0BGgbyuBgCCAISKgodQVVUSF9ST0xFX09SR0FOSVpBVElPTl9WSUVXRVIQzgEaBvK4GAIIAhIsCh9BVVRIX1JPTEVfV09SS1NQQUNFX1NVUEVSX0FETUlOEKwCGgbyuBgCCAMSLwoiQVVUSF9ST0xFX1dPUktTUEFDRV9DTElOSUNBTF9BRE1JThCtAhoG8rgYAggDEi4KIUFVVEhfUk9MRV9XT1JLU1BBQ0VfQklMTElOR19BRE1JThCuAhoG8rgYAggDEiYKGUFVVEhfUk9MRV9XT1JLU1BBQ0VfQURNSU4QrwIaBvK4GAIIAxIoChtBVVRIX1JPTEVfV09SS1NQQUNFX01BTkFHRVIQsAIaBvK4GAIIAxIlChhBVVRIX1JPTEVfV09SS1NQQUNFX1VTRVIQsQIaBvK4GAIIAxInChpBVVRIX1JPTEVfV09SS1NQQUNFX1ZJRVdFUhCyAhoG8rgYAggDEiQKIEFVVEhfUk9MRV9DT05ORUNURURfVEVTVF9QQVRJRU5UEA8SJQohQVVUSF9ST0xFX0NPTk5FQ1RFRF9URVNUX1BST1ZJREVSEBASIgoeQVVUSF9ST0xFX0NPTk5FQ1RFRF9URVNUX1BST1hZEBESIwofQVVUSF9ST0xFX0NPTk5FQ1RFRF9URVNUX1ZJRVdFUhASKlQKCUdyYXBoVHlwZRIaChZHUkFQSF9UWVBFX1VOU1BFQ0lGSUVEEAASFAoQR1JBUEhfVFlQRV9JTlBVVBABEhUKEUdSQVBIX1RZUEVfT1VUUFVUEAIqjwIKDUZpZWxkQmVoYXZpb3ISHgoaRklFTERfQkVIQVZJT1JfVU5TUEVDSUZJRUQQABIbChdGSUVMRF9CRUhBVklPUl9PUFRJT05BTBABEhsKF0ZJRUxEX0JFSEFWSU9SX1JFUVVJUkVEEAISHgoaRklFTERfQkVIQVZJT1JfT1VUUFVUX09OTFkQAxIdChlGSUVMRF9CRUhBVklPUl9JTlBVVF9PTkxZEAQSHAoYRklFTERfQkVIQVZJT1JfSU1NVVRBQkxFEAUSIQodRklFTERfQkVIQVZJT1JfVU5PUkRFUkVEX0xJU1QQBhIkCiBGSUVMRF9CRUhBVklPUl9OT05fRU1QVFlfREVGQVVMVBAHKswBCg1TeW50aGV0aWNUeXBlEh4KGlNZTlRIRVRJQ19UWVBFX1VOU1BFQ0lGSUVEEAASKQolU1lOVEhFVElDX1RZUEVfRElSRUNUX0ZST01fRElDVElPTkFSWRABEioKJlNZTlRIRVRJQ19UWVBFX1NFTEVDVF9SQU5ET01fRlJPTV9MSVNUEAISIQodU1lOVEhFVElDX1RZUEVfTElTVF9GUk9NX0xJU1QQAxIhCh1TWU5USEVUSUNfVFlQRV9HRU5FUkFURURfTE9HTxAGKsQCChJDbGFzc2lmaWNhdGlvblR5cGUSIwofQ0xBU1NJRklDQVRJT05fVFlQRV9VTlNQRUNJRklFRBAAEicKI0NMQVNTSUZJQ0FUSU9OX1RZUEVfREVSSVZBVElWRV9EQVRBEAESJQohQ0xBU1NJRklDQVRJT05fVFlQRV9ERV9JREVOVElGSUVEEAISHgoaQ0xBU1NJRklDQVRJT05fVFlQRV9QVUJMSUMQAxIkCiBDTEFTU0lGSUNBVElPTl9UWVBFX0lOVEVSTkFMX1VTRRAEEiQKIENMQVNTSUZJQ0FUSU9OX1RZUEVfQ09ORklERU5USUFMEAUSIgoeQ0xBU1NJRklDQVRJT05fVFlQRV9SRVNUUklDVEVEEAYSKQolQ0xBU1NJRklDQVRJT05fVFlQRV9ISUdITFlfUkVTVFJJQ1RFRBAHKsABCgxTcGVjRW51bVR5cGUSHgoaU1BFQ19FTlVNX1RZUEVfVU5TUEVDSUZJRUQQABIXChNTUEVDX0VOVU1fVFlQRV9OT05FEAESGQoVU1BFQ19FTlVNX1RZUEVfVE9QSUNTEAISGwoXU1BFQ19FTlVNX1RZUEVfQ09NTUFORFMQAxIZChVTUEVDX0VOVU1fVFlQRV9FVkVOVFMQBBIkCiBTUEVDX0VOVU1fVFlQRV9ST1VUSU5FX0xJU1RFTkVSUxAFKngKCkV2ZW50U2NvcGUSGwoXRVZFTlRfU0NPUEVfVU5TUEVDSUZJRUQQABIUChBFVkVOVF9TQ09QRV9VU0VSEAESGQoVRVZFTlRfU0NPUEVfV09SS1NQQUNFEAISHAoYRVZFTlRfU0NPUEVfT1JHQU5JWkFUSU9OEAM6ZgoMbmV0d29ya19maWxlEhwuZ29vZ2xlLnByb3RvYnVmLkZpbGVPcHRpb25zGMC4AiABKAsyIy5wbGF0Zm9ybS5vcHRpb25zLnYyLk5ldHdvcmtPcHRpb25zUgtuZXR3b3JrRmlsZTpaCghhcGlfZmlsZRIcLmdvb2dsZS5wcm90b2J1Zi5GaWxlT3B0aW9ucxjQhgMgASgLMh8ucGxhdGZvcm0ub3B0aW9ucy52Mi5BcGlPcHRpb25zUgdhcGlGaWxlOloKBmVudGl0eRIcLmdvb2dsZS5wcm90b2J1Zi5GaWxlT3B0aW9ucxjRhgMgASgLMiIucGxhdGZvcm0ub3B0aW9ucy52Mi5FbnRpdHlPcHRpb25zUgZlbnRpdHk6YAoIbGFuZ3VhZ2USHC5nb29nbGUucHJvdG9idWYuRmlsZU9wdGlvbnMY0oYDIAEoCzIkLnBsYXRmb3JtLm9wdGlvbnMudjIuTGFuZ3VhZ2VPcHRpb25zUghsYW5ndWFnZTpdCgdncmFwaHFsEhwuZ29vZ2xlLnByb3RvYnVmLkZpbGVPcHRpb25zGNOGAyABKAsyIy5wbGF0Zm9ybS5vcHRpb25zLnYyLkdyYXBocWxPcHRpb25zUgdncmFwaHFsOm8KDWNvbmZpZ3VyYXRpb24SHC5nb29nbGUucHJvdG9idWYuRmlsZU9wdGlvbnMY1IYDIAEoCzIpLnBsYXRmb3JtLm9wdGlvbnMudjIuQ29uZmlndXJhdGlvbk9wdGlvbnNSDWNvbmZpZ3VyYXRpb246YAocaGFzX211bHRpcGxlX2ltcGxlbWVudGF0aW9ucxIcLmdvb2dsZS5wcm90b2J1Zi5GaWxlT3B0aW9ucxjVhgMgASgIUhpoYXNNdWx0aXBsZUltcGxlbWVudGF0aW9uczpjCgthcGlfc2VydmljZRIfLmdvb2dsZS5wcm90b2J1Zi5TZXJ2aWNlT3B0aW9ucxjahgMgASgLMh8ucGxhdGZvcm0ub3B0aW9ucy52Mi5BcGlPcHRpb25zUgphcGlTZXJ2aWNlOmQKB3NlcnZpY2USHy5nb29nbGUucHJvdG9idWYuU2VydmljZU9wdGlvbnMY24YDIAEoCzInLnBsYXRmb3JtLm9wdGlvbnMudjIuU3BlY1NlcnZpY2VPcHRpb25zUgdzZXJ2aWNlOloKBXByb3h5Eh8uZ29vZ2xlLnByb3RvYnVmLlNlcnZpY2VPcHRpb25zGNyGAyABKAsyIS5wbGF0Zm9ybS5vcHRpb25zLnYyLlByb3h5T3B0aW9uc1IFcHJveHk6ZgoJY29ubmVjdG9yEh8uZ29vZ2xlLnByb3RvYnVmLlNlcnZpY2VPcHRpb25zGN2GAyABKAsyJS5wbGF0Zm9ybS5vcHRpb25zLnYyLkNvbm5lY3Rvck9wdGlvbnNSCWNvbm5lY3RvcjpgCgphcGlfbWV0aG9kEh4uZ29vZ2xlLnByb3RvYnVmLk1ldGhvZE9wdGlvbnMY5IYDIAEoCzIfLnBsYXRmb3JtLm9wdGlvbnMudjIuQXBpT3B0aW9uc1IJYXBpTWV0aG9kOlYKBGNxcnMSHi5nb29nbGUucHJvdG9idWYuTWV0aG9kT3B0aW9ucxjlhgMgASgLMiAucGxhdGZvcm0ub3B0aW9ucy52Mi5DUVJTT3B0aW9uc1IEY3FyczpoCgpwZXJtaXNzaW9uEh4uZ29vZ2xlLnByb3RvYnVmLk1ldGhvZE9wdGlvbnMY5oYDIAEoCzImLnBsYXRmb3JtLm9wdGlvbnMudjIuUGVybWlzc2lvbk9wdGlvbnNSCnBlcm1pc3Npb246WwoEcmF0ZRIeLmdvb2dsZS5wcm90b2J1Zi5NZXRob2RPcHRpb25zGOeGAyABKAsyJS5wbGF0Zm9ybS5vcHRpb25zLnYyLlJhdGVMaW1pdE9wdGlvbnNSBHJhdGU6YwoLYXBpX21lc3NhZ2USHy5nb29nbGUucHJvdG9idWYuTWVzc2FnZU9wdGlvbnMY7oYDIAEoCzIfLnBsYXRmb3JtLm9wdGlvbnMudjIuQXBpT3B0aW9uc1IKYXBpTWVzc2FnZTpaCgVncmFwaBIfLmdvb2dsZS5wcm90b2J1Zi5NZXNzYWdlT3B0aW9ucxjvhgMgASgLMiEucGxhdGZvcm0ub3B0aW9ucy52Mi5HcmFwaE9wdGlvbnNSBWdyYXBoOmAKB3JvdXRpbmUSHy5nb29nbGUucHJvdG9idWYuTWVzc2FnZU9wdGlvbnMY8IYDIAEoCzIjLnBsYXRmb3JtLm9wdGlvbnMudjIuUm91dGluZU9wdGlvbnNSB3JvdXRpbmU6XQoJYXBpX2ZpZWxkEh0uZ29vZ2xlLnByb3RvYnVmLkZpZWxkT3B0aW9ucxj4hgMgASgLMh8ucGxhdGZvcm0ub3B0aW9ucy52Mi5BcGlPcHRpb25zUghhcGlGaWVsZDprCgxlbnRpdHlfZmllbGQSHS5nb29nbGUucHJvdG9idWYuRmllbGRPcHRpb25zGPmGAyABKAsyJy5wbGF0Zm9ybS5vcHRpb25zLnYyLkVudGl0eUZpZWxkT3B0aW9uc1ILZW50aXR5RmllbGQ6gAEKE2NvbmZpZ3VyYXRpb25fZmllbGQSHS5nb29nbGUucHJvdG9idWYuRmllbGRPcHRpb25zGPqGAyABKAsyLi5wbGF0Zm9ybS5vcHRpb25zLnYyLkNvbmZpZ3VyYXRpb25GaWVsZE9wdGlvbnNSEmNvbmZpZ3VyYXRpb25GaWVsZDpkCglzeW50aGV0aWMSHS5nb29nbGUucHJvdG9idWYuRmllbGRPcHRpb25zGPuGAyABKAsyJS5wbGF0Zm9ybS5vcHRpb25zLnYyLlN5bnRoZXRpY09wdGlvbnNSCXN5bnRoZXRpYzpUCgRzcGVjEhwuZ29vZ2xlLnByb3RvYnVmLkVudW1PcHRpb25zGIKHAyABKAsyIC5wbGF0Zm9ybS5vcHRpb25zLnYyLlNwZWNPcHRpb25zUgRzcGVjOmIKB2JpbGxpbmcSIS5nb29nbGUucHJvdG9idWYuRW51bVZhbHVlT3B0aW9ucxiMhwMgASgLMiMucGxhdGZvcm0ub3B0aW9ucy52Mi5CaWxsaW5nT3B0aW9uc1IHYmlsbGluZzpsCgtldmVudF9zY29wZRIhLmdvb2dsZS5wcm90b2J1Zi5FbnVtVmFsdWVPcHRpb25zGI2HAyABKAsyJi5wbGF0Zm9ybS5vcHRpb25zLnYyLkV2ZW50U2NvcGVPcHRpb25zUgpldmVudFNjb3BlOmYKCWF1dGhfcm9sZRIhLmdvb2dsZS5wcm90b2J1Zi5FbnVtVmFsdWVPcHRpb25zGI6HAyABKAsyJC5wbGF0Zm9ybS5vcHRpb25zLnYyLkF1dGhSb2xlT3B0aW9uc1IIYXV0aFJvbGU6XAoFZXZlbnQSIS5nb29nbGUucHJvdG9idWYuRW51bVZhbHVlT3B0aW9ucxiPhwMgASgLMiEucGxhdGZvcm0ub3B0aW9ucy52Mi5FdmVudE9wdGlvbnNSBWV2ZW50QmJaYGdpdGh1Yi5jb20vb3BlbmVjb3N5c3RlbXMvZWNvc3lzdGVtL2xpYnMvcHJvdG9idWYvZ28vcHJvdG9idWYvZ2VuL3BsYXRmb3JtL29wdGlvbnMvdjI7b3B0aW9udjJwYmIGcHJvdG8z", [file_google_protobuf_descriptor]);

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

