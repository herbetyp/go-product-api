package handlers

import (
	"github.com/herbetyp/go-product-api/internal/models"
	userModel "github.com/herbetyp/go-product-api/internal/models/user"
	service "github.com/herbetyp/go-product-api/pkg/services"
	"github.com/herbetyp/go-product-api/utils"
)

func CreateUser(data models.UserDTO) (models.User, error) {
	user := models.NewUser(data.Username, data.Email, data.Password)

	user.Password, _ = service.HashPassword(user.Password)

	u, err := userModel.Create(*user)
	if err != nil {
		return models.User{}, err
	}

	cacheKeys := []string{utils.USER_AUTHORIZATION_PREFIX + "all"}
	service.DeleteCache(cacheKeys, false)
	return u, nil
}

func GetUser(id uint, tokenUUID string) (models.User, error) {
	var user models.User

	cacheKey := utils.USER_AUTHORIZATION_PREFIX + utils.UintToString(id)
	if service.GetCache(cacheKey, &user) == "" {
		u, err := userModel.Get(id)
		if err != nil {
			return models.User{}, err
		}
		if u.ID == 0 {
			cacheKey = utils.USER_AUTHORIZATION_PREFIX + "null"
		}
		service.SetCache(cacheKey, &u)
		user = u
	}
	return user, nil
}

func GetUsers() ([]models.User, error) {
	var users []models.User

	cacheKey := utils.USER_AUTHORIZATION_PREFIX + "all"
	if service.GetCache(cacheKey, &users) == "" {
		us, err := userModel.GetAll()
		if err != nil {
			return []models.User{}, err
		}
		service.SetCache(cacheKey, &us)
		users = us
	}
	return users, nil
}

func UpdateUser(id uint, tokenUID string, data models.UserDTO) (models.User, error) {
	user := models.NewUserWithID(id, data.Username, data.Password)

	user.Password, _ = service.HashPassword(user.Password)
	u, err := userModel.Update(*user)
	if err != nil {
		return models.User{}, err
	}
	cacheKeys := []string{
		utils.USER_AUTHENTICATION_PREFIX + u.Email,
		utils.USER_AUTHORIZATION_PREFIX + utils.UintToString(id),
		utils.USER_AUTHORIZATION_PREFIX + "all",
	}
	service.DeleteCache(cacheKeys, false)
	return u, nil
}

func DeleteUser(id uint, tokenUID string, hardDelete string) (bool, error) {
	deleted, email, err := userModel.Delete(id, hardDelete)

	if err != nil {
		return deleted, err
	}
	cacheKeys := []string{
		utils.USER_AUTHENTICATION_PREFIX + email,
		utils.USER_AUTHORIZATION_PREFIX + utils.UintToString(id),
		utils.USER_AUTHORIZATION_PREFIX + "all",
	}
	service.DeleteCache(cacheKeys, false)
	return deleted, nil
}

func RecoveryUser(id uint, tokenUID string) (models.User, error) {
	user := models.NewUserWithID(id, "", "")

	u, err := userModel.Recovery(*user)
	if err != nil {
		return models.User{}, err
	}
	cacheKeys := []string{
		utils.USER_AUTHENTICATION_PREFIX + u.Email,
		utils.USER_AUTHORIZATION_PREFIX + utils.UintToString(id),
		utils.USER_AUTHORIZATION_PREFIX + "all",
	}
	service.DeleteCache(cacheKeys, false)
	return u, nil
}

func UpdateUserStatus(id uint, tokenUID string, active bool) (bool, error) {
	updatedStatus, email, err := userModel.UpdateStatus(id, active)

	if err != nil {
		return updatedStatus, err
	}
	cacheKeys := []string{
		utils.USER_AUTHENTICATION_PREFIX + email,
		utils.USER_AUTHORIZATION_PREFIX + utils.UintToString(id),
		utils.USER_AUTHORIZATION_PREFIX + "all",
	}
	service.DeleteCache(cacheKeys, false)
	return updatedStatus, nil
}
