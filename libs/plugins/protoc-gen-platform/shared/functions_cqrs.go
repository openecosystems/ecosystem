package shared

import (
	"strings"

	options "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/options/v2"

	pgs "github.com/lyft/protoc-gen-star/v2"
)

// CQRSOptions retrieves the CQRSOptions extension from the specified proto method and returns it.
func (fns Functions) CQRSOptions(method pgs.Method) options.CQRSOptions {
	var cqrs options.CQRSOptions

	_, err := method.Extension(options.E_Cqrs, &cqrs)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}

	return cqrs
}

// QueryMethods extracts and returns all method descriptors from the given file that are identified as query methods.
func (fns Functions) QueryMethods(file pgs.File) []pgs.Method {
	var methods []pgs.Method
	for _, service := range file.Services() {
		methods = service.Methods()
		break
	}

	var queryMethods []pgs.Method

	for _, method := range methods {
		if strings.HasPrefix(fns.IsCQRSType(method), "query") {
			queryMethods = append(queryMethods, method)
		}
	}

	return queryMethods
}

// IsMethodQuery determines if a given pgs.Method is categorized as a "query" based on its CQRSType prefix.
func (fns Functions) IsMethodQuery(method pgs.Method) bool {
	if strings.HasPrefix(fns.IsCQRSType(method), "query") {
		return true
	}

	return false
}

// HasAnyCQRSMethods checks if the given file contains at least one method classified as a Query or Mutation method.
func (fns Functions) HasAnyCQRSMethods(file pgs.File) bool {
	if fns.HasQueryMethods(file) {
		return true
	}

	if fns.HasMutationMethods(file) {
		return true
	}

	return false
}

// HasQueryMethods checks if the given file contains any query methods and returns true if at least one is found.
func (fns Functions) HasQueryMethods(file pgs.File) bool {
	methods := fns.QueryMethods(file)

	if len(methods) > 0 {
		return true
	}

	return false
}

// MutationMethods extracts and returns methods that are classified as "mutation" based on their CQRS type from the provided file.
func (fns Functions) MutationMethods(file pgs.File) []pgs.Method {
	var methods []pgs.Method
	for _, service := range file.Services() {
		methods = service.Methods()
		break
	}

	var mutationMethods []pgs.Method

	for _, method := range methods {
		if strings.HasPrefix(fns.IsCQRSType(method), "mutation") {
			mutationMethods = append(mutationMethods, method)
		}
	}

	return mutationMethods
}

// StreamingMethods extracts and returns methods that are classified as "mutation" based on their CQRS type from the provided file.
func (fns Functions) StreamingMethods(file pgs.File) []pgs.Method {
	var methods []pgs.Method
	for _, service := range file.Services() {
		methods = service.Methods()
		break
	}

	var mutationMethods []pgs.Method

	for _, method := range methods {
		if strings.Contains(fns.IsCQRSType(method), "stream") {
			mutationMethods = append(mutationMethods, method)
		}
	}

	return mutationMethods
}

// IsMethodMutation determines if the given method is a mutation based on its CQRS type prefix.
func (fns Functions) IsMethodMutation(method pgs.Method) bool {
	if strings.HasPrefix(fns.IsCQRSType(method), "mutation") {
		return true
	}

	return false
}

// HasMutationMethods determines if the given file contains any mutation methods by checking the length of retrieved methods.
func (fns Functions) HasMutationMethods(file pgs.File) bool {
	methods := fns.MutationMethods(file)

	if len(methods) > 0 {
		return true
	}

	return false
}

// IsCQRSType determines the CQRS type of the given gRPC method and returns it as a string representation.
func (fns Functions) IsCQRSType(method pgs.Method) string {
	var cqrs options.CQRSOptions

	_, err := method.Extension(options.E_Cqrs, &cqrs)
	if err != nil {
		panic(err.Error() + "unable to read method extension from method")
	}

	switch cqrs.Type {
	case options.CQRSType_CQRS_TYPE_NONE:
		return "none"
	case options.CQRSType_CQRS_TYPE_MUTATION_CREATE:
		return "mutation"
	case options.CQRSType_CQRS_TYPE_MUTATION_UPDATE:
		return "mutation"
	case options.CQRSType_CQRS_TYPE_MUTATION_DELETE:
		return "mutation"
	case options.CQRSType_CQRS_TYPE_MUTATION_CLIENT_STREAM:
		return "mutation-client-stream"
	case options.CQRSType_CQRS_TYPE_MUTATION_SERVER_STREAM:
		return "mutation-server-stream"
	case options.CQRSType_CQRS_TYPE_MUTATION_BIDI_STREAM:
		return "mutation-bidi-stream"
	case options.CQRSType_CQRS_TYPE_QUERY_LIST:
		return "query"
	case options.CQRSType_CQRS_TYPE_QUERY_GET:
		return "query"
	case options.CQRSType_CQRS_TYPE_QUERY_EXISTS:
		return "query"
	case options.CQRSType_CQRS_TYPE_QUERY_CLIENT_STREAM:
		return "query-client-stream"
	case options.CQRSType_CQRS_TYPE_QUERY_SERVER_STREAM:
		return "query-server-stream"
	case options.CQRSType_CQRS_TYPE_QUERY_BIDI_STREAM:
		return "query-bidi-stream"
	default:
		return "none"
	}
}

// IsCQRSList determines if the given method is of the CQRS type "list" by evaluating its associated CQRS type.
func (fns Functions) IsCQRSList(method pgs.Method) bool {
	t := fns.GetCQRSType(method)
	if t == "list" {
		return true
	}

	return false
}

// GetCQRSType determines the CQRS type of a given method based on its CQRS extension options and returns it as a string.
func (fns Functions) GetCQRSType(method pgs.Method) string {
	var cqrs options.CQRSOptions

	_, err := method.Extension(options.E_Cqrs, &cqrs)
	if err != nil {
		panic(err.Error() + "unable to read method extension from method")
	}

	switch cqrs.Type {
	case options.CQRSType_CQRS_TYPE_NONE:
		return "none"
	case options.CQRSType_CQRS_TYPE_MUTATION_CREATE:
		return "create"
	case options.CQRSType_CQRS_TYPE_MUTATION_UPDATE:
		return "update"
	case options.CQRSType_CQRS_TYPE_MUTATION_DELETE:
		return "delete"
	case options.CQRSType_CQRS_TYPE_MUTATION_CLIENT_STREAM:
		return "stream-client"
	case options.CQRSType_CQRS_TYPE_MUTATION_SERVER_STREAM:
		return "stream-server"
	case options.CQRSType_CQRS_TYPE_MUTATION_BIDI_STREAM:
		return "stream-bidi"
	case options.CQRSType_CQRS_TYPE_QUERY_LIST:
		return "list"
	case options.CQRSType_CQRS_TYPE_QUERY_GET:
		return "get"
	case options.CQRSType_CQRS_TYPE_QUERY_EXISTS:
		return "exists"
	case options.CQRSType_CQRS_TYPE_QUERY_CLIENT_STREAM:
		return "stream-client"
	case options.CQRSType_CQRS_TYPE_QUERY_SERVER_STREAM:
		return "stream-server"
	case options.CQRSType_CQRS_TYPE_QUERY_BIDI_STREAM:
		return "stream-bidi"
	default:
		return "none"
	}
}

// GetCQRSTypeEnumName determines the CQRS type of a given method based on its CQRS extension options and returns it as a string.
func (fns Functions) GetCQRSTypeEnumName(method pgs.Method) string {
	var cqrs options.CQRSOptions

	_, err := method.Extension(options.E_Cqrs, &cqrs)
	if err != nil {
		panic(err.Error() + "unable to read method extension from method")
	}

	return cqrs.Type.String()
}

// ConvertCQRSTypeToString converts a given CQRSType value to its corresponding string representation such as "Mutation" or "Query".
func (fns Functions) ConvertCQRSTypeToString(t options.CQRSType) string {
	switch t {
	case options.CQRSType_CQRS_TYPE_MUTATION_CREATE:
		fallthrough
	case options.CQRSType_CQRS_TYPE_MUTATION_UPDATE:
		fallthrough
	case options.CQRSType_CQRS_TYPE_MUTATION_CLIENT_STREAM:
		fallthrough
	case options.CQRSType_CQRS_TYPE_MUTATION_SERVER_STREAM:
		fallthrough
	case options.CQRSType_CQRS_TYPE_MUTATION_BIDI_STREAM:
		fallthrough
	case options.CQRSType_CQRS_TYPE_MUTATION_DELETE:
		return "Mutation"
	case options.CQRSType_CQRS_TYPE_QUERY_LIST:
		fallthrough
	case options.CQRSType_CQRS_TYPE_QUERY_EXISTS:
		fallthrough
	case options.CQRSType_CQRS_TYPE_QUERY_STREAM:
		fallthrough
	case options.CQRSType_CQRS_TYPE_QUERY_CLIENT_STREAM:
		fallthrough
	case options.CQRSType_CQRS_TYPE_QUERY_SERVER_STREAM:
		fallthrough
	case options.CQRSType_CQRS_TYPE_QUERY_BIDI_STREAM:
		fallthrough
	case options.CQRSType_CQRS_TYPE_QUERY_GET:
		return "Query"
	case options.CQRSType_CQRS_TYPE_NONE:
		fallthrough
	case options.CQRSType_CQRS_TYPE_UNSPECIFIED:
		return ""
	default:
		return ""
	}
}
