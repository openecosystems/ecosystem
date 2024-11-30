package data_catalog

import (
	"embed"
	"sort"
	"strings"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	_protobuf "libs/plugins/protoc-gen-platform/languages/protobuf"
	"libs/plugins/protoc-gen-platform/shared"
)

var (
	//go:embed templates/*.tmpl
	templates       embed.FS
	protobufOutPath = pgs.JoinPaths("public", "platform", "ontology", "v2alpha")
	outPath         = &protobufOutPath
)

const (
	language   = "protobuf"
	pluginName = "data-catalog"
)

type ProtobufDataCatalogModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	Tpl *template.Template
}

func ProtobufDataCatalogPlugin() *ProtobufDataCatalogModule {
	return &ProtobufDataCatalogModule{ModuleBase: &pgs.ModuleBase{}}
}

func (m *ProtobufDataCatalogModule) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())
}

func (m *ProtobufDataCatalogModule) Name() string { return pluginName }

func (m *ProtobufDataCatalogModule) Execute(targets map[string]pgs.File, _ map[string]pgs.Package) []pgs.Artifact {
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
	for k, _ := range targets {
		version := fns.GetPackageVersion(targets[k])
		versionedKeys[version] = append(versionedKeys[version], k)
	}

	_systemNames := make([]string, 0)
	for k, _ := range targets {
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

	for _, keys := range versionedKeys {
		for _, k := range keys {
			t := targets[k]
			msg := fns.Entity(t)
			if msg == nil {
				continue
			}
		}
	}

	complete = false
	for _, keys := range versionedKeys {
		for _, k := range keys {
			t := targets[k]
			msg := fns.Entity(t)
			if msg == nil {
				continue
			}
			m.GeneratePartialFileOptions(t)
			complete = true
			break
		}
		if complete {
			break
		}
	}

	complete = false
	for _, keys := range versionedKeys {
		for _, k := range keys {
			t := targets[k]
			msg := fns.Entity(t)
			if msg == nil {
				continue
			}
			m.GeneratePartialMessageOpen(t)
			complete = true
			break
		}
		if complete {
			break
		}
	}

	total := 0
	for _, s := range systemNames {
		total++
		m.GeneratePartialSystemField(s, total)
	}

	complete = false
	for _, keys := range versionedKeys {
		for _, k := range keys {
			t := targets[k]
			msg := fns.Entity(t)
			if msg == nil {
				continue
			}
			m.GeneratePartialMessageClose(t)
			complete = true
			break
		}
		if complete {
			break
		}
	}

	total = 0
	for _, s := range systemNames {
		m.GeneratePartialSystemOpen(s)
		for _, keys := range versionedKeys {
			for _, k := range keys {
				t := targets[k]
				if fns.DomainSystemName2(t).LowerCamelCase().String() != s {
					continue
				}

				msg := fns.Entity(t)
				if msg == nil {
					continue
				}
				total++
				m.GeneratePartialMessageField(t, total)
			}
		}
		m.GeneratePartialSystemClose(s)
	}

	for _, keys := range versionedKeys {
		for _, k := range keys {
			t := targets[k]
			msg := fns.Entity(t)
			if msg == nil {
				continue
			}
			m.GeneratePartialMessage(t)
		}
	}

	return m.Artifacts()
}

func (m ProtobufDataCatalogModule) GeneratePartialFileOpen(file pgs.File) {
	templateName := "file_open.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _protobuf.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"getPackageVersion": fns.GetPackageVersion,
		"service":           fns.Service,
		"parentService":     fns.ParentService,
		"queries":           fns.QueryMethods,
		"mutations":         fns.MutationMethods,
		"getCqrsType":       fns.GetCQRSType,
	})
	template.Must(tpl.ParseFS(templates, "templates/"+templateName))
	m.Tpl = tpl

	name := outPath.SetExt("/spec_data_catalog." + l.FileExtension())
	m.OverwriteGeneratorTemplateFile(name.String(), m.Tpl, file)
}

func (m ProtobufDataCatalogModule) GeneratePartialFileOptions(file pgs.File) {
	templateName := "file_options.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _protobuf.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"getGithubRepositoryConstant": fns.GetGithubRepositoryConstant,
		"getPackageVersion":           fns.GetPackageVersion,
		"getApiOptionsType":           fns.GetApiOptionsType,
		"getApiOptionsTypeName":       fns.GetApiOptionsTypeName,
		"service":                     fns.Service,
		"parentService":               fns.ParentService,
		"queries":                     fns.QueryMethods,
		"mutations":                   fns.MutationMethods,
		"getCqrsType":                 fns.GetCQRSType,
	})
	template.Must(tpl.ParseFS(templates, "templates/"+templateName))
	m.Tpl = tpl

	name := outPath.SetExt("/spec_data_catalog." + l.FileExtension())
	m.AddGeneratorTemplateAppend(name.String(), m.Tpl, file)
}

func (m ProtobufDataCatalogModule) GeneratePartialSystemField(system string, index int) {
	templateName := "system_field.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _protobuf.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"fieldPosition":         func() int { return index },
		"getApiOptionsTypeName": fns.GetApiOptionsTypeName,
		"getPackageVersion":     fns.GetPackageVersion,
	})
	template.Must(tpl.ParseFS(templates, "templates/"+templateName))
	m.Tpl = tpl

	name := outPath.SetExt("/spec_data_catalog." + l.FileExtension())
	m.AddGeneratorTemplateAppend(name.String(), m.Tpl, pgs.Name(system))
}

func (m ProtobufDataCatalogModule) GeneratePartialSystemOpen(system string) {
	templateName := "system_open.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _protobuf.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"getApiOptionsTypeName": fns.GetApiOptionsTypeName,
	})
	template.Must(tpl.ParseFS(templates, "templates/"+templateName))
	m.Tpl = tpl

	name := outPath.SetExt("/spec_data_catalog." + l.FileExtension())
	m.AddGeneratorTemplateAppend(name.String(), m.Tpl, pgs.Name(system))
}

func (m ProtobufDataCatalogModule) GeneratePartialMessageOpen(file pgs.File) {
	templateName := "message_open.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _protobuf.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"getApiOptionsTypeName": fns.GetApiOptionsTypeName,
	})
	template.Must(tpl.ParseFS(templates, "templates/"+templateName))
	m.Tpl = tpl

	name := outPath.SetExt("/spec_data_catalog." + l.FileExtension())
	m.AddGeneratorTemplateAppend(name.String(), m.Tpl, file)
}

func (m ProtobufDataCatalogModule) GeneratePartialMessageField(file pgs.File, index int) {
	templateName := "message_field.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _protobuf.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"fieldPosition":         func() int { return index },
		"service":               fns.Service,
		"parentService":         fns.ParentService,
		"queries":               fns.QueryMethods,
		"mutations":             fns.MutationMethods,
		"getCqrsType":           fns.GetCQRSType,
		"entity":                fns.Entity,
		"entityName":            fns.EntityName,
		"entityKeyName":         fns.EntityKeyName,
		"getPackageVersion":     fns.GetPackageVersion,
		"getPackageVersionName": fns.GetPackageVersionName,
		"getApiOptionsTypeName": fns.GetApiOptionsTypeName,
	})
	template.Must(tpl.ParseFS(templates, "templates/"+templateName))
	m.Tpl = tpl

	name := outPath.SetExt("/spec_data_catalog." + l.FileExtension())
	m.AddGeneratorTemplateAppend(name.String(), m.Tpl, file)
}

func (m ProtobufDataCatalogModule) GeneratePartialMessageClose(file pgs.File) {
	templateName := "message_close.go.tmpl"
	l := _protobuf.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	template.Must(tpl.ParseFS(templates, "templates/"+templateName))
	m.Tpl = tpl

	name := outPath.SetExt("/spec_data_catalog." + l.FileExtension())
	m.AddGeneratorTemplateAppend(name.String(), m.Tpl, file)
}

func (m ProtobufDataCatalogModule) GeneratePartialSystemClose(system string) {
	templateName := "system_close.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _protobuf.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"getApiOptionsTypeName": fns.GetApiOptionsTypeName,
	})
	template.Must(tpl.ParseFS(templates, "templates/"+templateName))
	m.Tpl = tpl

	name := outPath.SetExt("/spec_data_catalog." + l.FileExtension())
	m.AddGeneratorTemplateAppend(name.String(), m.Tpl, pgs.Name(system))
}

func (m ProtobufDataCatalogModule) GeneratePartialMessage(file pgs.File) {
	templateName := "message.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _protobuf.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"getApiOptionsTypeName": fns.GetApiOptionsTypeName,
		"getPackageVersion":     fns.GetPackageVersion,
		"entity":                fns.Entity,
		"service":               fns.Service,
		"entityName":            fns.EntityName,
		"entityKeyName":         fns.EntityKeyName,
		"getPackageVersionName": fns.GetPackageVersionName,
	})

	template.Must(tpl.ParseFS(templates, "templates/"+templateName))
	template.Must(tpl.ParseFS(templates, "templates/entity_field.go.tmpl"))
	m.Tpl = tpl

	name := outPath.SetExt("/spec_data_catalog." + l.FileExtension())
	m.AddGeneratorTemplateAppend(name.String(), m.Tpl, file)
}
