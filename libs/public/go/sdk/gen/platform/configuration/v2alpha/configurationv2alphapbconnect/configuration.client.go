// Code generated by protoc-gen-platform go/sdk. DO NOT EDIT.
// source: platform/configuration/v2alpha/configuration.proto

package configurationv2alphapbconnect

import (
	//"net/http"

	"connectrpc.com/connect"

	nebulav1 "github.com/openecosystems/ecosystem/libs/partner/go/nebula"
	specv2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/spec/v2"
)

func NewConfigurationServiceSpecClient(config *specv2pb.Platform, baseURL string, opts ...connect.ClientOption) *ConfigurationServiceClient {
	nebula := nebulav1.Binding{}
	httpClient := nebula.GetMeshHTTPClient(config, baseURL)
	//httpClient := http.DefaultClient

	c := NewConfigurationServiceClient(httpClient, baseURL, opts...)
	return &c
}
