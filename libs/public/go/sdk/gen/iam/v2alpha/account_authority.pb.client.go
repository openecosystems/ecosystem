// Code generated by protoc-gen-platform go/sdk. DO NOT EDIT.
// source: platform/iam/v2alpha/account_authority.proto

package iamv2alphapbsdk

import (
	"connectrpc.com/connect"

	nebulav1 "github.com/openecosystems/ecosystem/libs/partner/go/nebula"
	specv2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/spec/v2"
	iamv2alphapbconnect "github.com/openecosystems/ecosystem/libs/public/go/protobuf/gen/platform/iam/v2alpha/iamv2alphapbconnect"
)

func NewAccountAuthorityServiceSpecClient(config *specv2pb.Platform, baseURL string, opts ...connect.ClientOption) *iamv2alphapbconnect.AccountAuthorityServiceClient {
	nebula := nebulav1.Binding{}
	httpClient := nebula.GetMeshHTTPClient(config, baseURL)
	c := iamv2alphapbconnect.NewAccountAuthorityServiceClient(httpClient, baseURL, opts...)
	return &c
}
