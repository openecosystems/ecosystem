package protovalidatev0

import (
	"context"
	"fmt"
	"sync"

	sdkv2alphalib "libs/public/go/sdk/v2alpha"

	"github.com/bufbuild/protovalidate-go"
)

// Binding represents a structure for managing a validator and associated configuration.
type Binding struct {
	Validator *protovalidate.Validator

	configuration *Configuration
}

// Bound represents the initialized instance of the Binding object, used globally for protovalidation.
// BindingName defines the unique name identifier for the Protovalidate Binding.
var (
	Bound       *Binding
	BindingName = "PROTOVALIDATE_BINDING"
)

// Name returns the unique name identifier for the binding.
func (b *Binding) Name() string {
	return BindingName
}

// Validate performs validation checks for the binding within the provided context and bindings.
func (b *Binding) Validate(_ context.Context, _ *sdkv2alphalib.Bindings) error {
	// Verify any requirements

	return nil
}

// Bind attaches the binding to the given Bindings instance, ensuring a singleton binding registration with validation support.
func (b *Binding) Bind(_ context.Context, bindings *sdkv2alphalib.Bindings) *sdkv2alphalib.Bindings {
	if Bound == nil {
		var once sync.Once
		once.Do(
			func() {
				v, err := protovalidate.New()
				if err != nil {
					fmt.Println("failed to initialize validator:", err)
					panic(err)
				}

				Bound = &Binding{
					Validator: v,
				}
				bindings.Registered[b.Name()] = Bound
			})
	} else {
		bindings.Registered[b.Name()] = Bound
		fmt.Println("Protovalidate already bound")
	}

	return bindings
}

// GetBinding retrieves the currently bound instance of the Binding struct.
func (b *Binding) GetBinding() interface{} {
	return Bound
}

// Close closes the Protovalidate binding and performs any necessary cleanup operations.
func (b *Binding) Close() error {
	fmt.Println("Closing the Protovalidate binding")
	return nil
}
