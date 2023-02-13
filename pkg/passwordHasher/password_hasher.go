package passwordHasher

import (
	"Diploma/internal/customErrors"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type PasswordHasher struct {

}

func NewPasswordHasher() *PasswordHasher {
	return &PasswordHasher{}
}

type params struct {
	memory uint32
	iterations uint32
	parallelism uint8
	saltLength uint32
	keyLength uint32
}

func(ph *PasswordHasher) GenerateHashFromPassword(password string) (encodedHash string, err error) {
	var p = &params{
		memory: 64 * 1024, //64 mb
		iterations: 3,
		parallelism: 1,
		saltLength: 16,
		keyLength: 32,
	}

	salt, err := generateRandomBytes(p.saltLength)
	if err != nil {
		return "", customErrors.ErrHashingProblems
	}

	hash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash = fmt.Sprintf("_argon2id_v=%d_m=%d,t=%d,p=%d_%s_%s", argon2.Version, p.memory, p.iterations, p.parallelism, b64Salt, b64Hash)
	return encodedHash, nil
}

func generateRandomBytes(n uint32) ([]byte, error) {
    b := make([]byte, n)
    _, err := rand.Read(b)
    if err != nil {
        return nil, err
    }

    return b, nil
}

func(ph *PasswordHasher) VerifyPassword(password, encodedHash string) (err error) {
    // Extract the parameters, salt and derived key from the encoded password
    // hash.
    p, salt, hash, err := decodeHash(encodedHash)
    if err != nil {
        return customErrors.ErrWrongPassword
    }

    // Derive the key from the other password using the same parameters.
    otherHash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

    // Check that the contents of the hashed passwords are identical. Note
    // that we are using the subtle.ConstantTimeCompare() function for this
    // to help prevent timing attacks.
    if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
        return nil
    }
    return customErrors.ErrWrongPassword
}

func decodeHash(encodedHash string) (p *params, salt, hash []byte, err error) {
    vals := strings.Split(encodedHash, "_")
    if len(vals) != 6 {
        return nil, nil, nil, errors.New("the encoded hash is not in the correct format")
    }

    var version int
    _, err = fmt.Sscanf(vals[2], "v=%d", &version)
    if err != nil {
        return nil, nil, nil, err
    }
    if version != argon2.Version {
        return nil, nil, nil, errors.New("incompatible version of argon2")
    }

    p = &params{}
    _, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
    if err != nil {
        return nil, nil, nil, err
    }

    salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
    if err != nil {
        return nil, nil, nil, err
    }
    p.saltLength = uint32(len(salt))

    hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
    if err != nil {
        return nil, nil, nil, err
    }
    p.keyLength = uint32(len(hash))

    return p, salt, hash, nil
}