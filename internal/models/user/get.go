package user

import (
	"github.com/herbetyp/go-product-api/internal/database"
	model "github.com/herbetyp/go-product-api/internal/models"
)

func Get(id string) (model.User, error) {
	db := database.GetDatabase()

	var u model.User

	err := db.Model(u).Where("id = ?", id).First(&u).Error

	u = *model.FilterResult(u)
	return u, err
}
