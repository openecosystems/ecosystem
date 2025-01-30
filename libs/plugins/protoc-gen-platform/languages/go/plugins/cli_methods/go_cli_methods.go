package cli_methods

import (
	"embed"
	"libs/plugins/protoc-gen-platform/shared"
	"sort"
	"strings"
	"text/template"

	_go "libs/plugins/protoc-gen-platform/languages/go"

	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
)

//go:embed templates/*.tmpl
var templates embed.FS

const (
	language   = "go"
	pluginName = "cli-methods"
)

type GoCliMethodsModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	Tpl *template.Template
}

func GoCliMethodsPlugin() *GoCliMethodsModule {
	return &GoCliMethodsModule{ModuleBase: &pgs.ModuleBase{}}
}

func (m *GoCliMethodsModule) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())
}

func (m *GoCliMethodsModule) Name() string { return language + "/" + pluginName }

func (m *GoCliMethodsModule) Execute(targets map[string]pgs.File, _ map[string]pgs.Package) []pgs.Artifact {
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
		if !fns.IsSupportedLanguage(fns.LanguageOptions(t), paramLanguage) {
			continue
		}
		m.GenerateFile(t)
	}

	return m.Artifacts()
}

func (m GoCliMethodsModule) GenerateFile(file pgs.File) {
	templateName := "file.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"service":                       fns.Service,
		"parentService":                 fns.ParentService,
		"queries":                       fns.QueryMethods,
		"isMethodQuery":                 fns.IsMethodQuery,
		"mutations":                     fns.MutationMethods,
		"isMethodMutation":              fns.IsMethodMutation,
		"getCqrsType":                   fns.GetCQRSType,
		"getSpecCommands":               fns.GetSpecCommands,
		"getSpecEvents":                 fns.GetSpecEvents,
		"getSpecTopics":                 fns.GetSpecTopics,
		"goPackageOverwrite":            fns.GoPackageOverwrite,
		"goPath":                        fns.GoPath,
		"getImportPackages":             fns.GetGoImportPackagesCLI,
		"getImportName":                 fns.GetImportPackageMessageDirectlyFromGoPackage,
		"getMethodVerb":                 fns.GetMethodVerb,
		"dotNotationToFilePath":         fns.DotNotationToFilePath,
		"methodTrailingComment":         fns.MethodTrailingComment,
		"methodLeadingComment":          fns.MethodLeadingComment,
		"methodLeadingDetachedComments": fns.MethodLeadingDetachedComments,
		"getPackageVersion":             fns.GetPackageVersion,
		"getPackageVersionName":         fns.GetPackageVersionName,
		"getApiOptionsTypeName":         fns.GetApiOptionsTypeName,
		"domainSystemName2":             fns.DomainSystemName2,
		"getMethodShortName":            fns.GetMethodShortName,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	for _, s := range file.Services() {
		for _, method := range s.Methods() {
			name := m.ctx.OutputPath(file).SetBase(method.Name().LowerSnakeCase().String()).SetExt(".pb." + l.FileExtension())
			m.OverwriteGeneratorTemplateFile(name.String(), m.Tpl, method)
		}
	}
}
