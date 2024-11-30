package main

import (
	"encoding/json"
	"errors"
	"go/format"
	"path/filepath"
	"strings"

	"github.com/golang/protobuf/proto"
	plugins "github.com/googleapis/gnostic/plugins"
	surface "github.com/googleapis/gnostic/surface"
)

// This is the main function for the code generation plugin.
func main() {
	env, err := plugins.NewEnvironment()
	env.RespondAndExitIfError(err)

	packageName, err := resolvePackageName(env.Request.OutputPath)
	env.RespondAndExitIfError(err)

	// Use the name used to run the plugin to decide which files to generate.
	var files []string
	switch {
	case strings.Contains(env.Invocation, "gnostic-go-client"):
		files = []string{"client.go", "types.go", "constants.go"}
	case strings.Contains(env.Invocation, "gnostic-go-server"):
		files = []string{"server.go", "provider.go", "types.go", "constants.go"}
	default:
		files = []string{"client.go", "server.go", "provider.go", "types.go", "constants.go"}
	}

	inputDocumentType := env.Request.Models[0].TypeUrl
	for _, model := range env.Request.Models {
		switch model.TypeUrl {
		case "surface.v1.Model":
			surfaceModel := &surface.Model{}
			err = proto.Unmarshal(model.Value, surfaceModel)
			if err == nil {
				// Customize the code surface model for Go
				NewGoLanguageModel().Prepare(surfaceModel, inputDocumentType)

				modelJSON, _ := json.MarshalIndent(surfaceModel, "", "  ")
				modelFile := &plugins.File{Name: "model.json", Data: modelJSON}
				env.Response.Files = append(env.Response.Files, modelFile)

				// Create the renderer.
				renderer, err := NewServiceRenderer(surfaceModel)
				renderer.Package = packageName
				env.RespondAndExitIfError(err)

				// Run the renderer to generate files and add them to the response object.
				err = renderer.Render(env.Response, files)
				env.RespondAndExitIfError(err)

				// Return with success.
				env.RespondAndExit()
			}
		}
	}
	err = errors.New("No generated code surface model is available.")
	env.RespondAndExitIfError(err)
}

// resolvePackageName converts a path to a valid package name or
// error if path can't be resolved or resolves to an invalid package name.
func resolvePackageName(p string) (string, error) {
	p, err := filepath.Abs(p)
	if err == nil {
		p = filepath.Base(p)
		_, err = format.Source([]byte("package " + p))
	}
	if err != nil {
		return "", errors.New("invalid package name " + p)
	}
	return p, nil
}
