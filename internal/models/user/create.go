package user

import (
	"github.com/herbetyp/go-product-api/internal/database"
	model "github.com/herbetyp/go-product-api/internal/models"
	"gorm.io/gorm/clause"
)

func Create(u model.User) (model.User, error) {
	db := database.GetDatabase()

	result := db.Model(&u).Clauses(clause.Returning{}).Create(&u)
	if result.RowsAffected == 0 {
		return model.User{}, result.Error
	}

	u = *model.FilterUserResult(u)
	return u, nil
}
