package nebulav1ca

import (
	typev2pb "libs/protobuf/go/protobuf/gen/platform/type/v2"

	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"

	"connectrpc.com/connect"
)

// ErrFailedToRunCommand represents a custom error indicating a failure to execute a command, typically due to internal issues.
var ErrFailedToRunCommand sdkv2alphalib.SpecError = &sdkv2alphalib.Error{SpecApiErr: &sdkv2alphalib.SpecAPIError{Error: *sdkv2alphalib.NewConnectError(connect.CodeInternal, &FailedToRunCommandErrorDetail, "Error on our side with nebula")}}

// FailedToRunCommandErrorDetail is a predefined SpecErrorDetail indicating an internal error during command execution.
var FailedToRunCommandErrorDetail typev2pb.SpecErrorDetail = typev2pb.SpecErrorDetail{UserMessage: "Looks like there is an error on our side. We have reported it to the team."}
