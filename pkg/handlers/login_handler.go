package handlers

import (
	"fmt"
	"log"

	model "github.com/herbetyp/go-product-api/internal/models/login"
	service "github.com/herbetyp/go-product-api/pkg/services"
	"github.com/herbetyp/go-product-api/pkg/services/helpers"
)

func NewLogin(data model.LoginDTO) (string, error) {
	user, err := model.Get(data.Email)
	if err != nil {
		log.Printf("error on get user from database: %s", err)
		return "", fmt.Errorf("error on get user from database")
	}

	if user.Password != helpers.HashPassword(data.Password) {
		log.Printf("invalid password")
		return "", fmt.Errorf("invalid password")
	}

	token, err := service.GenerateToken(user.ID, user.Active)
	if err != nil {
		log.Printf("error on generate token: %s", err)
		return "", fmt.Errorf("error on generate token")
	}
	return token, nil
}
