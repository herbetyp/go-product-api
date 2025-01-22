package user

import (
	"github.com/herbetyp/go-product-api/internal/database"
	model "github.com/herbetyp/go-product-api/internal/models"
)

func Delete(u model.User, hardDelete string) (model.User, error) {
	db := database.GetDatabase()

	if hardDelete == "true" {
		result := db.Model(&u).Where("id", u.ID).Unscoped().Delete(&u)
		if result.RowsAffected == 0 {
			return model.User{}, result.Error
		}
		return u, result.Error
	}

	result := db.Model(&u).Where("id", u.ID).Delete(&u)
	if result.RowsAffected == 0 {
		return model.User{}, result.Error
	}
	return u, result.Error
}
