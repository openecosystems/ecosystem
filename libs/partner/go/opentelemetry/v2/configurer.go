package opentelemetryv2

import (
	"fmt"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

var (
	ResolvedConfiguration *Configuration
)

type Configuration struct {
	Opentelemetry struct {
		TraceProviderEnabled  bool
		MeterProviderEnabled  bool
		LoggerProviderEnabled bool
	}
}

func (b *Binding) ResolveConfiguration() {
	var c Configuration
	dc := b.GetDefaultConfiguration().(Configuration)
	sdkv2alphalib.Resolve(&c, dc)
	b.configuration = &c
	ResolvedConfiguration = &c
}

func (b *Binding) ValidateConfiguration() error {

	if !b.configuration.Opentelemetry.TraceProviderEnabled {
		fmt.Println("warn: opentelemtry trace is disabled. This may cause errors if you have other bindings that depend on it. Binding dependency management is on the roadmap.")
	}

	if !b.configuration.Opentelemetry.MeterProviderEnabled {
		//fmt.Println("warn: opentelemtry meter is disabled. This may cause errors if you have other bindings that depend on it. Binding dependency management is on the roadmap.")
	}

	if !b.configuration.Opentelemetry.LoggerProviderEnabled {
		//fmt.Println("warn: opentelemtry logger is disabled. This may cause errors if you have other bindings that depend on it. Binding dependency management is on the roadmap.")
	}
	return nil
}

func (b *Binding) GetDefaultConfiguration() interface{} {

	return Configuration{
		Opentelemetry: struct {
			TraceProviderEnabled  bool
			MeterProviderEnabled  bool
			LoggerProviderEnabled bool
		}{
			TraceProviderEnabled:  false,
			MeterProviderEnabled:  false,
			LoggerProviderEnabled: false,
		},
	}
}
