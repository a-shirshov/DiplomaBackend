package customErrors

import "errors"

var (
	ErrUserExists = errors.New("user already exists")
	ErrPostgres = errors.New("database problems")
	ErrWrongJson = errors.New("input json is not correct")
	ErrNoTokenInContext = errors.New("no authorization token")
	ErrWrongExtension = errors.New("file extension is not supported")
	ErrWrongEmail = errors.New("email is not found")
	ErrWrongPassword = errors.New("wrong password")
	ErrUserNotFound = errors.New("user not found")
	ErrBadRequest = errors.New("bad request")
	ErrHashingProblems = errors.New("problems during hashing")
	ErrSmthWentWrong = errors.New("something went wrong")
)