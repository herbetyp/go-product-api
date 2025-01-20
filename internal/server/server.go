package server

import (
	"log"

	"github.com/gin-gonic/gin"
	config "github.com/herbetyp/go-product-api/internal/configs"
	router "github.com/herbetyp/go-product-api/internal/server/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func RunServer() Server {
	APIConf := config.GetConfig().API

	return Server{
		port:   APIConf.Port,
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := router.ConfigRoutes(s.server)

	log.Print("Server is running on port: ", s.port)

	log.Fatal(router.Run(":" + s.port))
}
