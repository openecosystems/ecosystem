package shared

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
)

func (fns Functions) ServiceTrailingComment(file pgs.File) string {
	for _, service := range file.Services() {
		return service.SourceCodeInfo().Location().GetTrailingComments()
	}
	return ""
}

func (fns Functions) ServiceLeadingComment(file pgs.File) string {
	for _, service := range file.Services() {
		return service.SourceCodeInfo().Location().GetLeadingComments()
	}
	return ""
}

func (fns Functions) ServiceLeadingDetachedComments(file pgs.File) []string {
	for _, service := range file.Services() {
		return service.SourceCodeInfo().Location().GetLeadingDetachedComments()
	}
	return nil
}

func (fns Functions) MethodLeadingComment(method pgs.Method) string {
	return method.SourceCodeInfo().Location().GetLeadingComments()
}

func (fns Functions) MethodTrailingComment(method pgs.Method) string {
	return method.SourceCodeInfo().Location().GetTrailingComments()
}

func (fns Functions) MethodLeadingDetachedComments(method pgs.Method) []string {
	return method.SourceCodeInfo().Location().GetLeadingDetachedComments()
}
