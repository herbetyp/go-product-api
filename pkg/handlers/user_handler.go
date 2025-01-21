package handlers

import (
	"fmt"
	"log"

	"github.com/herbetyp/go-product-api/internal/models"
	model "github.com/herbetyp/go-product-api/internal/models/user"
	"github.com/herbetyp/go-product-api/pkg/services"
)

func NewUser(data models.UserDTO) (models.User, error) {
	user := models.NewUser(data.Username, data.Email, data.Password)

	user.Password = services.SHA512Encoder(user.Password)

	u, err := model.Create(*user)

	if err != nil {
		log.Printf("cannot create user: %s", err)
		return models.User{}, fmt.Errorf("cannot create user")
	}
	return u, nil
}

func GetUser(id string) (models.User, error) {
	user, err := model.Get(id)
	if err != nil {
		return models.User{}, fmt.Errorf("cannot find user: %v", err)

	}
	return user, nil
}
