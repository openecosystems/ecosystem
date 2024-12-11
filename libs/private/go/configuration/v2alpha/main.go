package configurationv2alphalib

import (
	"context"
	"fmt"
	"sync"

	"github.com/nats-io/nats.go/jetstream"
	"libs/partner/go/nats/v2"
	"libs/public/go/sdk/v2alpha"
)

// Binding struct that holds binding specific fields
type Binding struct {
	ConfigStore                  *jetstream.KeyValue
	AdaptiveConfigurationControl *AdaptiveConfigurationControl

	configuration *Configuration
}

var (
	Bound       *Binding
	BindingName = "CONFIGURATION_LIB_BINDING"
)

func (b *Binding) Name() string {
	return BindingName
}

func (b *Binding) Validate(_ context.Context, _ *sdkv2alphalib.Bindings) error {
	if natsnodev2.Bound == nil {
		fmt.Println("Please bind the Nats Node module to use this binding")
		panic("")
	}

	return nil
}

func (b *Binding) Bind(ctx context.Context, bindings *sdkv2alphalib.Bindings) *sdkv2alphalib.Bindings {
	if Bound == nil {
		var once sync.Once
		once.Do(
			func() {
				bn := sdkv2alphalib.ResolvedConfiguration.App.EnvironmentName + "-configuration"

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
				}
				bindings.Registered[b.Name()] = Bound
			})
	} else {
		bindings.Registered[b.Name()] = Bound
		fmt.Println("Configuration Library already bound")
	}

	return bindings
}

func (b *Binding) GetBinding() interface{} {
	return Bound
}

func (b *Binding) Close() error {
	fmt.Println("Closing the Configuration Library Binding")
	return nil
}

// Listen for file system changes
//     Update RAM with Mutex Lock
// Periodically check (Pull) for updates from ACC Server
// Listen for push changes using a PushPin listener
