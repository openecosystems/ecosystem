package pushpinv1

import (
	"context"
	"errors"
	"fmt"
	nebulaConfig "github.com/slackhq/nebula/config"
	"github.com/slackhq/nebula/service"
	"gopkg.in/yaml.v2"
	"libs/public/go/sdk/v2alpha"
	"net"
	"sync"
)

type Binding struct {
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

	// Verify any requirements

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
