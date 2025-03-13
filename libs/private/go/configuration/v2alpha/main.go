package configurationv2alphalib

import (
	"context"
	"fmt"
	"sync"

	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"
	natsnodev2 "libs/partner/go/nats/v2"

	"github.com/nats-io/nats.go/jetstream"
)

// Binding is the central struct for managing configuration storage and adaptive controls within the application context.
type Binding struct {
	ConfigStore                  *jetstream.KeyValue
	AdaptiveConfigurationControl *AdaptiveConfigurationControl

	configuration *Configuration
}

// Bound represents the singleton instance of Binding for managing and accessing configuration-related functionality.
// BindingName is the unique identifier assigned to the configuration library binding.
var (

	// Bound represents the global instance of the Binding used to manage and access configuration-related functionality.
	Bound *Binding

	// BindingName is a constant that represents the unique name for the configuration library binding.
	BindingName = "CONFIGURATION_LIB_BINDING"
)

// Name returns the unique name of the Binding, which is used for registry and identification purposes.
func (b *Binding) Name() string {
	return BindingName
}

// Validate checks if the Nats Node module is bound, ensuring it is required for the binding to function properly.
func (b *Binding) Validate(_ context.Context, _ *sdkv2alphalib.Bindings) error {
	if natsnodev2.Bound == nil {
		fmt.Println("Please bind the Nats Node module to use this binding")
		panic("")
	}

	return nil
}

// Bind initializes and registers the configuration binding in the provided Bindings object and ensures it is only bound once.
func (b *Binding) Bind(ctx context.Context, bindings *sdkv2alphalib.Bindings) *sdkv2alphalib.Bindings {
	if Bound == nil {
		var once sync.Once
		once.Do(
			func() {
				bn := b.configuration.App.EnvironmentName + "-configuration"

				js := *natsnodev2.Bound.JetStream
				cs, err := js.CreateOrUpdateKeyValue(ctx, jetstream.KeyValueConfig{
					Bucket:       bn,
					Description:  "",
					MaxValueSize: -1,
					History:      1,
					MaxBytes:     -1,
					Storage:      0, // File storage
					Replicas:     1,
					Placement: &jetstream.Placement{
						Cluster: "",
						Tags:    nil,
					},
					Compression: false,
				})
				if err != nil {
					return
				}

				Bound = &Binding{
					ConfigStore:                  &cs,
					AdaptiveConfigurationControl: NewAdaptiveConfigurationControl(&cs),

					configuration: b.configuration,
				}
				bindings.Registered[b.Name()] = Bound
			})
	} else {
		bindings.Registered[b.Name()] = Bound
		fmt.Println("Configuration Library already bound")
	}

	return bindings
}

// GetBinding returns the global Binding instance used for configuration-related functionality management.
func (b *Binding) GetBinding() interface{} {
	return Bound
}

// Close safely releases resources and performs cleanup for the Configuration Library Binding.
func (b *Binding) Close() error {
	fmt.Println("Closing the Configuration Library Binding")
	return nil
}

// Listen for file system changes
//     Update RAM with Mutex Lock
// Periodically check (Pull) for updates from ACC Server
// Listen for push changes using a PushPin listener
