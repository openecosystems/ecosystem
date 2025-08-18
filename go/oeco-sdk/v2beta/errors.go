package sdkv2betalib

import (
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
