package user

import (
	"github.com/herbetyp/go-product-api/internal/database"
	model "github.com/herbetyp/go-product-api/internal/models"
	"gorm.io/gorm/clause"
)

func Update(u model.User) (model.User, error) {
	db := database.GetDatabase()

	result := db.Model(&u).Clauses(clause.Returning{}).
		Where("id = ?", u.ID).Updates(map[string]interface{}{"username": u.Username, "password": u.Password})

	if result.RowsAffected == 0 {
		return model.User{}, result.Error
	}

	u = *model.FilterResult(u)
	return u, result.Error
}
