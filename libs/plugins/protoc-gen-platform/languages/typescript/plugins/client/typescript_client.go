package clienttypescript

import (
	"embed"
	"sort"
	"strings"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	_typescript "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/typescript"
	"github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/shared"
)

//go:embed templates/*.tmpl
var templates embed.FS

const (
	language   = "typescript"
	pluginName = "typescript-client"
)

type TypeScriptClientModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	Tpl *template.Template
}

func TypeScriptClientPlugin() *TypeScriptClientModule {
	return &TypeScriptClientModule{ModuleBase: &pgs.ModuleBase{}}
}

func (m *TypeScriptClientModule) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())
}

func (m *TypeScriptClientModule) Name() string { return language + "/" + pluginName }

func (m *TypeScriptClientModule) Execute(targets map[string]pgs.File, _ map[string]pgs.Package) []pgs.Artifact {
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

	for _, keys = range versionedKeys {
		for _, k := range keys {
			t := targets[k]
			msg := fns.Entity(t)
			if msg == nil {
				continue
			}
			m.GeneratePartialFileBody(t)
		}
	}

	complete = false
	for _, keys = range versionedKeys {
		for _, k := range keys {
			t := targets[k]
			msg := fns.Entity(t)
			if msg == nil {
				continue
			}

			m.GeneratePartialClassOpen(t)
			complete = true
			break
		}
		if complete {
			break
		}
	}

	for _, keys = range versionedKeys {
		for _, k := range keys {
			t := targets[k]
			msg := fns.Entity(t)
			if msg == nil {
				continue
			}
			m.GeneratePartialClassPrivateFields(t)
		}
	}

	for _, s := range systemNames {
		m.GeneratePartialSystemOpen(s)
		for _, keys = range versionedKeys {
			for _, k := range keys {
				t := targets[k]
				if fns.DomainSystemName2(t).LowerCamelCase().String() != s {
					continue
				}

				msg := fns.Entity(t)
				if msg == nil {
					continue
				}
				version := fns.GetPackageVersionName(t)
				combined := pgs.Name(msg.Name().String() + version.UpperCamelCase().String())

				m.GeneratePartialEntityOpen(msg)
				for _, service := range t.Services() {
					for _, method := range service.Methods() {
						if method.ServerStreaming() || method.ClientStreaming() {
							m.GeneratePartialMethodStreaming(method, msg.Name(), combined)
						} else {
							m.GeneratePartialMethod(method, combined)
						}
					}
				}
				m.GeneratePartialEntityClose(t)
			}
		}
		m.GeneratePartialSystemClose(s)
	}

	m.GeneratePartialFileClose()

	return m.Artifacts()
}

func (m *TypeScriptClientModule) GeneratePartialFileOpen(file pgs.File) {
	templateName := "file.ts.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _typescript.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"entity":                         fns.Entity,
		"entityName":                     fns.EntityName,
		"getPackageVersionName":          fns.GetPackageVersionName,
		"protoPathWithoutProtoExtension": fns.ProtoPathWithoutProtoExtension,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	name := "client.ts"
	m.OverwriteGeneratorTemplateFile(name, m.Tpl, file)
}

func (m *TypeScriptClientModule) GeneratePartialFileBody(file pgs.File) {
	templateName := "body.ts.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _typescript.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"entity":                         fns.Entity,
		"entityName":                     fns.EntityName,
		"getPackageVersionName":          fns.GetPackageVersionName,
		"protoPathWithoutProtoExtension": fns.ProtoPathWithoutProtoExtension,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	msg := fns.Entity(file)
	if msg == nil {
		return
	}

	name := "client.ts"
	m.AddGeneratorTemplateAppend(name, m.Tpl, file)
}

func (m *TypeScriptClientModule) GeneratePartialClassOpen(file pgs.File) {
	templateName := "class_open.ts.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _typescript.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"entity":                         fns.Entity,
		"entityName":                     fns.EntityName,
		"getPackageVersionName":          fns.GetPackageVersionName,
		"protoPathWithoutProtoExtension": fns.ProtoPathWithoutProtoExtension,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	msg := fns.Entity(file)
	if msg == nil {
		return
	}

	name := "client.ts"
	m.AddGeneratorTemplateAppend(name, m.Tpl, file)
}

func (m *TypeScriptClientModule) GeneratePartialClassPrivateFields(file pgs.File) {
	templateName := "class_private_fields.ts.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _typescript.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"entity":                         fns.Entity,
		"entityName":                     fns.EntityName,
		"getPackageVersionName":          fns.GetPackageVersionName,
		"protoPathWithoutProtoExtension": fns.ProtoPathWithoutProtoExtension,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	msg := fns.Entity(file)
	if msg == nil {
		return
	}

	name := "client.ts"
	m.AddGeneratorTemplateAppend(name, m.Tpl, file)
}

func (m *TypeScriptClientModule) GeneratePartialSystemOpen(system string) {
	templateName := "system_open.ts.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _typescript.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"entity":                         fns.Entity,
		"entityName":                     fns.EntityName,
		"getPackageVersionName":          fns.GetPackageVersionName,
		"protoPathWithoutProtoExtension": fns.ProtoPathWithoutProtoExtension,
		"domainSystemName2":              fns.DomainSystemName2,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	name := "client.ts"
	m.AddGeneratorTemplateAppend(name, m.Tpl, pgs.Name(system))
}

func (m *TypeScriptClientModule) GeneratePartialEntityOpen(entity pgs.Message) {
	templateName := "entity_open.ts.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _typescript.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"entity":                         fns.Entity,
		"entityName":                     fns.EntityName,
		"getPackageVersionName":          fns.GetPackageVersionName,
		"protoPathWithoutProtoExtension": fns.ProtoPathWithoutProtoExtension,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	name := "client.ts"
	m.AddGeneratorTemplateAppend(name, m.Tpl, entity)
}

func (m *TypeScriptClientModule) GeneratePartialMethod(method pgs.Method, entity pgs.Name) {
	templateName := "method.ts.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _typescript.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"entity":                         fns.Entity,
		"entityName":                     fns.EntityName,
		"getPackageVersionName":          fns.GetPackageVersionName,
		"protoPathWithoutProtoExtension": fns.ProtoPathWithoutProtoExtension,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	name := "client.ts"

	type Data struct {
		Method pgs.Method
		Entity pgs.Name
	}

	data := Data{
		Method: method,
		Entity: entity,
	}

	m.AddGeneratorTemplateAppend(name, m.Tpl, data)
}

func (m *TypeScriptClientModule) GeneratePartialMethodStreaming(method pgs.Method, entity pgs.Name, combined pgs.Name) {
	templateName := "method_streaming.ts.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _typescript.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"entity":                         fns.Entity,
		"entityName":                     fns.EntityName,
		"getPackageVersionName":          fns.GetPackageVersionName,
		"protoPathWithoutProtoExtension": fns.ProtoPathWithoutProtoExtension,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	name := "client.ts"

	type Data struct {
		Method   pgs.Method
		Entity   pgs.Name
		Combined pgs.Name
	}

	data := Data{
		Method:   method,
		Entity:   entity,
		Combined: combined,
	}

	m.AddGeneratorTemplateAppend(name, m.Tpl, data)
}

func (m *TypeScriptClientModule) GeneratePartialEntityClose(file pgs.File) {
	templateName := "entity_close.ts.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _typescript.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"entity":                         fns.Entity,
		"entityName":                     fns.EntityName,
		"getPackageVersionName":          fns.GetPackageVersionName,
		"protoPathWithoutProtoExtension": fns.ProtoPathWithoutProtoExtension,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	name := "client.ts"
	m.AddGeneratorTemplateAppend(name, m.Tpl, file)
}

func (m *TypeScriptClientModule) GeneratePartialSystemClose(system string) {
	templateName := "system_close.ts.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _typescript.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"entity":                         fns.Entity,
		"entityName":                     fns.EntityName,
		"getPackageVersionName":          fns.GetPackageVersionName,
		"protoPathWithoutProtoExtension": fns.ProtoPathWithoutProtoExtension,
		"domainSystemName2":              fns.DomainSystemName2,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	name := "client.ts"
	m.AddGeneratorTemplateAppend(name, m.Tpl, pgs.Name(system))
}

func (m *TypeScriptClientModule) GeneratePartialFileClose() {
	templateName := "close.ts.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _typescript.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"entity":                         fns.Entity,
		"entityName":                     fns.EntityName,
		"getPackageVersionName":          fns.GetPackageVersionName,
		"protoPathWithoutProtoExtension": fns.ProtoPathWithoutProtoExtension,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	name := "client.ts"
	m.AddGeneratorTemplateAppend(name, m.Tpl, "")
}
