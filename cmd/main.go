package main

import (
	config "github.com/herbetyp/go-product-api/configs"
	"github.com/herbetyp/go-product-api/internal/database"
	"github.com/herbetyp/go-product-api/internal/server"
)

func main() {
	// Loading App Config
	config.InitConfig()

	// Connecting on Database
	database.StartDatabase()

	// Starting Server
	runServer := server.RunServer()
	runServer.Run()
}
