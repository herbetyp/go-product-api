package handlers

import (
	"fmt"

	"github.com/herbetyp/go-product-api/internal/models"
	model "github.com/herbetyp/go-product-api/internal/models/login"
	"github.com/herbetyp/go-product-api/pkg/services"
	service "github.com/herbetyp/go-product-api/pkg/services"
	"github.com/herbetyp/go-product-api/utils"
)

func NewLogin(data model.LoginDTO) (string, string, uint, error) {
	var user models.User
	cacheKey := utils.USER_PREFIX + data.Email

	if services.GetCache(cacheKey, &user) == "" {
		u, err := model.Get(data.Email)
		if err != nil {
			return "", "", 0, err
		}
		services.SetCache(cacheKey, &u)
		user = u
	}

	if !user.Active {
		return "", "", 0, fmt.Errorf("user is not active")
	}

	if !service.CheckPasswordHash(data.Password, user.Password) {
		return "", "", 0, fmt.Errorf("invalid password")
	}

	token, jti, userId, err := service.GenerateToken(user.ID)
	if err != nil {
		return "", "", 0, err
	}
	return token, jti, userId, nil
}
