//nolint:revive
package sdkv2betalib

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"time"

	"connectrpc.com/connect"
	apexlog "github.com/apex/log"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	specv2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/spec/v2"
	typev2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/type/v2"
)

// Using guidance from: https://google.aip.dev/193

// Reason is the stable logical identity for a SpecError.
type Reason string

// HasReason check if the Error has a reason
type HasReason interface {
	error
	SpecReason() Reason
}

// SpecErrorable is an interface for defining API-based errors, providing methods to access and modify error details.
// SpecError represents a custom error type containing SpecApiError and optional internal error details.
// SpecApiError extends connect.Error and defines API-specific error handling functionality.
type (
	// SpecErrorable is an error result that happens when using an API.
	SpecErrorable interface {
		WithRequestInfo(info *errdetails.RequestInfo) SpecErrorable
		WithResourceInfo(info *errdetails.ResourceInfo) SpecErrorable
		WithErrorInfo(info *errdetails.ErrorInfo) SpecErrorable
		WithRetryInfo(info *errdetails.RetryInfo) SpecErrorable
		WithDebugInfo(info *errdetails.DebugInfo) SpecErrorable
		WithQuotaFailure(failure *errdetails.QuotaFailure) SpecErrorable
		WithPreconditionFailure(failure *errdetails.PreconditionFailure) SpecErrorable
		WithBadRequest(request *errdetails.BadRequest) SpecErrorable
		WithHelp(help *errdetails.Help) SpecErrorable
		WithSpecDetail(spec *specv2pb.Spec) SpecErrorable
		WithLocalizedMessage(message *errdetails.LocalizedMessage) SpecErrorable
		WithInternalErrorDetail(errs ...error) SpecErrorable
		// WithDebugDetail(ctx context.Context, spec *specv2pb.Spec, errs ...error) SpecErrorable
		SpecReason() Reason
		ToStatus() *status.Status
		ToConnectError() *connect.Error
		error
	}

	// SpecError the main Error type
	SpecError struct {
		reason     Reason
		ConnectErr connect.Error
	}
)

// NewSpecError creates a new connect.Error with a specified code, detail, and message, adding the detail to the error.
func NewSpecError(code connect.Code, message string) *SpecError {
	ee := SpecError{
		ConnectErr: *connect.NewError(code, errors.New(message)),
	}

	return &ee
}

func NewSpecErrorFromStatus(status *status.Status) SpecErrorable {
	if status == nil {
		return ErrServerInternal.WithInternalErrorDetail(errors.New("status is nil when attempting to create a new spec error"))
	}
	ee := SpecError{
		ConnectErr: *connect.NewError(connect.Code(status.Code), errors.New(status.Message)), // nolint:gosec
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

	return &ee
}

// WithRequestInfo with request information
func (se *SpecError) WithRequestInfo(info *errdetails.RequestInfo) SpecErrorable {
	d, err := connect.NewErrorDetail(info)
	if err != nil {
		apexlog.Error("server: SpecError creating new SpecError detail")
		return se
	}

	se.ConnectErr.AddDetail(d)
	return se
}

// WithResourceInfo resource information
func (se *SpecError) WithResourceInfo(info *errdetails.ResourceInfo) SpecErrorable {
	d, err := connect.NewErrorDetail(info)
	if err != nil {
		apexlog.Error("server: SpecError creating new ResourceInfo")
		return se
	}

	se.ConnectErr.AddDetail(d)
	return se
}

func (se *SpecError) WithErrorInfo(info *errdetails.ErrorInfo) SpecErrorable {
	d, err := connect.NewErrorDetail(info)
	if err != nil {
		apexlog.Error("server: SpecError creating new ErrorInfo")
		return se
	}

	if info.Reason != "" {
		se.reason = Reason(info.Reason)
	}

	se.ConnectErr.AddDetail(d)
	return se
}

func (se *SpecError) WithRetryInfo(info *errdetails.RetryInfo) SpecErrorable {
	d, err := connect.NewErrorDetail(info)
	if err != nil {
		apexlog.Error("server: SpecError creating new RetryInfo")
		return se
	}

	se.ConnectErr.AddDetail(d)
	return se
}

func (se *SpecError) WithDebugInfo(info *errdetails.DebugInfo) SpecErrorable {
	d, err := connect.NewErrorDetail(info)
	if err != nil {
		apexlog.Error("server: SpecError creating new DebugInfo")
		return se
	}

	se.ConnectErr.AddDetail(d)
	return se
}

func (se *SpecError) WithQuotaFailure(failure *errdetails.QuotaFailure) SpecErrorable {
	d, err := connect.NewErrorDetail(failure)
	if err != nil {
		apexlog.Error("server: SpecError creating new QuotaFailure")
		return se
	}

	se.ConnectErr.AddDetail(d)
	return se
}

func (se *SpecError) WithPreconditionFailure(failure *errdetails.PreconditionFailure) SpecErrorable {
	d, err := connect.NewErrorDetail(failure)
	if err != nil {
		apexlog.Error("server: SpecError creating new PreconditionFailure")
		return se
	}

	se.ConnectErr.AddDetail(d)
	return se
}

func (se *SpecError) WithBadRequest(request *errdetails.BadRequest) SpecErrorable {
	d, err := connect.NewErrorDetail(request)
	if err != nil {
		apexlog.Error("server: SpecError creating new BadRequest")
		return se
	}

	if request != nil && len(request.FieldViolations) > 0 {
		msgs := make([]string, 0, len(request.FieldViolations))
		for _, v := range request.FieldViolations {
			if v != nil {
				msgs = append(msgs, v.String())
			}
		}
		if len(msgs) > 0 {
			apexlog.WithField("bad_request_errors", strings.Join(msgs, "; ")).Info("captured bad request error details")
		}
	}

	newErr := *se
	t := connect.NewError(se.ConnectErr.Code(), fmt.Errorf(se.ConnectErr.Message()))
	newErr.ConnectErr = *t

	for _, detail := range se.ConnectErr.Details() {
		newErr.ConnectErr.AddDetail(detail)
	}

	newErr.ConnectErr.AddDetail(d)

	return &newErr
}

func (se *SpecError) WithHelp(help *errdetails.Help) SpecErrorable {
	d, err := connect.NewErrorDetail(help)
	if err != nil {
		apexlog.Error("server: SpecError creating new Help")
		return se
	}

	se.ConnectErr.AddDetail(d)
	return se
}

func (se *SpecError) WithLocalizedMessage(message *errdetails.LocalizedMessage) SpecErrorable {
	d, err := connect.NewErrorDetail(message)
	if err != nil {
		apexlog.Error("server: SpecError creating new LocalizedMessage")
		return se
	}

	se.ConnectErr.AddDetail(d)
	return se
}

func (se *SpecError) WithSpecDetail(spec *specv2pb.Spec) SpecErrorable {
	if spec == nil {
		return se
	}

	s := typev2pb.SpecErrorDetail{}

	if spec.SpanContext != nil && spec.SpanContext.TraceId != "" {
		s.CorrelationId = spec.SpanContext.TraceId
	}

	if spec.ReceivedAt != nil {
		s.ReceivedAt = spec.ReceivedAt
	}

	if spec.SentAt != nil && spec.SentAt.AsTime().IsZero() || spec.SentAt.AsTime().Equal(time.Unix(0, 0).UTC()) {
		s.SentAt = spec.ReceivedAt
	} else {
		s.SentAt = spec.SentAt
	}

	if spec.CompletedAt != nil && spec.CompletedAt.AsTime().IsZero() || spec.CompletedAt.AsTime().Equal(time.Unix(0, 0).UTC()) {
		s.CompletedAt = timestamppb.Now()
	} else {
		s.CompletedAt = spec.CompletedAt
	}

	if spec.SpecType != "" {
		s.SpecType = spec.SpecType
	}

	d, err := connect.NewErrorDetail(&s)
	if err != nil {
		apexlog.Error("server: SpecError creating SpecErrorDetail error detail")
		return se
	}

	newErr := *se // shallow copy
	newErr.ConnectErr = *connect.NewError(se.ConnectErr.Code(), se.ConnectErr.Unwrap())
	newErr.ConnectErr.AddDetail(d)
	// se.ConnectErr.AddDetail(d)
	return &newErr
}

// WithInternalErrorDetail sets internal error details for the SpecError instance and returns the updated SpecError object.
func (se *SpecError) WithInternalErrorDetail(errs ...error) SpecErrorable {
	msgs := make([]string, 0, len(errs))
	for _, e := range errs {
		if e != nil {
			msgs = append(msgs, e.Error())
		}
	}
	if len(msgs) > 0 {
		e := apexlog.WithField("internal_errors", strings.Join(msgs, "; "))

		if se.ToConnectError() != nil && len(se.ToConnectError().Details()) > 0 {
			for _, detail := range se.ToConnectError().Details() {
				_, d := detail.Value()
				e.WithError(d)
			}
		}
		e.Error("captured internal error details")
	}

	return se
}

// SpecError implements the error interface for the SpecError type, constructing and returning a formatted error message string.
// It includes details from both internalApiErr and ConnectErr if they are present.
func (se *SpecError) Error() string {
	var buffer bytes.Buffer

	if se.ConnectErr.Message() != "" {
		buffer.WriteString(se.ConnectErr.Error())
		return buffer.String()
	}

	buffer.WriteString(fmt.Sprintf("server: %s", se.ConnectErr.Unwrap()))
	return buffer.String()
}

// Is Check if this is a specific error
func (se *SpecError) Is(target error) bool {
	if target == nil {
		return false
	}

	// Direct type assertion for SpecError
	if t, ok := target.(*SpecError); ok {
		return se.reason != "" && se.reason == t.reason
	}

	// Avoid calling errors.As to prevent recursion
	if hr, ok := target.(HasReason); ok {
		return se.reason != "" && se.reason == hr.SpecReason()
	}

	// Direct type assertion for connect.Error
	if ce, ok := target.(*connect.Error); ok {
		return se.ConnectErr.Code() == ce.Code()
	}

	return false
}

func (se *SpecError) Code() connect.Code {
	return se.ConnectErr.Code()
}

func (se *SpecError) ToStatus() *status.Status {
	s := status.Status{
		Code:    int32(se.Code()), //nolint:gosec
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

func (se *SpecError) Unwrap() error {
	return &se.ConnectErr
}

func (se *SpecError) SpecReason() Reason { return se.reason }

func (se *SpecError) ToConnectError() *connect.Error {
	return &se.ConnectErr
}

// IsSpecErrorable checks if the given error implements SpecErrorable.
// Returns the casted SpecErrorable and a bool indicating success.
func IsSpecErrorable(err error) (SpecErrorable, bool) {
	if err == nil {
		return nil, false
	}
	var se SpecErrorable
	ok := errors.As(err, &se)
	return se, ok
}

// IsReason Quick check by reason, anywhere in the error chain
func IsReason(err error, reason Reason) bool {
	var hr HasReason
	return errors.As(err, &hr) && hr.SpecReason() == reason
}
