package shared

import (
	options "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/options/v2"

	pgs "github.com/lyft/protoc-gen-star/v2"
)

// HasService checks if the provided file contains one or more service definitions and returns true if it does.
func (fns Functions) HasService(file pgs.File) bool {
	if len(file.Services()) > 0 {
		return true
	}

	return false
}

// Service returns the first service found in the specified proto file or nil if no services are defined.
func (fns Functions) Service(file pgs.File) pgs.Service {
	for _, service := range file.Services() {
		return service
	}

	return nil
}

// ParentService returns the parent service for the given method by using its associated file. Transfers scope to `Service`.
func (fns Functions) ParentService(method pgs.Method) pgs.Service {
	file := method.File()
	return fns.Service(file)
}

// ServiceOptions extracts and returns SpecServiceOptions from a proto file using the Service extension.
// It panics if unable to read the service extension.
func (fns Functions) ServiceOptions(file pgs.File) options.SpecServiceOptions {
	var service options.SpecServiceOptions

	_, err := fns.Service(file).Extension(options.E_Service, &service)
	if err != nil {
		panic(err.Error() + "unable to read service extension from proto")
	}

	return service
}

// SpecServiceGrpcPort retrieves the gRPC port for a given service, returning 50000 if no specific port is defined in the service options.
func (fns Functions) SpecServiceGrpcPort(service pgs.Service) int32 {
	var serviceOption options.SpecServiceOptions
	hasOption, err := service.Extension(options.E_Service, &serviceOption)
	if err != nil {
		panic(err.Error() + "unable to read service extension from proto")
	}
	if !hasOption {
		return 50000
	}

	return serviceOption.GrpcPort
}

// SpecServiceHttpPort retrieves the HTTP port for the given service based on its SpecServiceOptions extension.
// If no SpecServiceOptions are defined, it defaults to port 8080. An error is raised for extension read failures.
func (fns Functions) SpecServiceHttpPort(service pgs.Service) int32 {
	var serviceOption options.SpecServiceOptions
	hasOption, err := service.Extension(options.E_Service, &serviceOption)
	if err != nil {
		panic(err.Error() + "unable to read service extension from proto")
	}
	if !hasOption {
		return 8080
	}

	return serviceOption.HttpPort
}
