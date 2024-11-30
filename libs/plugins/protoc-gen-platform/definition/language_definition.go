package definition

import (
	"text/template"

	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
)

type LanguageDefinition interface {
	FileExtension() (ext string)

	FilePath(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath

	MultiFilePath(f pgs.File, m pgs.Message) pgs.FilePath

	MultiFilePathService(f pgs.File, m pgs.Service) pgs.FilePath

	Template() *template.Template
}

type Language struct {
	Tpl    *template.Template
	Params pgs.Parameters
}

func (l *Language) FileExtension() string {
	return l.FileExtension()
}

func (l *Language) RegisterEntityFunctions() {
	l.RegisterEntityFunctions()
}

func (l *Language) RegisterServiceFunctions() {
	l.RegisterServiceFunctions()
}

func (l *Language) RegisterListenerFunctions() {
	l.RegisterListenerFunctions()
}

func (l *Language) RegisterSpecFunctions() {
	l.RegisterSpecFunctions()
}
