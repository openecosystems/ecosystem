package protovalidatev0

import (
	"context"
	"fmt"
	"sync"

	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"

	"github.com/bufbuild/protovalidate-go"
	bufcel "github.com/bufbuild/protovalidate-go/cel"
	"github.com/google/cel-go/cel"
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
					Validator:     &v,
					configuration: b.configuration,
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

// GetLibraryRules compiles and evaluates a predefined CEL expression using a library's compile and program options.
func (b *Binding) GetLibraryRules() string {
	library := bufcel.NewLibrary()

	env, err := cel.NewEnv(
		library.CompileOptions()...,
	)
	if err != nil {
		return ""
	}

	// Compile the CEL expression
	ast, issues := env.Compile("'1.2.3.0/24'.isIpPrefix(6)")
	if issues != nil && issues.Err() != nil {
		return ""
	}

	// Create the CEL program
	program, err := env.Program(ast, library.ProgramOptions()...)
	if err != nil {
		return ""
	}

	// eval, details, err := program.Eval("")
	_, _, err = program.Eval("")
	if err != nil {
		return ""
	}

	return ""
}
