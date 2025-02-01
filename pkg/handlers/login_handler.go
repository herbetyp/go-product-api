package handlers

import (
	"fmt"

	"github.com/herbetyp/go-product-api/internal/models"
	model "github.com/herbetyp/go-product-api/internal/models/login"
	"github.com/herbetyp/go-product-api/pkg/services"
	"github.com/herbetyp/go-product-api/utils"
)

func NewLogin(data model.LoginDTO) (string, string, uint, error) {
	var user models.User
	cacheKey := utils.USER_UID_PREFIX + utils.UseSHA256Hash(data.Email)

	cacheKeys := []string{cacheKey}
	if services.GetCache(cacheKeys, &user, []string{}) == "" {
		u, err := model.Get(data.Email)
		if err != nil {
			return "", "", 0, err
		}
		if u.ID == 0 {
			cacheKey = utils.USER_UID_PREFIX + "null"
		}
		services.SetCache(cacheKey, &u)
		user = u
	}

	if !user.Active {
		return "", "", 0, fmt.Errorf("user is not active")
	}

	if !services.CheckPasswordHash(data.Password, user.Password) {
		return "", "", 0, fmt.Errorf("invalid password")
	}

	token, jti, userId, err := services.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", "", 0, err
	}
	return token, jti, userId, nil
}
