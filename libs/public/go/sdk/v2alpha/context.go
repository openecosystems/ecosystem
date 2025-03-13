package sdkv2alphalib

// ContextKeyType defines a custom type for context keys used to avoid collisions in context values across packages.
type ContextKeyType string

// SettingsContextKey is a context key used to store and retrieve settings-related data in a context.
// SpecContextKey is a context key used to store and retrieve specification-related data in a context.
const (
	SettingsContextKey = ContextKeyType("settings")
	SpecContextKey     = ContextKeyType("spec")
	LoggerContextKey   = ContextKeyType("logger")
	NebulaCAContextKey = ContextKeyType("nebulav1ca")
)
