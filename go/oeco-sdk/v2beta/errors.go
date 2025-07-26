package sdkv2betalib

import (
	typev2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/type/v2"

	"connectrpc.com/connect"
)

// Sentinel instances for fast error checks

// ErrServerInternal represents an internal server error, indicating an issue on the server's side.
// ErrServerAuthentication represents an authentication error due to invalid or incorrect credentials.
// ErrServerRequest represents a client-side error caused by an invalid request.
// ErrServerAlreadyExists represents an error indicating that the resource already exists.
// ErrServerNotSupported represents an error indicating that the requested operation is not supported.
// ErrServerPreconditionFailed represents an error indicating that preconditions for the operation have failed.
// ErrServerUnimplemented represents an error indicating that the requested API call is not implemented.
// ErrServerUnknownResource represents an error indicating that the requested resource could not be found.
var (
	ErrServerInternal           SpecErrorable = NewSpecError(connect.CodeInternal, "An Error occurred on our side")
	ErrServerAuthentication     SpecErrorable = NewSpecError(connect.CodeUnauthenticated, "Invalid or incorrect credentials")
	ErrServerRequest            SpecErrorable = NewSpecError(connect.CodeInvalidArgument, "There was an error with your request")
	ErrServerAlreadyExists      SpecErrorable = NewSpecError(connect.CodeAlreadyExists, "Resource already exist")
	ErrServerNotSupported       SpecErrorable = NewSpecError(connect.CodeUnimplemented, "Not Supported")
	ErrServerPreconditionFailed SpecErrorable = NewSpecError(connect.CodeFailedPrecondition, "Preconditions failed")
	ErrServerUnimplemented      SpecErrorable = NewSpecError(connect.CodeUnimplemented, "API call unimplemented")
	ErrServerUnknownResource    SpecErrorable = NewSpecError(connect.CodeNotFound, "Unknown resource")
	ErrServerCanceled           SpecErrorable = NewSpecError(connect.CodeCanceled, "Request was cancelled")
	ErrServerUnknown            SpecErrorable = NewSpecError(connect.CodeUnknown, "An unknown error occurred")
	ErrServerDeadlineExceeded   SpecErrorable = NewSpecError(connect.CodeDeadlineExceeded, "Request deadline exceeded")
	ErrServerPermissionDenied   SpecErrorable = NewSpecError(connect.CodePermissionDenied, "Permission denied")
	ErrServerResourceExhausted  SpecErrorable = NewSpecError(connect.CodeResourceExhausted, "Resource limit exceeded")
	ErrServerOutOfRange         SpecErrorable = NewSpecError(connect.CodeOutOfRange, "Value out of range")
	ErrServerUnimplementedAlt   SpecErrorable = NewSpecError(connect.CodeUnimplemented, "This method is not implemented")
	ErrServerUnavailable        SpecErrorable = NewSpecError(connect.CodeUnavailable, "Service temporarily unavailable")
	ErrServerDataLoss           SpecErrorable = NewSpecError(connect.CodeDataLoss, "Data loss or corruption detected")
	ErrServerAborted            SpecErrorable = NewSpecError(connect.CodeAborted, "Operation aborted")
	ErrServerUnauthenticated    SpecErrorable = NewSpecError(connect.CodeUnauthenticated, "Authentication required or failed")
	ErrServerConflict           SpecErrorable = NewSpecError(connect.CodeAlreadyExists, "Conflict: resource already exists")
	ErrServerNotFound           SpecErrorable = NewSpecError(connect.CodeNotFound, "Resource not found")
)

// SpecServerInternalErrorDetail represents a server-side error with a default user message indicating an internal issue.
// SpecServerAuthenticationErrorDetail represents an authentication error with a user message prompting credential verification.
// SpecServerRequestErrorDetail represents a request-related error with a user message indicating a problem with the request.
// SpecServerAlreadyExistsErrorDetail represents a conflict error with a user message indicating the resource already exists.
// SpecServerNotSupportedErrorDetail represents a feature not supported error with a default user message.
// SpecServerPreconditionFailedErrorDetail represents a precondition failure error with an appropriate user message.
// SpecServerUnimplementedErrorDetail represents an unimplemented feature error with a message indicating ongoing development.
// SpecServerUnknownResourceErrorDetail represents an unknown resource error with a user message indicating the resource cannot be found.
var (
	SpecServerInternalErrorDetail           = typev2pb.SpecErrorDetail{UserMessage: "Looks like there is an error on our side. We have reported it to the team."}
	SpecServerAuthenticationErrorDetail     = typev2pb.SpecErrorDetail{UserMessage: "Please check your username and password or API Key then try again"}
	SpecServerRequestErrorDetail            = typev2pb.SpecErrorDetail{UserMessage: "There was an error with your request"}
	SpecServerAlreadyExistsErrorDetail      = typev2pb.SpecErrorDetail{UserMessage: "The resource you requested is already in use"}
	SpecServerNotSupportedErrorDetail       = typev2pb.SpecErrorDetail{UserMessage: "Not yet supported"}
	SpecServerPreconditionFailedErrorDetail = typev2pb.SpecErrorDetail{UserMessage: "Certain preconditions have failed"}
	SpecServerUnimplementedErrorDetail      = typev2pb.SpecErrorDetail{UserMessage: "We are hard at work implementing this feature"}
	SpecServerUnknownResourceErrorDetail    = typev2pb.SpecErrorDetail{UserMessage: "We could not find the resource you requested"}
)
