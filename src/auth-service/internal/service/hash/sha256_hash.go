package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

type Sha256Hasher struct {
}

func NewSha256Hasher() *Sha256Hasher {
	return &Sha256Hasher{}
}

func (b *Sha256Hasher) Hash(str string) (string, error) {
	checkSum := sha256.Sum256([]byte(str))
	return hex.EncodeToString(checkSum[:]), nil
}
