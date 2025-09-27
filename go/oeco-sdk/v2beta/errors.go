package sdkv2betalib

import (
	"connectrpc.com/connect"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

const (
	Domain = "https://platform.adino.system"
)

const (
	ReasonServerInternal           = "server_internal"
	ReasonServerAuthentication     = "server_authentication"
	ReasonServerRequest            = "server_request"
	ReasonServerAlreadyExists      = "server_already_exists"
	ReasonServerNotSupported       = "server_not_supported"
	ReasonServerPreconditionFailed = "server_precondition_failed"
	ReasonServerUnimplemented      = "server_unimplemented"
	ReasonServerUnknownResource    = "server_unknown_resource"
	ReasonServerCanceled           = "server_canceled"
	ReasonServerUnknown            = "server_unknown"
	ReasonServerDeadlineExceeded   = "server_deadline_exceeded"
	ReasonServerPermissionDenied   = "server_permission_denied"
	ReasonServerResourceExhausted  = "server_resource_exhausted"
	ReasonServerOutOfRange         = "server_out_of_range"
	ReasonServerUnimplementedAlt   = "server_unimplemented_alt"
	ReasonServerUnavailable        = "server_unavailable"
	ReasonServerDataLoss           = "server_data_loss"
	ReasonServerAborted            = "server_aborted"
	ReasonServerUnauthenticated    = "server_unauthenticated"
	ReasonServerConflict           = "server_conflict"
	ReasonServerNotFound           = "server_not_found"
)

// AllServerErrors slice of all server errors
var AllServerErrors = []SpecErrorable{
	ErrServerInternal,
	ErrServerAuthentication,
	ErrServerRequest,
	ErrServerAlreadyExists,
	ErrServerNotSupported,
	ErrServerPreconditionFailed,
	ErrServerUnimplemented,
	ErrServerUnknownResource,
	ErrServerCanceled,
	ErrServerUnknown,
	ErrServerDeadlineExceeded,
	ErrServerPermissionDenied,
	ErrServerResourceExhausted,
	ErrServerOutOfRange,
	ErrServerUnimplementedAlt,
	ErrServerUnavailable,
	ErrServerDataLoss,
	ErrServerAborted,
	ErrServerUnauthenticated,
	ErrServerConflict,
	ErrServerNotFound,
}

type BaseErrorRegistry struct{}

func (BaseErrorRegistry) AllErrors() map[Reason]SpecErrorable {
	return ProcessErrors(AllServerErrors)
}

// Sentinel instances for fast error checks
var (
	ErrServerInternal = NewSpecError(connect.CodeInternal, "An Error occurred on our side").WithErrorInfo(&errdetails.ErrorInfo{
		Reason:   ReasonServerInternal,
		Domain:   Domain,
		Metadata: nil,
	})

	ErrServerAuthentication = NewSpecError(connect.CodeUnauthenticated, "Invalid or incorrect credentials").WithErrorInfo(&errdetails.ErrorInfo{
		Reason:   ReasonServerAuthentication,
		Domain:   Domain,
		Metadata: nil,
	})

	ErrServerRequest = NewSpecError(connect.CodeInvalidArgument, "There was an error with your request").WithErrorInfo(&errdetails.ErrorInfo{
		Reason:   ReasonServerRequest,
		Domain:   Domain,
		Metadata: nil,
	})

	ErrServerAlreadyExists = NewSpecError(connect.CodeAlreadyExists, "Resource already exist").WithErrorInfo(&errdetails.ErrorInfo{
		Reason:   ReasonServerAlreadyExists,
		Domain:   Domain,
		Metadata: nil,
	})

	ErrServerNotSupported = NewSpecError(connect.CodeUnimplemented, "Not Supported").WithErrorInfo(&errdetails.ErrorInfo{
		Reason:   ReasonServerNotSupported,
		Domain:   Domain,
		Metadata: nil,
	})

	ErrServerPreconditionFailed = NewSpecError(connect.CodeFailedPrecondition, "Preconditions failed").WithErrorInfo(&errdetails.ErrorInfo{
		Reason:   ReasonServerPreconditionFailed,
		Domain:   Domain,
		Metadata: nil,
	})

	ErrServerUnimplemented = NewSpecError(connect.CodeUnimplemented, "API call unimplemented").WithErrorInfo(&errdetails.ErrorInfo{
		Reason:   ReasonServerUnimplemented,
		Domain:   Domain,
		Metadata: nil,
	})

	ErrServerUnknownResource = NewSpecError(connect.CodeNotFound, "Unknown resource").WithErrorInfo(&errdetails.ErrorInfo{
		Reason:   ReasonServerUnknownResource,
		Domain:   Domain,
		Metadata: nil,
	})

	ErrServerCanceled = NewSpecError(connect.CodeCanceled, "Request was cancelled").WithErrorInfo(&errdetails.ErrorInfo{
		Reason:   ReasonServerCanceled,
		Domain:   Domain,
		Metadata: nil,
	})

	ErrServerUnknown = NewSpecError(connect.CodeUnknown, "An unknown error occurred").WithErrorInfo(&errdetails.ErrorInfo{
		Reason:   ReasonServerUnknown,
		Domain:   Domain,
		Metadata: nil,
	})

	ErrServerDeadlineExceeded = NewSpecError(connect.CodeDeadlineExceeded, "Request deadline exceeded").WithErrorInfo(&errdetails.ErrorInfo{
		Reason:   ReasonServerDeadlineExceeded,
		Domain:   Domain,
		Metadata: nil,
	})

	ErrServerPermissionDenied = NewSpecError(connect.CodePermissionDenied, "Permission denied").WithErrorInfo(&errdetails.ErrorInfo{
		Reason:   ReasonServerPermissionDenied,
		Domain:   Domain,
		Metadata: nil,
	})

	ErrServerResourceExhausted = NewSpecError(connect.CodeResourceExhausted, "Resource limit exceeded").WithErrorInfo(&errdetails.ErrorInfo{
		Reason:   ReasonServerResourceExhausted,
		Domain:   Domain,
		Metadata: nil,
	})

	ErrServerOutOfRange = NewSpecError(connect.CodeOutOfRange, "Value out of range").WithErrorInfo(&errdetails.ErrorInfo{
		Reason:   ReasonServerOutOfRange,
		Domain:   Domain,
		Metadata: nil,
	})

	ErrServerUnimplementedAlt = NewSpecError(connect.CodeUnimplemented, "This method is not implemented").WithErrorInfo(&errdetails.ErrorInfo{
		Reason:   ReasonServerUnimplementedAlt,
		Domain:   Domain,
		Metadata: nil,
	})

	ErrServerUnavailable = NewSpecError(connect.CodeUnavailable, "Service temporarily unavailable").WithErrorInfo(&errdetails.ErrorInfo{
		Reason:   ReasonServerUnavailable,
		Domain:   Domain,
		Metadata: nil,
	})

	ErrServerDataLoss = NewSpecError(connect.CodeDataLoss, "Data loss or corruption detected").WithErrorInfo(&errdetails.ErrorInfo{
		Reason:   ReasonServerDataLoss,
		Domain:   Domain,
		Metadata: nil,
	})

	ErrServerAborted = NewSpecError(connect.CodeAborted, "Operation aborted").WithErrorInfo(&errdetails.ErrorInfo{
		Reason:   ReasonServerAborted,
		Domain:   Domain,
		Metadata: nil,
	})

	ErrServerUnauthenticated = NewSpecError(connect.CodeUnauthenticated, "Authentication required or failed").WithErrorInfo(&errdetails.ErrorInfo{
		Reason:   ReasonServerUnauthenticated,
		Domain:   Domain,
		Metadata: nil,
	})

	ErrServerConflict = NewSpecError(connect.CodeAlreadyExists, "Conflict: resource already exists").WithErrorInfo(&errdetails.ErrorInfo{
		Reason:   ReasonServerConflict,
		Domain:   Domain,
		Metadata: nil,
	})

	ErrServerNotFound = NewSpecError(connect.CodeNotFound, "Resource not found").WithErrorInfo(&errdetails.ErrorInfo{
		Reason:   ReasonServerNotFound,
		Domain:   Domain,
		Metadata: nil,
	})
)
