package shared

import (
	"fmt"
	"strings"

	pgs "github.com/lyft/protoc-gen-star/v2"
	options "libs/protobuf/go/protobuf/gen/platform/options/v2"
)

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

func (fns Functions) GetRoutineMessageFieldName(enumValue pgs.EnumValue, fieldName string) string {
	msg := fns.GetRoutineMessageField(enumValue, fieldName)
	if msg == nil {
		panic(fmt.Sprintf("Routine Message Field not found for enum: %s", enumValue.Name().String()))
	}
	tokens := strings.Split(msg.Descriptor().GetTypeName(), ".")
	return tokens[len(tokens)-1]
}
