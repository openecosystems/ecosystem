package shared

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	"strings"
)

// GetGoImportPathAndPackageForAnyFieldTypeElement handles Embedded Types
func (fns Functions) GetGoImportPathAndPackageForAnyFieldTypeElement(e pgs.FieldTypeElem) (string, string) {
	if e.IsEmbed() {
		return fns.GetImportPathAndPackageForAnyMessage(e.Embed())
	} else if e.IsEnum() {
		return fns.GetImportPathAndPackageForAnyEntity(e.Enum())
	}

	return "", ""
}

func (fns Functions) GetImportPathAndPackageForAnyMessage(msg pgs.Message) (string, string) {
	if msg.IsWellKnown() {
		switch msg.WellKnownType() {
		case pgs.AnyWKT:
			return "google.golang.org/protobuf/types/known/anypb", "anypb"
		case pgs.DurationWKT:
			return "google.golang.org/protobuf/types/known/durationpb", "durationpb"
		case pgs.TimestampWKT:
			return "google.golang.org/protobuf/types/known/timestamppb", "timestamppb"
		}
	}

	return fns.GetImportPathAndPackageForAnyEntity(msg)
}

func (fns Functions) GetImportPathAndPackageForAnyEntity(e pgs.Entity) (string, string) {

	o := e.File().Descriptor().GetOptions()
	path := strings.Split(*o.GoPackage, ";")
	return path[0], path[1]
}

func (fns Functions) GetAllGoFileLevelImportPaths(file pgs.File) map[string]string {

	o := file.Descriptor().GetOptions()

	p := o.GoPackage
	path := strings.Split(*p, ";")

	imports := make(map[string]string)

	imports[path[0]+path[1]] = path[1]
	imports[path[0]+"/"+path[1]+"grpc"] = path[1] + "grpc"

	// serviceProto
	// serviceGrpc

	//imports["specv2"] = "anypb"

	return imports
}

// GetAllGoFieldLevelImportPaths loops through file and identifies all imports
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

func (fns Functions) SelectivelyAddImport(key string, file pgs.File) bool {

	path, _ := fns.GetImportPathAndPackageForAnyEntity(file)
	if key != "" && !strings.Contains(key, path) {
		return true
	}

	return false
}

// GetGoImportPathForMessage given a Message, file the go path of that Message type
func (fns Functions) GetGoImportPathForMessage(msg pgs.Message) string {
	if msg.IsWellKnown() {
		switch msg.WellKnownType() {
		case pgs.AnyWKT:
			return "google.golang.org/protobuf/types/known/anypb"
		case pgs.DurationWKT:
			return "google.golang.org/protobuf/types/known/durationpb"
		case pgs.TimestampWKT:
			return "google.golang.org/protobuf/types/known/timestamppb"
		}

	}

	return fns.GetGoImportPackageAndPathFromAnyEntity(msg)
}

// GetGoImportPackageAndPathFromAnyEntity given a pgs.Entity, find the Go path, including remote libraries
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

// GetGoImportPathsElement handles Embedded Types
func (fns Functions) GetGoImportPathsElement(e pgs.FieldTypeElem) string {
	if e.IsEmbed() {
		return fns.GetImportPackageMessage(e.Embed())
	} else if e.IsEnum() {
		return fns.GetImportPackageEnum(e.Enum())
	}

	return ""
}

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

func (fns Functions) DoesImportPathContainAnyPb(file pgs.File) bool {

	paths := fns.GetAllGoFieldLevelImportPaths(file)
	for _, key := range paths {
		if key == "anypb" {
			return true
		}
	}

	return false
}
