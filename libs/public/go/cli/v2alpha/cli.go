package cliv2alphalib

import (
	"context"

	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

// CLI represents the core structure for managing CLI operations, bindings, and their configuration.
type CLI struct {
	Bindings *sdkv2alphalib.Bindings
	Bounds   []sdkv2alphalib.Binding
}

// NewCLI initializes a new CLI instance by registering the provided bindings and returns the CLI object.
func NewCLI(bounds []sdkv2alphalib.Binding) *CLI {
	ctx := context.Background()

	bindings := sdkv2alphalib.RegisterBindings(ctx, bounds)

	cli := &CLI{
		Bindings: bindings,
		Bounds:   bounds,
	}

	return cli
}

// GracefulShutdown gracefully shuts down the CLI by cleaning up resources and closing bindings within a timeout context.
func (cli *CLI) GracefulShutdown() {
	_, cancel := context.WithTimeout(context.Background(), 30)
	defer cancel()

	sdkv2alphalib.ShutdownBindings(cli.Bindings)
}
