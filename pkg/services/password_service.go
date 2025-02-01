package services

import "golang.org/x/crypto/bcrypt"

func HashPassword(s string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), 8)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
