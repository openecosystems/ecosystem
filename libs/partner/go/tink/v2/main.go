package tinkv2

import (
	"context"
	"fmt"
	"github.com/tink-crypto/tink-go-gcpkms/v2/integration/gcpkms"
	"github.com/tink-crypto/tink-go/v2/core/registry"
	"github.com/tink-crypto/tink-go/v2/tink"
	"libs/public/go/sdk/v2alpha"
	"sync"
)

type Binding struct {
	KMSClient registry.KMSClient
	KEAD      tink.AEAD

	configuration *Configuration
}

var (
	Bound       *Binding
	BindingName = "TINK_BINDING"
)

func (b *Binding) Name() string {
	return BindingName
}

func (b *Binding) Validate(_ context.Context, _ *sdkv2alphalib.Bindings) error {

	// Verify any log requirements

	return nil
}

func (b *Binding) Bind(ctx context.Context, bindings *sdkv2alphalib.Bindings) *sdkv2alphalib.Bindings {
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
				}
				bindings.Registered[b.Name()] = Bound
			})
	} else {
		bindings.Registered[b.Name()] = Bound
		fmt.Println("Tink already bound")
	}

	return bindings
}

func (b *Binding) GetBinding() interface{} {
	return Bound
}

func (b *Binding) Close() error {
	fmt.Println("Closing the Tink Binding")
	return nil
}
