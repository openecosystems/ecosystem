package shared

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	options "libs/protobuf/go/protobuf/gen/platform/options/v2"
)

func (fns Functions) ApiOptions(file pgs.File) options.ApiOptions {

	var apiFile options.ApiOptions

	_, err := file.Extension(options.E_ApiFile, &apiFile)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}

	return apiFile
}

func (fns Functions) GetApiOptionsTypeName(file pgs.File) pgs.Name {

	var apiType pgs.Name
	apiFile := fns.ApiOptions(file)
	switch apiFile.Type {
	case options.ApiType_API_TYPE_POC:
		apiType = "poc"
	case options.ApiType_API_TYPE_PUBLIC:
		apiType = "public"
	case options.ApiType_API_TYPE_PARTNER:
		apiType = "partner"
	case options.ApiType_API_TYPE_PRIVATE:
		apiType = "private"
	case options.ApiType_API_TYPE_UNSPECIFIED:
		apiType = "private"
	}

	return apiType
}

func (fns Functions) GetApiOptionsType(file pgs.File) string {
	return fns.GetApiOptionsTypeName(file).LowerCamelCase().String()
}
