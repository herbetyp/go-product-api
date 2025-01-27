package migrations

import (
	"log"

	"github.com/herbetyp/go-product-api/internal/models"
	"gorm.io/gorm"
)

func AutoMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{}, &models.Product{})

	if err != nil {
		log.Printf("cannot migrate database: %s", err)
		panic(err)
	}
}
