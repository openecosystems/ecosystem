//nolint:revive
package sdkv2alphalib

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"

	"connectrpc.com/connect"

	typev2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/type/v2"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// SpecError is an interface for defining API-based errors, providing methods to access and modify error details.
// Error represents a custom error type containing SpecApiError and optional internal error details.
// SpecApiError extends connect.Error and defines API-specific error handling functionality.
type (
	// SpecError is an error result that happens when using an API.
	SpecError interface {
		SpecApiError() *SpecAPIError
		WithErrorDetail(detail *typev2pb.SpecErrorDetail) *Error
		WithInternalErrorDetail(errs ...error) *Error
		error
	}

	Error struct {
		SpecApiErr     *SpecAPIError
		internalApiErr []error
	}

	SpecAPIError struct {
		connect.Error
	}
)

// printError formats and returns the API error details including the error code and message as a string.
func (e *SpecAPIError) printError() string {
	err := fmt.Sprintf("server: API error: code=%d message=%s", e.Code(), e.Message())
	return err
}

// SpecAPIError initializes and returns the SpecAPIError instance for further use.
func (e *SpecAPIError) SpecAPIError() *SpecAPIError {
	return e
}

// Is checks whether the provided error matches the current SpecApiError instance based on its internal type and code.
func (e *SpecAPIError) Is(err error) bool {
	if e == nil {
		return false
	}
	// Extract internal SpecApiError to match against.
	var sae *SpecAPIError
	ok := errors.As(err, &sae) //nolint:govet
	if !ok {
		return ok
	}
	return reflect.TypeOf(e.Code) == reflect.TypeOf(sae.Code)
}

// NewConnectError creates a new connect.Error with a specified code, detail, and message, adding the detail to the error.
func NewConnectError(code connect.Code, detail *typev2pb.SpecErrorDetail, message string) *connect.Error {
	ee := connect.NewError(code, errors.New(message))

	d, err := connect.NewErrorDetail(detail)
	if err != nil {
		return connect.NewError(code, err)
	}

	ee.AddDetail(d)

	return ee
}

// SpecApiError returns the SpecApiError instance associated with the Error.
func (se *Error) SpecApiError() *SpecAPIError {
	return se.SpecApiErr
}

// WithErrorDetail adds a SpecErrorDetail to the SpecApiError details of the Error object and returns the updated object.
func (se *Error) WithErrorDetail(detail *typev2pb.SpecErrorDetail) *Error {
	b, err1 := proto.Marshal(detail)
	if err1 != nil {
		fmt.Println("server: Error marshalling SpecErrorDetail while adding to SpecError")
		return se
	}

	detailAny := anypb.Any{
		TypeUrl: "type.platformapis.com/" + string(detail.ProtoReflect().Descriptor().FullName()),
		Value:   b,
	}

	d, err2 := connect.NewErrorDetail(&detailAny)
	if err2 != nil {
		fmt.Println("server: Error creating new Error detail")
		return se
	}

	se.SpecApiErr.AddDetail(d)

	return se
}

// WithInternalErrorDetail sets internal error details for the Error instance and returns the updated Error object.
func (se *Error) WithInternalErrorDetail(errs ...error) *Error {
	fmt.Println("internal error details: ", errs)
	se.internalApiErr = errs
	return se
}

// Error implements the error interface for the Error type, constructing and returning a formatted error message string.
// It includes details from both internalApiErr and SpecApiErr if they are present.
func (se *Error) Error() string {
	var buffer bytes.Buffer

	if len(se.internalApiErr) > 0 {
		for _, e := range se.internalApiErr {
			buffer.WriteString("server: " + e.Error())
		}
	}

	if se.SpecApiErr != nil && se.SpecApiErr.Message() != "" {
		buffer.WriteString(se.SpecApiErr.printError())
		return buffer.String()
	}

	buffer.WriteString(fmt.Sprintf("server: %s", se.SpecApiErr.Unwrap()))
	return buffer.String()
}

//func (se *Error) Unwrap() error {
//	// Allow matching embedded SpecApiError in case there is one.
//	if se.SpecApiErr == nil {
//		return nil
//	}
//	return se.SpecApiErr
//}
