package errors

import "errors"

var (
	ErrPostgres = errors.New("Database Problems")
	ErrWrongJson = errors.New("Input json is not correct")
)