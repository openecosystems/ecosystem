package clicommandsv2beta

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
	goOutPath = pgs.JoinPaths("platform", "cmd")
	outPath   = &goOutPath
)

const (
	language   = "go"
	pluginName = "cli-commands-v2beta"
)

type GoCliCommandsModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	Tpl *template.Template
}

func GoCliCommandsPlugin() *GoCliCommandsModule {
	return &GoCliCommandsModule{ModuleBase: &pgs.ModuleBase{}}
}

func (m *GoCliCommandsModule) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())
}

func (m *GoCliCommandsModule) Name() string { return language + "/" + pluginName }

func (m *GoCliCommandsModule) Execute(targets map[string]pgs.File, _ map[string]pgs.Package) []pgs.Artifact {
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
			if !strings.HasPrefix(t.FullyQualifiedName(), ".platform") {
				continue
			}
			m.GeneratePartialFileOpen(t)
			break
		}
	}

	sv := make(map[string]bool)
	for _, keys := range versionedKeys {
		for _, k := range keys {
			t := targets[k]
			_sv := fns.DomainSystemName2(t).LowerCamelCase().String() + fns.GetPackageVersion(t)
			if _, ok := sv[_sv]; ok {
				continue
			}
			sv[_sv] = true

			if !strings.HasPrefix(t.FullyQualifiedName(), ".platform") {
				continue
			}
			m.GeneratePartialImport(t)
		}
	}

	for _, keys := range versionedKeys {
		for _, k := range keys {
			t := targets[k]
			m.GeneratePartialFunctionOpen(t)
			break
		}
		break
	}

	sv2 := make(map[string]bool)
	for _, keys := range versionedKeys {
		for _, k := range keys {
			t := targets[k]
			_sv := fns.DomainSystemName2(t).LowerCamelCase().String() + fns.GetPackageVersion(t)
			if _, ok := sv2[_sv]; ok {
				continue
			}
			sv2[_sv] = true

			msg := fns.Entity(t)
			if msg == nil {
				continue
			}
			m.GeneratePartialFunctionBody(t)
		}
	}

	for _, keys := range versionedKeys {
		for _, k := range keys {
			t := targets[k]
			m.GeneratePartialFunctionClose(t)
			break
		}
		break
	}

	return m.Artifacts()
}

func (m GoCliCommandsModule) GeneratePartialFileOpen(file pgs.File) {
	templateName := "file_open.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"service":       fns.Service,
		"parentService": fns.ParentService,
		"queries":       fns.QueryMethods,
		"mutations":     fns.MutationMethods,
		"getCqrsType":   fns.GetCQRSType,
	})
	template.Must(tpl.ParseFS(templates, "templates/"+templateName))
	m.Tpl = tpl

	name := outPath.SetExt("/cli.pb." + l.FileExtension())
	m.OverwriteGeneratorTemplateFile(name.String(), m.Tpl, file)
}

func (m GoCliCommandsModule) GeneratePartialImport(file pgs.File) {
	templateName := "import.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"service":               fns.Service,
		"parentService":         fns.ParentService,
		"queries":               fns.QueryMethods,
		"mutations":             fns.MutationMethods,
		"getCqrsType":           fns.GetCQRSType,
		"getGoPath":             fns.GoPath,
		"dotNotationToFilePath": fns.DotNotationToFilePath,
		"goPackage":             fns.GoPackage,
		"getPackageVersionName": fns.GetPackageVersionName,
		"getPackageVersion":     fns.GetPackageVersion,
		"domainSystemName2":     fns.DomainSystemName2,
		"getApiOptionsTypeName": fns.GetApiOptionsTypeName,
	})
	template.Must(tpl.ParseFS(templates, "templates/"+templateName))
	m.Tpl = tpl

	name := outPath.SetExt("/cli.pb." + l.FileExtension())
	m.AddGeneratorTemplateAppend(name.String(), m.Tpl, file)
}

func (m GoCliCommandsModule) GeneratePartialFunctionOpen(file pgs.File) {
	templateName := "function_open.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"service":       fns.Service,
		"parentService": fns.ParentService,
		"queries":       fns.QueryMethods,
		"mutations":     fns.MutationMethods,
		"getCqrsType":   fns.GetCQRSType,
	})
	template.Must(tpl.ParseFS(templates, "templates/"+templateName))
	m.Tpl = tpl

	name := outPath.SetExt("/cli.pb." + l.FileExtension())
	m.AddGeneratorTemplateAppend(name.String(), m.Tpl, file)
}

func (m GoCliCommandsModule) GeneratePartialFunctionBody(file pgs.File) {
	templateName := "function_body.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"service":               fns.Service,
		"parentService":         fns.ParentService,
		"queries":               fns.QueryMethods,
		"mutations":             fns.MutationMethods,
		"getCqrsType":           fns.GetCQRSType,
		"getPackageVersionName": fns.GetPackageVersionName,
		"getPackageVersion":     fns.GetPackageVersion,
		"domainSystemName2":     fns.DomainSystemName2,
	})
	template.Must(tpl.ParseFS(templates, "templates/"+templateName))
	m.Tpl = tpl

	name := outPath.SetExt("/cli.pb." + l.FileExtension())
	m.AddGeneratorTemplateAppend(name.String(), m.Tpl, file)
}

func (m GoCliCommandsModule) GeneratePartialFunctionClose(file pgs.File) {
	templateName := "function_close.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"service":       fns.Service,
		"parentService": fns.ParentService,
		"queries":       fns.QueryMethods,
		"mutations":     fns.MutationMethods,
		"getCqrsType":   fns.GetCQRSType,
	})
	template.Must(tpl.ParseFS(templates, "templates/"+templateName))
	m.Tpl = tpl

	name := outPath.SetExt("/cli.pb." + l.FileExtension())
	m.AddGeneratorTemplateAppend(name.String(), m.Tpl, file)
}
