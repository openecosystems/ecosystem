package shared

// TypeParam is a constant representing the parameter name for specifying the type of implementation.
// TypeParamError is an error message for invalid `type` parameter values.
// LanguageParam is a constant representing the parameter name for specifying the programming language.
// LanguageParamError is an error message for invalid `language` parameter values.
// GoLanguageOnlyParamError is an error message indicating the `language` parameter must be set to "go".
const (
	TypeParam                = "type"
	TypeParamError           = "`type` parameter must be set to either: entity, spec, server, client, client-properties, service, listener, cqrs, configuration, billing"
	LanguageParam            = "language"
	LanguageParamError       = "`language` parameter must be set to either: go, java, csharp, swift, android, python, typescript, graphql, graphql-binding, or graphql-resolver"
	GoLanguageOnlyParamError = "`language` parameter must be set to: go"
)

// SpecTypePathPrefix defines the prefix path for platform specification types.
// PlatformTypesPathPrefix defines the prefix path for platform type definitions.
// GithubRepository specifies the GitHub repository URL of the jeannotcompany project.
// RemoteGeneratedRepository specifies the remote repository for generated code in buf.build.
const (
	SpecTypePathPrefix        = "platform/spec"
	PlatformTypesPathPrefix   = "platform/type"
	GithubRepository          = "github.com/jeannotcompany"
	RemoteGeneratedRepository = "buf.build/gen/go/jeannotcompany"
)

// GetGithubRepositoryConstant returns the constant value of the predefined GitHub repository URL.
func (fns Functions) GetGithubRepositoryConstant() string {
	return GithubRepository
}

// GetRemoteGeneratedRepositoryConstant returns the constant value of the remote generated repository identifier.
func (fns Functions) GetRemoteGeneratedRepositoryConstant() string {
	return RemoteGeneratedRepository
}
