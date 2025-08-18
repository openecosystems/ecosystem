package pushpinv1

import (
	sdkv2betalib "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta"

	"connectrpc.com/connect"
)

// ErrFailedToRunCommand indicates an internal server error that occurred while attempting to execute a command.
var ErrFailedToRunCommand sdkv2betalib.SpecErrorable = sdkv2betalib.NewSpecError(connect.CodeInternal, "SpecError on our side pushpin")
