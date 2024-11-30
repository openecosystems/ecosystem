package shared

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	options "libs/protobuf/go/protobuf/gen/platform/options/v2"
)

func (fns Functions) ConfigurationOptions(file pgs.File) options.ConfigurationOptions {
	var config options.ConfigurationOptions

	_, err := file.Extension(options.E_Configuration, &config)
	if err != nil {
		panic(err.Error() + "unable to read extension from proto")
	}

	return config
}

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

func (fns Functions) ConfigurationName(file pgs.File) pgs.Name {
	msg := fns.Configuration(file)
	return pgs.Name(msg.Name())
}

func (fns Functions) ConfigurationNumber(file pgs.File) int32 {
	config := fns.ConfigurationOptions(file)
	return config.FieldNumber
}

func (fns Functions) IsConfiguration(file pgs.File) bool {
	msg := fns.Configuration(file)
	return msg != nil
}
