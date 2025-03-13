package shared

import (
	"regexp"
	"sort"
	"strings"

	pgs "github.com/lyft/protoc-gen-star/v2"
)

// GetAnyGoFieldPackage returns the Go package name for a given protobuf field based on its type and attributes.
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

// GetSpecTypePackage returns the package name of the first imported file that matches the SpecTypePathPrefix.
// If no match is found, it returns an empty string.
func (fns Functions) GetSpecTypePackage(file pgs.File) string {
	for _, imp := range file.Imports() {
		if strings.Contains(imp.Name().String(), SpecTypePathPrefix) {
			_, pkg := fns.GetImportPathAndPackageForAnyEntity(imp.File())
			return pkg + "."
		}
	}
	return ""
}

// GetPlatformTypePackage extracts and returns the package name of a platform/type import with a trailing dot if found.
func (fns Functions) GetPlatformTypePackage(file pgs.File) string {
	for _, imp := range file.Imports() {
		if strings.Contains(imp.Name().String(), PlatformTypesPathPrefix) {
			_, pkg := fns.GetImportPathAndPackageForAnyEntity(imp.File())
			return pkg + "."
		}
	}
	return ""
}

// GetGoMapKeyTypePackage returns the Go import path and package name for the key type of a map field.
// If the field is not a map, it panics.
// Utilizes SelectivelyAddPackageName to construct the formatted package string.
func (fns Functions) GetGoMapKeyTypePackage(field pgs.Field) string {
	if !field.Type().IsMap() {
		panic("Field must be a map to determine map key")
	}

	path, pkg := fns.GetGoImportPathAndPackageForAnyFieldTypeElement(field.Type().Key())
	return fns.SelectivelyAddPackageName(path, pkg, field)
}

// GetGoMapValueTypePackage determines the Go package path for the value type of a map field in a protobuf definition.
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

// GetEnumTypePackage retrieves the package name of the Enum type associated with the given field.
func (fns Functions) GetEnumTypePackage(field pgs.Field) string {
	_, pkg := fns.GetImportPathAndPackageForAnyEntity(field.Type().Enum())
	return pkg
	// return fns.SelectivelyAddPackageName(path, pkg, field)
}

// GetStructTypePackage returns the package name of the Go struct type associated with the provided field.
// Panics if the field is not a struct.
func (fns Functions) GetStructTypePackage(field pgs.Field) string {
	if !fns.IsGoStruct(field) {
		panic("Field must be a struct to determine struct type")
	}

	_, pkg := fns.GetImportPathAndPackageForAnyMessage(field.Type().Embed())

	return pkg
	// return fns.SelectivelyAddPackageName(path, pkg, field)
}

// GetGoSliceValueTypePackage returns the type package for a repeated field's element, panics if the field is not repeated.
// It determines the import path and package for the element type and conditionally adds the package name.
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

// SelectivelyAddPackageName determines whether to prepend the provided package name to a type based on its context.
// It checks if the field's import path matches the specified path, and selectively includes the package prefix.
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

// GetGoPackageFromFile extracts and returns the Go package name from the provided protobuf file.
func (fns Functions) GetGoPackageFromFile(file pgs.File) string {
	goPkg := fns.GetGoPathFromFile(file)

	return fns.GetPackageNameFromGoPath(goPkg)
}

// GetGoPathFromFile extracts and returns the Go package path from the provided Protobuf file representation.
func (fns Functions) GetGoPathFromFile(file pgs.File) string {
	entity := fns.Entity(file)

	goPkg := fns.GetGoPathFromMessage(entity)

	return goPkg
}

// GetGoPathFromMessage returns the Go import path for a given Protocol Buffers message.
func (fns Functions) GetGoPathFromMessage(msg pgs.Message) string {
	return fns.GetGoImportPathForMessage(msg)
}

// GetGoPathFromAnyPgsEntity determines the Go import path for a given Protocol Buffers entity based on its file properties.
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

// GetPackageNameFromGoPath extracts the Go package name from a given fully qualified package import path.
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

// getImportNameFromPackage extracts the import name from a Go package string, optionally considering alias usage.
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

// getImportNameElem retrieves the import name associated with a given field type element within the context of a file.
func (fns Functions) getImportNameElem(e pgs.FieldTypeElem, file pgs.File) string {
	goPkg := fns.GetImportPackageElem(e)
	return fns.getImportNameFromPackage(goPkg, file, true)
}

// GetGoImportNameMessageDirectlyFromGoPackage retrieves the Go import name for the given Protobuf message directly from its package.
func (fns Functions) GetGoImportNameMessageDirectlyFromGoPackage(msg pgs.Message, file pgs.File) string {
	goPkg := fns.GetImportPackageMessage(msg)
	return fns.getImportNameFromPackage(goPkg, file, true)
}

// GetGoImportNameMessage returns the fully qualified Go import name for a given Protocol Buffers message in a file.
func (fns Functions) GetGoImportNameMessage(msg pgs.Message, file pgs.File) string {
	goPkg := fns.GetImportPackageMessage(msg)
	return fns.getImportNameFromPackage(goPkg, file, true)
}

// GetGoImportNameMessageNoAlias returns the Go import name for a message without including aliasing for the import.
func (fns Functions) GetGoImportNameMessageNoAlias(msg pgs.Message, file pgs.File) string {
	goPkg := fns.GetImportPackageMessage(msg)
	return fns.getImportNameFromPackage(goPkg, file, false)
}

// GetImportPackageElem returns the import package path for a given FieldTypeElem based on whether it is an embed or an enum.
func (fns Functions) GetImportPackageElem(e pgs.FieldTypeElem) string {
	if e.IsEmbed() {
		return fns.GetImportPackageMessage(e.Embed())
	} else if e.IsEnum() {
		return fns.GetImportPackageEntity(e.Enum())
	}

	return ""
}

// GetImportPackageMessageDirectlyFromGoPackage determines the Go package to import for a protobuf message.
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

// GetImportPackageMessage returns the Go import path for a given Protobuf message based on its type or entity information.
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

// GetImportPackageEntity returns the Go import package path for the provided protocol buffer entity.
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

// GetGoImportPackages extracts and returns a sorted list of unique Go import package paths required by the given proto file.
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
		if strings.Contains(k, "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/spec/v2") ||
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

// GetGoImportPackagesServer retrieves a list of Go import packages required for the server generation of the given proto file.
func (fns Functions) GetGoImportPackagesServer(file pgs.File) []string {
	return fns.getGoImportPackages(file, "server", true)
}

// GetGoImportPackagesClient returns a list of Go package imports required for the client based on the provided file.
func (fns Functions) GetGoImportPackagesClient(file pgs.File) []string {
	return fns.getGoImportPackages(file, "client", false)
}

// GetGoImportPackagesCLI returns a list of Go import packages for CLI client generation based on the provided file.
func (fns Functions) GetGoImportPackagesCLI(file pgs.File) []string {
	return fns.getGoImportPackages(file, "cli-client", false)
}

// getGoImportPackages extracts a list of unique importable Go packages based on input/output message dependencies in a file.
// It filters out specific package patterns and ensures the result is sorted alphabetically.
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
		if strings.Contains(k, "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/spec/v2") ||
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
