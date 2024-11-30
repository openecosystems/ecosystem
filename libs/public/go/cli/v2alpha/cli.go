package cliv2alphalib

import (
	"context"
	"libs/public/go/sdk/v2alpha"
)

type CLI struct {
	Bindings *sdkv2alphalib.Bindings
	Bounds   []sdkv2alphalib.Binding
}

func NewCLI(bounds []sdkv2alphalib.Binding) *CLI {

	ctx := context.Background()

	bindings := sdkv2alphalib.RegisterBindings(ctx, bounds)

	cli := &CLI{
		Bindings: bindings,
		Bounds:   bounds,
	}

	return cli
}

func (cli *CLI) GracefulShutdown() {

	_, cancel := context.WithTimeout(context.Background(), 30)
	defer cancel()

	sdkv2alphalib.ShutdownBindings(cli.Bindings)
}
