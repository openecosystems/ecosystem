package shared

import (
	"strings"

	options "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/options/v2"

	pgs "github.com/lyft/protoc-gen-star/v2"
)

// GetSpecCommands retrieves the enum from the file that matches the "SPEC_ENUM_TYPE_COMMANDS" type in its spec options.
// If no matching enum is found, it returns nil.
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

// GetSpecEvents extracts and returns the first Enum descriptor with a SpecOption of type SPEC_ENUM_TYPE_EVENTS.
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

// GetSpecTopics retrieves the enum with a "SPEC_ENUM_TYPE_TOPICS" type from the provided file's enums.
// If no such enum is found, it returns nil.
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

// GetSpecEnumSuffix extracts and returns the last part of an EnumValue's name, delimited by underscores.
func (fns Functions) GetSpecEnumSuffix(enum pgs.EnumValue) pgs.Name {
	parts := strings.Split(enum.Name().String(), "_")

	return pgs.Name(parts[len(parts)-1])
}

// GetChannelPrefix generates a channel prefix string based on the protobuf package name in lowercase dot notation.
func (fns Functions) GetChannelPrefix(file pgs.File) string {
	return file.Package().ProtoName().LowerDotNotation().String()
}
