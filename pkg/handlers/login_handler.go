package handlers

import (
	"fmt"

	model "github.com/herbetyp/go-product-api/internal/models/login"
	service "github.com/herbetyp/go-product-api/pkg/services"
	"github.com/herbetyp/go-product-api/utils"
)

func NewLogin(data model.LoginDTO) (string, error) {
	user, err := model.Get(data.Email)
	if err != nil {
		return "", fmt.Errorf("error on get user: %s", err)
	}

	if !user.Active {
		return "", fmt.Errorf("user is not active")
	}

	if user.Password != utils.HashPassword(data.Password) {
		return "", fmt.Errorf("invalid password")
	}

	token, err := service.GenerateToken(user.ID)
	if err != nil {
		return "", fmt.Errorf("error on generate token: %s", err)
	}
	return token, nil
}
