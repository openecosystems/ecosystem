package serverv2alphalib

import (
	"connectrpc.com/connect"
	"libs/protobuf/go/protobuf/gen/platform/type/v2"
	"libs/public/go/sdk/v2alpha"
)

var (
	ErrServerInternal sdkv2alphalib.SpecError = &sdkv2alphalib.Error{SpecApiErr: &sdkv2alphalib.SpecApiError{Error: *sdkv2alphalib.NewConnectError(connect.CodeInternal, &SpecServerInternalErrorDetail, "Error on our side")}}
)

var (
	SpecServerInternalErrorDetail = typev2pb.SpecErrorDetail{UserMessage: "Looks like there is an error on our side. We have reported it to the team."}
)
