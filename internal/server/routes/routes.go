package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/herbetyp/go-product-api/internal/server/middlewares"
	"github.com/herbetyp/go-product-api/pkg/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	base_url := router.Group("/v1")

	base_url.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	// Login
	base_url.POST("oauth2/token", controllers.NewLogin)

	// Users
	base_url.POST("/users", controllers.CreateUser)

	users := base_url.Group("/users", middlewares.AuthMiddleware())
	users.GET("", controllers.GetUsers)

	user_id := users.Group("/:user_id", middlewares.UserMiddleware())
	user_id.GET("", controllers.GetUser)
	user_id.PATCH("", controllers.UpdateUser)
	user_id.DELETE("", controllers.DeleteUser)
	user_id.POST("/recovery", controllers.RecoveryUser)

	// Products
	products := base_url.Group("/products", middlewares.AuthMiddleware())
	products.POST("", controllers.CreateProduct)
	products.GET("", controllers.GetProducts)

	product_id := products.Group("/:product_id")
	product_id.GET("", controllers.GetProduct)

	return router
}
