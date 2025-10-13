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

	complete := false
	for _, keys := range versionedKeys {
		for _, k := range keys {
			t := targets[k]
			msg := fns.Entity(t)
			if msg == nil {
				continue
			}

			m.GeneratePartialFileOpen(t)
			complete = true
			break
		}
		if complete {
			break
		}
	}

	sv := make(map[string]bool)
	for _, keys := range versionedKeys {
		for _, k := range keys {
			t := targets[k]
			_sv := fns.DomainSystemName2(t).LowerCamelCase().String() + fns.GetPackageVersion(t)
			msg := fns.Entity(t)
			if msg == nil {
				continue
			}
			if _, ok := sv[_sv]; ok {
				continue
			}
			sv[_sv] = true

			m.GeneratePartialImports(t)
		}
	}

	m.GeneratePartialStructOpen()
	for _, s := range systemNames {
		m.GeneratePartialStructBody(s)
	}
	m.GeneratePartialStructClose()

	complete = false
	for _, keys = range versionedKeys {
		for _, k := range keys {
			t := targets[k]
			msg := fns.Entity(t)
			if msg == nil {
				continue
			}

			m.GeneratePartialClientOpen(t)
			complete = true
			break
		}
		if complete {
			break
		}
	}

	for _, s := range systemNames {
		m.GeneratePartialSystemOpen(s)
		for _, keys = range versionedKeys {
			for _, k := range keys {
				t := targets[k]
				// m.Log(s, k, fns.DomainSystemName2(t).LowerCamelCase().String())
				if fns.DomainSystemName2(t).LowerCamelCase().String() != s {
					continue
				}

				msg := fns.Entity(t)
				if msg == nil {
					continue
				}

				for _, service := range t.Services() {
					m.GeneratePartialService(service)
				}
			}
		}
		m.GeneratePartialSystemClose(s)
	}

	complete = false
	for _, keys = range versionedKeys {
		for _, k := range keys {
			t := targets[k]
			msg := fns.Entity(t)
			if msg == nil {
				continue
			}

			m.GeneratePartialClientClose(t)
			complete = true
			break
		}
		if complete {
			break
		}
	}

	for _, s := range systemNames {
		m.GeneratePartialSystemStructOpen(s)
		for _, keys = range versionedKeys {
			for _, k := range keys {
				t := targets[k]
				// m.Log(s, k, fns.DomainSystemName2(t).LowerCamelCase().String())
				if fns.DomainSystemName2(t).LowerCamelCase().String() != s {
					continue
				}

				msg := fns.Entity(t)
				if msg == nil {
					continue
				}

				for _, service := range t.Services() {
					m.GeneratePartialSystemStructService(service)
				}
			}
		}
		m.GeneratePartialSystemStructClose(s)
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

func (m *GoSpecClientModule) GeneratePartialStructOpen() {
	templateName := "struct_open.go.tmpl"
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
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	name := getSpecClientFileName(file, fns)
	m.AddGeneratorTemplateAppend(name, m.Tpl, "")
}

func (m *GoSpecClientModule) GeneratePartialStructBody(system string) {
	templateName := "struct_body.go.tmpl"
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
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	name := getSpecClientFileName(file, fns)
	m.AddGeneratorTemplateAppend(name, m.Tpl, pgs.Name(system))
}

func (m *GoSpecClientModule) GeneratePartialStructClose() {
	templateName := "struct_close.go.tmpl"
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
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	name := getSpecClientFileName(file, fns)
	m.AddGeneratorTemplateAppend(name, m.Tpl, "")
}

func (m *GoSpecClientModule) GeneratePartialClientOpen(file pgs.File) {
	templateName := "client_open.go.tmpl"
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

func (m *GoSpecClientModule) GeneratePartialSystemOpen(system string) {
	templateName := "system_open.go.tmpl"
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
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	name := getSpecClientFileName(file, fns)
	m.AddGeneratorTemplateAppend(name, m.Tpl, pgs.Name(system))
}

func (m *GoSpecClientModule) GeneratePartialService(service pgs.Service) {
	templateName := "service.go.tmpl"
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
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	name := getSpecClientFileName(file, fns)

	m.AddGeneratorTemplateAppend(name, m.Tpl, service)
}

func (m *GoSpecClientModule) GeneratePartialSystemClose(system string) {
	templateName := "system_close.go.tmpl"
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
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	name := getSpecClientFileName(file, fns)
	m.AddGeneratorTemplateAppend(name, m.Tpl, pgs.Name(system))
}

func (m *GoSpecClientModule) GeneratePartialClientClose(file pgs.File) {
	templateName := "client_close.go.tmpl"
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

func (m *GoSpecClientModule) GeneratePartialSystemStructOpen(system string) {
	templateName := "system_struct_open.go.tmpl"
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
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	name := getSpecClientFileName(file, fns)
	m.AddGeneratorTemplateAppend(name, m.Tpl, pgs.Name(system))
}

func (m *GoSpecClientModule) GeneratePartialSystemStructService(service pgs.Service) {
	templateName := "system_struct_service.go.tmpl"
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
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	name := getSpecClientFileName(file, fns)

	m.AddGeneratorTemplateAppend(name, m.Tpl, service)
}

func (m *GoSpecClientModule) GeneratePartialSystemStructClose(system string) {
	templateName := "system_struct_close.go.tmpl"
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
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	name := getSpecClientFileName(file, fns)
	m.AddGeneratorTemplateAppend(name, m.Tpl, pgs.Name(system))
}

func (m *GoSpecClientModule) GeneratePartialFileClose() {
	templateName := "close.go.tmpl"
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
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	name := getSpecClientFileName(file, fns)
	m.AddGeneratorTemplateAppend(name, m.Tpl, "")
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
