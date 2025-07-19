package sendgridlistv1

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"

	sdkv2betalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2beta"

	"github.com/apex/log"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
)

// Binding represents a structure containing a client for interacting with SendGrid services.
type Binding struct {
	Client *ClientWithResponses
}

// Bound is a singleton instance of Binding used to manage connection to the SendGrid API.
// BindingName is the constant name used to identify the Binding type for SendGrid lists.
// SendGridAPIEndpoint is the base URL for SendGrid API requests.
// SendGridAPIKey retrieves the SendGrid API key from the environment variable "SENDGRID_API_KEY".
var (
	Bound               *Binding
	BindingName         = "SEND_GRID_LISTS_BINDING"
	SendGridAPIEndpoint = "https://api.sendgrid.com"
	SendGridAPIKey      = os.Getenv("SENDGRID_API_KEY")
)

// Name returns the name of the binding as a string identifier.
func (b *Binding) Name() string {
	return BindingName
}

// Validate checks if the required SendGrid API key is set in the environment variables. Returns an error if not found.
func (b *Binding) Validate(_ context.Context, _ *sdkv2betalib.Bindings) error {
	if SendGridAPIKey == "" {
		return errors.New("SENDGRID_API_KEY environment variable is required")
	}

	return nil
}

// Bind initializes and registers the SendGrid API binding if not already bound, or uses the existing binding instance.
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

// GetBinding returns the currently bound instance of the Binding or nil if not initialized.
func (b *Binding) GetBinding() interface{} {
	return Bound
}

// Close gracefully shuts down the Sendgrid List API Binding and releases any associated resources.
func (b *Binding) Close() error {
	fmt.Println("Closing Sendgrid List API Binding")
	return nil
}
