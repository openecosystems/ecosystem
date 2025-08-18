package nebulav1ca

import (
	sdkv2betalib "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta"

	"connectrpc.com/connect"
)

// ErrFailedToRunCommand represents a custom error indicating a failure to execute a command, typically due to internal issues.
var ErrFailedToRunCommand sdkv2betalib.SpecErrorable = sdkv2betalib.NewSpecError(connect.CodeInternal, "SpecError on our side with nebula")
