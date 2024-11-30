package shared

// CLI Parameter constants
const (
	TypeParam                = "type"
	TypeParamError           = "`type` parameter must be set to either: entity, spec, server, client, client-properties, service, listener, cqrs, configuration, billing"
	LanguageParam            = "language"
	LanguageParamError       = "`language` parameter must be set to either: go, java, csharp, swift, android, python, typescript, graphql, graphql-binding, or graphql-resolver"
	GoLanguageOnlyParamError = "`language` parameter must be set to: go"
)

const (
	SpecTypePathPrefix        = "platform/spec"
	PlatformTypesPathPrefix   = "platform/type"
	GithubRepository          = "github.com/jeannotcompany"
	RemoteGeneratedRepository = "buf.build/gen/go/jeannotcompany"
)

func (fns Functions) GetGithubRepositoryConstant() string {
	return GithubRepository
}

func (fns Functions) GetRemoteGeneratedRepositoryConstant() string {
	return RemoteGeneratedRepository
}
