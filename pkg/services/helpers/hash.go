package helpers

import (
	"crypto/sha512"
	"fmt"
)

func HashPassword(s string) string {
	stringHash := sha512.Sum512([]byte(s))

	return fmt.Sprintf("%x", stringHash)
}
