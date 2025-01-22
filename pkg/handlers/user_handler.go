package handlers

import (
	"log"

	"github.com/herbetyp/go-product-api/internal/models"
	model "github.com/herbetyp/go-product-api/internal/models/user"
	"github.com/herbetyp/go-product-api/pkg/services"
)

func CreateUser(data models.UserDTO) (models.User, error) {
	user := models.NewUser(data.Username, data.Email, data.Password)

	user.Password = services.SHA512Crypto(user.Password)

	u, err := model.Create(*user)
	if err != nil {
		log.Printf("cannot create user: %s", err)
		return models.User{}, err
	}
	return u, nil
}

func GetUser(id string) (models.User, error) {
	user, err := model.Get(id)
	if err != nil {
		log.Printf("cannot find user: %v", err)
		return models.User{}, err

	}
	return user, nil
}

func GetUsers() ([]models.User, error) {
	users, err := model.GetAll()
	if err != nil {
		log.Printf("cannot find users: %v", err)
		return []models.User{}, err
	}
	return users, nil
}

func UpdateUser(id string, data models.UserDTO) (models.User, error) {
	user := models.NewUserWithID(id, data.Username, data.Email, data.Password)

	user.Password = services.SHA512Crypto(user.Password)

	u, err := model.Update(*user)
	if err != nil {
		log.Printf("cannot update user: %v", err)
		return models.User{}, err
	}
	return u, nil
}

func DeleteUser(id string, hardDelete string) (models.User, error) {
	user := models.NewUserWithID(id, "", "", "")

	u, err := model.Delete(*user, hardDelete)
	if err != nil {
		log.Printf("cannot delete user: %v", err)
		return models.User{}, err
	}
	return u, nil
}

func RecoveryUser(id string) (models.User, error) {
	user := models.NewUserWithID(id, "", "", "")

	u, err := model.Recovery(*user)

	if err != nil {
		return models.User{}, err
	}

	return u, nil
}
