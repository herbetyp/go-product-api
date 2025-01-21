package login

import (
	"time"

	"github.com/herbetyp/go-product-api/internal/database"
	model "github.com/herbetyp/go-product-api/internal/models"
)

func Get(email string) (model.User, error) {
	db := database.GetDatabase()

	var u model.User

	u.LastLogin = time.Now().Local()

	err := db.Model(u).Where("email = ?", email).Find(&u).First(&u).Error

	u = *model.FilterResponse(u)
	return u, err
}
