//nolint:revive
package sdkv2betalib

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"connectrpc.com/connect"
	apexlog "github.com/apex/log"
	"google.golang.org/genproto/googleapis/rpc/errdetails"

	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/types/known/anypb"
)

// Using guidance from: https://google.aip.dev/193

// SpecErrorable is an interface for defining API-based errors, providing methods to access and modify error details.
// SpecError represents a custom error type containing SpecApiError and optional internal error details.
// SpecApiError extends connect.Error and defines API-specific error handling functionality.
type (
	// SpecErrorable is an error result that happens when using an API.
	SpecErrorable interface {
		WithRequestInfo(info *errdetails.RequestInfo) SpecError
		WithResourceInfo(info *errdetails.ResourceInfo) SpecError
		WithErrorInfo(info *errdetails.ErrorInfo) SpecError
		WithRetryInfo(info *errdetails.RetryInfo) SpecError
		WithDebugInfo(info *errdetails.DebugInfo) SpecError
		WithQuotaFailure(failure *errdetails.QuotaFailure) SpecError
		WithPreconditionFailure(failure *errdetails.PreconditionFailure) SpecError
		WithBadRequest(request *errdetails.BadRequest) SpecError
		WithHelp(help *errdetails.Help) SpecError
		WithLocalizedMessage(message *errdetails.LocalizedMessage) SpecError
		WithInternalErrorDetail(errs ...error) SpecError
		// WithDebugDetail(ctx context.Context, spec *specv2pb.Spec, errs ...error) SpecError
		ToStatus() *status.Status
		ToConnectError() *connect.Error
		error
	}

	// SpecError the main Error type
	SpecError struct {
		ConnectErr connect.Error
	}
)

// NewSpecError creates a new connect.Error with a specified code, detail, and message, adding the detail to the error.
func NewSpecError(code connect.Code, message string) SpecError {
	ee := SpecError{
		ConnectErr: *connect.NewError(code, errors.New(message)),
	}

	return ee
}

func NewSpecErrorFromStatus(status *status.Status) SpecError {
	if status == nil {
		return ErrServerInternal.WithInternalErrorDetail(errors.New("status is nil when attempting to create a new spec error"))
	}
	ee := SpecError{
		ConnectErr: *connect.NewError(connect.Code(status.Code), errors.New(status.Message)),
	}

	if status.Details != nil {
		for _, detail := range status.Details {
			d, err := connect.NewErrorDetail(detail)
			if err != nil {
				apexlog.Error("Could not parse the error detail from that status: " + detail.TypeUrl)
				continue
			}
			ee.ConnectErr.AddDetail(d)
		}
	}

	return ee
}

// WithRequestInfo with request information
func (se SpecError) WithRequestInfo(info *errdetails.RequestInfo) SpecError {
	d, err := connect.NewErrorDetail(info)
	if err != nil {
		apexlog.Error("server: SpecError creating new SpecError detail")
		return se
	}

	se.ConnectErr.AddDetail(d)
	return se
}

// WithResourceInfo resource information
func (se SpecError) WithResourceInfo(info *errdetails.ResourceInfo) SpecError {
	d, err := connect.NewErrorDetail(info)
	if err != nil {
		apexlog.Error("server: SpecError creating new ResourceInfo")
		return se
	}

	se.ConnectErr.AddDetail(d)
	return se
}

func (se SpecError) WithErrorInfo(info *errdetails.ErrorInfo) SpecError {
	d, err := connect.NewErrorDetail(info)
	if err != nil {
		apexlog.Error("server: SpecError creating new ErrorInfo")
		return se
	}

	se.ConnectErr.AddDetail(d)
	return se
}

func (se SpecError) WithRetryInfo(info *errdetails.RetryInfo) SpecError {
	d, err := connect.NewErrorDetail(info)
	if err != nil {
		apexlog.Error("server: SpecError creating new RetryInfo")
		return se
	}

	se.ConnectErr.AddDetail(d)
	return se
}

func (se SpecError) WithDebugInfo(info *errdetails.DebugInfo) SpecError {
	d, err := connect.NewErrorDetail(info)
	if err != nil {
		apexlog.Error("server: SpecError creating new DebugInfo")
		return se
	}

	se.ConnectErr.AddDetail(d)
	return se
}

func (se SpecError) WithQuotaFailure(failure *errdetails.QuotaFailure) SpecError {
	d, err := connect.NewErrorDetail(failure)
	if err != nil {
		apexlog.Error("server: SpecError creating new QuotaFailure")
		return se
	}

	se.ConnectErr.AddDetail(d)
	return se
}

func (se SpecError) WithPreconditionFailure(failure *errdetails.PreconditionFailure) SpecError {
	d, err := connect.NewErrorDetail(failure)
	if err != nil {
		apexlog.Error("server: SpecError creating new PreconditionFailure")
		return se
	}

	se.ConnectErr.AddDetail(d)
	return se
}

func (se SpecError) WithBadRequest(request *errdetails.BadRequest) SpecError {
	d, err := connect.NewErrorDetail(request)
	if err != nil {
		apexlog.Error("server: SpecError creating new BadRequest")
		return se
	}

	se.ConnectErr.AddDetail(d)
	return se
}

func (se SpecError) WithHelp(help *errdetails.Help) SpecError {
	d, err := connect.NewErrorDetail(help)
	if err != nil {
		apexlog.Error("server: SpecError creating new Help")
		return se
	}

	se.ConnectErr.AddDetail(d)
	return se
}

func (se SpecError) WithLocalizedMessage(message *errdetails.LocalizedMessage) SpecError {
	d, err := connect.NewErrorDetail(message)
	if err != nil {
		apexlog.Error("server: SpecError creating new LocalizedMessage")
		return se
	}

	se.ConnectErr.AddDetail(d)
	return se
}

// WithInternalErrorDetail sets internal error details for the SpecError instance and returns the updated SpecError object.
func (se SpecError) WithInternalErrorDetail(errs ...error) SpecError {
	var errStrings []string
	for _, err := range errs {
		errStrings = append(errStrings, err.Error())
	}
	apexlog.WithField("internal_errors", strings.Join(errStrings, "; ")).Error("captured internal errors")
	return se
}

// SpecError implements the error interface for the SpecError type, constructing and returning a formatted error message string.
// It includes details from both internalApiErr and ConnectErr if they are present.
func (se SpecError) Error() string {
	var buffer bytes.Buffer

	if se.ConnectErr.Message() != "" {
		buffer.WriteString(se.ConnectErr.Error())
		return buffer.String()
	}

	buffer.WriteString(fmt.Sprintf("server: %s", se.ConnectErr.Unwrap()))
	return buffer.String()
}

func (se SpecError) Is(target error) bool {
	if target == nil {
		return false
	}

	// Case 1: Check if target is a SpecError to compare types
	// var se2 SpecError
	if errors.As(target, &se) {
		return se.ConnectErr.Code() == se.ConnectErr.Code()
	}

	// Case 2: Check if target is a connect.Error
	var ce *connect.Error
	if errors.As(target, &ce) {
		return se.ConnectErr.Code() == ce.Code()
	}

	return false
}

func (se SpecError) Code() connect.Code {
	return se.ConnectErr.Code()
}

func (se SpecError) ToStatus() *status.Status {
	s := status.Status{
		Code:    int32(se.Code()),
		Message: se.ConnectErr.Message(),
	}

	for _, detail := range se.ConnectErr.Details() {
		pb, err := detail.Value()
		if err != nil {
			apexlog.Warn("server: SpecError adding detail to Status; skipping but continuing")
			continue
		}

		anyMsg, err := anypb.New(pb)
		if err != nil {
			apexlog.Warn("server: SpecError adding protobuffed detail to Status; skipping but continuing")
			continue
		}

		s.Details = append(s.Details, anyMsg)
	}

	return &s
}

func (se SpecError) ToConnectError() *connect.Error {
	return &se.ConnectErr
}
