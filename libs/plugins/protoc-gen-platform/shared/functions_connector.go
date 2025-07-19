package shared

import (
	"strings"

	options "github.com/openecosystems/ecosystem/libs/protobuf/go/sdk/v2beta/gen/platform/options/v2"

	pgs "github.com/lyft/protoc-gen-star/v2"
)

// ConnectorOptions retrieves the ConnectorOptions extension from the provided protobuf service descriptor.
// It panics if the extension cannot be read.
func (fns Functions) ConnectorOptions(service pgs.Service) options.ConnectorOptions {
	var connector options.ConnectorOptions

	_, err := service.Extension(options.E_Connector, &connector)
	if err != nil {
		panic(err.Error() + "unable to read service extension from proto")
	}

	return connector
}

// Connectors extracts and returns a slice of services in a file that have the Connector extension applied.
func (fns Functions) Connectors(file pgs.File) []pgs.Service {
	var connector options.ConnectorOptions
	connectors := make([]pgs.Service, 0)

	for _, service := range file.Services() {
		ok, err := service.Extension(options.E_Connector, &connector)
		if err != nil {
			panic(err.Error() + "unable to read extension from proto")
		}

		if ok {
			connectors = append(connectors, service)
		}
	}

	return connectors
}

// ConnectorType maps the connector type of a service into its corresponding name, stripping standard prefixes.
func (fns Functions) ConnectorType(svc pgs.Service) pgs.Name {
	connector := fns.ConnectorOptions(svc)
	name, ok := options.ConnectorType_name[int32(connector.Type)]
	if !ok {
		panic("invalid connector type " + connector.Type.String())
	}

	// This assumes all EnumValues are prefixed with "CONNECTOR_TYPE_" which
	// is the current standard naming convention.
	name = strings.Replace(name, "CONNECTOR_TYPE_", "", 1)
	return pgs.Name(name)
}
