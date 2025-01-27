package utils

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

func GetTokenFromHeader(header string) (string, error) {
	const BearerScheme = "Bearer "

	if header == "" {
		msg := "missing Authorization header"
		log.Print(msg)
		return "", fmt.Errorf("%s", msg)
	}

	if len(header) <= len(BearerScheme) {
		msg := "invalid Authorization header format"
		log.Printf("%s: %s", msg, header)
		return "", fmt.Errorf("%s", msg)
	}

	tokenString := header[len(BearerScheme):]
	return tokenString, nil
}

func GetJwtClaims(tokenString string) (jwt.MapClaims, error) {
	token, _, err := jwt.NewParser().ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
