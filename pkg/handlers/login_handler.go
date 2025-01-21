package handlers

import (
	"fmt"
	"log"

	model "github.com/herbetyp/go-product-api/internal/models/login"
	service "github.com/herbetyp/go-product-api/pkg/services"
)

func NewLogin(data model.LoginDTO) (string, error) {
	user, err := model.Get(data.Email)
	if err != nil {
		log.Printf("error on get user from database: %s", err)
		return "", fmt.Errorf("error on get user from database")
	}

	if user.Password != service.SHA512Encoder(data.Password) {
		log.Print("invalid password")
		return "", fmt.Errorf("invalid password")
	}

	token, err := service.GenerateToken(user.ID)
	if err != nil {
		log.Printf("error on generate token: %s", err)
		return "", fmt.Errorf("error on generate token")
	}
	return token, nil
}
