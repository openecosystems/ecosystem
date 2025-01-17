package protobufindextypescript

import (
	"embed"
	"sort"
	"strings"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	_typescript "libs/plugins/protoc-gen-platform/languages/typescript"
	"libs/plugins/protoc-gen-platform/shared"
)

//go:embed templates/*.tmpl
var templates embed.FS

const (
	language   = "typescript"
	pluginName = "protobuf-index"
)

type TypeScriptProtobufIndexModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	Tpl *template.Template
}

func TypeScriptProtobufIndexPlugin() *TypeScriptProtobufIndexModule {
	return &TypeScriptProtobufIndexModule{ModuleBase: &pgs.ModuleBase{}}
}

func (m *TypeScriptProtobufIndexModule) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())
}

func (m *TypeScriptProtobufIndexModule) Name() string { return language + "/" + pluginName }

func (m *TypeScriptProtobufIndexModule) Execute(targets map[string]pgs.File, _ map[string]pgs.Package) []pgs.Artifact {
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
	for k, _ := range targets {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		t := targets[k]
		m.GeneratePartialFileOpen(t)
		break
	}

	for _, k := range keys {
		t := targets[k]
		m.GeneratePartialFileBody(t)
	}

	return m.Artifacts()
}

func (m TypeScriptProtobufIndexModule) GeneratePartialFileOpen(file pgs.File) {
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

	name := "index.ts"
	m.OverwriteGeneratorTemplateFile(name, m.Tpl, file)
}

func (m TypeScriptProtobufIndexModule) GeneratePartialFileBody(file pgs.File) {
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

	name := "index.ts"
	m.AddGeneratorTemplateAppend(name, m.Tpl, file)
}
