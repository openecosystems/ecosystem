package spec

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

//go:embed templates/*.tmpl
var templates embed.FS

const (
	language   = "go"
	pluginName = "spec"
)

type GoSpecModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	Tpl *template.Template
}

func GoSpecPlugin() *GoSpecModule { return &GoSpecModule{ModuleBase: &pgs.ModuleBase{}} }

func (m *GoSpecModule) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())
}

func (m *GoSpecModule) Name() string { return language + "/" + pluginName }

func (m *GoSpecModule) Execute(targets map[string]pgs.File, _ map[string]pgs.Package) []pgs.Artifact {
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
		m.GenerateFile(t)
	}

	return m.Artifacts()
}

func (m GoSpecModule) GenerateFile(file pgs.File) {
	templateName := "file.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"service":         fns.Service,
		"parentService":   fns.ParentService,
		"queries":         fns.QueryMethods,
		"mutations":       fns.MutationMethods,
		"getCqrsType":     fns.GetCQRSType,
		"getSpecCommands": fns.GetSpecCommands,
		"getSpecEvents":   fns.GetSpecEvents,
		"getSpecTopics":   fns.GetSpecTopics,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	msg := fns.Entity(file)
	if msg == nil {
		return
	}

	name := m.ctx.OutputPath(file).SetExt(".spec." + l.FileExtension())
	m.OverwriteGeneratorTemplateFile(name.String(), m.Tpl, file)
}
