package sendgridlistv3

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/apex/log"
	"github.com/oapi-codegen/oapi-codegen/v2/pkg/securityprovider"
	"libs/public/go/sdk/v2alpha"
)

type Binding struct {
	Client *ClientWithResponses
}

var (
	Bound               *Binding
	BindingName         = "SEND_GRID_LISTS_BINDING"
	SendGridAPIEndpoint = "https://api.sendgrid.com"
	SendGridAPIKey      = os.Getenv("SENDGRID_API_KEY")
)

func (b *Binding) Name() string {
	return BindingName
}

func (b *Binding) Validate(_ context.Context, c *sdkv2alphalib.Configuration, _ *sdkv2alphalib.Bindings) error {
	if SendGridAPIKey == "" {
		return errors.New("SENDGRID_API_KEY environment variable is required")
	}

	return nil
}

func (b *Binding) Bind(ctx context.Context, c *sdkv2alphalib.Configuration, bindings *sdkv2alphalib.Bindings) *sdkv2alphalib.Bindings {
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
					log.Fatal("Could not connect to the Send Grid Client ")
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

func (b *Binding) GetBinding() interface{} {
	return Bound
}

func (b *Binding) Close() error {
	fmt.Println("Closing Sendgrid List API Binding")
	return nil
}
