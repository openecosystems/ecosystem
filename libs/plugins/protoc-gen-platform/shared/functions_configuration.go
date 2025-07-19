package shared

import (
	options "github.com/openecosystems/ecosystem/libs/protobuf/go/sdk/v2beta/gen/platform/options/v2"

	pgs "github.com/lyft/protoc-gen-star/v2"
)

// ConfigurationOptions retrieves configuration options from the given proto file by accessing its associated extension.
// Panics if the extension cannot be read or retrieved successfully.
// Returns an instance of options.ConfigurationOptions.
func (fns Functions) ConfigurationOptions(file pgs.File) options.ConfigurationOptions {
	var config options.ConfigurationOptions

	_, err := file.Extension(options.E_Configuration, &config)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}

	return config
}

// Configuration locates and returns a specific protobuf message in a file based on a given configuration.
// If the configuration is enabled, it searches for a message with an expected name derived from the file name.
// Panics if the expected message name is not found and the configuration is enabled. Returns nil if not enabled.
func (fns Functions) Configuration(file pgs.File) pgs.Message {
	config := fns.ConfigurationOptions(file)

	expectedName := pgs.Name(fns.ProtoName(file)).UpperCamelCase().String() + "Configuration"

	if config.Enabled {
		for _, msg := range file.AllMessages() {
			if msg.Name().String() == expectedName {
				return msg
			}
		}
		panic("unable to find message with name " + expectedName)
	}

	return nil
}

// ConfigurationName generates a name for the configuration by extracting it from the provided file.
func (fns Functions) ConfigurationName(file pgs.File) pgs.Name {
	msg := fns.Configuration(file)
	return pgs.Name(msg.Name())
}

// ConfigurationNumber extracts and returns the field number of the configuration option for the given protobuf file.
func (fns Functions) ConfigurationNumber(file pgs.File) int32 {
	config := fns.ConfigurationOptions(file)
	return config.FieldNumber
}

// IsConfiguration checks if the provided file contains a configuration message. Returns true if such a message exists.
func (fns Functions) IsConfiguration(file pgs.File) bool {
	msg := fns.Configuration(file)
	return msg != nil
}
