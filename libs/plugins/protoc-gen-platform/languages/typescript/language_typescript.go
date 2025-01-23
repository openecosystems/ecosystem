package _typescript

import (
	"os"
	"strings"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	"libs/plugins/protoc-gen-platform/definition"
	"libs/plugins/protoc-gen-platform/shared"
)

type LanguageTypescript struct {
	Language definition.Language
	Params   *pgs.Parameters
}

func GetLanguage(templateName string, ctx pgsgo.Context, params pgs.Parameters) definition.LanguageDefinition {
	fns := shared.Functions{Pctx: pgsgo.InitContext(params)}

	tpl := template.New(templateName).Funcs(map[string]interface{}{
		"package":            typeScriptPackage,
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

	return &LanguageTypescript{
		Language: definition.Language{
			Tpl:    tpl,
			Params: params,
		},
		Params: &params,
	}
}

func (languageTypeScript *LanguageTypescript) FileExtension() (ext string) {
	return "ts"
}

func (languageTypeScript *LanguageTypescript) FilePath(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath {
	fullPath := strings.Replace(typeScriptPackage(f), ".", string(os.PathSeparator), -1)
	filePath := pgs.JoinPaths(fullPath)
	return &filePath
}

func (languageTypeScript *LanguageTypescript) MultiFilePath(f pgs.File, m pgs.Message) pgs.FilePath {
	fullPath := strings.Replace(typeScriptPackage(f), ".", string(os.PathSeparator), -1)
	filePath := pgs.JoinPaths(fullPath)
	return filePath
}

func (languageTypeScript *LanguageTypescript) MultiFilePathService(f pgs.File, s pgs.Service) pgs.FilePath {
	fullPath := strings.Replace(typeScriptPackage(f), ".", string(os.PathSeparator), -1)
	filePath := pgs.JoinPaths(fullPath)
	return filePath
}

func (languageTypeScript *LanguageTypescript) Template() *template.Template {
	return languageTypeScript.Language.Tpl
}

func typeScriptPackage(file pgs.File) string {
	return file.Package().ProtoName().String()
}
