package shared

import (
	"strings"

	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
)

// Functions provides a collection of utility methods for processing Protobuf data and generating corresponding outputs.
type Functions struct{ Pctx pgsgo.Context }

// IsLast checks if the given index `i` is the last index in a collection of size `size`. Returns true if it is, else false.
func (fns Functions) IsLast(i, size int) bool {
	return i == size-1
}

// PluginName returns a function that combines the language and type parameters into a formatted string.
func (fns Functions) PluginName(params pgs.Parameters) func() string {
	return func() string {
		return params.Str(LanguageParam) + "/" + params.Str(TypeParam)
	}
}

// DotNotationToFilePath converts a dot-separated string into a file path by replacing dots with path separators.
func (fns Functions) DotNotationToFilePath(dot string) pgs.FilePath {
	path := strings.Split(dot, ".")
	return pgs.JoinPaths(path...)
}

// InputCondition determines an input condition string based on a field's type, such as "string", "no", or "nil".
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

// GetMethodVerb extracts the first part of a method's name as its "verb" and returns it as a pgs.Name.
// Panics if the provided method is nil.
func (fns Functions) GetMethodVerb(method pgs.Method) pgs.Name {
	if method == nil {
		panic("Method cannot be nil")
	}

	return pgs.Name(method.Name().Split()[0])
}

// AllMethods returns a slice of all methods from the first service defined in the given Protobuf file.
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

// ProtoName extracts the base name of a proto file, removing its directory path and ".proto" extension.
func (fns Functions) ProtoName(file pgs.File) string {
	split := strings.Split(file.Name().String(), "/")
	fileName := split[len(split)-1]
	return strings.Split(fileName, ".proto")[0]
}

// ProtoPathWithoutProtoExtension returns the file path of a Protobuf file without its ".proto" extension.
func (fns Functions) ProtoPathWithoutProtoExtension(file pgs.File) string {
	split := strings.Split(file.Name().String(), ".")
	fileName := split[0]
	return fileName
}

// SnakeCaseToDashCase converts a snake_case string to a dash-case string by replacing underscores with hyphens.
func (fns Functions) SnakeCaseToDashCase(name string) string {
	return strings.Replace(name, "_", "-", -1)
}

// DashCase converts a snake_case string to a dash-case string by replacing underscores with hyphens.
func (fns Functions) DashCase(name pgs.Name) string {
	snake := name.LowerSnakeCase()
	return fns.SnakeCaseToDashCase(snake.String())
}

// DescriptorPackage extracts and returns the package name as a string from the given Protocol Buffers file descriptor.
func (fns Functions) DescriptorPackage(file pgs.File) string {
	if file.Descriptor().Package == nil {
		return ""
	}
	return *file.Descriptor().Package
}
