package sdkv2betalib

type ErrorRegistry interface {
	AllErrors() map[Reason]SpecErrorable
}

func ProcessErrors(errors []SpecErrorable) map[Reason]SpecErrorable {
	registry := make(map[Reason]SpecErrorable)
	for _, err := range errors {
		registry[err.SpecReason()] = err
	}

	return registry
}
