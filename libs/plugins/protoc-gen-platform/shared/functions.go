package shared

import (
	"strings"

	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
)

type Functions struct{ Pctx pgsgo.Context }

func (fns Functions) IsLast(i, size int) bool {
	return i == size-1
}

func (fns Functions) PluginName(params pgs.Parameters) func() string {
	return func() string {
		return params.Str(LanguageParam) + "/" + params.Str(TypeParam)
	}
}

func (fns Functions) DotNotationToFilePath(dot string) pgs.FilePath {

	path := strings.Split(dot, ".")
	return pgs.JoinPaths(path...)

}

func (fns Functions) InputCondition(field pgs.Field) string {

	if field.Type().IsRepeated() {
		return "nil"
	}

	switch field.Type().ProtoType() {
	case pgs.Int64T, pgs.UInt64T, pgs.Int32T, pgs.Fixed64T, pgs.Fixed32T, pgs.UInt32T, pgs.SFixed32, pgs.SFixed64, pgs.SInt32, pgs.SInt64:
		return "no"
	case pgs.DoubleT, pgs.FloatT:
		return "no"
	case pgs.StringT:
		return "string"
	case pgs.BytesT:
		return "nil"
	case pgs.BoolT:
		return "no"
	case pgs.MessageT:
		if field.Type().IsEmbed() {
			return "no"
		}
	case pgs.GroupT:
		panic("Group type is not supported")
	case pgs.EnumT:

		if field.Type().IsRepeated() {

			if field.Type().Element().IsEmbed() {
				return "nil"
			}

		} else {
			return "no"
		}

	default:
		panic("unreachable: invalid scalar type")
	}

	panic("Could not process")

}

func (fns Functions) GetMethodVerb(method pgs.Method) pgs.Name {

	if method == nil {
		panic("Method cannot be nil")
	}

	return pgs.Name(method.Name().Split()[0])

}

func (fns Functions) AllMethods(file pgs.File) []pgs.Method {

	var methods []pgs.Method
	for _, service := range file.Services() {
		methods = service.Methods()
		break
	}

	var allMethods []pgs.Method

	for _, method := range methods {
		allMethods = append(allMethods, method)
	}

	return allMethods
}

func (fns Functions) ProtoName(file pgs.File) string {
	split := strings.Split(file.Name().String(), "/")
	fileName := split[len(split)-1]
	return strings.Split(fileName, ".proto")[0]
}

func (fns Functions) SnakeCaseToDashCase(name string) string {
	return strings.Replace(name, "_", "-", -1)
}

func (fns Functions) DescriptorPackage(file pgs.File) string {
	if file.Descriptor().Package == nil {
		return ""
	}
	return *file.Descriptor().Package
}
