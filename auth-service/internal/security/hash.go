package security

import (
	"github.com/alexedwards/argon2id"
)

type ArgonHasher struct{}

func NewArgonHasher() *ArgonHasher {
	return &ArgonHasher{}
}

func (hasher *ArgonHasher) Hash(secret string) (string, error) {
	hash, err := argon2id.CreateHash(secret, argon2id.DefaultParams)

	if err != nil {
		return "", err
	}

	return hash, nil
}

func (hasher *ArgonHasher) Verify(secret string, hash string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(secret, hash)

	if err != nil {
		return false, err
	}

	return match, nil
}
