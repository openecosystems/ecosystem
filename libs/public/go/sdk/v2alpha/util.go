package sdkv2alphalib

import (
	"context"
	"errors"
	"reflect"

	specproto "libs/protobuf/go/protobuf/gen/platform/spec/v2"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
)

// GetDataFromSpec extracts and unmarshals the data field of a Spec object into a provided ProtoMessage instance.
// Returns an error if the Spec or its Data field is nil, or if unmarshaling fails.
func GetDataFromSpec[D protoreflect.ProtoMessage](_ context.Context, s *specproto.Spec, data D) error {
	if s == nil {
		// return errors.NewSpecError(ctx, errors.SpecInternalError(), "Cannot create object from nil spec")
		return ErrServerInternal
	}

	if s.Data == nil {
		// return errors.NewSpecError(ctx, errors.SpecPreconditionFailedError(), "Data object is not provided on the spec")
		return ErrServerPreconditionFailed
	}

	err := anypb.UnmarshalTo(s.Data, data, proto.UnmarshalOptions{
		Merge:          false,
		AllowPartial:   false,
		DiscardUnknown: false,
		Resolver:       nil,
		RecursionLimit: 0,
	})
	if err != nil {
		// return errors.NewSpecError(ctx, errors.SpecInternalError(), "failed to unmarshal data: "+err.Error())
		return ErrServerInternal
	}
	return nil
}

// UpdateSpecFromContext updates the fields in the given spec's context with values from the provided spec context.
func UpdateSpecFromContext[C any](spec *specproto.Spec, specContext C) {
	organizationSlug, err := GetField(specContext, "OrganizationSlug")
	if err == nil {
		spec.Context.OrganizationSlug = organizationSlug.String()
	}
	workspaceSlug, err := GetField(specContext, "WorkspaceSlug")
	if err == nil {
		spec.Context.WorkspaceSlug = workspaceSlug.String()
	}
}

// GetField retrieves the specified field value from a given struct using reflection.
// Returns an error if the field does not exist or is invalid.
func GetField(item interface{}, fieldName string) (*reflect.Value, error) {
	values := reflect.ValueOf(item)
	value := values.FieldByName(fieldName)
	if !value.IsValid() || value.IsZero() {
		return nil, errors.New("field not found")
	}
	return &value, nil
}

// GetContextBinValue generates a concatenated string of OrganizationSlug and WorkspaceSlug from SpecContext if both are non-empty.
// If WorkspaceSlug is empty, returns only OrganizationSlug. Returns an error if both fields are empty.
func GetContextBinValue(specContext *specproto.SpecContext) (string, error) {
	if specContext.OrganizationSlug != "" && specContext.WorkspaceSlug != "" {
		return specContext.OrganizationSlug + ":" + specContext.WorkspaceSlug, nil
	} else if specContext.WorkspaceSlug == "" {
		return specContext.OrganizationSlug, nil
	}
	return "", errors.New("empty spec context bin value")
}

// ConvertToJSON converts complex nested structures to JSON-compatible formats (e.g., maps with string keys, arrays).
func ConvertToJSON(v interface{}) (r interface{}) {
	switch v := v.(type) {
	case []interface{}:
		for i, e := range v {
			v[i] = ConvertToJSON(e)
		}
		// r = []interface{}(v)
	case map[interface{}]interface{}:
		newMap := make(map[string]interface{}, len(v))
		for k, e := range v {
			newMap[k.(string)] = ConvertToJSON(e)
		}
		// r = map[string]interface{}(newMap)
	default:
		r = v
	}
	return
}
