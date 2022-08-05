package errors

import "errors"

var (
	ErrPostgres = errors.New("Database Problems")
	ErrWrongJson = errors.New("Input json is not correct")
	ErrNoTokenInContext = errors.New("Token was lost on server")
)