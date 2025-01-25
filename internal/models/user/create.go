package user

import (
	"github.com/herbetyp/go-product-api/internal/database"
	model "github.com/herbetyp/go-product-api/internal/models"
)

func Create(u model.User) (model.User, error) {
	db := database.GetDatabase()

	err := db.Model(&u).Create(&u).Error

	u = *model.FilterUserResult(u)
	return u, err
}
