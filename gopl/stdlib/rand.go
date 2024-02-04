package stdlib

import (
	"crypto/rand"
	"encoding/hex"
)

func RandHex(len int) (string, error) {
	b := make([]byte, len)

	_, err := rand.Read(b)

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b[:]), nil
}

func MustRandHex(len int) string {
	s, err := RandHex(len)
	if err != nil {
		panic(err)
	}

	return s
}
