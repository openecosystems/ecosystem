package nebulav1ca

import (
	"connectrpc.com/connect"
	typev2pb "libs/protobuf/go/protobuf/gen/platform/type/v2"
	"libs/public/go/sdk/v2alpha"
)

var ErrFailedToRunCommand sdkv2alphalib.SpecError = &sdkv2alphalib.Error{SpecApiErr: &sdkv2alphalib.SpecApiError{Error: *sdkv2alphalib.NewConnectError(connect.CodeInternal, &FailedToRunCommandErrorDetail, "Error on our side")}}

var FailedToRunCommandErrorDetail typev2pb.SpecErrorDetail = typev2pb.SpecErrorDetail{UserMessage: "Looks like there is an error on our side. We have reported it to the team."}
