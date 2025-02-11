package pushpinv1

import (
	"context"
	"fmt"
	"sync"

	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

// Binding represents a struct that manages configuration and binding lifecycle interactions.
type Binding struct {
	configuration *Configuration
}

// Bound represents the current instance of Binding, shared globally within the application.
// BindingName specifies the unique name identifier for the Binding instance.
// IsBound indicates the binding status of the Binding instance.
var (
	Bound       *Binding
	BindingName = "PUSHPIN_BINDING"
	IsBound     = false
)

// Name returns the name of the binding as a string.
func (b *Binding) Name() string {
	return BindingName
}

// Validate checks the validity of the binding and ensures it meets required conditions or constraints.
func (b *Binding) Validate(_ context.Context, _ *sdkv2alphalib.Bindings) error {
	// Verify any requirements

	return nil
}

// Bind registers the Binding instance into the Bindings registry if not already registered, ensuring it's bound once.
func (b *Binding) Bind(_ context.Context, bindings *sdkv2alphalib.Bindings) *sdkv2alphalib.Bindings {
	if Bound == nil {
		var once sync.Once
		once.Do(
			func() {
				Bound = &Binding{
					configuration: b.configuration,
				}

				bindings.Registered[b.Name()] = Bound
				IsBound = true
			})
	} else {
		bindings.Registered[b.Name()] = Bound
		IsBound = true
		fmt.Println("Pushpin already bound")
	}

	return bindings
}

// GetBinding returns the currently bound Binding instance.
func (b *Binding) GetBinding() interface{} {
	return Bound
}

// Close releases any resources or performs cleanup tasks before the Binding instance is discarded. Returns an error if any.
func (b *Binding) Close() error {
	return nil
}
