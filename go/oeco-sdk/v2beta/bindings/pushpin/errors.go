package pushpinv1

import (
	typev2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/type/v2"

	sdkv2betalib "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta"

	"connectrpc.com/connect"
)

// ErrFailedToRunCommand indicates an internal server error that occurred while attempting to execute a command.
var ErrFailedToRunCommand sdkv2betalib.SpecErrorable = sdkv2betalib.NewSpecError(connect.CodeInternal, "SpecError on our side pushpin")

// FailedToRunCommandErrorDetail represents a predefined SpecErrorDetail with a user-friendly error message for failures.
var FailedToRunCommandErrorDetail = typev2pb.SpecErrorDetail{UserMessage: "Looks like there is an error on our side. We have reported it to the team."}
