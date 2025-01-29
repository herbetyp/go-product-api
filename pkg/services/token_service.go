package services

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	config "github.com/herbetyp/go-product-api/configs"
	"github.com/herbetyp/go-product-api/utils"
)

func GenerateToken(id uint) (string, string, uint, error) {
	JWTConf := config.GetConfig().JWT
	jti := utils.NewUUID()

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub":     fmt.Sprint(id),
		"iss":     "auth-product-api",
		"aud":     "api://go-product-api",
		"exp":     time.Now().Add(time.Duration(JWTConf.ExpiresIn) * time.Second).Unix(),
		"iat":     time.Now().Unix(),
		"jti":     jti,
		"version": JWTConf.Version,
	})

	t, err := token.SignedString([]byte(JWTConf.SecretKey))
	if err != nil {
		log.Printf("error generating token: %s", err)
		return "", "", 0, err
	}

	return t, jti, id, nil
}

func ValidateToken(token string) (bool, jwt.MapClaims, error) {
	conf := config.GetConfig()

	// Validate token signature
	tokenDecoded, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return false, nil
		}
		return []byte(conf.JWT.SecretKey), nil
	})

	if err != nil {
		log.Printf("%s", err)
		return false, jwt.MapClaims{}, err
	}

	claims, err := utils.GetJwtClaims(tokenDecoded.Raw)
	if err != nil {
		log.Printf("not get claim: %s", err)
	}

	// Validate expiration
	if exp, ok := claims["exp"].(float64); ok {
		if int64(exp) < time.Now().Unix() {
			log.Print("expired token")
			return false, jwt.MapClaims{}, fmt.Errorf("expired token")
		}
	} else {
		log.Print("missing exp claim")
		return false, jwt.MapClaims{}, fmt.Errorf("missing exp claim")
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
