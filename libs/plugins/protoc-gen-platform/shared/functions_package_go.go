package shared

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	"regexp"
	"sort"
	"strings"
)

func (fns Functions) GetAnyGoFieldPackage(f pgs.Field) string {

	//if f.Type().IsMap() {
	//
	//  if f.Type().Key().ProtoType() == pgs. {
	//    return MapEnumField
	//  }
	//
	//  keyPath, keyPkg := fns.GetGoImportPathAndPackageForAnyFieldTypeElement(f.Type().Key())
	//
	//  valuePath, valuePkg := fns.GetGoImportPathAndPackageForAnyFieldTypeElement(f.Type().Element())
	//  if fns.SelectivelyAddImport(valuePath, file) {
	//    imports[valuePath] = valuePkg
	//  }
	//} else if f.Type().IsRepeated() {
	//  repeatedPath, repeatedPkg := fns.GetGoImportPathAndPackageForAnyFieldTypeElement(f.Type().Element())
	//  if fns.SelectivelyAddImport(repeatedPath, file) {
	//    imports[repeatedPath] = repeatedPkg
	//  }
	//} else if f.Type().IsEmbed() {
	//  embedPath, embedPkg := fns.GetImportPathAndPackageForAnyMessage(f.Type().Embed())
	//  if fns.SelectivelyAddImport(embedPath, file) {
	//    imports[embedPath] = embedPkg
	//  }
	//} else if f.Type().IsEnum() {
	//  enumPath, enumPkg := fns.GetImportPathAndPackageForAnyEntity(f.Type().Enum())
	//  if fns.SelectivelyAddImport(enumPath, file) {
	//    imports[enumPath] = enumPkg
	//  }
	//}

	return ""
}

func (fns Functions) GetSpecTypePackage(file pgs.File) string {

	for _, imp := range file.Imports() {
		if strings.Contains(imp.Name().String(), SpecTypePathPrefix) {
			_, pkg := fns.GetImportPathAndPackageForAnyEntity(imp.File())
			return pkg + "."
		}
	}
	return ""
}

func (fns Functions) GetPlatformTypePackage(file pgs.File) string {

	for _, imp := range file.Imports() {
		if strings.Contains(imp.Name().String(), PlatformTypesPathPrefix) {
			_, pkg := fns.GetImportPathAndPackageForAnyEntity(imp.File())
			return pkg + "."
		}
	}
	return ""
}

func (fns Functions) GetGoMapKeyTypePackage(field pgs.Field) string {

	if !field.Type().IsMap() {
		panic("Field must be a map to determine map key")
	}

	path, pkg := fns.GetGoImportPathAndPackageForAnyFieldTypeElement(field.Type().Key())
	return fns.SelectivelyAddPackageName(path, pkg, field)
}

func (fns Functions) GetGoMapValueTypePackage(field pgs.Field) string {

	if !field.Type().IsMap() {
		panic("Field must be a map to determine map value")
	}

	path, pkg := fns.GetGoImportPathAndPackageForAnyFieldTypeElement(field.Type().Element())
	if field.Type().Element().IsEmbed() {
		pkg = "*" + pkg
	}
	return fns.SelectivelyAddPackageName(path, pkg, field)
}

func (fns Functions) GetEnumTypePackage(field pgs.Field) string {
	_, pkg := fns.GetImportPathAndPackageForAnyEntity(field.Type().Enum())
	return pkg
	//return fns.SelectivelyAddPackageName(path, pkg, field)
}

func (fns Functions) GetStructTypePackage(field pgs.Field) string {

	if !fns.IsGoStruct(field) {
		panic("Field must be a struct to determine struct type")
	}

	_, pkg := fns.GetImportPathAndPackageForAnyMessage(field.Type().Embed())

	return pkg
	//return fns.SelectivelyAddPackageName(path, pkg, field)
}

func (fns Functions) GetGoSliceValueTypePackage(field pgs.Field) string {

	if !field.Type().IsRepeated() {
		panic("Field must be a list to determine list value")
	}

	path, pkg := fns.GetGoImportPathAndPackageForAnyFieldTypeElement(field.Type().Element())
	if field.Type().Element().IsEmbed() {
		// These are pointers for some reason
		pkg = "*" + pkg
	}

	return fns.SelectivelyAddPackageName(path, pkg, field)
}

func (fns Functions) SelectivelyAddPackageName(path string, pkg string, field pgs.Field) string {
	p, _ := fns.GetImportPathAndPackageForAnyEntity(field.File())
	if p != "" && strings.Contains(p, path) {
		if field.Type().IsRepeated() && field.Type().Element().IsEmbed() {
			// These are pointers for some reason
			return "*"
		}
		return ""
	}

	return pkg + "."
}

// GetGoPackageFromFile if the File contains a Spec Entity, then we can get the package name. For example, eventv2alpha
func (fns Functions) GetGoPackageFromFile(file pgs.File) string {

	goPkg := fns.GetGoPathFromFile(file)

	return fns.GetPackageNameFromGoPath(goPkg)
}

// GetGoPathFromFile if the File contains a Spec Entity, then we can get the go path. For example,
func (fns Functions) GetGoPathFromFile(file pgs.File) string {

	entity := fns.Entity(file)

	goPkg := fns.GetGoPathFromMessage(entity)

	return goPkg
}

// GetGoPathFromMessage given a Message, file the go path of that Message type
func (fns Functions) GetGoPathFromMessage(msg pgs.Message) string {
	return fns.GetGoImportPathForMessage(msg)
}

// GetGoPathFromAnyPgsEntity given a pgs.Entity, find the Go path, including remote libraries
func (fns Functions) GetGoPathFromAnyPgsEntity(e pgs.Entity) string {
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

// GetPackageNameFromGoPath Split the package name from the fully qualified go path
func (fns Functions) GetPackageNameFromGoPath(goPkg string) string {
	if goPkg == "" {
		return ""
	}

	// assume package naming convention based on fully qualified package import
	parts := strings.Split(goPkg, "/")
	finalPart := parts[len(parts)-1]
	re := regexp.MustCompile(`^v\d+\w*$`)
	if re.MatchString(finalPart) {
		return parts[len(parts)-2] + finalPart
	}

	return finalPart
}

// /////////
// /////////
// ///////////////////////////
func (fns Functions) getImportNameFromPackage(goPkg string, file pgs.File, useAlias bool) string {
	if goPkg == "" {
		return ""
	}

	// assume package naming convention based on fully qualified package import
	parts := strings.Split(goPkg, "/")
	finalPart := parts[len(parts)-1]
	re := regexp.MustCompile(`^v\d+\w*$`)
	if re.MatchString(finalPart) {
		return parts[len(parts)-2] + finalPart + "."
	}

	return finalPart + "."
}

func (fns Functions) getImportNameElem(e pgs.FieldTypeElem, file pgs.File) string {
	goPkg := fns.GetImportPackageElem(e)
	return fns.getImportNameFromPackage(goPkg, file, true)

}

func (fns Functions) GetGoImportNameMessageDirectlyFromGoPackage(msg pgs.Message, file pgs.File) string {
	goPkg := fns.GetImportPackageMessage(msg)
	return fns.getImportNameFromPackage(goPkg, file, true)
}

// Returns the eventv2alpha name of a service
func (fns Functions) GetGoImportNameMessage(msg pgs.Message, file pgs.File) string {
	goPkg := fns.GetImportPackageMessage(msg)
	return fns.getImportNameFromPackage(goPkg, file, true)
}

func (fns Functions) GetGoImportNameMessageNoAlias(msg pgs.Message, file pgs.File) string {
	goPkg := fns.GetImportPackageMessage(msg)
	return fns.getImportNameFromPackage(goPkg, file, false)
}

func (fns Functions) GetImportPackageElem(e pgs.FieldTypeElem) string {
	if e.IsEmbed() {
		return fns.GetImportPackageMessage(e.Embed())
	} else if e.IsEnum() {
		return fns.GetImportPackageEntity(e.Enum())
	}

	return ""
}

func (fns Functions) GetImportPackageMessageDirectlyFromGoPackage(msg pgs.Message) string {
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

	o := msg.File().Descriptor().GetOptions()
	path := strings.Split(*o.GoPackage, ";")
	return path[1]
}

func (fns Functions) GetImportPackageMessage(msg pgs.Message) string {
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

	return fns.GetImportPackageEntity(msg)
}

func (fns Functions) GetImportPackageEntity(e pgs.Entity) string {
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

func (fns Functions) GetGoImportPackages(file pgs.File) []string {

	packageSet := make(map[string]bool)
	for _, msg := range file.AllMessages() {
		for _, f := range msg.Fields() {
			if f.Type().IsMap() {
				keyPkg := fns.GetImportPackageElem(f.Type().Key())
				if keyPkg != "" {
					packageSet[keyPkg] = true
				}
				valuePkg := fns.GetImportPackageElem(f.Type().Element())
				if valuePkg != "" {
					packageSet[valuePkg] = true
				}
			} else if f.Type().IsRepeated() {
				pkg := fns.GetImportPackageElem(f.Type().Element())
				if pkg != "" {
					packageSet[pkg] = true
				}
			} else if f.Type().IsEmbed() {
				pkg := fns.GetImportPackageMessage(f.Type().Embed())
				if pkg != "" {
					packageSet[pkg] = true
				}
			} else if f.Type().IsEnum() {
				pkg := fns.GetImportPackageEntity(f.Type().Enum())
				if pkg != "" {
					packageSet[pkg] = true
				}
			}
		}
	}

	packages := make([]string, 0, len(packageSet))
	for k := range packageSet {
		if strings.Contains(k, "libs/protobuf/go/protobuf/gen/platform/spec/v2") ||
			strings.Contains(k, "jeannotcompany/public/grpc/go/platform") ||
			strings.Contains(k, "jeannotcompany/poc/grpc/go/platform") ||
			strings.Contains(k, "jeannotcompany/partner/grpc/go/platform") ||
			strings.Contains(k, "jeannotcompany/private/grpc/go/platform") {

			continue
		}

		packages = append(packages, k)
	}

	sort.Strings(packages)
	return packages

}

func (fns Functions) GetGoImportPackagesServer(file pgs.File) []string {
	return fns.getGoImportPackages(file, "server", true)
}

func (fns Functions) GetGoImportPackagesClient(file pgs.File) []string {
	return fns.getGoImportPackages(file, "client", false)
}

func (fns Functions) GetGoImportPackagesCLI(file pgs.File) []string {
	return fns.getGoImportPackages(file, "cli-client", false)
}

func (fns Functions) getGoImportPackages(file pgs.File, packageOverwrite string, skipAlias bool) []string {

	packageSet := make(map[string]bool)
	for _, svc := range file.Services() {
		for _, m := range svc.Methods() {
			pkg := fns.GetImportPackageMessage(m.Input())
			if pkg != "" {
				packageSet[pkg] = true
			}
			pkg = fns.GetImportPackageMessage(m.Output())
			if pkg != "" {
				packageSet[pkg] = true
			}
		}
	}

	packages := make([]string, 0, len(packageSet))
	for k := range packageSet {
		if strings.Contains(k, "libs/protobuf/go/protobuf/gen/platform/spec/v2") ||
			strings.Contains(k, "jeannotcompany/public/grpc/go/platform") ||
			strings.Contains(k, "jeannotcompany/poc/grpc/go/platform") ||
			strings.Contains(k, "jeannotcompany/partner/grpc/go/platform") ||
			strings.Contains(k, "jeannotcompany/private/grpc/go/platform") {

			continue
		}

		packages = append(packages, k)
	}

	sort.Strings(packages)
	return packages

}
