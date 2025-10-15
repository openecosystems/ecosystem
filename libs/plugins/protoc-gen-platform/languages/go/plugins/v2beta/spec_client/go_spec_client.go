package specclientv2beta

import (
	"embed"
	"sort"
	"strings"
	"text/template"

	shared "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/shared"

	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	_go "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go"
)

var (
	//go:embed templates/*.tmpl
	templates embed.FS
	goOutPath = pgs.JoinPaths("")
	outPath   = &goOutPath
)

const (
	language   = "go"
	pluginName = "go-spec-client"
)

type GoSpecClientModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	Tpl *template.Template
}

func GoSpecClientPlugin() *GoSpecClientModule {
	return &GoSpecClientModule{ModuleBase: &pgs.ModuleBase{}}
}

func (m *GoSpecClientModule) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())
}

func (m *GoSpecClientModule) Name() string { return language + "/" + pluginName }

func (m *GoSpecClientModule) Execute(targets map[string]pgs.File, _ map[string]pgs.Package) []pgs.Artifact {
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

		p := targets[k].Descriptor().GetPackage()
		s := strings.Split(p, ".")

		if len(s) != 3 {
			continue
		}

		version := fns.GetPackageVersion(targets[k])
		versionedKeys[version] = append(versionedKeys[version], k)
	}

	_systemNames := make([]string, 0)
	for k := range targets {
		systemName := fns.DomainSystemName2(targets[k])
		_systemNames = append(_systemNames, systemName.LowerCamelCase().String())
	}

	systemNames := make([]string, 0, len(_systemNames))
	m1 := make(map[string]bool)

	for _, val := range _systemNames {
		if _, ok := m1[val]; !ok {
			m1[val] = true
			systemNames = append(systemNames, val)
		}
	}
	sort.Strings(systemNames)

	for _, v := range versionedKeys {
		sort.Strings(v)
	}

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

		m.GeneratePartialFileOpen(t)
		m.GeneratePartialImports(t)
		m.GeneratePartialInterfaceOpen(t.Services()[0], t)

		for _, method := range t.Services()[0].Methods() {
			m.GeneratePartialMethod(method, t)
		}

		m.GeneratePartialInterfaceClose(t.Services()[0], t)

		for _, method := range t.Services()[0].Methods() {
			m.GeneratePartialMethodImplementation(method, t)
		}
	}

	return m.Artifacts()
}

func (m *GoSpecClientModule) GeneratePartialFileOpen(file pgs.File) {
	templateName := "file.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"entity":                         fns.Entity,
		"entityName":                     fns.EntityName,
		"getPackageVersionName":          fns.GetPackageVersionName,
		"protoPathWithoutProtoExtension": fns.ProtoPathWithoutProtoExtension,
		"getGoPackageAlias":              fns.GetGoPackageAlias,
		"getImportPackageEntity":         fns.GetImportPackageEntity,
		"getCqrsType":                    fns.GetCQRSType,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	name := getSpecClientFileName(file, fns)

	m.OverwriteGeneratorTemplateFile(name, m.Tpl, file)
}

func (m *GoSpecClientModule) GeneratePartialImports(file pgs.File) {
	templateName := "imports.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"entity":                         fns.Entity,
		"entityName":                     fns.EntityName,
		"getPackageVersionName":          fns.GetPackageVersionName,
		"protoPathWithoutProtoExtension": fns.ProtoPathWithoutProtoExtension,
		"getGoPackageAlias":              fns.GetGoPackageAlias,
		"getImportPackageEntity":         fns.GetImportPackageEntity,
		"getCqrsType":                    fns.GetCQRSType,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	msg := fns.Entity(file)
	if msg == nil {
		return
	}

	name := getSpecClientFileName(file, fns)
	m.AddGeneratorTemplateAppend(name, m.Tpl, file)
}

func (m *GoSpecClientModule) GeneratePartialInterfaceOpen(service pgs.Service, file pgs.File) {
	templateName := "interface_open.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"entity":                         fns.Entity,
		"entityName":                     fns.EntityName,
		"getPackageVersionName":          fns.GetPackageVersionName,
		"protoPathWithoutProtoExtension": fns.ProtoPathWithoutProtoExtension,
		"domainSystemName2":              fns.DomainSystemName2,
		"getGoPackageAlias":              fns.GetGoPackageAlias,
		"getImportPackageEntity":         fns.GetImportPackageEntity,
		"getCqrsType":                    fns.GetCQRSType,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	name := getSpecClientFileName(file, fns)

	m.AddGeneratorTemplateAppend(name, m.Tpl, service.Name())
}

func (m *GoSpecClientModule) GeneratePartialMethod(method pgs.Method, file pgs.File) {
	templateName := "method.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"entity":                         fns.Entity,
		"entityName":                     fns.EntityName,
		"getPackageVersionName":          fns.GetPackageVersionName,
		"protoPathWithoutProtoExtension": fns.ProtoPathWithoutProtoExtension,
		"getGoPackageAlias":              fns.GetGoPackageAlias,
		"getImportPackageEntity":         fns.GetImportPackageEntity,
		"getCqrsType":                    fns.GetCQRSType,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	name := getSpecClientFileName(file, fns)

	m.AddGeneratorTemplateAppend(name, m.Tpl, method)
}

func (m *GoSpecClientModule) GeneratePartialInterfaceClose(service pgs.Service, file pgs.File) {
	templateName := "interface_close.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"entity":                         fns.Entity,
		"entityName":                     fns.EntityName,
		"getPackageVersionName":          fns.GetPackageVersionName,
		"protoPathWithoutProtoExtension": fns.ProtoPathWithoutProtoExtension,
		"domainSystemName2":              fns.DomainSystemName2,
		"getGoPackageAlias":              fns.GetGoPackageAlias,
		"getImportPackageEntity":         fns.GetImportPackageEntity,
		"getCqrsType":                    fns.GetCQRSType,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	name := getSpecClientFileName(file, fns)
	m.AddGeneratorTemplateAppend(name, m.Tpl, service.Name())
}

func (m *GoSpecClientModule) GeneratePartialMethodImplementation(method pgs.Method, file pgs.File) {
	templateName := "method_implementation.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"entity":                         fns.Entity,
		"entityName":                     fns.EntityName,
		"getPackageVersionName":          fns.GetPackageVersionName,
		"protoPathWithoutProtoExtension": fns.ProtoPathWithoutProtoExtension,
		"getGoPackageAlias":              fns.GetGoPackageAlias,
		"getImportPackageEntity":         fns.GetImportPackageEntity,
		"getCqrsType":                    fns.GetCQRSType,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	name := getSpecClientFileName(file, fns)

	m.AddGeneratorTemplateAppend(name, m.Tpl, method)
}

func getSpecClientFileName(file pgs.File, fns shared.Functions) string {
	_scope, _system, _version, _ := fns.GetPackageScopeSystemAndVersion(file)
	scope := _scope.LowerSnakeCase().String()
	system := _system.LowerSnakeCase().String()
	version := _version.String()
	gopackage := fns.GetGoPackageAlias(file)
	fileName := fns.ProtoName(file)

	clientFileName := outPath.SetExt("/" + scope + "/" + system + "/" + version + "/" + gopackage + "connect" + "/" + fileName + ".client.go").String()
	return clientFileName
}
