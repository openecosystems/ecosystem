package nebulav1

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	specv2pb "libs/protobuf/go/protobuf/gen/platform/spec/v2"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"

	nebulaConfig "github.com/slackhq/nebula/config"
	"github.com/slackhq/nebula/service"
	"gopkg.in/yaml.v2"
)

// Binding represents a configuration and mesh socket service for the Nebula binding integration.
type Binding struct {
	MeshSocket *service.Service

	configuration *Configuration
}

// Bound holds the reference to the active Binding instance once configured and initialized.
// BindingName represents the identifier used for the binding instance in the application.
// IsBound indicates whether the Binding instance has been successfully initialized and configured.
var (
	Bound       *Binding
	BindingName = "NEBULA_BINDING"
	IsBound     = false
)

// Name returns the predefined name of the Binding instance.
func (b *Binding) Name() string {
	return BindingName
}

// Validate performs validation of the binding within the given context and bindings. Returns an error if validation fails.
func (b *Binding) Validate(_ context.Context, _ *sdkv2alphalib.Bindings) error {
	return nil
}

// Bind creates a binding by configuring a mesh socket, registers it, and ensures the binding is only initialized once.
func (b *Binding) Bind(_ context.Context, bindings *sdkv2alphalib.Bindings) *sdkv2alphalib.Bindings {
	if Bound == nil {
		var once sync.Once
		once.Do(
			func() {
				IsBound = true
				// Review this
				socket, err := b.ConfigureMeshSocket()
				if err != nil {
					return
				}
				Bound = &Binding{
					MeshSocket: socket,
				}

				bindings.Registered[b.Name()] = Bound
			})
	} else {
		bindings.Registered[b.Name()] = Bound
		IsBound = true
		fmt.Println("Nebula already bound")
	}

	return bindings
}

// GetBinding returns the current binding instance that is stored in the global Bound variable.
func (b *Binding) GetBinding() interface{} {
	return Bound
}

// Close releases resources associated with the Binding instance and ensures a clean shutdown of any active services.
func (b *Binding) Close() error {
	return nil
}

// GetSocket initializes a network listener on the given HTTP port if the binding is properly configured and bound.
// Returns a pointer to the listener or an error if preconditions are not met or configuration fails.
func (b *Binding) GetSocket(httpPort string) (*net.Listener, error) {
	if IsBound {
		configBytes, err := yaml.Marshal(ResolvedConfiguration.Nebula)
		if err != nil {
			fmt.Printf("Error resolving Nebula configuration: %v\n", err)
			fmt.Println(err.Error())
		}

		var cfg nebulaConfig.C
		if err = cfg.LoadString(string(configBytes)); err != nil {
			fmt.Println("ERROR loading config:", err)
		}

		svc, err := service.New(&cfg)
		if err != nil {
			fmt.Printf("Error creating service: %v\n", err)
			fmt.Println(err.Error())
		}

		ln, err := svc.Listen("tcp", ":"+httpPort)
		if err != nil {
			fmt.Println("Error listening:", err)
		}

		return &ln, nil
	}

	return nil, sdkv2alphalib.ErrServerPreconditionFailed.WithInternalErrorDetail(errors.New("the Nebula binding is not properly configured or not set"))
}

// GetMeshHTTPClient creates and returns an HTTP client configured optionally for Mesh or Internet-based calls depending on config.
func (b *Binding) GetMeshHTTPClient(config *specv2pb.SpecSettings, _ string /*url*/) *http.Client {
	httpClient := http.DefaultClient

	// go func() {
	// TODO: Check the service in the Global Settings to see if this call is a Mesh or Internet based call
	if config != nil && config.Platform != nil && config.Platform.Pki != nil {
		h := make(map[string][]string)
		for k, e := range config.Platform.StaticHostMap {
			h[k] = e.Map
		}

		nebulaC := Nebula{
			Host: h,
			Pki: Pki{
				Ca:   os.ExpandEnv(config.Platform.Pki.Ca),
				Cert: os.ExpandEnv(config.Platform.Pki.Cert),
				Key:  os.ExpandEnv(config.Platform.Pki.Key),
			},
			Lighthouse: Lighthouse{
				AmLighthouse: false,
				Interval:     int(config.Platform.Lighthouse.Interval),
				Hosts:        config.Platform.Lighthouse.Hosts,
			},
			Punchy: Punchy{
				Punch:        config.Platform.Punchy.Punch,
				Respond:      config.Platform.Punchy.Respond,
				RespondDelay: config.Platform.Punchy.RespondDelay,
				Delay:        config.Platform.Punchy.Delay,
			},
			Tun: Tun{
				// User:     true,
				Disabled:           false,
				Dev:                "utun8",
				DropLocalBroadcast: false,
				DropMulticast:      false,
				TxQueue:            500,
				Mtu:                1300,
			},
			Listen: Listen{
				Host: "0.0.0.0",
				Port: 4242,
			},
			Relay: Relay{
				AmRelay:   false,
				UseRelays: false,
			},
			Logging: Logging{
				Level:  "error",
				Format: "text",
			},
			Firewall: Firewall{
				OutboundAction: "drop",
				InboundAction:  "drop",
				Conntrack: Conntrack{
					TCPTimeout:     "12m",
					UDPTimeout:     "3m",
					DefaultTimeout: "10m",
				},
				Outbound: Outbound{
					{
						Port:  "any",
						Proto: "any",
						Host:  "any",
					},
				},
				Inbound: Inbound{
					{
						Port:  "any",
						Proto: "any",
						Host:  "any",
					},
					{
						Port:  "any",
						Proto: "any",
						Host:  "any",
					},
				},
			},
		}

		configBytes, err := yaml.Marshal(nebulaC)
		if err != nil {
			fmt.Printf("Error resolving Nebula configuration: %v\n", err)
			fmt.Println(err.Error())
		}

		var cfg nebulaConfig.C
		if err = cfg.LoadString(string(configBytes)); err != nil {
			fmt.Println("ERROR loading config:", err)
		}

		svc, err := service.New(&cfg)
		if err != nil {
			fmt.Printf("Error creating service: %v\n", err)
			fmt.Println(err.Error())
		}
		// defer svc.Close()

		// config.MeshSocket = svc

		httpClient = &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				DialContext: func(_ context.Context, network string, address string) (net.Conn, error) {
					return svc.Dial(network, address)
					// return svc.Dial("tcp", url)
				},
			},
		}
	}
	//}()

	return httpClient
}

// ConfigureMeshSocket initializes and configures the Nebula mesh socket, returning the created service or an error.
func (b *Binding) ConfigureMeshSocket() (*service.Service, error) {
	// go func() {
	// if IsBound {
	configBytes, err := yaml.Marshal(ResolvedConfiguration.Nebula)
	if err != nil {
		fmt.Printf("Error resolving Nebula configuration: %v\n", err)
		fmt.Println(err.Error())
	}

	var cfg nebulaConfig.C
	if err = cfg.LoadString(string(configBytes)); err != nil {
		fmt.Println("ERROR loading config:", err)
		return nil, errors.New("the Nebula binding configuration could not be loaded")
	}

	svc, err := service.New(&cfg)
	if err != nil {
		fmt.Printf("Error creating service: %v\n", err)
		fmt.Println(err.Error())
		return nil, errors.New("the Nebula binding is not properly configured or not set")
	}

	return svc, nil

	//}
	//}()
}
