package errs

import "errors"

// Context error.
var (
	ErrBadContext = errors.New("context err")
)

// General validate error.
var (
	ErrValidate = errors.New("validate")
)
