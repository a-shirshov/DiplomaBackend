package pkg

type PasswordHasher interface {
	GenerateHashFromPassword(password string) (encodedHash string, err error)
	VerifyPassword(password, encodedHash string) (match bool, err error)
}