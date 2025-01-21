package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/herbetyp/go-product-api/pkg/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	base_url := router.Group("/v1")

	base_url.GET("/ping", func(ctx *gin.Context) { ctx.JSON(200, gin.H{"message": "pong"}) })

	base_url.POST("oauth2/token", controllers.NewLogin)

	users := base_url.Group("/users")
	users.POST("", controllers.NewUser)

	return router
}
