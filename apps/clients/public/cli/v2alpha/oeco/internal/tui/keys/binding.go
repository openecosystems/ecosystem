package keys

type KeyBindingType int64

const (
	Undefined KeyBindingType = iota
	Enclave
	Context
	Organization
	Package
	Connector
	Api
	Ecosystem
)

type KeyBinding interface {
	Name() KeyBindingType
}

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
	case Api:
		return "api"
	case Ecosystem:
		return "ecosystem"
	case Undefined:
		return "undefined"
	default:
		return "undefined"
	}
}
