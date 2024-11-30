package shared

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	"strings"
)

func (fns Functions) GetRemoteProtobufGoPathFromFile(file pgs.File) string {
	o := file.Descriptor().GetOptions()

	p := o.GoPackage
	path := strings.Split(*p, ";")

	u := strings.TrimPrefix(path[0], GithubRepository+"/generated-public")

	return string(pgs.JoinPaths(RemoteGeneratedRepository+"/protobuf/protocolbuffers/go", u))
}

func (fns Functions) GetRemoteProtoGoPathFromFile(file pgs.File) string {
	o := file.Descriptor().GetOptions()

	p := o.GoPackage
	path := strings.Split(*p, ";")

	apiType := fns.GetApiOptionsType(file)

	u := strings.TrimPrefix(path[0], GithubRepository+"/generated-"+apiType)

	return string(pgs.JoinPaths(RemoteGeneratedRepository+"/"+apiType+"/protocolbuffers/go", u))
}

func (fns Functions) GetRemoteGrpcGoPathFromFile(file pgs.File) string {
	o := file.Descriptor().GetOptions()

	p := o.GoPackage
	path := strings.Split(*p, ";")

	apiType := fns.GetApiOptionsType(file)

	u := strings.TrimPrefix(path[0], GithubRepository+"/generated-"+apiType)

	return string(pgs.JoinPaths(RemoteGeneratedRepository+"/"+apiType+"/grpc/go", u, path[1]+"grpc"))
}
