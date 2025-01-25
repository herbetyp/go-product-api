package product

import (
	"github.com/herbetyp/go-product-api/internal/database"
	"github.com/herbetyp/go-product-api/internal/models"
)

func GetAll() ([]models.Product, error) {
	db := database.GetDatabase()

	var p []models.Product

	result := db.Model(&p).Find(&p)
	if result.RowsAffected == 0 {
		return []models.Product{}, result.Error
	}
	return p, nil
}
