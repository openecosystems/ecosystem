package cli_system

import (
	"embed"
	"sort"
	"strings"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	_go "libs/plugins/protoc-gen-platform/languages/go"
	"libs/plugins/protoc-gen-platform/shared"
)

var (
	//go:embed templates/*.tmpl
	templates embed.FS
	goOutPath = pgs.JoinPaths("platform")
	outPath   = &goOutPath
)

const (
	language   = "go"
	pluginName = "cli-system"
)

type GoCliSystemModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	Tpl *template.Template
}

func GoCliSystemPlugin() *GoCliSystemModule {
	return &GoCliSystemModule{ModuleBase: &pgs.ModuleBase{}}
}

func (m *GoCliSystemModule) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())
}

func (m *GoCliSystemModule) Name() string { return language + "/" + pluginName }

func (m *GoCliSystemModule) Execute(targets map[string]pgs.File, _ map[string]pgs.Package) []pgs.Artifact {
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
	for k, _ := range targets {
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

func (m GoCliSystemModule) GenerateFile(file pgs.File) {
	templateName := "file.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())
	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"service":                        fns.Service,
		"serviceLeadingComment":          fns.ServiceLeadingComment,
		"serviceTrailingComment":         fns.ServiceTrailingComment,
		"serviceLeadingDetachedComments": fns.ServiceLeadingDetachedComments,
		"parentService":                  fns.ParentService,
		"queries":                        fns.QueryMethods,
		"isMethodQuery":                  fns.IsMethodQuery,
		"mutations":                      fns.MutationMethods,
		"isMethodMutation":               fns.IsMethodMutation,
		"getCqrsType":                    fns.GetCQRSType,
		"getSpecCommands":                fns.GetSpecCommands,
		"getSpecEvents":                  fns.GetSpecEvents,
		"getSpecTopics":                  fns.GetSpecTopics,
		"goPackageOverwrite":             fns.GoPackageOverwrite,
		"goPackage":                      fns.GoPackage,
		"goPath":                         fns.GoPath,
		"allMethods":                     fns.AllMethods,
		"dotNotationToFilePath":          fns.DotNotationToFilePath,
		"getPackageVersion":              fns.GetPackageVersion,
		"getPackageVersionName":          fns.GetPackageVersionName,
		"domainSystemName2":              fns.DomainSystemName2,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	msg := fns.Entity(file)
	if msg == nil {
		return
	}

	//name := m.ctx.OutputPath(file).SetExt(".cmd." + l.FileExtension())
	name := outPath.SetExt("/" + fns.DomainSystemName2(file).LowerCamelCase().String() + "/" + fns.GetPackageVersion(file) + "/" + fns.DomainSystemName2(file).LowerCamelCase().String() + ".cmd." + l.FileExtension())
	m.OverwriteGeneratorTemplateFile(name.String(), m.Tpl, file)
}
