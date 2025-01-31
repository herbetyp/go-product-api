package main

import (
	config "github.com/herbetyp/go-product-api/configs"
	"github.com/herbetyp/go-product-api/internal/database"
	"github.com/herbetyp/go-product-api/internal/database/migrations"
	"github.com/herbetyp/go-product-api/internal/server"
	"github.com/herbetyp/go-product-api/pkg/services"
)

func main() {
	// Loading App Config
	config.InitConfig()

	// Connecting on Database
	database.StartDatabase()

	// Running Migrations
	db := database.GetDatabase()
	migrations.AutoMigrations(db)

	// Starting Cache
	services.StartCache()

	// Starting Server
	runServer := server.RunServer()
	runServer.Run()
}
