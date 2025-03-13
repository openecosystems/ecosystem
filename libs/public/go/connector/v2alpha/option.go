package connectorv2alphalib

import (
	"fmt"

	typev2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/type/v2"

	"connectrpc.com/connect"

	"google.golang.org/protobuf/reflect/protoreflect"
)

// ConnectorOption defines an interface for applying custom configuration to a connectorOptions object.
type ConnectorOption interface {
	apply(*connectorOptions)
}

// WithTargetProtocols sets the allowed target protocols for the connector using the provided list of protocols.
func WithTargetProtocols(protocols ...typev2pb.Protocol) ConnectorOption {
	return connectorOptionFunc(func(opts *connectorOptions) {
		opts.protocols = make(map[typev2pb.Protocol]struct{}, len(protocols))
		for _, p := range protocols {
			opts.protocols[p] = struct{}{}
		}
	})
}

// connectorOptions defines the configuration options for a connector, including supported protocols and codecs.
type connectorOptions struct {
	protocols map[typev2pb.Protocol]struct{}
	// codecNames     map[string]struct{}
	// preferredCodec string
}

// descKind returns a string describing the kind of protoreflect.Descriptor instance provided as input.
func descKind(desc protoreflect.Descriptor) string {
	switch desc := desc.(type) {
	case protoreflect.FileDescriptor:
		return "a file"
	case protoreflect.MessageDescriptor:
		return "a message"
	case protoreflect.FieldDescriptor:
		if desc.IsExtension() {
			return "an extension"
		}
		return "a field"
	case protoreflect.OneofDescriptor:
		return "a oneof"
	case protoreflect.EnumDescriptor:
		return "an enum"
	case protoreflect.EnumValueDescriptor:
		return "an enum value"
	case protoreflect.ServiceDescriptor:
		return "a service"
	case protoreflect.MethodDescriptor:
		return "a method"
	default:
		return fmt.Sprintf("%T", desc)
	}
}

// newConnectorOptions creates and configures a new connectorOptions instance using the provided ConnectorOption slice.
// Returns the configured connectorOptions and an error if validation fails.
func newConnectorOptions(options []ConnectorOption) (*connectorOptions, *connect.Error) {
	config := connectorOptions{
		protocols: nil,
	}

	for _, opt := range options {
		opt.apply(&config)
	}

	if err := config.validate(); err != nil {
		return nil, err
	}

	return &config, nil
}

// validate checks the integrity and consistency of the connectorOptions fields.
// Returns a *connect.Error if validation fails or nil if successful.
func (c *connectorOptions) validate() *connect.Error {
	return nil
}

// connectorOptionFunc is a function type that modifies the settings of a connectorOptions instance.
type connectorOptionFunc func(*connectorOptions)

// apply applies the connectorOptionFunc to the given connectorOptions.
func (f connectorOptionFunc) apply(opts *connectorOptions) {
	f(opts)
}
