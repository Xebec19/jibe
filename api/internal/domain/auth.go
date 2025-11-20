package domain

import (
	"crypto/rand"
	"encoding/hex"
)

type Auth struct{}

// GenerateNonce returns a random nonce of given length
func (auth *Auth) GenerateNonce(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}
