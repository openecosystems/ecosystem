package sdkv2alphalib

import (
	typev2pb "libs/protobuf/go/protobuf/gen/platform/type/v2"

	"connectrpc.com/connect"
)

// ErrServerInternal represents an internal server error, indicating an issue on the server's side.
// ErrServerAuthentication represents an authentication error due to invalid or incorrect credentials.
// ErrServerRequest represents a client-side error caused by an invalid request.
// ErrServerAlreadyExists represents an error indicating that the resource already exists.
// ErrServerNotSupported represents an error indicating that the requested operation is not supported.
// ErrServerPreconditionFailed represents an error indicating that preconditions for the operation have failed.
// ErrServerUnimplemented represents an error indicating that the requested API call is not implemented.
// ErrServerUnknownResource represents an error indicating that the requested resource could not be found.
var (
	ErrServerInternal           SpecError = &Error{SpecApiErr: &SpecAPIError{*NewConnectError(connect.CodeInternal, &SpecServerInternalErrorDetail, "Error on our side")}}
	ErrServerAuthentication     SpecError = &Error{SpecApiErr: &SpecAPIError{*NewConnectError(connect.CodeUnauthenticated, &SpecServerAuthenticationErrorDetail, "Invalid or incorrect credentials")}}
	ErrServerRequest            SpecError = &Error{SpecApiErr: &SpecAPIError{*NewConnectError(connect.CodeInvalidArgument, &SpecServerRequestErrorDetail, "There was an error with your request")}}
	ErrServerAlreadyExists      SpecError = &Error{SpecApiErr: &SpecAPIError{*NewConnectError(connect.CodeAlreadyExists, &SpecServerAlreadyExistsErrorDetail, "Resource already exist")}}
	ErrServerNotSupported       SpecError = &Error{SpecApiErr: &SpecAPIError{*NewConnectError(connect.CodeUnimplemented, &SpecServerNotSupportedErrorDetail, "Not Supported")}}
	ErrServerPreconditionFailed SpecError = &Error{SpecApiErr: &SpecAPIError{*NewConnectError(connect.CodeFailedPrecondition, &SpecServerPreconditionFailedErrorDetail, "Preconditions failed")}}
	ErrServerUnimplemented      SpecError = &Error{SpecApiErr: &SpecAPIError{*NewConnectError(connect.CodeUnimplemented, &SpecServerUnimplementedErrorDetail, "API call unimplemented")}}
	ErrServerUnknownResource    SpecError = &Error{SpecApiErr: &SpecAPIError{*NewConnectError(connect.CodeNotFound, &SpecServerUnknownResourceErrorDetail, "Unknown resource")}}
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
