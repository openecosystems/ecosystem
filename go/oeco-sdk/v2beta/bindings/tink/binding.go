package tinkv1

import (
	"context"
	"fmt"
	"sync"

	sdkv2betalib "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta"

	"github.com/tink-crypto/tink-go-gcpkms/v2/integration/gcpkms"
	"github.com/tink-crypto/tink-go/v2/core/registry"
	"github.com/tink-crypto/tink-go/v2/tink"
)

// Binding struct that holds binding specific fields
type Binding struct {
	KMSClient registry.KMSClient
	KEAD      tink.AEAD

	configuration *Configuration
}

// Bound is a global reference to the Binding instance used for cryptographic operations.
// BindingName is the identifier for the Tink binding.
var (
	Bound       *Binding
	BindingName = "TINK_BINDING"
)

// Name returns the name identifier of the binding, defined as the constant `BindingName`.
func (b *Binding) Name() string {
	return BindingName
}

// Validate checks the validity of the current binding's setup and ensures any required logging conditions are met.
func (b *Binding) Validate(_ context.Context, _ *sdkv2betalib.Bindings) error {
	// Verify any log requirements

	return nil
}

// Bind initializes and registers the Binding with provided configurations and returns the updated Bindings instance.
func (b *Binding) Bind(ctx context.Context, bindings *sdkv2betalib.Bindings) *sdkv2betalib.Bindings {
	if Bound == nil {
		var once sync.Once
		once.Do(
			func() {
				keyURI := "gcp-kms://projects/*/locations/*/keyRings/*/cryptoKeys/*"

				// Get a KEK (key encryption key) AEAD.
				client, err := gcpkms.NewClientWithOptions(ctx, keyURI)
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
				kekAEAD, err := client.GetAEAD(keyURI)
				if err != nil {
					fmt.Println(err)
					panic(err)
				}

				Bound = &Binding{
					KMSClient: client,
					KEAD:      kekAEAD,

					configuration: b.configuration,
				}
				bindings.Registered[b.Name()] = Bound
			})
	} else {
		bindings.Registered[b.Name()] = Bound
		fmt.Println("Tink already bound")
	}

	return bindings
}

// GetBinding retrieves the currently initialized Binding instance stored in the global Bound variable.
func (b *Binding) GetBinding() interface{} {
	return Bound
}

// Close releases any resources held by the Binding and performs necessary cleanup operations.
func (b *Binding) Close() error {
	fmt.Println("Closing the Tink Binding")
	return nil
}
