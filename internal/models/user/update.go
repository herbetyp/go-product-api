package user

import (
	"github.com/herbetyp/go-product-api/internal/database"
	model "github.com/herbetyp/go-product-api/internal/models"
	"gorm.io/gorm/clause"
)

func Update(u model.User) (model.User, error) {
	db := database.GetDatabase()

	result := db.Model(&u).Clauses(clause.Returning{}).Where("id = ?", u.ID)
	if u.Username != "" {
		result = result.Update("username", u.Username)
	} else if u.Password != "" {
		result = result.Update("password", u.Password)
	} else {
		result = result.Updates(map[string]interface{}{
			"username": u.Username, "password": u.Password})
	}

	if result.RowsAffected == 0 {
		return model.User{}, nil
	} else if result.Error != nil {
		return model.User{}, result.Error
	}

	u = *model.FilterUserResult(u)
	return u, nil
}

func UpdateStatus(id uint, active bool) (bool, string, error) {
	db := database.GetDatabase()
	var u model.User

	result := db.Model(&u).Clauses(clause.Returning{}).
		Where("id = ?", id).Update("active", active)
	if result.RowsAffected == 0 {
		return false, "", nil
	} else if result.Error != nil {
		return false, "", result.Error
	}
	return true, u.Email, nil
}
