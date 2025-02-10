package sdkv2alphalib

import (
	"context"
	"fmt"
)

// Bindings represents a collection of registered bindings and their associated listenable channels.
type Bindings struct {
	Registered                   map[string]Binding
	RegisteredListenableChannels map[string]Listenable
}

// Binding represents an interface for implementing components required for validation, binding, and resource management.
// Name returns the unique name identifier for the binding.
// Validate performs validation checks for the binding within the given context and bindings.
// Bind attaches the binding to the provided bindings structure and returns an updated instance.
// GetBinding retrieves the underlying concrete binding implementation.
// Close releases resources or performs cleanup tasks associated with the binding.
type Binding interface {
	Name() string

	Validate(ctx context.Context, bindings *Bindings) error

	Bind(ctx context.Context, bindings *Bindings) *Bindings

	GetBinding() interface{}

	Close() error
}

// SpecListenableErr represents an error that can be sent via a channel while listening on a SpecListenable interface.
type SpecListenableErr struct {
	Error error
}

// Listenable represents an entity capable of listening to events or data streams in a contextual environment.
// Listen begins listening using the provided context and sends errors through the specified SpecListenableErr channel.
type Listenable interface {
	Listen(ctx context.Context, listenerErr chan SpecListenableErr)
}

// Bounds is a global variable holding the active *Bindings instance containing registered bindings and listenable channels.
var Bounds *Bindings

// RegisterBindings initializes and validates bindings, resolving configurations and handling errors during the process.
// It returns a populated *Bindings object and updates the global Bounds variable.
func RegisterBindings(ctx context.Context, bounds []Binding) *Bindings {
	b := make(map[string]Binding)
	c := make(map[string]Listenable)
	bindingsInstance := &Bindings{
		Registered:                   b,
		RegisteredListenableChannels: c,
	}

	var errs []error
	for _, b := range bounds {
		fmt.Println("binding: ", b.Name())
		if c, ok := b.(SpecConfigurationProvider); ok {
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

// ShutdownBindings gracefully closes all registered bindings by invoking their Close method. Errors during closure are printed.
func ShutdownBindings(bindings *Bindings) {
	if bindings.Registered != nil {
		for _, b := range bindings.Registered {
			if err := b.Close(); err != nil {
				fmt.Print(err)
			}
		}
	}
}
