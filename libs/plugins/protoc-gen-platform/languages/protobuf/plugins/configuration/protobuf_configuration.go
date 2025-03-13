package configuration

import (
	"embed"
	"sort"
	"strings"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	_protobuf "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/protobuf"
	"github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/shared"
)

var (
	//go:embed templates/*.tmpl
	templates       embed.FS
	protobufOutPath = pgs.JoinPaths("public", "platform", "configuration", "v2alpha")
	outPath         = &protobufOutPath
)

const (
	language   = "protobuf"
	pluginName = "configuration"
)

type ProtobufConfigurationModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	Tpl *template.Template
}

func ProtobufConfigurationPlugin() *ProtobufConfigurationModule {
	return &ProtobufConfigurationModule{ModuleBase: &pgs.ModuleBase{}}
}

func (m *ProtobufConfigurationModule) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())
}

func (m *ProtobufConfigurationModule) Name() string { return pluginName }

func (m *ProtobufConfigurationModule) Execute(targets map[string]pgs.File, _ map[string]pgs.Package) []pgs.Artifact {
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

	complete := false
	for _, keys := range versionedKeys {
		for _, k := range keys {
			t := targets[k]
			msg := fns.Configuration(t)
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
			msg := fns.Configuration(t)
			if msg == nil {
				continue
			}
			m.GeneratePartialImport(t)
		}
	}

	complete = false
	for _, keys := range versionedKeys {
		for _, k := range keys {
			t := targets[k]
			msg := fns.Configuration(t)
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
			m.GeneratePartialMessageOpen(t)
			complete = true
			break
		}
		if complete {
			break
		}
	}

	total := 0
	for _, keys := range versionedKeys {
		for _, k := range keys {
			t := targets[k]
			msg := fns.Configuration(t)
			if msg == nil {
				continue
			}
			total++
			m.GeneratePartialMessageField(t, total)
		}
	}

	complete = false
	for _, keys := range versionedKeys {
		for _, k := range keys {
			t := targets[k]
			m.GeneratePartialMessageClose(t)
			complete = true
			break
		}
		if complete {
			break
		}
	}

	return m.Artifacts()
}

func (m ProtobufConfigurationModule) GeneratePartialFileOpen(file pgs.File) {
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

	name := outPath.SetExt("/spec_configuration." + l.FileExtension())
	m.OverwriteGeneratorTemplateFile(name.String(), m.Tpl, file)
}

func (m ProtobufConfigurationModule) GeneratePartialImport(file pgs.File) {
	templateName := "import.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _protobuf.GetLanguage(templateName, m.ctx, m.Parameters())

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

	name := outPath.SetExt("/spec_configuration." + l.FileExtension())
	m.AddGeneratorTemplateAppend(name.String(), m.Tpl, file)
}

func (m ProtobufConfigurationModule) GeneratePartialFileOptions(file pgs.File) {
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

	name := outPath.SetExt("/spec_configuration." + l.FileExtension())
	m.AddGeneratorTemplateAppend(name.String(), m.Tpl, file)
}

func (m ProtobufConfigurationModule) GeneratePartialMessageOpen(file pgs.File) {
	templateName := "message_open.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _protobuf.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"getApiOptionsTypeName": fns.GetApiOptionsTypeName,
	})
	template.Must(tpl.ParseFS(templates, "templates/"+templateName))
	m.Tpl = tpl

	name := outPath.SetExt("/spec_configuration." + l.FileExtension())
	m.AddGeneratorTemplateAppend(name.String(), m.Tpl, file)
}

func (m ProtobufConfigurationModule) GeneratePartialMessageField(file pgs.File, index int) {
	templateName := "message_field.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _protobuf.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"fieldPosition":         func() int { return index + 1 },
		"service":               fns.Service,
		"parentService":         fns.ParentService,
		"queries":               fns.QueryMethods,
		"mutations":             fns.MutationMethods,
		"getCqrsType":           fns.GetCQRSType,
		"configurationName":     fns.ConfigurationName,
		"configurationNumber":   fns.ConfigurationNumber,
		"getApiOptionsTypeName": fns.GetApiOptionsTypeName,
		"getPackageVersion":     fns.GetPackageVersion,
	})
	template.Must(tpl.ParseFS(templates, "templates/"+templateName))
	m.Tpl = tpl

	name := outPath.SetExt("/spec_configuration." + l.FileExtension())
	m.AddGeneratorTemplateAppend(name.String(), m.Tpl, file)
}

func (m ProtobufConfigurationModule) GeneratePartialMessageClose(file pgs.File) {
	templateName := "message_close.go.tmpl"
	l := _protobuf.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	template.Must(tpl.ParseFS(templates, "templates/"+templateName))
	m.Tpl = tpl

	name := outPath.SetExt("/spec_configuration." + l.FileExtension())
	m.AddGeneratorTemplateAppend(name.String(), m.Tpl, file)
}
