package nebulav1ca

import (
	typev2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/type/v2"

	sdkv2betalib "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta"

	"connectrpc.com/connect"
)

// ErrFailedToRunCommand represents a custom error indicating a failure to execute a command, typically due to internal issues.
var ErrFailedToRunCommand sdkv2betalib.SpecError = &sdkv2betalib.Error{SpecApiErr: &sdkv2betalib.SpecAPIError{Error: *sdkv2betalib.NewConnectError(connect.CodeInternal, &FailedToRunCommandErrorDetail, "Error on our side with nebula")}}

// FailedToRunCommandErrorDetail is a predefined SpecErrorDetail indicating an internal error during command execution.
var FailedToRunCommandErrorDetail typev2pb.SpecErrorDetail = typev2pb.SpecErrorDetail{UserMessage: "Looks like there is an error on our side. We have reported it to the team."}
