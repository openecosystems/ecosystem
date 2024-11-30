package main

import (
	"apps/clients/public/cli/v2alpha/oeco/cmd"
	nebulav1ca "libs/partner/go/nebula/v1/ca"
	cliv2alphalib "libs/public/go/cli/v2alpha"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
	"runtime"
)

func main() {

	bounds := []sdkv2alphalib.Binding{
		//&zaploggerv1.Binding{},
		//&natsnodev2.Binding{SpecEventListeners: []natsnodev2.SpecEventListener{
		//
		//}},
		&nebulav1ca.Binding{},
	}

	cli := cliv2alphalib.NewCLI(bounds)

	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute(cli)
}
