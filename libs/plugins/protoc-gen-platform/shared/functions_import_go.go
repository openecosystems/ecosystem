package shared

import (
	"strings"

	pgs "github.com/lyft/protoc-gen-star/v2"
)

// GetGoImportPathAndPackageForAnyFieldTypeElement returns the Go import path and package for a FieldTypeElem.
// It handles embedded messages and enums, returning empty strings for unsupported types.
func (fns Functions) GetGoImportPathAndPackageForAnyFieldTypeElement(e pgs.FieldTypeElem) (string, string) {
	if e.IsEmbed() {
		return fns.GetImportPathAndPackageForAnyMessage(e.Embed())
	} else if e.IsEnum() {
		return fns.GetImportPathAndPackageForAnyEntity(e.Enum())
	}

	return "", ""
}

// GetImportPathAndPackageForAnyMessage determines the import path and package name for a given Protobuf message.
// It handles well-known types explicitly and delegates others to GetImportPathAndPackageForAnyEntity.
func (fns Functions) GetImportPathAndPackageForAnyMessage(msg pgs.Message) (string, string) {
	if msg.IsWellKnown() {
		switch msg.WellKnownType() {
		case pgs.AnyWKT:
			return "google.golang.org/protobuf/types/known/anypb", "anypb"
		case pgs.DurationWKT:
			return "google.golang.org/protobuf/types/known/durationpb", "durationpb"
		case pgs.TimestampWKT:
			return "google.golang.org/protobuf/types/known/timestamppb", "timestamppb"
		case pgs.EmptyWKT:
			return "google.golang.org/protobuf/types/known/emptypb", "emptypb"
		}
	}

	return fns.GetImportPathAndPackageForAnyEntity(msg)
}

// GetImportPathAndPackageForAnyEntity retrieves the Go import path and package name for a given protobuf entity.
func (fns Functions) GetImportPathAndPackageForAnyEntity(e pgs.Entity) (string, string) {
	o := e.File().Descriptor().GetOptions()
	path := strings.Split(*o.GoPackage, ";")
	return path[0], path[1]
}

// GetAllGoFileLevelImportPaths extracts and returns a map of Go import paths and their corresponding package aliases.
func (fns Functions) GetAllGoFileLevelImportPaths(file pgs.File) map[string]string {
	o := file.Descriptor().GetOptions()

	p := o.GoPackage
	path := strings.Split(*p, ";")

	imports := make(map[string]string)

	imports[path[0]+path[1]] = path[1]
	imports[path[0]+"/"+path[1]+"grpc"] = path[1] + "grpc"

	// serviceProto
	// serviceGrpc

	// imports["specv2"] = "anypb"

	return imports
}

// GetAllGoFieldLevelImportPaths returns a map of Go import paths and their corresponding packages for all field types in a file.
func (fns Functions) GetAllGoFieldLevelImportPaths(file pgs.File) map[string]string {
	imports := make(map[string]string)
	for _, msg := range file.AllMessages() {
		for _, f := range msg.Fields() {
			if f.Type().IsMap() {
				keyPath, keyPkg := fns.GetGoImportPathAndPackageForAnyFieldTypeElement(f.Type().Key())
				if fns.SelectivelyAddImport(keyPath, file) {
					imports[keyPath] = keyPkg
				}

				valuePath, valuePkg := fns.GetGoImportPathAndPackageForAnyFieldTypeElement(f.Type().Element())
				if fns.SelectivelyAddImport(valuePath, file) {
					imports[valuePath] = valuePkg
				}
			} else if f.Type().IsRepeated() {
				repeatedPath, repeatedPkg := fns.GetGoImportPathAndPackageForAnyFieldTypeElement(f.Type().Element())
				if fns.SelectivelyAddImport(repeatedPath, file) {
					imports[repeatedPath] = repeatedPkg
				}
			} else if f.Type().IsEmbed() {
				embedPath, embedPkg := fns.GetImportPathAndPackageForAnyMessage(f.Type().Embed())
				if fns.SelectivelyAddImport(embedPath, file) {
					imports[embedPath] = embedPkg
				}
			} else if f.Type().IsEnum() {
				enumPath, enumPkg := fns.GetImportPathAndPackageForAnyEntity(f.Type().Enum())
				if fns.SelectivelyAddImport(enumPath, file) {
					imports[enumPath] = enumPkg
				}
			}
		}
	}

	return imports
}

// SelectivelyAddImport determines whether an import path should be added based on its presence in the provided key.
func (fns Functions) SelectivelyAddImport(key string, file pgs.File) bool {
	path, _ := fns.GetImportPathAndPackageForAnyEntity(file)
	if key != "" && !strings.Contains(key, path) {
		return true
	}

	return false
}

// GetGoImportPathForMessage returns the Go import path for the given protobuf Message, including handling well-known types.
func (fns Functions) GetGoImportPathForMessage(msg pgs.Message) string {
	if msg.IsWellKnown() {
		switch msg.WellKnownType() {
		case pgs.AnyWKT:
			return "google.golang.org/protobuf/types/known/anypb"
		case pgs.DurationWKT:
			return "google.golang.org/protobuf/types/known/durationpb"
		case pgs.TimestampWKT:
			return "google.golang.org/protobuf/types/known/timestamppb"
		case pgs.EmptyWKT:
			return "google.golang.org/protobuf/types/known/emptypb"
		}
	}

	return fns.GetGoImportPackageAndPathFromAnyEntity(msg)
}

// GetGoImportPackageAndPathFromAnyEntity determines the Go import package and path based on the provided Protobuf entity.
func (fns Functions) GetGoImportPackageAndPathFromAnyEntity(e pgs.Entity) string {
	if strings.HasPrefix(e.File().FullyQualifiedName(), ".platform.spec.") {
		return fns.GetRemoteProtobufGoPathFromFile(e.File())
	}
	if strings.HasPrefix(e.File().FullyQualifiedName(), ".platform.type.") {
		return fns.GetRemoteProtobufGoPathFromFile(e.File())
	}
	if strings.HasPrefix(e.File().FullyQualifiedName(), ".platform.") {
		return fns.GetRemoteGrpcGoPathFromFile(e.File())
	}
	o := e.File().Descriptor().GetOptions()
	path := strings.Split(*o.GoPackage, ";")
	return path[0]
}

// GetGoImportPathsElement returns the Go import path for a given FieldTypeElem based on whether it's an embed or enum.
func (fns Functions) GetGoImportPathsElement(e pgs.FieldTypeElem) string {
	if e.IsEmbed() {
		return fns.GetImportPackageMessage(e.Embed())
	} else if e.IsEnum() {
		return fns.GetImportPackageEnum(e.Enum())
	}

	return ""
}

// GetImportPackageEnum returns the Go package path of a given protocol buffer entity based on its file descriptor.
func (fns Functions) GetImportPackageEnum(e pgs.Entity) string {
	//if strings.HasPrefix(e.File().FullyQualifiedName(), ".platform.spec.") {
	//	return fns.GetRemoteProtoGoPathFromFile(e.File())
	//}
	//if strings.HasPrefix(e.File().FullyQualifiedName(), ".platform.") {
	//	return fns.GetRemoteGrpcGoPathFromFile(e.File())
	//}
	o := e.File().Descriptor().GetOptions()
	path := strings.Split(*o.GoPackage, ";")
	return path[0]
}

// DoesImportPathContainAnyPb checks if the Go import paths for the given proto file contain the "anypb" package.
func (fns Functions) DoesImportPathContainAnyPb(file pgs.File) bool {
	paths := fns.GetAllGoFieldLevelImportPaths(file)
	for _, key := range paths {
		if key == "anypb" {
			return true
		}
	}

	return false
}
