package clisystemsv2beta

import (
	"embed"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	_go "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go"
	shared "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/shared"

	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
)

var (
	//go:embed templates/*.tmpl
	templates embed.FS
	goOutPath = pgs.JoinPaths("")
	outPath   = &goOutPath
)

const (
	language   = "go"
	pluginName = "cli-systems-v2beta"
)

// GoCliSystemsModule represents a Protoc plugin module for generating Go CLI system files.
// It extends pgs.ModuleBase and utilizes a custom PGS context and Go templates for file generation.
type GoCliSystemsModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	Tpl *template.Template
}

// GoCliSystemsPlugin creates and returns a new instance of GoCliSystemsModule for use with Protobuf code generation.
func GoCliSystemsPlugin() *GoCliSystemsModule {
	return &GoCliSystemsModule{ModuleBase: &pgs.ModuleBase{}}
}

// InitContext initializes the module's context by setting the base context and initializing the pgsgo context using parameters.
func (m *GoCliSystemsModule) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())
}

// Name returns the constructed name of the module composed of the defined language and plugin name constants.
func (m *GoCliSystemsModule) Name() string { return language + "/" + pluginName }

// Execute processes protocol buffer targets, applies filters based on plugin and language parameters, and generates artifacts.
func (m *GoCliSystemsModule) Execute(targets map[string]pgs.File, _ map[string]pgs.Package) []pgs.Artifact {
	paramLanguage := m.Parameters().Str(shared.LanguageParam)
	m.Assert(paramLanguage != "", shared.LanguageParamError)

	if paramLanguage != strings.ToLower(language) {
		return nil
	}

	paramType := m.Parameters().Str(shared.TypeParam)
	m.Assert(paramType != "", shared.TypeParamError)

	if paramType != pluginName {
		return nil
	}

	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}

	// Idempotent looping, use keys for range NOT targets
	keys := make([]string, 0)
	for k := range targets {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		t := targets[k]
		msg := fns.Entity(t)
		if msg == nil {
			continue
		}
		if !fns.IsSupportedLanguage(fns.LanguageOptions(t), paramLanguage) {
			continue
		}
		m.GenerateFile(t)
	}

	return m.Artifacts()
}

// GenerateFile generates the output file for the provided proto file using a specified template and settings.
// It initializes a template with required functions, parses template files, and writes the generated content.
// The output path is dynamically created based on domain system name and package version.
// If no valid message is found in the proto file, the method terminates early without generating a file.
func (m GoCliSystemsModule) GenerateFile(file pgs.File) {
	templateName := "file.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())
	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"service":                        fns.Service,
		"serviceLeadingComment":          fns.ServiceLeadingComment,
		"serviceTrailingComment":         fns.ServiceTrailingComment,
		"serviceLeadingDetachedComments": fns.ServiceLeadingDetachedComments,
		"parentService":                  fns.ParentService,
		"queries":                        fns.QueryMethods,
		"isMethodQuery":                  fns.IsMethodQuery,
		"mutations":                      fns.MutationMethods,
		"isMethodMutation":               fns.IsMethodMutation,
		"getCqrsType":                    fns.GetCQRSType,
		"getSpecCommands":                fns.GetSpecCommands,
		"getSpecEvents":                  fns.GetSpecEvents,
		"getSpecTopics":                  fns.GetSpecTopics,
		"goPackageOverwrite":             fns.GoPackageOverwrite,
		"goPackage":                      fns.GoPackage,
		"goPath":                         fns.GoPath,
		"allMethods":                     fns.AllMethods,
		"dotNotationToFilePath":          fns.DotNotationToFilePath,
		"getPackageVersion":              fns.GetPackageVersion,
		"getPackageVersionName":          fns.GetPackageVersionName,
		"domainSystemName2":              fns.DomainSystemName2,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	msg := fns.Entity(file)
	if msg == nil {
		return
	}

	system := fns.DomainSystemName2(file).LowerCamelCase().String()
	version := fns.GetPackageVersion(file)
	fullPath := file.InputPath().String()
	dirPath := filepath.Dir(fullPath)

	name := outPath.SetExt("/" + dirPath + "/" + system + version + "pbcli" + "/" + "systems.cmd.go").String()
	m.OverwriteGeneratorTemplateFile(name, m.Tpl, file)
}
