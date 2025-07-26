package sdkconnectorv2beta

import (
	"embed"
	"sort"
	"strings"
	"text/template"

	_go "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go"
	shared "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/shared"

	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
)

// templates contains embedded file system for template files matching the pattern "templates/*.tmpl".
// goOutPath stores the joined output path for generated files.
// outPath is a pointer to the goOutPath variable.
var (
	//go:embed templates/*.tmpl
	templates embed.FS
	goOutPath = pgs.JoinPaths("")
	outPath   = &goOutPath
)

// language represents the programming language being used.
// pluginName is the name of the plugin.
// rootFolder defines the name of the root folder.
// gitHubRepository specifies the GitHub repository associated with the project.
// cdToRoot indicates how many levels to traverse to reach the root directory.
const (
	language         = "go"
	pluginName       = "sdk-connector-v2beta"
	rootFolder       = "root-folder"
	gitHubRepository = "github-repository"

	// How many levels should we go up to find the root. For example: "../../../../../"
	cdToRoot = "cd-to-root"
)

// GoSdkConnectorModule provides functionalities to integrate and customize SDK connectors in Go code generation workflows.
type GoSdkConnectorModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	Tpl *template.Template

	data data
}

// data represents a structure containing information related to file handling and repository details.
type data struct {
	CDToRootFolder   string
	PgsFile          *pgs.File
	GitHubRepository string

	RelativePath string
}

// GoSdkConnectorPlugin initializes and returns a new instance of GoSdkConnectorModule with a default ModuleBase configuration.
func GoSdkConnectorPlugin() *GoSdkConnectorModule {
	return &GoSdkConnectorModule{ModuleBase: &pgs.ModuleBase{}}
}

// InitContext initializes the module's context with the provided BuildContext and parameters.
func (m *GoSdkConnectorModule) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())
}

// Name returns the concatenation of the language and pluginName values as a single string.
func (m *GoSdkConnectorModule) Name() string { return language + "/" + pluginName }

// Execute processes the provided targets and generates necessary files based on module parameters and configurations.
func (m *GoSdkConnectorModule) Execute(targets map[string]pgs.File, _ map[string]pgs.Package) []pgs.Artifact {
	paramLanguage := m.Parameters().Str(shared.LanguageParam)
	m.Assert(paramLanguage != "", shared.LanguageParamError)
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}

	if paramLanguage != strings.ToLower(language) {
		return nil
	}

	paramType := m.Parameters().Str(shared.TypeParam)
	m.Assert(paramType != "", shared.TypeParamError)

	if paramType != pluginName {
		return nil
	}

	// Idempotent looping, use keys for range NOT targets
	keys := make([]string, 0)
	for k := range targets {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		t := targets[k]
		m.data = data{
			PgsFile:          &t,
			GitHubRepository: m.Parameters().Str(gitHubRepository),
			CDToRootFolder:   m.Parameters().Str(cdToRoot),
			RelativePath:     strings.TrimPrefix(fns.GoPackage(t), m.Parameters().Str(gitHubRepository)+"/"),
		}
		m.GenerateClientFile(t)
		// m.GenerateProjectJsonFile(t)
		// m.GeneratePackageJsonFile(t)
		m.GenerateGoModFile(t)
		m.GenerateGoReleaserFile(t)
		m.GenerateReadmeFile(t)
	}

	return m.Artifacts()
}

// GenerateClientFile generates a client file for the provided protobuf file using the defined template and functions.
func (m GoSdkConnectorModule) GenerateClientFile(file pgs.File) {
	templateName := "file.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"getGithubRepositoryConstant": fns.GetGithubRepositoryConstant,
		"service":                     fns.Service,
		"getRoutineMessage":           fns.GetRoutineMessage,
		"getRoutineMessageFieldName":  fns.GetRoutineMessageFieldName,
		"parentService":               fns.ParentService,
		"queries":                     fns.QueryMethods,
		"mutations":                   fns.MutationMethods,
		"getRoutines":                 fns.GetRoutines,
		"getImportPackages":           fns.GetGoImportPackagesServer,
		"goPackageOverwrite":          fns.GoPackageOverwrite,
		"goPackageRemote":             fns.GetRemoteProtoGoPathFromFile,
		"getPackageVersion":           fns.GetPackageVersion,
		"getApiOptionsTypeName":       fns.GetApiOptionsTypeName,
		"domainSystemName2":           fns.DomainSystemName2,
		"getApiOptionsNetwork":        fns.GetApiOptionsNetwork,
		"goPackage":                   fns.GoPackage,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	if !fns.HasService(file) {
		return
	}

	_scope, _system, _version, _ := fns.GetPackageScopeSystemAndVersion(file)
	scope := _scope.LowerCamelCase().String()
	system := _system.LowerCamelCase().String()
	systemFlattened := strings.ReplaceAll(system, "_", "")
	version := _version.LowerCamelCase().String()
	fileName := fns.ProtoName(file)

	// packagePath := fns.DotNotationToFilePath(file.Package().ProtoName().String())

	clientFileName := outPath.SetExt("/" + scope + "/" + system + "/" + version + "/" + systemFlattened + version + "pbconnect" + "/" + fileName + ".client.go").String()
	// clientFileName := strings.TrimPrefix(m.ctx.OutputPath(file).SetExt(".client.go").String(), "platform/")
	m.OverwriteGeneratorTemplateFile(clientFileName, m.Tpl, file)
}

// GenerateProjectJsonFile generates a "project.json" file for the provided protobuf file using a predefined template.
// It initializes and processes the Go template with custom functions and file context, creating output in a structured path.
func (m GoSdkConnectorModule) GenerateProjectJsonFile(file pgs.File) {
	templateName := "project.json.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"getGithubRepositoryConstant": fns.GetGithubRepositoryConstant,
		"service":                     fns.Service,
		"getRoutineMessage":           fns.GetRoutineMessage,
		"getRoutineMessageFieldName":  fns.GetRoutineMessageFieldName,
		"parentService":               fns.ParentService,
		"queries":                     fns.QueryMethods,
		"mutations":                   fns.MutationMethods,
		"getRoutines":                 fns.GetRoutines,
		"getImportPackages":           fns.GetGoImportPackagesServer,
		"goPackageOverwrite":          fns.GoPackageOverwrite,
		"goPackageRemote":             fns.GetRemoteProtoGoPathFromFile,
		"getPackageVersion":           fns.GetPackageVersion,
		"getApiOptionsTypeName":       fns.GetApiOptionsTypeName,
		"domainSystemName2":           fns.DomainSystemName2,
		"getApiOptionsNetwork":        fns.GetApiOptionsNetwork,
		"goPackage":                   fns.GoPackage,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	_scope, _system, _version, _ := fns.GetPackageScopeSystemAndVersion(file)
	scope := _scope.LowerCamelCase().String()
	system := _system.LowerCamelCase().String()
	version := _version.LowerCamelCase().String()

	projectJsonFileName := outPath.SetExt("/" + scope + "/" + system + "/" + version + "/" + "project.json").String()
	m.OverwriteGeneratorTemplateFile(projectJsonFileName, m.Tpl, m.data)
}

// GeneratePackageJsonFile generates a package.json file using the provided file, templates, and context for configuration.
func (m GoSdkConnectorModule) GeneratePackageJsonFile(file pgs.File) {
	templateName := "package.json.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"getGithubRepositoryConstant": fns.GetGithubRepositoryConstant,
		"service":                     fns.Service,
		"getRoutineMessage":           fns.GetRoutineMessage,
		"getRoutineMessageFieldName":  fns.GetRoutineMessageFieldName,
		"parentService":               fns.ParentService,
		"queries":                     fns.QueryMethods,
		"mutations":                   fns.MutationMethods,
		"getRoutines":                 fns.GetRoutines,
		"getImportPackages":           fns.GetGoImportPackagesServer,
		"goPackageOverwrite":          fns.GoPackageOverwrite,
		"goPackageRemote":             fns.GetRemoteProtoGoPathFromFile,
		"getPackageVersion":           fns.GetPackageVersion,
		"getApiOptionsTypeName":       fns.GetApiOptionsTypeName,
		"domainSystemName2":           fns.DomainSystemName2,
		"getApiOptionsNetwork":        fns.GetApiOptionsNetwork,
		"goPackage":                   fns.GoPackage,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	_scope, _system, _version, _ := fns.GetPackageScopeSystemAndVersion(file)
	scope := _scope.LowerCamelCase().String()
	system := _system.LowerCamelCase().String()
	version := _version.LowerCamelCase().String()

	projectJsonFileName := outPath.SetExt("/" + scope + "/" + system + "/" + version + "/" + "package.json").String()
	m.OverwriteGeneratorTemplateFile(projectJsonFileName, m.Tpl, m.data)
}

// GenerateGoModFile generates a Go module file (go.mod) using a template and file-specific context data.
// It dynamically sets up template functions and renders the content to the appropriate file path.
func (m GoSdkConnectorModule) GenerateGoModFile(file pgs.File) {
	templateName := "go.mod.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"getGithubRepositoryConstant": fns.GetGithubRepositoryConstant,
		"service":                     fns.Service,
		"getRoutineMessage":           fns.GetRoutineMessage,
		"getRoutineMessageFieldName":  fns.GetRoutineMessageFieldName,
		"parentService":               fns.ParentService,
		"queries":                     fns.QueryMethods,
		"mutations":                   fns.MutationMethods,
		"getRoutines":                 fns.GetRoutines,
		"getImportPackages":           fns.GetGoImportPackagesServer,
		"goPackageOverwrite":          fns.GoPackageOverwrite,
		"goPackageRemote":             fns.GetRemoteProtoGoPathFromFile,
		"getPackageVersion":           fns.GetPackageVersion,
		"getApiOptionsTypeName":       fns.GetApiOptionsTypeName,
		"domainSystemName2":           fns.DomainSystemName2,
		"getApiOptionsNetwork":        fns.GetApiOptionsNetwork,
		"goPackage":                   fns.GoPackage,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	_scope, _system, _version, _ := fns.GetPackageScopeSystemAndVersion(file)
	scope := _scope.LowerCamelCase().String()
	system := _system.LowerCamelCase().String()
	version := _version.LowerCamelCase().String()

	goModFileName := outPath.SetExt("/" + scope + "/" + system + "/" + version + "/" + "go.mod").String()
	m.OverwriteGeneratorTemplateFile(goModFileName, m.Tpl, m.data)
}

// GenerateGoReleaserFile generates a GoReleaser configuration file using the provided protobuf file and predefined templates.
func (m GoSdkConnectorModule) GenerateGoReleaserFile(file pgs.File) {
	templateName := "goreleaser.yaml.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"getGithubRepositoryConstant": fns.GetGithubRepositoryConstant,
		"service":                     fns.Service,
		"getRoutineMessage":           fns.GetRoutineMessage,
		"getRoutineMessageFieldName":  fns.GetRoutineMessageFieldName,
		"parentService":               fns.ParentService,
		"queries":                     fns.QueryMethods,
		"mutations":                   fns.MutationMethods,
		"getRoutines":                 fns.GetRoutines,
		"getImportPackages":           fns.GetGoImportPackagesServer,
		"goPackageOverwrite":          fns.GoPackageOverwrite,
		"goPackageRemote":             fns.GetRemoteProtoGoPathFromFile,
		"getPackageVersion":           fns.GetPackageVersion,
		"getApiOptionsTypeName":       fns.GetApiOptionsTypeName,
		"domainSystemName2":           fns.DomainSystemName2,
		"getApiOptionsNetwork":        fns.GetApiOptionsNetwork,
		"goPackage":                   fns.GoPackage,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	_scope, _system, _version, _ := fns.GetPackageScopeSystemAndVersion(file)
	scope := _scope.LowerCamelCase().String()
	system := _system.LowerCamelCase().String()
	version := _version.LowerCamelCase().String()

	name := outPath.SetExt("/" + scope + "/" + system + "/" + version + "/" + ".goreleaser.sdk.yaml").String()
	m.OverwriteGeneratorTemplateFile(name, m.Tpl, m.data)
}

// GenerateReadmeFile generates a README.md file based on a predefined template and file-specific parameters.
// It initializes template functions and uses them to populate the content dynamically.
// The resulting README file includes metadata and documentation tailored to the provided file context.
func (m GoSdkConnectorModule) GenerateReadmeFile(file pgs.File) {
	templateName := "readme.md.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"getGithubRepositoryConstant": fns.GetGithubRepositoryConstant,
		"service":                     fns.Service,
		"getRoutineMessage":           fns.GetRoutineMessage,
		"getRoutineMessageFieldName":  fns.GetRoutineMessageFieldName,
		"parentService":               fns.ParentService,
		"queries":                     fns.QueryMethods,
		"mutations":                   fns.MutationMethods,
		"getRoutines":                 fns.GetRoutines,
		"getImportPackages":           fns.GetGoImportPackagesServer,
		"goPackageOverwrite":          fns.GoPackageOverwrite,
		"goPackageRemote":             fns.GetRemoteProtoGoPathFromFile,
		"getPackageVersion":           fns.GetPackageVersion,
		"getApiOptionsTypeName":       fns.GetApiOptionsTypeName,
		"domainSystemName2":           fns.DomainSystemName2,
		"getApiOptionsNetwork":        fns.GetApiOptionsNetwork,
		"goPackage":                   fns.GoPackage,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	_scope, _system, _version, _ := fns.GetPackageScopeSystemAndVersion(file)
	scope := _scope.LowerCamelCase().String()
	system := _system.LowerCamelCase().String()
	version := _version.LowerCamelCase().String()

	name := outPath.SetExt("/" + scope + "/" + system + "/" + version + "/" + "README.md").String()
	m.OverwriteGeneratorTemplateFile(name, m.Tpl, file)
}
