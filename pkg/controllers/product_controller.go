package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/herbetyp/go-product-api/internal/models"
	model "github.com/herbetyp/go-product-api/internal/models/product"
	"github.com/herbetyp/go-product-api/pkg/handlers"
)

func CreateProduct(c *gin.Context) {
	var dto model.ProductDTO

	err := c.BindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "invalid request payload"})
		return
	}

	result, err := handlers.CreateProduct(dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "not created product",
		})
		return
	}

	c.JSON(200, result)
}

func GetProduct(c *gin.Context) {
	id := c.Param("product_id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "Missing product id")
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	result, err := handlers.GetProduct(uint(productId))

	if result == (models.Product{}) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "not found product",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "not get product",
		})
		return
	}

	c.JSON(200, result)
}

func GetProducts(c *gin.Context) {
	result, err := handlers.GetProducts()

	if len(result) == 0 {
		c.JSON(http.StatusNotFound,
			gin.H{"error": "not found products"})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "not get products"})
		return
	}

	c.JSON(200, result)
}
