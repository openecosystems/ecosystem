package natsnodev1

import (
	"strings"

	specv2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/spec/v2"
	"go.uber.org/zap/zapcore"
)

func baseFields(spec *specv2pb.Spec, subject string) []zapcore.Field {
	var fields []zapcore.Field

	if spec.Context != nil {
		if subject != "" {
			fields = append(fields, zapcore.Field{Key: "subject", Type: zapcore.StringType, String: subject})
		}

		if spec.SpanContext != nil && spec.SpanContext.TraceId != "" {
			fields = append(fields, zapcore.Field{Key: "trace_id", Type: zapcore.StringType, String: spec.SpanContext.TraceId})
		}

		if spec.Context.Validation != nil && spec.Context.Validation.ValidateOnly {
			fields = append(fields, zapcore.Field{Key: "validate_only", Type: zapcore.StringType, String: "true"})
		}

		if spec.Context.OrganizationId != "" {
			fields = append(fields, zapcore.Field{Key: "organization_id", Type: zapcore.StringType, String: spec.Context.OrganizationId})
		}

		if spec.Context.EcosystemId != "" {
			fields = append(fields, zapcore.Field{Key: "ecosystem_id", Type: zapcore.StringType, String: spec.Context.EcosystemId})
		}

		if spec.SpecData != nil && spec.SpecData.FieldMask != nil && len(spec.SpecData.FieldMask.GetPaths()) > 0 {
			fields = append(fields, zapcore.Field{Key: "field_mask", Type: zapcore.StringType, String: strings.Join(spec.SpecData.FieldMask.GetPaths(), ",")})
		}
	}

	return fields
}

func receivedFields(spec *specv2pb.Spec, subject string) []zapcore.Field {
	fields := baseFields(spec, subject)

	if spec != nil {
		if spec.ReceivedAt != nil {
			fields = append(fields, zapcore.Field{Key: "received_at", Type: zapcore.StringType, String: spec.ReceivedAt.AsTime().String()})
		}
	}

	return fields
}

func completedFields(spec *specv2pb.Spec, subject string) []zapcore.Field {
	fields := baseFields(spec, subject)

	if spec != nil {
		if spec.CompletedAt != nil {
			fields = append(fields, zapcore.Field{Key: "completed_at", Type: zapcore.StringType, String: spec.CompletedAt.AsTime().String()})
		}
	}

	return fields
}
