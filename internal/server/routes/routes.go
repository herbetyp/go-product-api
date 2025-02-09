package routes

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/herbetyp/go-product-api/internal/server/middlewares"
	"github.com/herbetyp/go-product-api/pkg/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	// Base v1 API
	base_url := router.Group("/v1", middlewares.RequestIDMiddleware())
	if os.Getenv("GIN_MODE") == "release" {
		base_url.Use(middlewares.RateLimitByIPMiddleware())
	}

	// Health check
	base_url.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "Healthy"})
	})

	// Login
	base_url.POST("oauth2/token", controllers.NewLogin)

	// Admin
	admin := base_url.Group("/admin", middlewares.AuthMiddleware(), middlewares.AdminMiddleware())
	admin.GET("/users", controllers.GetUsers)
	admin.GET("/users/:user-id", controllers.GetUser)
	admin.DELETE("/users/:user-id", controllers.DeleteUser)
	admin.PATCH("/users/:user-id/status", controllers.UpdateUserStatus)
	admin.POST("/users/:user-id/recovery", controllers.RecoveryUser)
	admin.DELETE("products/:product-id", controllers.DeleteProduct)
	admin.POST("products/:product-id", controllers.RecoveryProduct)

	// Users
	base_url.POST("/users", controllers.CreateUser)
	users := base_url.Group("/users", middlewares.AuthMiddleware(), middlewares.UserMiddleware())
	users.GET("/:user-id", controllers.GetUser)
	users.PATCH("/:user-id", controllers.UpdateUser)

	// Products
	products := base_url.Group("/products", middlewares.AuthMiddleware())
	products.POST("", controllers.CreateProduct)
	products.GET("", controllers.GetProducts)
	products.GET("/:product-id", controllers.GetProduct)
	products.PUT("/:product-id", controllers.UpdateProduct)

	return router
}
