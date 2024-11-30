package shared

import (
	"strings"

	pgs "github.com/lyft/protoc-gen-star/v2"
	options "libs/protobuf/go/protobuf/gen/platform/options/v2"
)

func (fns Functions) GetSpecCommands(file pgs.File) pgs.Enum {

	for _, enum := range file.AllEnums() {
		var spec options.SpecOptions

		_, err := enum.Extension(options.E_Spec, &spec)
		if err != nil {
			panic(err.Error() + "unable to read enum extension from enum")
		}

		if spec.GetType().String() == "SPEC_ENUM_TYPE_COMMANDS" {
			return enum
		}
	}

	return nil

}

func (fns Functions) GetSpecEvents(file pgs.File) pgs.Enum {

	for _, enum := range file.AllEnums() {
		var spec options.SpecOptions

		_, err := enum.Extension(options.E_Spec, &spec)
		if err != nil {
			panic(err.Error() + "unable to read enum extension from enum")
		}

		if spec.GetType().String() == "SPEC_ENUM_TYPE_EVENTS" {
			return enum
		}
	}

	return nil

}

func (fns Functions) GetSpecTopics(file pgs.File) pgs.Enum {

	for _, enum := range file.AllEnums() {
		var spec options.SpecOptions

		_, err := enum.Extension(options.E_Spec, &spec)
		if err != nil {
			panic(err.Error() + "unable to read enum extension from enum")
		}

		if spec.GetType().String() == "SPEC_ENUM_TYPE_TOPICS" {
			return enum
		}
	}

	return nil

}

func (fns Functions) GetSpecEnumSuffix(enum pgs.EnumValue) pgs.Name {

	var parts = strings.Split(enum.Name().String(), "_")

	return pgs.Name(parts[len(parts)-1])

}

func (fns Functions) GetChannelPrefix(file pgs.File) string {
	return file.Package().ProtoName().LowerDotNotation().String()
}
