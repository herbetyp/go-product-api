package services

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	config "github.com/herbetyp/go-product-api/configs"
	"github.com/herbetyp/go-product-api/pkg/services/helpers"
)

func GetJwtClaims(tokenString string) (jwt.MapClaims, error) {
	token, _, _ := jwt.NewParser().ParseUnverified(tokenString, jwt.MapClaims{})
	claims, _ := token.Claims.(jwt.MapClaims)

	return claims, nil
}

func GenerateToken(id uint) (string, error) {
	JWTConf := config.GetConfig().JWT

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub":     fmt.Sprint(id),
		"iss":     "auth-product-api",
		"aud":     "api://product-api",
		"exp":     time.Now().Add(time.Duration(JWTConf.ExpiresIn) * time.Second).Unix(),
		"iat":     time.Now().Unix(),
		"jti":     helpers.NewUUID(),
		"version": JWTConf.Version,
	})

	t, err := token.SignedString([]byte(JWTConf.SecretKey))

	if err != nil {
		log.Printf("error generating token: %s", err)
		return "", err
	}
	return t, nil
}

func ValidateToken(token string) (bool, jwt.MapClaims, error) {
	conf := config.GetConfig()

	// Validate token
	tokenDecoded, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return false, nil
		}
		return []byte(conf.JWT.SecretKey), nil
	})

	if err != nil {
		log.Printf("invalid token: %s", err)
		return false, jwt.MapClaims{}, err
	}

	claims, _ := GetJwtClaims(tokenDecoded.Raw)

	// Validate claims
	if claims["iss"] != "auth-product-api" || claims["aud"] != "api://product-api" {
		log.Printf("invalid claim: %s", err)
		return false, jwt.MapClaims{}, err
	}
	return true, claims, nil
}
