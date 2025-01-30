package configurationv2alphapbint

import (
	"connectrpc.com/connect"

	nebulav1 "libs/partner/go/nebula/v1"
	specv2pb "libs/protobuf/go/protobuf/gen/platform/spec/v2"
	configurationv2alphapbconnect "libs/public/go/protobuf/gen/platform/configuration/v2alpha/configurationv2alphapbconnect"
)

// NewConfigurationServiceSpecClient creates a new instance of ConfigurationServiceClient using a SpecSettings configuration.
func NewConfigurationServiceSpecClient(config *specv2pb.SpecSettings, baseURL string, opts ...connect.ClientOption) *configurationv2alphapbconnect.ConfigurationServiceClient {
	nebula := nebulav1.Binding{}
	httpClient := nebula.GetMeshHTTPClient(config, baseURL)

	c := configurationv2alphapbconnect.NewConfigurationServiceClient(httpClient, baseURL, opts...)

	// time.Sleep(3 * time.Second)
	// fmt.Println("Woke up!")

	return &c
}
