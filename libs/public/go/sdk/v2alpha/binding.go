package sdkv2alphalib

import (
	"context"
	"fmt"
)

type Bindings struct {
	Registered                   map[string]Binding
	RegisteredListenableChannels map[string]Listenable
}

type Binding interface {
	Name() string

	Validate(ctx context.Context, bindings *Bindings) error

	Bind(ctx context.Context, bindings *Bindings) *Bindings

	GetBinding() interface{}

	Close() error
}

type SpecListenableErr struct {
	Error error
}

type Listenable interface {
	Listen(ctx context.Context, listenerErr chan SpecListenableErr)
}

var Bounds *Bindings

func RegisterBindings(ctx context.Context, bounds []Binding) *Bindings {
	b := make(map[string]Binding)
	c := make(map[string]Listenable)
	bindingsInstance := &Bindings{
		Registered:                   b,
		RegisteredListenableChannels: c,
	}

	var errs []error
	for _, b := range bounds {
		if c, ok := b.(Configurable); ok {
			c.ResolveConfiguration()
			err := c.ValidateConfiguration()
			if err != nil {
				errs = append(errs, err)
			}
		}

		if err := b.Validate(ctx, bindingsInstance); err != nil {
			fmt.Println("validate error: ", err)
			errs = append(errs, err)
		}

		if len(errs) > 0 {
			fmt.Println("binding errors: ", errs)
			panic(errs)
		}

		bindingsInstance = b.Bind(ctx, bindingsInstance)
	}

	Bounds = bindingsInstance

	return bindingsInstance
}

func ShutdownBindings(bindings *Bindings) {
	if bindings.Registered != nil {
		for _, b := range bindings.Registered {
			if err := b.Close(); err != nil {
				fmt.Print(err)
			}
		}
	}
}
