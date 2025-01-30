package shared

import (
	"strings"

	pgs "github.com/lyft/protoc-gen-star/v2"
)

// GetRemoteProtobufGoPathFromFile derives the remote Protobuf Go import path for a given file based on its options.
func (fns Functions) GetRemoteProtobufGoPathFromFile(file pgs.File) string {
	o := file.Descriptor().GetOptions()

	p := o.GoPackage
	path := strings.Split(*p, ";")

	u := strings.TrimPrefix(path[0], GithubRepository+"/generated-public")

	return string(pgs.JoinPaths(RemoteGeneratedRepository+"/protobuf/protocolbuffers/go", u))
}

// GetRemoteProtoGoPathFromFile retrieves the remote Go package path for a given Protocol Buffers file.
func (fns Functions) GetRemoteProtoGoPathFromFile(file pgs.File) string {
	o := file.Descriptor().GetOptions()

	p := o.GoPackage
	path := strings.Split(*p, ";")

	apiType := fns.GetApiOptionsType(file)

	u := strings.TrimPrefix(path[0], GithubRepository+"/generated-"+apiType)

	return string(pgs.JoinPaths(RemoteGeneratedRepository+"/"+apiType+"/protocolbuffers/go", u))
}

// GetRemoteGrpcGoPathFromFile generates the remote GRPC Go path for the given Protocol Buffers file.
func (fns Functions) GetRemoteGrpcGoPathFromFile(file pgs.File) string {
	o := file.Descriptor().GetOptions()

	p := o.GoPackage
	path := strings.Split(*p, ";")

	apiType := fns.GetApiOptionsType(file)

	u := strings.TrimPrefix(path[0], GithubRepository+"/generated-"+apiType)

	return string(pgs.JoinPaths(RemoteGeneratedRepository+"/"+apiType+"/grpc/go", u, path[1]+"grpc"))
}
