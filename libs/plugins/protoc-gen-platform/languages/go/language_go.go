package _go

import (
	"os"
	"strings"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	"libs/plugins/protoc-gen-platform/definition"
	"libs/plugins/protoc-gen-platform/shared"
)

type LanguageGo struct {
	Language definition.Language
	Params   *pgs.Parameters
}

func GetLanguage(templateName string, ctx pgsgo.Context, params pgs.Parameters) definition.LanguageDefinition {
	fns := shared.Functions{Pctx: pgsgo.InitContext(params)}

	tpl := template.New(templateName).Funcs(map[string]interface{}{
		"package":            ctx.PackageName,
		"name":               ctx.Name,
		"entity":             fns.Entity,
		"parentEntity":       fns.ParentEntity,
		"entityName":         fns.EntityName,
		"entityNamePlural":   fns.EntityNamePlural,
		"entityNamespace":    fns.EntityNamespace,
		"entityType":         fns.EntityType,
		"entityTypeFromFile": fns.EntityTypeFromFile,
		"entityKeyName":      fns.EntityKeyName,
		"isEntity":           fns.IsEntity,
		"isLast":             fns.IsLast,
		"pluginName":         fns.PluginName(ctx.Params()),
	})

	return &LanguageGo{
		Language: definition.Language{
			Tpl:    tpl,
			Params: params,
		},
		Params: &params,
	}
}

func (languageGo *LanguageGo) FileExtension() (ext string) {
	return "go"
}

func (languageGo *LanguageGo) FilePath(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath {
	fullPath := strings.Replace("", ".", string(os.PathSeparator), -1)
	filePath := pgs.JoinPaths(fullPath)
	return &filePath
}

func (languageGo *LanguageGo) MultiFilePath(f pgs.File, m pgs.Message) pgs.FilePath {
	panic("Go doesn't support multi-file paths")
}

func (languageGo *LanguageGo) MultiFilePathService(f pgs.File, m pgs.Service) pgs.FilePath {
	panic("implement me")
}

func (languageGo *LanguageGo) Template() *template.Template {
	return languageGo.Language.Tpl
}
