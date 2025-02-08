package services

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	config "github.com/herbetyp/go-product-api/configs"
	"github.com/herbetyp/go-product-api/internal/helpers"
)

func GenerateToken(id uint) (string, string, uint, error) {
	JWTConf := config.GetConfig().JWT
	jti := helpers.NewUUID()

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub":     fmt.Sprint(id),
		"iss":     "auth-product-api",
		"aud":     "api://go-product-api",
		"exp":     time.Now().Add(time.Duration(JWTConf.ExpiresIn) * time.Second).Unix(),
		"iat":     time.Now().Unix(),
		"jti":     jti,
		"version": JWTConf.Version,
	})

	t, err := token.SignedString(helpers.GetPrivateKey())
	if err != nil {
		log.Printf("error generating token: %s", err)
		return "", "", 0, err
	}

	return t, jti, id, nil
}

func ValidateToken(token string) (bool, jwt.MapClaims, error) {
	tokenDecoded, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodRSA); !isValid {
			return nil, fmt.Errorf("invalid token signing method")
		}
		return helpers.GetPrivateKey().Public(), nil
	})

	if err != nil {
		log.Printf("%s", err)
		return false, jwt.MapClaims{}, err
	}

	claims, err := helpers.GetJwtClaims(tokenDecoded.Raw)
	if err != nil {
		log.Printf("not get claim: %s", err)
	}

	// Validate default claims
	if claims["iss"] != "auth-product-api" {
		log.Print("invalid iss claim")
		return false, jwt.MapClaims{}, fmt.Errorf("invalid iss claim")

	} else if claims["aud"] != "api://go-product-api" {
		log.Print("invalid version claim")
		return false, jwt.MapClaims{}, fmt.Errorf("invalid version claim")
	}
	return true, claims, nil
}
