package pushpinv1

import (
	"context"
	"fmt"
	"sync"

	"libs/public/go/sdk/v2alpha"
)

// Binding struct that holds binding specific fields
type Binding struct {
	configuration *Configuration
}

var (
	Bound       *Binding
	BindingName = "PUSHPIN_BINDING"
	IsBound     = false
)

func (b *Binding) Name() string {
	return BindingName
}

func (b *Binding) Validate(_ context.Context, _ *sdkv2alphalib.Bindings) error {
	// Verify any requirements

	return nil
}

func (b *Binding) Bind(_ context.Context, bindings *sdkv2alphalib.Bindings) *sdkv2alphalib.Bindings {
	if Bound == nil {
		var once sync.Once
		once.Do(
			func() {
				Bound = &Binding{}

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

func (b *Binding) GetBinding() interface{} {
	return Bound
}

func (b *Binding) Close() error {
	return nil
}
