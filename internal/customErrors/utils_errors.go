package customErrors

import "errors"

var (
	ErrSanitizing    = errors.New("data sanitizing error")
	ErrSanitizer     = errors.New("internal sanitizer error")
	ErrValidation = errors.New("data Validation error")
)
