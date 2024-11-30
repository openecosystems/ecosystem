package main

import (
	"strings"
)

func (renderer *Renderer) RenderProvider() ([]byte, error) {
	f := NewLineWriter()
	f.WriteLine("// GENERATED FILE: DO NOT EDIT!\n")
	f.WriteLine("package " + renderer.Package)
	f.WriteLine(``)
	f.WriteLine(`import "context"`)
	f.WriteLine(``)
	f.WriteLine(`// To create a server, first write a class that implements this interface.`)
	f.WriteLine(`// Then pass an instance of it to Initialize().`)
	f.WriteLine(`type Provider interface {`)
	for _, method := range renderer.Model.Methods {
		parametersType := renderer.Model.TypeWithTypeName(method.ParametersTypeName)
		responsesType := renderer.Model.TypeWithTypeName(method.ResponsesTypeName)
		f.WriteLine(``)
		f.WriteLine(commentForText(method.Description))
		if parametersType != nil {
			if responsesType != nil {
				f.WriteLine(method.ProcessorName +
					`(ctx context.Context, parameters *` + parametersType.TypeName +
					`, responses *` + responsesType.TypeName + `) (err error)`)
			} else {
				f.WriteLine(method.ProcessorName + `(ctx context.Context, parameters *` + parametersType.TypeName + `) (err error)`)
			}
		} else {
			if responsesType != nil {
				f.WriteLine(method.ProcessorName + `(ctx context.Context, responses *` + responsesType.TypeName + `) (err error)`)
			} else {
				f.WriteLine(method.ProcessorName + `(ctx context.Context, ) (err error)`)
			}
		}
	}
	f.WriteLine(`}`)
	return f.Bytes(), nil
}

func commentForText(text string) string {
	result := ""
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		if i > 0 {
			result += "\n"
		}
		result += "// " + line
	}
	return result
}
