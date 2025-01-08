package connectorv2alphalib

import (
	"fmt"

	"connectrpc.com/connect"

	"libs/protobuf/go/protobuf/gen/platform/type/v2"

	"google.golang.org/protobuf/reflect/protoreflect"
)

// A ConnectorOption configures a [Connector].
type ConnectorOption interface {
	apply(*connectorOptions)
}

func WithTargetProtocols(protocols ...typev2pb.Protocol) ConnectorOption {
	return connectorOptionFunc(func(opts *connectorOptions) {
		opts.protocols = make(map[typev2pb.Protocol]struct{}, len(protocols))
		for _, p := range protocols {
			opts.protocols[p] = struct{}{}
		}
	})
}

type connectorOptions struct {
	protocols      map[typev2pb.Protocol]struct{}
	codecNames     map[string]struct{}
	preferredCodec string
}

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

func (c *connectorOptions) validate() *connect.Error {
	return nil
}

type connectorOptionFunc func(*connectorOptions)

func (f connectorOptionFunc) apply(opts *connectorOptions) {
	f(opts)
}
