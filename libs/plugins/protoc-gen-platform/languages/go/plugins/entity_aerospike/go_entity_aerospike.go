package entity

import (
	"embed"
	"libs/plugins/protoc-gen-platform/shared"
	"sort"
	"strings"
	"text/template"

	_go "libs/plugins/protoc-gen-platform/languages/go"
	"libs/plugins/protoc-gen-platform/shared"

	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
	options "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/options/v2"
)

//go:embed templates/*.tmpl
var templates embed.FS

const (
	language   = "go"
	pluginName = "entity-aerospike"
)

type GoEntityModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	Tpl *template.Template
}

func GoEntityPlugin() *GoEntityModule {
	return &GoEntityModule{ModuleBase: &pgs.ModuleBase{}}
}

func (m *GoEntityModule) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())
}

func (m *GoEntityModule) Name() string { return language + "/" + pluginName }

func (m *GoEntityModule) Execute(targets map[string]pgs.File, _ map[string]pgs.Package) []pgs.Artifact {
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
		var entity options.EntityOptions
		_, err := t.Extension(options.E_Entity, &entity)
		if err != nil {
			panic(err.Error() + "unable to read extension from proto")
		}

		if !fns.IsSupportedLanguage(fns.LanguageOptions(t), paramLanguage) {
			continue
		}

		if entity.Type != options.EntityType_ENTITY_TYPE_AEROSPIKE {
			continue
		}

		m.GenerateFile(t)
	}

	return m.Artifacts()
}

func (m GoEntityModule) GenerateFile(file pgs.File) {
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
		"entityType":                    fns.GetEntityType,
		"isString":                      fns.IsGoString,
		"isDuration":                    fns.IsGoDuration,
		"isTimestamp":                   fns.IsGoTimestamp,
		"isInteger32":                   fns.IsGoInteger32,
		"isUnsignedInteger32":           fns.IsGoUnsignedInteger32,
		"isInteger64":                   fns.IsGoInteger64,
		"isUnsignedInteger64":           fns.IsGoUnsignedInteger64,
		"isFloat32":                     fns.IsGoFloat32,
		"isFloat64":                     fns.IsGoFloat64,
		"isBoolean":                     fns.IsBoolean,
		"isByte":                        fns.IsGoByte,
		"isMap":                         fns.IsGoMap,
		"isStruct":                      fns.IsGoStruct,
		"isStructPB":                    fns.IsGoStructPB,
		"structType":                    fns.GetStructType,
		"structTypePackage":             fns.GetStructTypePackage,
		"sliceValueType":                fns.GetGoSliceValueType,
		"sliceValueTypePackage":         fns.GetGoSliceValueTypePackage,
		"mapKeyType":                    fns.GetGoMapKeyType,
		"mapKeyTypePackage":             fns.GetGoMapKeyTypePackage,
		"mapValueType":                  fns.GetGoMapValueType,
		"mapValueTypePackage":           fns.GetGoMapValueTypePackage,
		"enumTypePackage":               fns.GetEnumTypePackage,
		"entityNameFromField":           fns.EntityNameFromField,
		"binName":                       fns.BinName,
		"goPackageOverwrite":            fns.GoPackageOverwrite,
		"getImportPackages":             fns.GetGoImportPackagesServer,
		"getImportName":                 fns.GetGoImportNameMessage,
		"getEntityGoPackage":            fns.GetGoPackageFromFile,
		"getAllGoFieldLevelImportPaths": fns.GetAllGoFieldLevelImportPaths,
		"doesImportPathContainAnyPb":    fns.DoesImportPathContainAnyPb,
		"domainSystemName2":             fns.DomainSystemName2,
	})
	template.Must(tpl.ParseFS(templates, "templates/*"))
	m.Tpl = tpl

	msg := fns.Entity(file)
	if msg == nil {
		return
	}

	name := m.ctx.OutputPath(file).SetExt(".entity-aerospike." + l.FileExtension())
	m.OverwriteGeneratorTemplateFile(name.String(), m.Tpl, file)
}
