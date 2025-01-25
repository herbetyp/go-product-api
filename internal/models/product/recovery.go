package product

import (
	"github.com/herbetyp/go-product-api/internal/database"
	"github.com/herbetyp/go-product-api/internal/models"
	"gorm.io/gorm/clause"
)

func Recovery(p models.Product) (models.Product, error) {
	db := database.GetDatabase()

	result := db.Model(&p).Clauses(clause.Returning{}).Unscoped().
		Where("id", p.ID).Update("deleted_at", nil)

	if result.RowsAffected == 0 {
		return models.Product{}, nil
	} else if result.Error != nil {
		return models.Product{}, result.Error
	}
	return p, nil
}
