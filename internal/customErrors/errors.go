package customErrors

import "errors"

var (
	ErrUserExists = errors.New("user already exists")
	ErrPostgres = errors.New("database Problems")
	ErrWrongJson = errors.New("input json is not correct")
	ErrNoTokenInContext = errors.New("token was lost on server")
	ErrWrongExtension = errors.New("file extension is not supported")
	ErrWrongEmail = errors.New("email is not found")
	ErrWrongPassword = errors.New("wrong password")
	ErrUserNotFound = errors.New("user not found")
)