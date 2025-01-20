package main

import (
	"github.com/herbetyp/go-product-api/internal/server"
	config "github.com/herbetyp/go-product-api/internal/configs"
)

func main() {
	// Loading App Config
	config.InitConfig()

	// Connecting on Database
	// database.StartDatabase()

	// Starting Server
	runServer := server.RunServer()
	runServer.Run()
}
