package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func Generate(size int) string {
	buf := make([]byte, size)

	if _, err := rand.Read(buf); err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.EncodeToString(buf)
}
