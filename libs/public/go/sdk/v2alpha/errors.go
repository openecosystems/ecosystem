package sdkv2alphalib

import (
	typev2pb "libs/protobuf/go/protobuf/gen/platform/type/v2"

	"connectrpc.com/connect"
)

var (
	ErrServerInternal           SpecError = &Error{SpecApiErr: &SpecApiError{*NewConnectError(connect.CodeInternal, &SpecServerInternalErrorDetail, "Error on our side")}}
	ErrServerAuthentication     SpecError = &Error{SpecApiErr: &SpecApiError{*NewConnectError(connect.CodeUnauthenticated, &SpecServerAuthenticationErrorDetail, "Invalid or incorrect credentials")}}
	ErrServerRequest            SpecError = &Error{SpecApiErr: &SpecApiError{*NewConnectError(connect.CodeInvalidArgument, &SpecServerRequestErrorDetail, "There was an error with your request")}}
	ErrServerAlreadyExists      SpecError = &Error{SpecApiErr: &SpecApiError{*NewConnectError(connect.CodeAlreadyExists, &SpecServerAlreadyExistsErrorDetail, "Resource already exist")}}
	ErrServerNotSupported       SpecError = &Error{SpecApiErr: &SpecApiError{*NewConnectError(connect.CodeUnimplemented, &SpecServerNotSupportedErrorDetail, "Not Supported")}}
	ErrServerPreconditionFailed SpecError = &Error{SpecApiErr: &SpecApiError{*NewConnectError(connect.CodeFailedPrecondition, &SpecServerPreconditionFailedErrorDetail, "Preconditions failed")}}
	ErrServerUnimplemented      SpecError = &Error{SpecApiErr: &SpecApiError{*NewConnectError(connect.CodeUnimplemented, &SpecServerUnimplementedErrorDetail, "API call unimplemented")}}
	ErrServerUnknownResource    SpecError = &Error{SpecApiErr: &SpecApiError{*NewConnectError(connect.CodeNotFound, &SpecServerUnknownResourceErrorDetail, "Unknown resource")}}
)

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
