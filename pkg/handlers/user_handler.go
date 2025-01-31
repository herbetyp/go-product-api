package handlers

import (
	"github.com/herbetyp/go-product-api/internal/models"
	model "github.com/herbetyp/go-product-api/internal/models/user"
	"github.com/herbetyp/go-product-api/pkg/services"
	service "github.com/herbetyp/go-product-api/pkg/services"
	"github.com/herbetyp/go-product-api/utils"
)

func CreateUser(data models.UserDTO) (models.User, error) {
	user := models.NewUser(data.Username, data.Email, data.Password)

	user.Password, _ = service.HashPassword(user.Password)

	u, err := model.Create(*user)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func GetUser(id uint) (models.User, error) {
	user, err := model.Get(id)

	if err != nil {
		return models.User{}, err

	}
	return user, nil
}

func GetUsers() ([]models.User, error) {
	users, err := model.GetAll()

	if err != nil {
		return []models.User{}, err
	}
	return users, nil
}

func UpdateUser(id uint, data models.UserDTO) (models.User, error) {
	user := models.NewUserWithID(id, data.Username, data.Password)

	user.Password, _ = service.HashPassword(user.Password)

	u, err := model.Update(*user)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func DeleteUser(id uint, hardDelete string) (bool, error) {
	deleted, err := model.Delete(id, hardDelete)

	if err != nil {
		return deleted, err
	}
	return deleted, nil
}

func RecoveryUser(id uint) (models.User, error) {
	user := models.NewUserWithID(id, "", "")

	u, err := model.Recovery(*user)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func UpdateUserStatus(id uint, active bool) (bool, error) {
	updatedStatus, err := model.UpdateStatus(id, active)

	u, _ := model.Get(id)
	services.DeleteCache(utils.USER_PREFIX, u.Email, false)
	if err != nil {
		return false, err
	}
	return updatedStatus, nil
}
