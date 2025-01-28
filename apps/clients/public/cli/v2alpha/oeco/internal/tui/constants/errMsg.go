package constants

// ErrMsg represents a custom error message containing an error object.
// It implements the error interface by providing an Error method.
type ErrMsg struct {
	Err error
}

// Error returns the string representation of the wrapped error in the ErrMsg struct.
func (e ErrMsg) Error() string { return e.Err.Error() }
