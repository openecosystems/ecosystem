package nebulav1

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"sync"

	nebulav1 "libs/partner/go/pushpin/v1"
	"libs/protobuf/go/protobuf/gen/platform/spec/v2"
	"libs/public/go/sdk/v2alpha"

	nebulaConfig "github.com/slackhq/nebula/config"
	"github.com/slackhq/nebula/service"
	"gopkg.in/yaml.v2"
)

// Binding struct that holds binding specific fields
type Binding struct {
	MeshSocket *service.Service

	configuration *Configuration
}

var (
	Bound       *Binding
	BindingName = "NEBULA_BINDING"
	IsBound     = false
)

func (b *Binding) Name() string {
	return BindingName
}

func (b *Binding) Validate(_ context.Context, _ *sdkv2alphalib.Bindings) error {
	return nil
}

func (b *Binding) Bind(_ context.Context, bindings *sdkv2alphalib.Bindings) *sdkv2alphalib.Bindings {
	if Bound == nil {
		var once sync.Once
		once.Do(
			func() {
				Bound = &Binding{}

				bindings.Registered[b.Name()] = Bound
				IsBound = true
			})
	} else {
		bindings.Registered[b.Name()] = Bound
		IsBound = true
		fmt.Println("Nebula already bound")
	}

	return bindings
}

func (b *Binding) GetBinding() interface{} {
	return Bound
}

func (b *Binding) Close() error {
	return nil
}

func (b *Binding) GetSocket(httpPort string) (*net.Listener, error) {
	if IsBound {

		configBytes, err := yaml.Marshal(ResolvedConfiguration.Nebula)
		if err != nil {
			fmt.Printf("Error resolving Nebula configuration: %v\n", err)
			fmt.Printf(err.Error())
		}

		var cfg nebulaConfig.C
		if err = cfg.LoadString(string(configBytes)); err != nil {
			fmt.Println("ERROR loading config:", err)
		}

		svc, err := service.New(&cfg)
		if err != nil {
			fmt.Printf("Error creating service: %v\n", err)
			fmt.Printf(err.Error())
		}

		fmt.Println(fmt.Sprintf(":%d", httpPort))
		ln, err := svc.Listen("tcp", fmt.Sprintf(":%d", httpPort))
		if err != nil {
			fmt.Println("Error listening:", err)
		}

		return &ln, nil

	}

	return nil, sdkv2alphalib.ErrServerPreconditionFailed.WithInternalErrorDetail(errors.New("the Nebula binding is not properly configured or not set"))
}

func (b *Binding) GetMeshHttpClient(config *specv2pb.SpecSettings, url string) *http.Client {
	httpClient := http.DefaultClient

	go func() {
		// TODO: Check the service in the Global Settings to see if this call is a Mesh or Internet based call
		if config != nil && config.Platform != nil && config.Platform.Pki != nil {

			h := make(map[string][]string)
			for k, e := range config.Platform.StaticHostMap {
				h[k] = e.Map
			}

			nebulaC := Nebula{
				Host: h,
				Pki: Pki{
					Ca:   config.Platform.Pki.Ca,
					Cert: config.Platform.Pki.Cert,
					Key:  config.Platform.Pki.Key,
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
					User: true,
				},
				Logging: Logging{
					Level:  "info",
					Format: "",
				},
				Firewall: Firewall{
					OutboundAction: "",
					InboundAction:  "",
					Conntrack: Conntrack{
						TCPTimeout:     "",
						UDPTimeout:     "",
						DefaultTimeout: "",
					},
					Outbound: nil,
					Inbound:  nil,
				},
			}

			configBytes, err := yaml.Marshal(nebulaC)
			if err != nil {
				fmt.Printf("Error resolving Nebula configuration: %v\n", err)
				fmt.Printf(err.Error())
			}

			var cfg nebulaConfig.C
			if err = cfg.LoadString(string(configBytes)); err != nil {
				fmt.Println("ERROR loading config:", err)
			}

			svc, err := service.New(&cfg)
			if err != nil {
				fmt.Printf("Error creating service: %v\n", err)
				fmt.Printf(err.Error())
			}
			// defer svc.Close()

			// config.MeshSocket = svc

			httpClient = &http.Client{
				Transport: &http.Transport{
					DialContext: func(ctx context.Context, network string, addr string) (net.Conn, error) {
						return svc.Dial("tcp", url)
					},
				},
			}
		}
	}()

	return httpClient
}

func (b *Binding) ConfigureMeshSocket() {
	go func() {
		if nebulav1.IsBound {

			configBytes, err := yaml.Marshal(nebulav1.ResolvedConfiguration.Nebula)
			if err != nil {
				fmt.Printf("Error resolving Nebula configuration: %v\n", err)
				fmt.Printf(err.Error())
			}

			var cfg nebulaConfig.C
			if err = cfg.LoadString(string(configBytes)); err != nil {
				fmt.Println("ERROR loading config:", err)
			}

			svc, err := service.New(&cfg)
			if err != nil {
				fmt.Printf("Error creating service: %v\n", err)
				fmt.Printf(err.Error())
			}

			b.MeshSocket = svc

		}
	}()
}
