package handlers

import (
	"fmt"

	model "github.com/herbetyp/go-product-api/internal/models/login"
	service "github.com/herbetyp/go-product-api/pkg/services"
	"github.com/herbetyp/go-product-api/utils"
)

func NewLogin(data model.LoginDTO) (string, string, uint, error) {
	user, err := model.Get(data.Email)

	if err != nil {
		return "", "", 0, err
	}

	if !user.Active {
		return "", "", 0, fmt.Errorf("user is not active")
	}

	if user.Password != utils.HashPassword(data.Password) {
		return "", "", 0, fmt.Errorf("invalid password")
	}

	token, jti, userId, err := service.GenerateToken(user.ID)
	if err != nil {
		return "", "", 0, err
	}
	return token, jti, userId, nil
}
