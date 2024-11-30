package shared

import (
	"strings"

	pgs "github.com/lyft/protoc-gen-star/v2"
	options "libs/protobuf/go/protobuf/gen/platform/options/v2"
)

func (fns Functions) ConnectorOptions(service pgs.Service) options.ConnectorOptions {
	var connector options.ConnectorOptions

	_, err := service.Extension(options.E_Connector, &connector)
	if err != nil {
		panic(err.Error() + "unable to read service extension from proto")
	}

	return connector
}

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
