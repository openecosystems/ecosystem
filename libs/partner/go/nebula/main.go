package nebulav1

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/slackhq/nebula/service"
	"gopkg.in/yaml.v2"

	specv2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/spec/v2"
	typev2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/type/v2"

	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"
	nebulaConfig "github.com/slackhq/nebula/config"
)

// Binding represents a configuration and mesh socket service for the Nebula binding integration.
type Binding struct {
	MeshSocket *service.Service

	configuration *Configuration
	cp            *sdkv2alphalib.CredentialProvider
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

				provider, err := sdkv2alphalib.NewCredentialProvider()
				if err != nil {
					return
				}
				b.cp = provider

				socket, err := b.ConfigureMeshSocket()
				if err != nil {
					return
				}

				Bound = &Binding{
					MeshSocket: socket,

					cp:            provider,
					configuration: b.configuration,
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
	if b.MeshSocket != nil {
		err := b.MeshSocket.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

// GetMeshListener initializes a network listener on the given HTTP port if the binding is properly configured and bound.
// Returns a pointer to the listener or an error if preconditions are not met or configuration fails.
func (b *Binding) GetMeshListener(endpoint string) (*net.Listener, error) {
	if IsBound {
		_, port, err := net.SplitHostPort(endpoint)
		if err != nil {
			fmt.Println("Error:", err)
			return nil, err
		}

		ln, err := b.MeshSocket.Listen("tcp", ":"+port)
		if err != nil {
			fmt.Println("Error listening:", err)
		}

		return &ln, nil
	}

	return nil, sdkv2alphalib.ErrServerPreconditionFailed.WithInternalErrorDetail(errors.New("the Nebula binding is not properly configured or not set"))
}

// GetMeshHTTPClient creates and returns an HTTP client configured optionally for Mesh or Internet-based calls depending on config.
func (b *Binding) GetMeshHTTPClient(config *specv2pb.Platform, _ string /*url*/) *http.Client {
	httpClient := http.DefaultClient

	b.cp, _ = sdkv2alphalib.NewCredentialProvider()
	socket, err := b.ConfigureMeshSocket()
	if err != nil {
		fmt.Println("Could not configure mesh socket", err)
		return nil
	}
	b.MeshSocket = socket

	if config != nil && config.Mesh != nil && config.Mesh.Enabled {
		httpClient = &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				DialContext: func(_ context.Context, network string, address string) (net.Conn, error) {
					return b.MeshSocket.Dial(network, address)
				},
			},
		}
	}

	return httpClient
}

// ConfigureMeshSocket initializes and configures the Nebula mesh socket, returning the created service or an error.
func (b *Binding) ConfigureMeshSocket() (*service.Service, error) {
	override := ""
	credentialType := typev2pb.CredentialType_CREDENTIAL_TYPE_MESH_ACCOUNT

	if b.configuration.Platform.Mesh == nil || b.configuration.Platform.Mesh.CredentialPath != "" {
		override = b.configuration.Platform.Mesh.CredentialPath
		credentialType = typev2pb.CredentialType_CREDENTIAL_TYPE_ACCOUNT_AUTHORITY
	}

	credential, err := b.cp.GetCredential(credentialType, override)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	h := make(map[string][]string)
	h[b.configuration.Platform.Mesh.DnsEndpoint] = b.configuration.Platform.DnsEndpoints

	udpHost := "0.0.0.0"
	udpPort := 0
	if b.configuration.Platform.Mesh.UdpEndpoint != "" {
		_udpPort := ""
		udpHost, _udpPort, err = net.SplitHostPort(b.configuration.Platform.Mesh.UdpEndpoint)
		if err != nil {
			fmt.Println("Error:", err)
			return nil, fmt.Errorf("invalid UDP endpoint: %s, %s", b.configuration.Platform.Mesh.UdpEndpoint, err)
		}
		udpPort, err = strconv.Atoi(_udpPort)
		if err != nil {
			fmt.Println("Error:", err)
			return nil, fmt.Errorf("cannot convert UDP port to string: %s, %s", b.configuration.Platform.Mesh.UdpEndpoint, err)
		}
	}

	nebulaC := Nebula{
		Host: h,
		Pki: Pki{
			Ca:   credential.AaCertX509,
			Cert: credential.CertX509,
			Key:  credential.PrivateKey,
		},
		Lighthouse: Lighthouse{
			AmLighthouse: false,
			Interval:     60,
			Hosts:        []string{b.configuration.Platform.Mesh.DnsEndpoint},
		},
		Punchy: Punchy{
			Punch:        b.configuration.Platform.Mesh.Punchy,
			Respond:      b.configuration.Platform.Mesh.Punchy,
			RespondDelay: "5s",
			Delay:        "1s",
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
			Host: udpHost,
			Port: udpPort,
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

	if b.configuration.Platform.Mesh.DnsServer {
		nebulaC.Lighthouse.AmLighthouse = true
		nebulaC.Lighthouse.Hosts = []string{}
		nebulaC.Host = nil
		nebulaC.Listen.Port = 4242
	}

	if b.configuration.App.Debug {
		nebulaC.Logging.Level = "info"
	}

	if b.configuration.App.Verbose {
		nebulaC.Logging.Level = "debug"
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

	return svc, nil
}
