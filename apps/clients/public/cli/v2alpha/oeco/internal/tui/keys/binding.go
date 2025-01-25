package keys

// KeyBindingType represents a custom type for defining various key binding categories in the system.
type KeyBindingType int64

// Undefined represents an uninitialized or unknown KeyBindingType value.
// Enclave represents a KeyBindingType scoped to an enclave.
// Context represents a KeyBindingType scoped to a context.
// Organization represents a KeyBindingType scoped to an organization.
// Package represents a KeyBindingType scoped to a package.
// Connector represents a KeyBindingType scoped to a connector.
// Api represents a KeyBindingType scoped to an API.
// Ecosystem represents a KeyBindingType scoped to an ecosystem.
const (
	Undefined KeyBindingType = iota
	Enclave
	Context
	Organization
	Package
	Connector
	API
	Ecosystem
)

// KeyBinding defines an interface for types that can provide a name represented as a KeyBindingType.
type KeyBinding interface {
	Name() KeyBindingType
}

// String returns the string representation of the KeyBindingType enumeration value.
func (k KeyBindingType) String() string {
	switch k {
	case Enclave:
		return "enclave"
	case Context:
		return "context"
	case Organization:
		return "organization"
	case Package:
		return "package"
	case Connector:
		return "connector"
	case API:
		return "api"
	case Ecosystem:
		return "ecosystem"
	case Undefined:
		return "undefined"
	default:
		return "undefined"
	}
}
