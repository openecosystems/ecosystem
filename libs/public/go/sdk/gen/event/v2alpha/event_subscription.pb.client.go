// Code generated by protoc-gen-platform go/sdk. DO NOT EDIT.
// source: platform/event/v2alpha/event_subscription.proto

package eventv2alphapbsdk

import (
	"connectrpc.com/connect"
	specv2pb "libs/protobuf/go/protobuf/gen/platform/spec/v2"
	"libs/public/go/protobuf/gen/platform/event/v2alpha/eventv2alphapbconnect"
	"net/http"
)

func NewEventSubscriptionServiceSpecClient(config *specv2pb.SpecSettings, baseURL string, opts ...connect.ClientOption) *eventv2alphapbconnect.EventSubscriptionServiceClient {

	//httpClient := clientv2alphalib.GetMeshHttpClient(config, baseURL)
	httpClient := http.DefaultClient
	c := eventv2alphapbconnect.NewEventSubscriptionServiceClient(httpClient, baseURL, opts...)
	return &c
}
