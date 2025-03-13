package shared

import (
	"strings"

	pgs "github.com/lyft/protoc-gen-star/v2"
)

// ServiceTrailingComment extracts the trailing comments of the first service in a given protobuf file.
func (fns Functions) ServiceTrailingComment(file pgs.File) string {
	for _, service := range file.Services() {
		return service.SourceCodeInfo().Location().GetTrailingComments()
	}
	return ""
}

// ServiceLeadingComment retrieves the leading comments associated with the first service in the provided proto file.
// If no services are present or comments are missing, it returns an empty string.
func (fns Functions) ServiceLeadingComment(file pgs.File) string {
	for _, service := range file.Services() {
		return service.SourceCodeInfo().Location().GetLeadingComments()
	}
	return ""
}

// ServiceLeadingDetachedComments retrieves leading detached comments from each service in a provided proto file.
func (fns Functions) ServiceLeadingDetachedComments(file pgs.File) []string {
	for _, service := range file.Services() {
		return service.SourceCodeInfo().Location().GetLeadingDetachedComments()
	}
	return nil
}

// MethodLeadingComment retrieves the leading comments associated with a given Protobuf method in its source code.
func (fns Functions) MethodLeadingComment(method pgs.Method) string {
	return strings.TrimSpace(method.SourceCodeInfo().Location().GetLeadingComments())
}

// MethodTrailingComment retrieves the trailing comments associated with the provided gRPC method definition.
func (fns Functions) MethodTrailingComment(method pgs.Method) string {
	return strings.TrimSpace(method.SourceCodeInfo().Location().GetTrailingComments())
}

// MethodLeadingDetachedComments retrieves leading detached comments associated with a given protobuf method.
func (fns Functions) MethodLeadingDetachedComments(method pgs.Method) []string {
	return method.SourceCodeInfo().Location().GetLeadingDetachedComments()
}
