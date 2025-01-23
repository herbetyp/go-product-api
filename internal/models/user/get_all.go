package user

import (
	"github.com/herbetyp/go-product-api/internal/database"
	model "github.com/herbetyp/go-product-api/internal/models"
)

func GetAll() ([]model.User, error) {
	db := database.GetDatabase()

	var u []model.User

	err := db.Model(&u).Omit("password").Order("created_at DESC").Find(&u).Error

	return u, err
}
