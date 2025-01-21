package migrations

import (
	"log"

	model "github.com/herbetyp/go-product-api/internal/models"
	"gorm.io/gorm"
)

func AutoMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&model.User{})

	if err != nil {
		log.Printf("cannot migrate database: %s", err)
		panic(err)
	}
}
