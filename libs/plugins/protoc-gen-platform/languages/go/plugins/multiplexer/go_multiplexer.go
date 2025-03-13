package multiplexer

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
	pluginName = "multiplexer"
)

type GoServerModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	Tpl *template.Template
}

func GoMultiplexerPlugin() *GoServerModule {
	return &GoServerModule{ModuleBase: &pgs.ModuleBase{}}
}

func (m *GoServerModule) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())
}

func (m *GoServerModule) Name() string { return language + "/" + pluginName }

func (m *GoServerModule) Execute(targets map[string]pgs.File, _ map[string]pgs.Package) []pgs.Artifact {
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

		if !fns.IsSupportedLanguage(fns.LanguageOptions(t), paramLanguage) {
			continue
		}

		if !fns.HasAnyCQRSMethods(t) {
			continue
		}

		m.GenerateFile(t)
	}

	return m.Artifacts()
}

func (m GoServerModule) GenerateFile(file pgs.File) {
	templateName := "file.go.tmpl"
	fns := shared.Functions{Pctx: pgsgo.InitContext(m.Parameters())}
	l := _go.GetLanguage(templateName, m.ctx, m.Parameters())

	tpl := l.Template()
	tpl.Funcs(map[string]interface{}{
		"getGithubRepositoryConstant":   fns.GetGithubRepositoryConstant,
		"service":                       fns.Service,
		"parentService":                 fns.ParentService,
		"queries":                       fns.QueryMethods,
		"mutations":                     fns.MutationMethods,
		"getCqrsType":                   fns.GetCQRSType,
		"goPath":                        fns.GoPath,
		"goPackage":                     fns.GoPackage,
		"goPackageOverwrite":            fns.GoPackageOverwrite,
		"getImportPackages":             fns.GetGoImportPackagesServer,
		"getImportName":                 fns.GetGoImportNameMessage,
		"hasField":                      fns.HasField,
		"goPackageRemote":               fns.GetRemoteProtoGoPathFromFile,
		"goSystemGrpcPackageRemote":     fns.GetRemoteGrpcGoPathFromFile,
		"getAllGoFieldLevelImportPaths": fns.GetAllGoFieldLevelImportPaths,
		"getSpecTypePackage":            fns.GetSpecTypePackage,
		"getPlatformTypePackage":        fns.GetPlatformTypePackage,
		"getPackageVersion":             fns.GetPackageVersion,
		"getApiOptionsTypeName":         fns.GetApiOptionsTypeName,
		"domainSystemName2":             fns.DomainSystemName2,
		"getTopLevelFolderFromFile":     fns.GetTopLevelFolderFromFile,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	msg := fns.Entity(file)
	if msg == nil {
		return
	}

	name := m.ctx.OutputPath(file).SetExt(".multiplexer." + l.FileExtension())
	m.OverwriteGeneratorTemplateFile(name.String(), m.Tpl, file)
}
