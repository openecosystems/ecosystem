package serverv2alphalib

import (
	typev2pb "libs/protobuf/go/protobuf/gen/platform/type/v2"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"

	"connectrpc.com/connect"
)

// ErrServerInternal represents a server-side internal error, typically indicating an issue that has been reported for resolution.
var ErrServerInternal sdkv2alphalib.SpecError = &sdkv2alphalib.Error{SpecApiErr: &sdkv2alphalib.SpecAPIError{Error: *sdkv2alphalib.NewConnectError(connect.CodeInternal, &SpecServerInternalErrorDetail, "Error on our side serverlib")}}

// SpecServerInternalErrorDetail represents a predefined error detail for internal server errors with a user-friendly message.
var SpecServerInternalErrorDetail = typev2pb.SpecErrorDetail{UserMessage: "Looks like there is an error on our side. We have reported it to the team."}
