package pushpinv1

import (
	typev2pb "libs/protobuf/go/protobuf/gen/platform/type/v2"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"

	"connectrpc.com/connect"
)

// ErrFailedToRunCommand indicates an internal server error that occurred while attempting to execute a command.
var ErrFailedToRunCommand sdkv2alphalib.SpecError = &sdkv2alphalib.Error{SpecApiErr: &sdkv2alphalib.SpecAPIError{Error: *sdkv2alphalib.NewConnectError(connect.CodeInternal, &FailedToRunCommandErrorDetail, "Error on our side pushpin")}}

// FailedToRunCommandErrorDetail represents a predefined SpecErrorDetail with a user-friendly error message for failures.
var FailedToRunCommandErrorDetail = typev2pb.SpecErrorDetail{UserMessage: "Looks like there is an error on our side. We have reported it to the team."}
