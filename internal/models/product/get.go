package product

import (
	"github.com/herbetyp/go-product-api/internal/database"
	model "github.com/herbetyp/go-product-api/internal/models"
)

func Get(id uint) (model.Product, error) {
	db := database.GetDatabase()

	var p model.Product

	result := db.Model(&p).First(&p, id)
	if result.RowsAffected == 0 {
		return model.Product{}, result.Error
	}
	return p, result.Error
}
