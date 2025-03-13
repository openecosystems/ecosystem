package shared

import (
	"fmt"
	"strings"

	options "libs/protobuf/go/protobuf/gen/platform/options/v2"

	pgs "github.com/lyft/protoc-gen-star/v2"
)

// DefaultLanguages is a map linking LanguageType enum values to their corresponding string representations.
var DefaultLanguages = map[options.LanguageType]string{
	options.LanguageType_LANGUAGE_TYPE_CPLUSPLUS:  "c++",
	options.LanguageType_LANGUAGE_TYPE_GOLANG:     "go",
	options.LanguageType_LANGUAGE_TYPE_JAVA:       "java",
	options.LanguageType_LANGUAGE_TYPE_PYTHON:     "python",
	options.LanguageType_LANGUAGE_TYPE_TYPESCRIPT: "typescript",
	options.LanguageType_LANGUAGE_TYPE_CSHARP:     "csharp",
	options.LanguageType_LANGUAGE_TYPE_SWIFT:      "swift",
	options.LanguageType_LANGUAGE_TYPE_ANDROID:    "android",
	options.LanguageType_LANGUAGE_TYPE_GRAPHQL:    "graphql",
	options.LanguageType_LANGUAGE_TYPE_LUA:        "lua",
}

// DefaultLanguageTypes represents a predefined list of commonly supported programming language types.
var DefaultLanguageTypes = []options.LanguageType{
	options.LanguageType_LANGUAGE_TYPE_CPLUSPLUS,
	options.LanguageType_LANGUAGE_TYPE_GOLANG,
	options.LanguageType_LANGUAGE_TYPE_JAVA,
	options.LanguageType_LANGUAGE_TYPE_PYTHON,
	options.LanguageType_LANGUAGE_TYPE_TYPESCRIPT,
	options.LanguageType_LANGUAGE_TYPE_CSHARP,
	options.LanguageType_LANGUAGE_TYPE_SWIFT,
	options.LanguageType_LANGUAGE_TYPE_ANDROID,
	options.LanguageType_LANGUAGE_TYPE_GRAPHQL,
	options.LanguageType_LANGUAGE_TYPE_LUA,
}

// IsJava checks if the provided file is configured for Java generation and matches the given language string "java".
func (fns Functions) IsJava(file pgs.File, mlanguage string) bool {
	if opts := file.Descriptor().GetOptions(); opts != nil && opts.GetJavaMultipleFiles() && mlanguage == "java" {
		return true
	}

	return false
}

// IsCSharp checks if the given file has C# as the specified language and meets the required options criteria.
func (fns Functions) IsCSharp(file pgs.File, mlanguage string) bool {
	if opts := file.Descriptor().GetOptions(); opts != nil && opts.GetJavaMultipleFiles() && mlanguage == "csharp" {
		return true
	}

	return false
}

// IsGolang checks if the provided file has Java multiple files option enabled and the specified language is "golang".
func (fns Functions) IsGolang(file pgs.File, mlanguage string) bool {
	if opts := file.Descriptor().GetOptions(); opts != nil && opts.GetJavaMultipleFiles() && mlanguage == "golang" {
		return true
	}

	return false
}

// GoPath extracts the second component of the GoPackage option from the given protobuf file's descriptor.
func (fns Functions) GoPath(file pgs.File) string {
	p := file.Descriptor().GetOptions().GoPackage
	path := strings.Split(*p, ";")

	return path[1]
}

// GoPackage extracts and returns the Go package path from the FileDescriptor's options.
func (fns Functions) GoPackage(file pgs.File) string {
	p := file.Descriptor().GetOptions().GoPackage
	path := strings.Split(*p, ";")

	return path[0]
}

// GoPackageOverwrite retrieves or constructs a custom Go package name for a given file and base name.
// It checks the file's options for `go_package` and `java_package`, then modifies or defaults values as necessary.
func (fns Functions) GoPackageOverwrite(file pgs.File, name string) string {
	o := file.Descriptor().GetOptions()
	if o == nil {
		return fmt.Sprintf("%s-%s line83", name, file.Name().String())
	}
	p := o.GoPackage
	if p == nil {
		if o.JavaPackage != nil {
			return *o.JavaPackage
		}
		return name
	}
	path := strings.Split(*p, ";")
	r := strings.Split(path[0], "/")
	return strings.Replace(path[0], r[2], name, 1)
}

// LanguageOptions retrieves language options for the given file or assigns default options if none are specified.
func (fns Functions) LanguageOptions(file pgs.File) options.LanguageOptions {
	var lOptions options.LanguageOptions

	_, err := file.Extension(options.E_Language, &lOptions)
	if err != nil {
		lOptions.Languages = DefaultLanguageTypes
	}

	if len(lOptions.GetLanguages()) == 0 {
		lOptions.Languages = DefaultLanguageTypes
	}

	return lOptions
}

// SupportedLanguages generates a map of supported languages and their corresponding LanguageType from a given list.
func (fns Functions) SupportedLanguages(lTypes []options.LanguageType) map[string]options.LanguageType {
	supportedLanguages := make(map[string]options.LanguageType)

	for _, l := range lTypes {
		supportedLanguages[DefaultLanguages[l]] = l
	}

	return supportedLanguages
}

// IsSupportedLanguage checks if the given mlanguage is included in the supported languages defined by lOptions.
func (fns Functions) IsSupportedLanguage(lOptions options.LanguageOptions, mlanguage string) bool {
	supportedLanguages := fns.SupportedLanguages(lOptions.Languages)

	if _, found := supportedLanguages[mlanguage]; found {
		return true
	}

	return false
}
