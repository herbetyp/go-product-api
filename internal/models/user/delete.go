package user

import (
	"github.com/herbetyp/go-product-api/internal/database"
	model "github.com/herbetyp/go-product-api/internal/models"
)

func Delete(id uint, hardDelete string) (bool, error) {
	db := database.GetDatabase()
	u := model.User{}

	if hardDelete == "true" {
		result := db.Model(&u).Where("id", id).Unscoped().Delete(&u)
		if result.Error != nil {
			return false, result.Error
		} else if result.RowsAffected == 0 {
			return false, nil
		}
		return true, nil
	}

	result := db.Model(&u).Where("id", id).Delete(&u)
	if result.RowsAffected == 0 {
		return false, nil
	} else if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
