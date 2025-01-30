package shared

import (
	"fmt"
	"strings"

	options "libs/protobuf/go/protobuf/gen/platform/options/v2"

	pgs "github.com/lyft/protoc-gen-star/v2"
)

// GetRoutines retrieves the first enum from the provided file that has the SPEC_ENUM_TYPE_ROUTINE_LISTENERS type in its extension.
// If no such enum is found, it returns nil. It panics if the enum extension cannot be read.
func (fns Functions) GetRoutines(file pgs.File) pgs.Enum {
	for _, enum := range file.AllEnums() {
		var spec options.SpecOptions

		_, err := enum.Extension(options.E_Spec, &spec)
		if err != nil {
			panic(err.Error() + "unable to read enum extension from enum")
		}
		if spec.GetType().String() == "SPEC_ENUM_TYPE_ROUTINE_LISTENERS" {
			return enum
		}
	}

	return nil
}

// GetRoutineMessage retrieves the `pgs.Message` associated with the given `pgs.EnumValue` by matching listener values.
func (fns Functions) GetRoutineMessage(enumValue pgs.EnumValue) pgs.Message {
	for _, msg := range enumValue.File().AllMessages() {
		var routine options.RoutineOptions

		_, err := msg.Extension(options.E_Routine, &routine)
		if err != nil {
			panic(err.Error() + "unable to read enum extension from enum")
		}
		if routine.GetListener() == enumValue.Name().String() {
			return msg
		}
	}

	return nil
}

// GetRoutineMessageField retrieves a specific field from the routine message associated with the given enum value.
// Panics if the routine message is not found or returns nil if no matching field is found.
func (fns Functions) GetRoutineMessageField(enumValue pgs.EnumValue, fieldName string) pgs.Field {
	routineMsg := fns.GetRoutineMessage(enumValue)
	if routineMsg == nil {
		panic(fmt.Sprintf("Routine Message not found for enum: %s", enumValue.Name().String()))
	}
	for _, msg := range routineMsg.Fields() {
		if msg.Name().String() == fieldName {
			return msg
		}
	}
	return nil
}

// GetRoutineMessageFieldName returns the name of a specific message field based on the provided enum value and field name.
// It splits the type name of the retrieved message by "." and returns the last token to capture the field's identifier.
// Panics if the message field cannot be found for the given enum value and field name.
func (fns Functions) GetRoutineMessageFieldName(enumValue pgs.EnumValue, fieldName string) string {
	msg := fns.GetRoutineMessageField(enumValue, fieldName)
	if msg == nil {
		panic(fmt.Sprintf("Routine Message Field not found for enum: %s", enumValue.Name().String()))
	}
	tokens := strings.Split(msg.Descriptor().GetTypeName(), ".")
	return tokens[len(tokens)-1]
}
