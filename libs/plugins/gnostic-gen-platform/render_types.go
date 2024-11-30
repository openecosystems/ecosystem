package main

import (
	surface "github.com/googleapis/gnostic/surface"
)

func (renderer *Renderer) RenderTypes() ([]byte, error) {
	f := NewLineWriter()
	f.WriteLine(`// GENERATED FILE: DO NOT EDIT!`)
	f.WriteLine(``)
	f.WriteLine(`package ` + renderer.Package)
	f.WriteLine(`// Types used by the API.`)
	for _, modelType := range renderer.Model.Types {
		f.WriteLine(`// ` + modelType.Description)
		if modelType.Kind == surface.TypeKind_STRUCT {
			f.WriteLine(`type ` + modelType.TypeName + ` struct {`)
			for _, field := range modelType.Fields {
				typ := field.NativeType
				if field.Kind == surface.FieldKind_REFERENCE {
					typ = "*" + typ
				} else if field.Kind == surface.FieldKind_ARRAY {
					typ = "[]" + typ
				} else if field.Kind == surface.FieldKind_MAP {
					typ = "map[string]" + typ
				} else if field.Kind == surface.FieldKind_ANY {
					typ = "interface{}"
				}
				f.WriteLine(field.FieldName + ` ` + typ + jsonTag(field))
			}
			f.WriteLine(`}`)
		} else if modelType.Kind == surface.TypeKind_OBJECT {
			f.WriteLine(`type ` + modelType.TypeName + ` map[string]` + modelType.ContentType)
		} else {
			f.WriteLine(`type ` + modelType.TypeName + ` interface {}`)
		}
	}
	return f.Bytes(), nil
}

func jsonTag(field *surface.Field) string {
	if field.Serialize {
		return " `json:" + `"` + field.Name + `,omitempty"` + "`"
	}
	return ""
}
