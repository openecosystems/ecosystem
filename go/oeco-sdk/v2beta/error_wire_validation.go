package sdkv2betalib

import (
	"errors"
	"fmt"
	"strings"

	"buf.build/go/protovalidate"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func ConvertValidationErrorToFieldValidations(validationErr error) ([]*errdetails.BadRequest_FieldViolation, SpecErrorable) {
	var err *protovalidate.ValidationError
	if !errors.As(validationErr, &err) {
		return nil, ErrServerInternal.WithInternalErrorDetail(
			fmt.Errorf("expected ValidationError, got: %T", validationErr),
		)
	}

	var fv []*errdetails.BadRequest_FieldViolation
	for _, violation := range err.Violations {
		fieldBuilder := &strings.Builder{}
		descriptionBuilder := &strings.Builder{}

		if fieldPath := protovalidate.FieldPathString(violation.Proto.GetField()); fieldPath != "" {
			fieldBuilder.WriteString(fieldPath)
		}

		_, _ = fmt.Fprintf(descriptionBuilder, "%s [%s]",
			violation.Proto.GetMessage(),
			violation.Proto.GetRuleId())

		fv = append(fv, &errdetails.BadRequest_FieldViolation{
			Field:       fieldBuilder.String(),
			Description: descriptionBuilder.String(),
			Reason:      "WIRE_PROTOCOL_VALIDATION_ERROR",
			LocalizedMessage: &errdetails.LocalizedMessage{
				Locale:  LocaleEnglishUS,
				Message: "Validation failed for this request",
			},
		})
	}

	return fv, nil
}
