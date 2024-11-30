package protovalidatev0

import (
	"context"
	"fmt"
	"sync"

	"github.com/bufbuild/protovalidate-go"
	"libs/public/go/sdk/v2alpha"
)

type Binding struct {
	Validator *protovalidate.Validator

	configuration *Configuration
}

var (
	Bound       *Binding
	BindingName = "PROTOVALIDATE_BINDING"
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

func (b *Binding) GetBinding() interface{} {
	return Bound
}

func (b *Binding) Close() error {
	fmt.Println("Closing the Protovalidate binding")
	return nil
}
