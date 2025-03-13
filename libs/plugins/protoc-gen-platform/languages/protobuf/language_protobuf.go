package _go

import (
	"os"
	"strings"
	"text/template"

	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	"github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/definition"
	"github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/shared"
)

type LanguageProtobuf struct {
	Language definition.Language
	Params   *pgs.Parameters
}

func GetLanguage(templateName string, ctx pgsgo.Context, params pgs.Parameters) definition.LanguageDefinition {
	fns := shared.Functions{Pctx: pgsgo.InitContext(params)}

	tpl := template.New(templateName).Funcs(map[string]interface{}{
		"package":           ctx.PackageName,
		"name":              ctx.Name,
		"configuration":     fns.Configuration,
		"configurationName": fns.ConfigurationName,
		"isConfiguration":   fns.IsConfiguration,
		"isLast":            fns.IsLast,
		"pluginName":        fns.PluginName(ctx.Params()),
	})

	return &LanguageProtobuf{
		Language: definition.Language{
			Tpl:    tpl,
			Params: params,
		},
		Params: &params,
	}
}

func (languageProtobuf *LanguageProtobuf) FileExtension() (ext string) {
	return "proto"
}

func (languageProtobuf *LanguageProtobuf) FilePath(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath {
	fullPath := strings.Replace("", ".", string(os.PathSeparator), -1)
	filePath := pgs.JoinPaths(fullPath)
	return &filePath
}

func (languageProtobuf *LanguageProtobuf) MultiFilePath(f pgs.File, m pgs.Message) pgs.FilePath {
	panic("not implemented")
}

func (languageProtobuf *LanguageProtobuf) MultiFilePathService(f pgs.File, m pgs.Service) pgs.FilePath {
	panic("not implemented")
}

func (languageProtobuf *LanguageProtobuf) Template() *template.Template {
	return languageProtobuf.Language.Tpl
}
