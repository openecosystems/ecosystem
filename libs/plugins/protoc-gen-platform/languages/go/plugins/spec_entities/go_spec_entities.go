package spec_entities

import (
	"embed"
	"sort"
	"strings"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	_go "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go"
	"github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/shared"
)

var (
	//go:embed templates/*.tmpl
	templates embed.FS
	fileName  = "combineFile"
	goOutPath = pgs.JoinPaths("platform", "spec")
	outPath   = &goOutPath
)

const (
	language   = "go"
	pluginName = "spec-entities"
)

type GoSpecEntitiesModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	Tpl *template.Template
}

func GoSpecEntitiesPlugin() *GoSpecEntitiesModule {
	return &GoSpecEntitiesModule{ModuleBase: &pgs.ModuleBase{}}
}

func (m *GoSpecEntitiesModule) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())
}

func (m *GoSpecEntitiesModule) Name() string { return language + "/" + pluginName }

func (m *GoSpecEntitiesModule) Execute(targets map[string]pgs.File, _ map[string]pgs.Package) []pgs.Artifact {
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
	versionedKeys := make(map[string][]string, 0)
	for k := range targets {
		version := fns.GetPackageVersion(targets[k])
		versionedKeys[version] = append(versionedKeys[version], k)
	}

	for _, v := range versionedKeys {
		sort.Strings(v)
	}

	for _, keys := range versionedKeys {
		for _, k := range keys {
			t := targets[k]
			if !fns.IsSupportedLanguage(fns.LanguageOptions(t), paramLanguage) {
				continue
			}

			msg := fns.Entity(t)
			if msg == nil {
				continue
			}

			m.GeneratePartialFileOpen(t)
			break
		}
	}

	for _, keys := range versionedKeys {
		for _, k := range keys {
			t := targets[k]
			if !fns.IsSupportedLanguage(fns.LanguageOptions(t), paramLanguage) {
				continue
			}

			msg := fns.Entity(t)
			if msg == nil {
				continue
			}
			m.GeneratePartialImportOpen(t)
			break
		}
	}

	for _, keys := range versionedKeys {
		for _, k := range keys {
			t := targets[k]
			if !fns.IsSupportedLanguage(fns.LanguageOptions(t), paramLanguage) {
				continue
			}

			msg := fns.Entity(t)
			if msg == nil {
				continue
			}
			m.GeneratePartialImport(t)
		}
	}

	for _, keys := range versionedKeys {
		for _, k := range keys {
			t := targets[k]
			if !fns.IsSupportedLanguage(fns.LanguageOptions(t), paramLanguage) {
				continue
			}

			msg := fns.Entity(t)
			if msg == nil {
				continue
			}
			m.GeneratePartialImportClose(t)
			break
		}
	}

	for _, keys := range versionedKeys {
		for _, k := range keys {
			t := targets[k]
			if !fns.IsSupportedLanguage(fns.LanguageOptions(t), paramLanguage) {
				continue
			}

			msg := fns.Entity(t)
			if msg == nil {
				continue
			}
			m.GeneratePartialFileBody(t)
		}
	}

	for _, keys := range versionedKeys {
		for _, k := range keys {
			t := targets[k]
			if !fns.IsSupportedLanguage(fns.LanguageOptions(t), paramLanguage) {
				continue
			}

			msg := fns.Entity(t)
			if msg == nil {
				continue
			}
			m.GeneratePartialFileBodyClose(t)
			break
		}
	}

	for _, keys := range versionedKeys {
		for _, k := range keys {
			t := targets[k]
			m.GenerateProjectJsonFile(t)
			m.GeneratePackageJsonFile(t)
			m.GenerateGoModFile(t)
			break
		}
	}

	return m.Artifacts()
}

func (m GoSpecEntitiesModule) GeneratePartialFileOpen(file pgs.File) {
	templateName := "file_open.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	template.Must(tpl.ParseFS(templates, "templates/"+templateName))
	m.Tpl = tpl

	name := outPath.SetExt("/" + fns.GetPackageVersion(file) + "/entities.pb.entities." + l.FileExtension())
	m.OverwriteGeneratorTemplateFile(name.String(), m.Tpl, file)
}

func (m GoSpecEntitiesModule) GeneratePartialImportOpen(file pgs.File) {
	templateName := "imports_open.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"getGithubRepositoryConstant": fns.GetGithubRepositoryConstant,
		"service":                     fns.Service,
		"parentService":               fns.ParentService,
		"queries":                     fns.QueryMethods,
		"mutations":                   fns.MutationMethods,
		"getCqrsType":                 fns.GetCQRSType,
		"getGoPath":                   fns.GoPath,
		"dotNotationToFilePath":       fns.DotNotationToFilePath,
		"goPackage":                   fns.GoPackage,
		"getPackageVersion":           fns.GetPackageVersion,
		"getApiOptionsTypeName":       fns.GetApiOptionsTypeName,
		"domainSystemName2":           fns.DomainSystemName2,
		"getImportPackageEntity":      fns.GetImportPackageEntity,
	})
	template.Must(tpl.ParseFS(templates, "templates/"+templateName))
	m.Tpl = tpl

	name := outPath.SetExt("/" + fns.GetPackageVersion(file) + "/entities.pb.entities." + l.FileExtension())
	m.AddGeneratorTemplateAppend(name.String(), m.Tpl, file)
}

func (m GoSpecEntitiesModule) GeneratePartialImport(file pgs.File) {
	templateName := "imports.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"getGithubRepositoryConstant": fns.GetGithubRepositoryConstant,
		"service":                     fns.Service,
		"parentService":               fns.ParentService,
		"queries":                     fns.QueryMethods,
		"mutations":                   fns.MutationMethods,
		"getCqrsType":                 fns.GetCQRSType,
		"getGoPath":                   fns.GoPath,
		"dotNotationToFilePath":       fns.DotNotationToFilePath,
		"goPackage":                   fns.GoPackage,
		"getPackageVersion":           fns.GetPackageVersion,
		"getApiOptionsTypeName":       fns.GetApiOptionsTypeName,
		"domainSystemName2":           fns.DomainSystemName2,
		"getImportPackageEntity":      fns.GetImportPackageEntity,
	})
	template.Must(tpl.ParseFS(templates, "templates/"+templateName))
	m.Tpl = tpl

	name := outPath.SetExt("/" + fns.GetPackageVersion(file) + "/entities.pb.entities." + l.FileExtension())
	m.AddGeneratorTemplateAppend(name.String(), m.Tpl, file)
}

func (m GoSpecEntitiesModule) GeneratePartialImportClose(file pgs.File) {
	templateName := "imports_close.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"getGithubRepositoryConstant": fns.GetGithubRepositoryConstant,
		"service":                     fns.Service,
		"parentService":               fns.ParentService,
		"queries":                     fns.QueryMethods,
		"mutations":                   fns.MutationMethods,
		"getCqrsType":                 fns.GetCQRSType,
		"getPackageVersion":           fns.GetPackageVersion,
		"getApiOptionsTypeName":       fns.GetApiOptionsTypeName,
		"domainSystemName2":           fns.DomainSystemName2,
		"getImportPackageEntity":      fns.GetImportPackageEntity,
	})
	template.Must(tpl.ParseFS(templates, "templates/"+templateName))
	m.Tpl = tpl

	name := outPath.SetExt("/" + fns.GetPackageVersion(file) + "/entities.pb.entities." + l.FileExtension())
	m.AddGeneratorTemplateAppend(name.String(), m.Tpl, file)
}

func (m GoSpecEntitiesModule) GeneratePartialFileBody(file pgs.File) {
	templateName := "body.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"getGithubRepositoryConstant": fns.GetGithubRepositoryConstant,
		"service":                     fns.Service,
		"parentService":               fns.ParentService,
		"queries":                     fns.QueryMethods,
		"mutations":                   fns.MutationMethods,
		"getCqrsType":                 fns.GetCQRSType,
		"getPackageVersion":           fns.GetPackageVersion,
		"getApiOptionsTypeName":       fns.GetApiOptionsTypeName,
		"domainSystemName2":           fns.DomainSystemName2,
		"getImportPackageEntity":      fns.GetImportPackageEntity,
	})
	template.Must(tpl.ParseFS(templates, "templates/"+templateName))
	m.Tpl = tpl

	name := outPath.SetExt("/" + fns.GetPackageVersion(file) + "/entities.pb.entities." + l.FileExtension())
	m.AddGeneratorTemplateAppend(name.String(), m.Tpl, file)
}

func (m GoSpecEntitiesModule) GeneratePartialFileBodyClose(file pgs.File) {
	templateName := "body_close.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"getGithubRepositoryConstant": fns.GetGithubRepositoryConstant,
		"service":                     fns.Service,
		"parentService":               fns.ParentService,
		"queries":                     fns.QueryMethods,
		"mutations":                   fns.MutationMethods,
		"getCqrsType":                 fns.GetCQRSType,
		"getPackageVersion":           fns.GetPackageVersion,
		"getApiOptionsTypeName":       fns.GetApiOptionsTypeName,
		"domainSystemName2":           fns.DomainSystemName2,
		"getImportPackageEntity":      fns.GetImportPackageEntity,
	})
	template.Must(tpl.ParseFS(templates, "templates/"+templateName))
	m.Tpl = tpl

	name := outPath.SetExt("/" + fns.GetPackageVersion(file) + "/entities.pb.entities." + l.FileExtension())
	m.AddGeneratorTemplateAppend(name.String(), m.Tpl, file)
}

func (m GoSpecEntitiesModule) GenerateGoModFile(file pgs.File) {
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

	name := outPath.SetExt("/" + fns.GetPackageVersion(file) + "/go.mod")
	m.OverwriteGeneratorTemplateFile(name.String(), m.Tpl, file)
}

func (m GoSpecEntitiesModule) GenerateProjectJsonFile(file pgs.File) {
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

	name := outPath.SetExt("/" + fns.GetPackageVersion(file) + "/project.json")
	m.OverwriteGeneratorTemplateFile(name.String(), m.Tpl, file)
}

func (m GoSpecEntitiesModule) GeneratePackageJsonFile(file pgs.File) {
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

	name := outPath.SetExt("/" + fns.GetPackageVersion(file) + "/package.json")
	m.OverwriteGeneratorTemplateFile(name.String(), m.Tpl, file)
}
