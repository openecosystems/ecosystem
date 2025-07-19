package sendgridcontactsv1

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"

	sdkv2betalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2beta"

	"github.com/apex/log"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
)

// Binding represents a structure that holds a reference to a ClientWithResponses instance.
type Binding struct {
	Client *ClientWithResponses
}

// Bound represents the global singleton instance of the SendGrid binding.
// BindingName is the constant name used to identify the SendGrid binding.
// SendGridAPIEndpoint defines the base endpoint for the SendGrid API.
// SendGridAPIKey retrieves the API key from the environment variable SENDGRID_API_KEY.
var (
	Bound               *Binding
	BindingName         = "SEND_GRID_BINDING"
	SendGridAPIEndpoint = "https://api.sendgrid.com"
	SendGridAPIKey      = os.Getenv("SENDGRID_API_KEY")
)

// Name returns the unique identifier name for the Binding instance.
func (b *Binding) Name() string {
	return BindingName
}

// Validate checks the binding's configuration for correctness and returns an error if any issues are found.
func (b *Binding) Validate(_ context.Context, _ *sdkv2betalib.Bindings) error {
	return nil
}

// Bind initializes and registers the binding with SendGrid's API client, ensuring it only binds once globally.
func (b *Binding) Bind(_ context.Context, bindings *sdkv2betalib.Bindings) *sdkv2betalib.Bindings {
	if Bound == nil {
		var once sync.Once
		once.Do(
			func() {
				hc := http.Client{}

				auth, err := securityprovider.NewSecurityProviderBearerToken(SendGridAPIKey)
				if err != nil {
					log.Fatal(err.Error())
				}

				client, err := NewClientWithResponses(SendGridAPIEndpoint, WithHTTPClient(&hc), WithRequestEditorFn(auth.Intercept))
				if err != nil {
					log.Fatal(err.Error())
				}

				if client == nil {
					panic(err)
				}

				Bound = &Binding{
					Client: client,
				}
				bindings.Registered[b.Name()] = Bound
			})
	} else {
		bindings.Registered[b.Name()] = Bound
		fmt.Println("Sendgrid already bound")
	}

	return bindings
}

// GetBinding returns the currently bound Binding instance.
func (b *Binding) GetBinding() interface{} {
	return Bound
}

// Close releases resources associated with the Sendgrid Contacts API Binding and performs necessary cleanup operations.
func (b *Binding) Close() error {
	fmt.Println("Closing Sendgrid Contacts API Binding")
	return nil
}
