package shared

import (
	"strings"

	options "libs/protobuf/go/protobuf/gen/platform/options/v2"

	pgs "github.com/lyft/protoc-gen-star/v2"
)

// ApiOptions retrieves the API options from the given protobuf file's custom extension and returns them as ApiOptions.
func (fns Functions) ApiOptions(file pgs.File) options.ApiOptions {
	var apiFile options.ApiOptions

	_, err := file.Extension(options.E_ApiFile, &apiFile)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}

	return apiFile
}

// GetApiOptionsTypeName maps a file's API type to its corresponding string representation (e.g., "poc", "public", etc.).
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

// GetApiOptionsType generates the API options type for the provided file in lower camel case format.
func (fns Functions) GetApiOptionsType(file pgs.File) string {
	return fns.GetApiOptionsTypeName(file).LowerCamelCase().String()
}

// GetTopLevelFolderFromFile extracts the top-level folder name from the provided file's path.
func (fns Functions) GetTopLevelFolderFromFile(file pgs.File) pgs.Name {
	path := file.Name().Split()
	v := strings.Split(path[0], "/")
	return pgs.Name(v[0])
}
