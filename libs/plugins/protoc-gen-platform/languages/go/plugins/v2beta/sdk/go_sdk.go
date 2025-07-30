package sdkv2beta

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

var (
	//go:embed templates/*.tmpl
	templates embed.FS
	goOutPath = pgs.JoinPaths("")
	outPath   = &goOutPath
)

const (
	language   = "go"
	pluginName = "sdk-v2beta"
)

type GoSdkModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	Tpl *template.Template
}

func GoSdkPlugin() *GoSdkModule {
	return &GoSdkModule{ModuleBase: &pgs.ModuleBase{}}
}

func (m *GoSdkModule) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())
}

func (m *GoSdkModule) Name() string { return language + "/" + pluginName }

func (m *GoSdkModule) Execute(targets map[string]pgs.File, _ map[string]pgs.Package) []pgs.Artifact {
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

	// Idempotent looping, use keys for range NOT targets
	keys := make([]string, 0)
	for k := range targets {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		t := targets[k]
		m.GenerateClientFile(t)
		// m.GenerateProjectJsonFile(t)
		// m.GeneratePackageJsonFile(t)
		// m.GenerateGoModFile(t)
		// m.GenerateGoReleaserFile(t)
		m.GenerateReadmeFile(t)
	}

	return m.Artifacts()
}

func (m GoSdkModule) GenerateClientFile(file pgs.File) {
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

	msg := fns.Entity(file)
	if msg == nil {
		return
	}

	_scope, _system, _version, _ := fns.GetPackageScopeSystemAndVersion(file)
	scope := _scope.LowerCamelCase().String()
	system := _system.LowerCamelCase().String()
	version := _version.String()
	gopackage := fns.GetGoPackageFromFile(file)
	fileName := fns.ProtoName(file)

	// system := fns.DomainSystemName2(file).LowerCamelCase().String()
	// version := fns.GetPackageVersion(file)

	clientFileName := outPath.SetExt("/" + scope + "/" + system + "/" + version + "/" + gopackage + "pbconnect" + "/" + fileName + ".client.go").String()
	// clientFileName := strings.TrimPrefix(m.ctx.OutputPath(file).SetExt(".client.go").String(), "platform/")
	m.OverwriteGeneratorTemplateFile(clientFileName, m.Tpl, file)
}

func (m GoSdkModule) GenerateProjectJsonFile(file pgs.File) {
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
	version := _version.String()

	projectJsonFileName := outPath.SetExt("/" + scope + "/" + system + "/" + version + "/" + "project.json").String()
	m.OverwriteGeneratorTemplateFile(projectJsonFileName, m.Tpl, file)
}

func (m GoSdkModule) GeneratePackageJsonFile(file pgs.File) {
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
	version := _version.String()

	projectJsonFileName := outPath.SetExt("/" + scope + "/" + system + "/" + version + "/" + "package.json").String()
	m.OverwriteGeneratorTemplateFile(projectJsonFileName, m.Tpl, file)
}

func (m GoSdkModule) GenerateGoModFile(file pgs.File) {
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
	version := _version.String()

	goModFileName := outPath.SetExt("/" + scope + "/" + system + "/" + version + "/" + "go.mod").String()
	m.OverwriteGeneratorTemplateFile(goModFileName, m.Tpl, file)
}

func (m GoSdkModule) GenerateGoReleaserFile(file pgs.File) {
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
	version := _version.String()

	name := outPath.SetExt("/" + scope + "/" + system + "/" + version + "/" + ".goreleaser.yaml").String()
	m.OverwriteGeneratorTemplateFile(name, m.Tpl, file)
}

func (m GoSdkModule) GenerateReadmeFile(file pgs.File) {
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

	msg := fns.Entity(file)
	if msg == nil {
		return
	}

	_scope, _system, _version, _ := fns.GetPackageScopeSystemAndVersion(file)
	scope := _scope.LowerCamelCase().String()
	system := _system.LowerCamelCase().String()
	version := _version.String()

	name := outPath.SetExt("/" + scope + "/" + system + "/" + version + "/" + "README.md").String()
	m.OverwriteGeneratorTemplateFile(name, m.Tpl, file)
}
