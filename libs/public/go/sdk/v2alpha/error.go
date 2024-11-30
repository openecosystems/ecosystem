package sdkv2alphalib

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"

	"connectrpc.com/connect"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	typev2pb "libs/protobuf/go/protobuf/gen/platform/type/v2"
)

type (
	// SpecError is an error result that happens when using an API.
	SpecError interface {
		SpecApiError() *SpecApiError
		WithErrorDetail(detail *typev2pb.SpecErrorDetail) *Error
		WithInternalErrorDetail(errs ...error) *Error
		error
	}

	Error struct {
		SpecApiErr     *SpecApiError
		internalApiErr []error
	}

	SpecApiError struct {
		connect.Error
	}
)

// PrintError prints the error code and message.
func (e *SpecApiError) printError() string {
	return fmt.Sprintf("server: API error: code=%d message=%s", e.Code, e.Message())
}

// SpecApiError implements the SpecError interface.
func (e *SpecApiError) SpecApiError() *SpecApiError {
	return e
}

// Is matches against a SpecApiError.
func (e *SpecApiError) Is(err error) bool {
	if e == nil {
		return false
	}
	// Extract internal SpecApiError to match against.
	var sae *SpecApiError
	ok := errors.As(err, &sae)
	if !ok {
		return ok
	}
	return reflect.TypeOf(e.Code) == reflect.TypeOf(sae.Code)
}

func NewConnectError(code connect.Code, detail *typev2pb.SpecErrorDetail, message string) *connect.Error {
	ee := connect.NewError(code, errors.New(message))

	d, err := connect.NewErrorDetail(detail)
	if err != nil {
		return nil
	}

	ee.AddDetail(d)

	return ee
}

func (se *Error) SpecApiError() *SpecApiError {
	return se.SpecApiErr
}

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

func (se *Error) WithInternalErrorDetail(errs ...error) *Error {
	fmt.Println("internal error details: ", errs)
	se.internalApiErr = errs
	return se
}

func (se *Error) Error() string {
	var buffer bytes.Buffer

	if se.internalApiErr != nil && len(se.internalApiErr) > 0 {
		for _, e := range se.internalApiErr {
			buffer.WriteString(fmt.Sprintf("server: %s", e.Error()))
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
