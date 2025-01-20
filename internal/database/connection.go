package database

import (
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	config "github.com/herbetyp/go-product-api/internal/configs"
)

var db *gorm.DB

func StartDatabase() {
	DBConf := config.GetConfig().DB

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		DBConf.Host, DBConf.Port, DBConf.User, DBConf.Password, DBConf.DBName, DBConf.SSLmode)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Print("Could not connect to the Postgres database", err)
	}

	db = database

	config, err := database.DB()
	if err != nil {
		log.Print("Could not get the database config", err)
	}

	config.SetMaxIdleConns(DBConf.SetMaxIdleConns)
	config.SetMaxOpenConns(DBConf.SetMaxOpenConns)
	config.SetConnMaxLifetime(DBConf.SetConnMaxLifetime)

	log.Print("Connected to the Postgres database: ", DBConf.Port)
}

func GetDatabase() *gorm.DB {
	return db
}
