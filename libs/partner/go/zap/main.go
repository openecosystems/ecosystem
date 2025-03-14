package zaploggerv1

import (
	"context"
	"fmt"
	"sync"

	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"

	"go.uber.org/zap"
)

// Binding represents a logging framework binding utilizing Uber Zap for structured logging.
// It includes both a standard logger and a sugared logger for flexible usage within applications.
// This type also manages the configuration for initializing the logger effectively.
type Binding struct {
	Logger        *zap.Logger
	SugaredLogger *zap.SugaredLogger

	configuration *Configuration
}

// Bound is a global variable representing an instance of the Binding structure for logging purposes.
// BindingName is a constant string used as the identifier for the logging binding.
var (
	Bound       *Binding
	BindingName = "ZAP_LOGGING_BINDING"
)

// Name returns the unique name identifier for the Binding.
func (b *Binding) Name() string {
	return BindingName
}

// Validate performs validation checks on the Binding within the provided context and bindings.
func (b *Binding) Validate(_ context.Context, _ *sdkv2alphalib.Bindings) error {
	// Verify any log requirements
	return nil
}

// Bind attaches the binding to the provided bindings structure, initializes logger if necessary, and returns updated bindings.
func (b *Binding) Bind(_ context.Context, bindings *sdkv2alphalib.Bindings) *sdkv2alphalib.Bindings {
	if Bound == nil {
		var once sync.Once
		once.Do(
			func() {
				var err error
				b.Logger, err = b.configuration.Zap.Build()
				if err != nil {
					fmt.Println(fmt.Errorf("could not build Zap logger: %v", err))
				}

				defer b.Logger.Sync() //nolint:errcheck

				Bound = &Binding{
					Logger:        b.Logger,
					SugaredLogger: b.Logger.Sugar(),

					configuration: b.configuration,
				}
				bindings.Registered[b.Name()] = Bound
			})
	} else {
		bindings.Registered[b.Name()] = Bound
		fmt.Println("Zap Logging already bound")
	}

	return bindings
}

// GetBinding returns the current instance of the Bound binding.
func (b *Binding) GetBinding() interface{} {
	return Bound
}

// Close shuts down the Uber Zap Logger Binding and performs necessary cleanup operations.
func (b *Binding) Close() error {
	fmt.Println("Closing the Uber Zap Logger Binding")
	return nil
}
