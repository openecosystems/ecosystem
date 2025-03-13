package sdkv2alphalib

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"reflect"

	specproto "libs/protobuf/go/protobuf/gen/platform/spec/v2"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"gopkg.in/yaml.v3"
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

// ProtoToYAML converts a protobuf message into its equivalent YAML representation and returns it as a string.
// Returns an error if the conversion to JSON or YAML fails.
func ProtoToYAML(pb proto.Message) (string, error) {
	jsonData, err := protojson.Marshal(pb) // Protobuf → JSON
	if err != nil {
		return "", err
	}

	var mapData map[string]interface{}
	if err := yaml.Unmarshal(jsonData, &mapData); err != nil {
		return "", err
	}

	yamlData, err := yaml.Marshal(mapData) // Convert map to YAML
	if err != nil {
		return "", err
	}

	return string(yamlData), nil
}

// YamlToProto converts a YAML string into a Protobuf message by first translating it to JSON and then unmarshalling it.
// It takes a YAML string (yamlStr) and a Protobuf message (pb) as input and returns an error if the conversion fails.
func YamlToProto(yamlStr string, pb proto.Message) error {
	var mapData map[string]interface{}
	if err := yaml.Unmarshal([]byte(yamlStr), &mapData); err != nil {
		return err
	}

	jsonData, err := yaml.Marshal(mapData) // Convert YAML to JSON
	if err != nil {
		return err
	}

	return protojson.Unmarshal(jsonData, pb) // JSON → Protobuf
}

// ProtobufStructToByteArray converts a Protobuf struct into a byte array using proto.Marshal.
// It takes an interface{} as input, which should be a valid proto.Message.
// Returns the marshaled byte array or an error if the input cannot be marshaled.
func ProtobufStructToByteArray(data interface{}) ([]byte, error) {
	return proto.Marshal(data.(proto.Message))
}

// StructToByteArray encodes a given struct into a byte array using gob encoding and returns the resulting bytes and an error.
func StructToByteArray(data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	return buf.Bytes(), err
}
