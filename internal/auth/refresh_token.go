package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
)

func MakeRefreshToken() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", errors.New("unable to generate random data")
	}

	refreshToken := hex.EncodeToString(key)

	return refreshToken, nil
}
