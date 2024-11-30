package shared

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	options "libs/protobuf/go/protobuf/gen/platform/options/v2"
)

func (fns Functions) SyntheticOptions(method pgs.Method) options.SyntheticOptions {

	var synthetic options.SyntheticOptions

	_, err := method.Extension(options.E_Synthetic, &synthetic)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}

	return synthetic
}

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

func (fns Functions) GetSyntheticDictionaryKeyOld(file pgs.Field) string {

	var synthetic options.SyntheticOptions

	_, err := file.Extension(options.E_Synthetic, &synthetic)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}

	return synthetic.GetDictionaryKey()

}

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

func (fns Functions) IsMethodSynthetic(method pgs.Method) bool {

	if fns.GetCQRSType(method) == "create" {
		return true
	}

	return false
}

func (fns Functions) HasSyntheticMethods(file pgs.File) bool {

	methods := fns.SyntheticMethods(file)

	if len(methods) > 0 {
		return true
	}

	return false

}
