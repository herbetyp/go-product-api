package product

import (
	"github.com/herbetyp/go-product-api/internal/database"
	model "github.com/herbetyp/go-product-api/internal/models"
	"gorm.io/gorm/clause"
)

func Create(p model.Product) (model.Product, error) {
	db := database.GetDatabase()

	result := db.Model(&p).Clauses(clause.Returning{}).Create(&p)

	if result.RowsAffected == 0 {
		return model.Product{}, result.Error
	}
	return p, nil
}
