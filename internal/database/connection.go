package database

import (
	"fmt"
	"log"

	config "github.com/herbetyp/go-product-api/configs"
	"github.com/herbetyp/go-product-api/internal/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDatabase() {
	DBConf := config.GetConfig().DB

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		DBConf.Host, DBConf.Port, DBConf.User, DBConf.Password, DBConf.DBName, DBConf.SSLmode)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Could not connect to database: %s", err)
		panic(err)
	}

	db = database

	config, err := database.DB()
	if err != nil {
		log.Printf("Could not get the database config: %s", err)
	}

	config.SetMaxIdleConns(DBConf.SetMaxIdleConns)
	config.SetMaxOpenConns(DBConf.SetMaxOpenConns)
	config.SetConnMaxLifetime(DBConf.SetConnMaxLifetime)

	log.Printf("Connected to database on port: %d", DBConf.Port)

	migrations.AutoMigrations(db)

}
func GetDatabase() *gorm.DB {
	return db
}
