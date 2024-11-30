package shared

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	options "libs/protobuf/go/protobuf/gen/platform/options/v2"
)

func (fns Functions) HasService(file pgs.File) bool {

	if len(file.Services()) > 0 {
		return true
	}

	return false

}

func (fns Functions) Service(file pgs.File) pgs.Service {

	for _, service := range file.Services() {

		return service

	}

	return nil

}

func (fns Functions) ParentService(method pgs.Method) pgs.Service {

	file := method.File()
	return fns.Service(file)

}

func (fns Functions) ServiceOptions(file pgs.File) options.SpecServiceOptions {

	var service options.SpecServiceOptions

	_, err := fns.Service(file).Extension(options.E_Service, &service)
	if err != nil {
		panic(err.Error() + "unable to read service extension from proto")
	}

	return service
}

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
