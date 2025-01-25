package user

import (
	"github.com/herbetyp/go-product-api/internal/database"
	model "github.com/herbetyp/go-product-api/internal/models"
)

func Get(id uint) (model.User, error) {
	db := database.GetDatabase()

	var u model.User

	result := db.Model(&u).Where("id", id).First(&u)
	if result.RowsAffected == 0 {
		return model.User{}, nil
	} else if result.Error != nil {
		return model.User{}, result.Error
	}

	u = *model.FilterUserResult(u)
	return u, nil
}
