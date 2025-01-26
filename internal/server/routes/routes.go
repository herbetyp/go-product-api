package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herbetyp/go-product-api/internal/server/middlewares"
	"github.com/herbetyp/go-product-api/pkg/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	base_url := router.Group("/v1", middlewares.RateLimitByIPMiddleware())

	base_url.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Login
	base_url.POST("oauth2/token", controllers.NewLogin)

	// Users
	base_url.POST("/users", controllers.CreateUser)
	users := base_url.Group("/users", middlewares.AuthMiddleware(), middlewares.UserMiddleware())
	users.GET("", middlewares.AdminMiddleware(), controllers.GetUsers)
	users.GET("/:user-id", controllers.GetUser)
	users.PATCH("/:user-id", controllers.UpdateUser)
	users.DELETE("/:user-id", controllers.DeleteUser)
	users.POST("/:user-id/recovery", controllers.RecoveryUser)

	// Products
	products := base_url.Group("/products", middlewares.AuthMiddleware())
	products.POST("", controllers.CreateProduct)
	products.GET("", controllers.GetProducts)
	products.GET("/:product-id", controllers.GetProduct)
	products.PUT("/:product-id", controllers.UpdateProduct)
	products.DELETE("/:product-id", controllers.DeleteProduct)
	products.POST("/:product-id", controllers.RecoveryProduct)

	return router
}
