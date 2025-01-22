package user

import (
	"github.com/herbetyp/go-product-api/internal/database"
	model "github.com/herbetyp/go-product-api/internal/models"
)

func Recovery(u model.User) (model.User, error) {
	db := database.GetDatabase()

	result := db.Model(&u).Unscoped().Where("id", u.ID).Update("deleted_at", nil)

	if result.RowsAffected == 0 {
		return model.User{}, result.Error
	}

	u = *model.FilterResult(u)
	return u, result.Error
}
