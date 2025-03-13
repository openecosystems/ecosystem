package shared

import (
	options "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/options/v2"

	pgs "github.com/lyft/protoc-gen-star/v2"
)

// SyntheticOptions retrieves and returns the SyntheticOptions extension associated with the provided protobuf method.
// It panics if the extension cannot be read from the proto definition.
func (fns Functions) SyntheticOptions(method pgs.Method) options.SyntheticOptions {
	var synthetic options.SyntheticOptions

	_, err := method.Extension(options.E_Synthetic, &synthetic)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}

	return synthetic
}

// Synthetic extracts and returns a message specified by the dictionary key in the SyntheticOptions of a gRPC method.
func (fns Functions) Synthetic(method pgs.Method) pgs.Message {
	var synthetic options.SyntheticOptions

	_, err := method.Extension(options.E_Synthetic, &synthetic)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}

	if synthetic.GetDictionaryKey() != "" {
		for _, msg := range method.Input().AllMessages() {
			if msg.Name().String() == pgs.Name(synthetic.GetDictionaryKey()).UpperCamelCase().String() {
				return msg
			}
		}
	}

	return nil
}

// GetSyntheticType retrieves the synthetic type from the provided protobuf Field extension.
// Returns a string representation of the synthetic type, defaulting to "SYNTHETIC_TYPE_UNSPECIFIED".
// Panics if unable to read the extension data from the provided Field.
func (fns Functions) GetSyntheticType(file pgs.Field) string {
	var synthetic options.SyntheticOptions

	_, err := file.Extension(options.E_Synthetic, &synthetic)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}
	switch synthetic.Type {
	case options.SyntheticType_SYNTHETIC_TYPE_UNSPECIFIED:
		return "SYNTHETIC_TYPE_UNSPECIFIED"
	case options.SyntheticType_SYNTHETIC_TYPE_DIRECT_FROM_DICTIONARY:
		return "SYNTHETIC_TYPE_DIRECT_FROM_DICTIONARY"
	case options.SyntheticType_SYNTHETIC_TYPE_LIST_FROM_LIST:
		return "SYNTHETIC_TYPE_LIST_FROM_LIST"
	case options.SyntheticType_SYNTHETIC_TYPE_SELECT_RANDOM_FROM_LIST:
		return "SYNTHETIC_TYPE_SELECT_RANDOM_FROM_LIST"
	case options.SyntheticType_SYNTHETIC_TYPE_GENERATED_LOGO:
		return "SYNTHETIC_TYPE_GENERATED_LOGO"
	default:
		return "SYNTHETIC_TYPE_UNSPECIFIED"
	}
}

// GetSyntheticDictionaryKey returns the string representation of the dictionary key derived from the given field.
func (fns Functions) GetSyntheticDictionaryKey(file pgs.Field) string {
	return file.Name().String()
	/*
		var synthetic options.SyntheticOptions

		_, err := file.Extension(options.E_Synthetic, &synthetic)
		if err != nil {
			panic(err.Error() + "unable to read extension from proto")
		}

		return synthetic.GetDictionaryKey()
	*/
}

// GetSyntheticDictionaryKeyOld retrieves the synthetic dictionary key from the given protobuf field extension.
// It panics if there is an error reading the extension.
func (fns Functions) GetSyntheticDictionaryKeyOld(file pgs.Field) string {
	var synthetic options.SyntheticOptions

	_, err := file.Extension(options.E_Synthetic, &synthetic)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}

	return synthetic.GetDictionaryKey()
}

// SyntheticMethods extracts and returns a list of "create" methods determined by the CQRS type from the given file's services.
func (fns Functions) SyntheticMethods(file pgs.File) []pgs.Method {
	var methods []pgs.Method
	for _, service := range file.Services() {
		methods = service.Methods()
		break
	}

	var mutationMethods []pgs.Method

	for _, method := range methods {
		if fns.GetCQRSType(method) == "create" {
			mutationMethods = append(mutationMethods, method)
		}
	}

	return mutationMethods
}

// IsMethodSynthetic checks if the provided method is synthetic based on its CQRSType. Returns true for "create" methods.
func (fns Functions) IsMethodSynthetic(method pgs.Method) bool {
	if fns.GetCQRSType(method) == "create" {
		return true
	}

	return false
}

// HasSyntheticMethods checks whether the given file contains any synthetic methods and returns true if found, otherwise false.
func (fns Functions) HasSyntheticMethods(file pgs.File) bool {
	methods := fns.SyntheticMethods(file)

	if len(methods) > 0 {
		return true
	}

	return false
}
