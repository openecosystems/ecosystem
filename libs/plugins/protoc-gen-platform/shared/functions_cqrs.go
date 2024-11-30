package shared

import (
	"strings"

	pgs "github.com/lyft/protoc-gen-star/v2"
	options "libs/protobuf/go/protobuf/gen/platform/options/v2"
)

func (fns Functions) CQRSOptions(method pgs.Method) options.CQRSOptions {

	var cqrs options.CQRSOptions

	_, err := method.Extension(options.E_Cqrs, &cqrs)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}

	return cqrs
}

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

func (fns Functions) IsMethodQuery(method pgs.Method) bool {

	if strings.HasPrefix(fns.IsCQRSType(method), "query") {
		return true
	}

	return false
}

func (fns Functions) HasAnyCQRSMethods(file pgs.File) bool {

	if fns.HasQueryMethods(file) {
		return true
	}

	if fns.HasMutationMethods(file) {
		return true
	}

	return false

}

func (fns Functions) HasQueryMethods(file pgs.File) bool {

	methods := fns.QueryMethods(file)

	if len(methods) > 0 {
		return true
	}

	return false

}

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

func (fns Functions) IsMethodMutation(method pgs.Method) bool {

	if strings.HasPrefix(fns.IsCQRSType(method), "mutation") {
		return true
	}

	return false
}

func (fns Functions) HasMutationMethods(file pgs.File) bool {

	methods := fns.MutationMethods(file)

	if len(methods) > 0 {
		return true
	}

	return false

}

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

func (fns Functions) IsCQRSList(method pgs.Method) bool {

	t := fns.GetCQRSType(method)
	if t == "list" {
		return true
	}

	return false
}

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
