package services

import (
	"crypto/sha512"
	"fmt"
)

func SHA512Crypto(s string) string {
	stringHash := sha512.Sum512([]byte(s))

	return fmt.Sprintf("%x", stringHash)
}
