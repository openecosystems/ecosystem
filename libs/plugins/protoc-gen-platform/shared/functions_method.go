package shared

import (
	pgs "github.com/lyft/protoc-gen-star/v2"

	options "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/options/v2"
)

// GetMethodShortName returns the short name of a method based on its ApiOptions or defaults to the method name if none exists.
func (fns Functions) GetMethodShortName(method pgs.Method) pgs.Name {
	shortName := method.Name()

	var api options.ApiOptions
	_, err := method.Extension(options.E_ApiMethod, &api)
	if err != nil {
		return shortName
	}

	if api.Shortname == "" {
		return shortName
	}

	return pgs.Name(api.Shortname)
}
