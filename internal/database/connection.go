package database

import (
	"fmt"
	"log"

	config "github.com/herbetyp/go-product-api/internal/configs"
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
		log.Print("Could not connect to database", err)
	}

	db = database

	config, err := database.DB()
	if err != nil {
		log.Print("Could not get the database config", err)
	}

	config.SetMaxIdleConns(DBConf.SetMaxIdleConns)
	config.SetMaxOpenConns(DBConf.SetMaxOpenConns)
	config.SetConnMaxLifetime(DBConf.SetConnMaxLifetime)

	log.Print("Connected to database on port: ", DBConf.Port)
}

func GetDatabase() *gorm.DB {
	return db
}
