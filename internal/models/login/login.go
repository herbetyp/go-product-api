package login

import (
	"time"

	"github.com/herbetyp/go-product-api/configs"
	"github.com/herbetyp/go-product-api/internal/database"
	model "github.com/herbetyp/go-product-api/internal/models"
)

func Get(email string) (model.User, error) {
	db := database.GetDatabase()
	jwtConfig := configs.GetConfig().JWT

	var u model.User

	result := db.Model(&u).Where("email", email).First(&u)
	if result.RowsAffected == 0 {
		return model.User{}, nil
	} else if result.Error != nil {
		return model.User{}, result.Error
	}

	// Update last login if it's older than 30 minutes
	then := time.Now().Add(-jwtConfig.ExpiresIn * time.Second / 2)
	if u.LastLoginAt.Before(then) {
		u.LastLoginAt = time.Now()
		db.Model(&u).Update("last_login_at", u.LastLoginAt)
	}
	return u, nil
}
