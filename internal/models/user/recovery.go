package user

import (
	"github.com/herbetyp/go-product-api/internal/database"
	model "github.com/herbetyp/go-product-api/internal/models"
	"gorm.io/gorm/clause"
)

func Recovery(u model.User) (model.User, error) {
	db := database.GetDatabase()

	result := db.Model(&u).Clauses(clause.Returning{}).Unscoped().
		Where("id", u.ID).Update("deleted_at", nil)

	if result.RowsAffected == 0 {
		return model.User{}, nil
	} else if result.Error != nil {
		return model.User{}, result.Error
	}

	u = *model.FilterUserResult(u)
	return u, nil
}
