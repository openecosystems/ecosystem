package zaploggerv1

import (
	"context"
	"fmt"
	"sync"

	"go.uber.org/zap"
	"libs/public/go/sdk/v2alpha"
)

type Binding struct {
	Logger        *zap.Logger
	SugaredLogger *zap.SugaredLogger

	configuration *Configuration
}

var (
	Bound       *Binding
	BindingName = "ZAP_LOGGING_BINDING"
)

func (b *Binding) Name() string {
	return BindingName
}

func (b *Binding) Validate(_ context.Context, _ *sdkv2alphalib.Bindings) error {
	// Verify any log requirements
	return nil
}

func (b *Binding) Bind(_ context.Context, bindings *sdkv2alphalib.Bindings) *sdkv2alphalib.Bindings {
	if Bound == nil {
		var once sync.Once
		once.Do(
			func() {
				var err error
				b.Logger, err = ResolvedConfiguration.Zap.Build()
				if err != nil {
					fmt.Println(fmt.Errorf("could not build Zap logger: %v", err))
				}

				defer func(Logger *zap.Logger) {
					err := Logger.Sync()
					if err != nil {
					}
				}(b.Logger)

				Bound = &Binding{
					Logger:        b.Logger,
					SugaredLogger: b.Logger.Sugar(),
				}
				bindings.Registered[b.Name()] = Bound
			})
	} else {
		bindings.Registered[b.Name()] = Bound
		fmt.Println("Zap Logging already bound")
	}

	return bindings
}

func (b *Binding) GetBinding() interface{} {
	return Bound
}

func (b *Binding) Close() error {
	fmt.Println("Closing the Uber Zap Logger Binding")
	return nil
}
