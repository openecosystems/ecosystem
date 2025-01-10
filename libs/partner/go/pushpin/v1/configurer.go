package pushpinv1

import sdkv2alphalib "libs/public/go/sdk/v2alpha"

// ResolvedConfiguration holds the configuration for this binding
var ResolvedConfiguration *Configuration

type Configuration struct{}

func (b *Binding) ResolveConfiguration() {
	var c Configuration
	dc := b.GetDefaultConfiguration().(Configuration)
	sdkv2alphalib.Resolve(&c, dc)
	b.configuration = &c
	ResolvedConfiguration = &c
}

func (b *Binding) ValidateConfiguration() error {
	return nil
}

func (b *Binding) GetDefaultConfiguration() interface{} {
	return Configuration{}
}
