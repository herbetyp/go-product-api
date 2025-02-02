package user

import (
	"github.com/herbetyp/go-product-api/internal/database"
	model "github.com/herbetyp/go-product-api/internal/models"
	"gorm.io/gorm/clause"
)

func Delete(id uint, hardDelete string) (bool, string, error) {
	db := database.GetDatabase()
	var u model.User

	if hardDelete == "true" {
		result := db.Model(&u).Clauses(clause.Returning{}).Where("id", id).Unscoped().Delete(&u)
		if result.Error != nil {
			return false, "", result.Error
		} else if result.RowsAffected == 0 {
			return false, "", nil
		}
		return true, "", nil
	}

	result := db.Model(&u).Clauses(clause.Returning{}).Where("id", id).Delete(&u)
	if result.RowsAffected == 0 {
		return false, "", nil
	} else if result.Error != nil {
		return false, "", result.Error
	}
	return true, u.Email, nil
}
