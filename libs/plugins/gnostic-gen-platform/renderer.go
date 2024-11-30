package main

import (
	"fmt"
	_ "os"
	"path/filepath"

	plugins "github.com/googleapis/gnostic/plugins"
	surface "github.com/googleapis/gnostic/surface"
	"golang.org/x/tools/imports"
)

// Renderer generates code for a surface.Model.
type Renderer struct {
	Model   *surface.Model
	Package string // package name
}

// NewServiceRenderer creates a renderer.
func NewServiceRenderer(model *surface.Model) (renderer *Renderer, err error) {
	renderer = &Renderer{}
	renderer.Model = model
	return renderer, nil
}

// Generate runs the renderer to generate the named files.
func (renderer *Renderer) Render(response *plugins.Response, files []string) (err error) {
	for _, filename := range files {
		file := &plugins.File{Name: filename}
		switch filename {
		case "client.go":
			file.Data, err = renderer.RenderClient()
		case "types.go":
			file.Data, err = renderer.RenderTypes()
		case "provider.go":
			file.Data, err = renderer.RenderProvider()
		case "server.go":
			file.Data, err = renderer.RenderServer()
		case "constants.go":
			file.Data, err = renderer.RenderConstants()
		default:
			file.Data = nil
		}
		if err != nil {
			response.Errors = append(response.Errors, fmt.Sprintf("ERROR %v", err))
		}
		// run generated Go files through imports pkg
		if filepath.Ext(file.Name) == ".go" {
			file.Data, err = imports.Process(file.Name, file.Data, nil)
			if err != nil {
				response.Errors = append(response.Errors, err.Error())
			}
		}
		response.Files = append(response.Files, file)
	}
	return
}
