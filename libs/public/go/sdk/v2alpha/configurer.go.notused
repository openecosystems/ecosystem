package sdkv2alphalib

import specv2pb "libs/protobuf/go/protobuf/gen/platform/spec/v2"

// ResolvedConfiguration holds the final resolved and unified configuration settings for the application at runtime.
var ResolvedConfiguration *Configuration

// Configurable represents an interface for components capable of resolving, validating, and providing default configurations.
type Configurable interface {
	ResolveConfiguration()
	GetDefaultConfiguration() interface{}
	ValidateConfiguration() error
}

// Configuration represents the main structure for application-specific configuration settings.
type Configuration struct {
	App specv2pb.App `yaml:"app,omitempty"`
}
