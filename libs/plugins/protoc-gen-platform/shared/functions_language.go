package shared

import (
	"fmt"
	"strings"

	pgs "github.com/lyft/protoc-gen-star/v2"
	options "libs/protobuf/go/protobuf/gen/platform/options/v2"
)

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

func (fns Functions) IsJava(file pgs.File, mlanguage string) bool {

	if opts := file.Descriptor().GetOptions(); opts != nil && opts.GetJavaMultipleFiles() && mlanguage == "java" {
		return true
	}

	return false
}

func (fns Functions) IsCSharp(file pgs.File, mlanguage string) bool {

	if opts := file.Descriptor().GetOptions(); opts != nil && opts.GetJavaMultipleFiles() && mlanguage == "csharp" {
		return true
	}

	return false
}

func (fns Functions) IsGolang(file pgs.File, mlanguage string) bool {

	if opts := file.Descriptor().GetOptions(); opts != nil && opts.GetJavaMultipleFiles() && mlanguage == "golang" {
		return true
	}

	return false
}

func (fns Functions) GoPath(file pgs.File) string {

	p := file.Descriptor().GetOptions().GoPackage
	path := strings.Split(*p, ";")

	return path[1]

}

func (fns Functions) GoPackage(file pgs.File) string {

	p := file.Descriptor().GetOptions().GoPackage
	path := strings.Split(*p, ";")

	return path[0]

}

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

func (fns Functions) SupportedLanguages(lTypes []options.LanguageType) map[string]options.LanguageType {

	supportedLanguages := make(map[string]options.LanguageType)

	for _, l := range lTypes {
		supportedLanguages[DefaultLanguages[l]] = l
	}

	return supportedLanguages
}

func (fns Functions) IsSupportedLanguage(lOptions options.LanguageOptions, mlanguage string) bool {

	supportedLanguages := fns.SupportedLanguages(lOptions.Languages)

	if _, found := supportedLanguages[mlanguage]; found {
		return true
	}

	return false
}
