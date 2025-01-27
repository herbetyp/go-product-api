package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/herbetyp/go-product-api/internal/models"
	model "github.com/herbetyp/go-product-api/internal/models/product"
	"github.com/herbetyp/go-product-api/pkg/handlers"
	"github.com/herbetyp/go-product-api/utils"
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

	c.JSON(http.StatusOK, result)
}

func GetProduct(c *gin.Context) {
	id := c.Param("product-id")
	if id == "" {
		c.JSON(http.StatusBadRequest, "missing product id")
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id has to be integer",
		})
		return
	}

	result, err := handlers.GetProduct(uint(productId))

	if result == (models.Product{}) && err == nil {
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

	c.JSON(http.StatusOK, result)
}

func GetProducts(c *gin.Context) {
	result, err := handlers.GetProducts()

	if len(result) == 0 && err == nil {
		c.JSON(http.StatusNotFound,
			gin.H{"error": "not found products"})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "not get products"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("product-id")

	if id == "" {
		log.Print("missing product id")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "missing product id",
		})
		return
	}

	prodID, err := strconv.Atoi(id)

	if err != nil {
		log.Printf("id has to be integer: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id has to be integer",
		})
		return
	}

	var dto model.ProductDTO

	err = c.BindJSON(&dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid request payload")
		return
	}

	result, err := handlers.UpdateProduct(uint(prodID), dto)
	if result == (models.Product{}) && err == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "not found product",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "not updated product",
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("product-id")
	hardDelete := c.Query("hard-delete")

	if id == "" {
		log.Print("Missing product id")
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing product id"})
		return
	}

	uintID, err := utils.StringToUint(id)
	if err != nil {
		log.Printf("invalid product id: %s", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	result, err := handlers.DeleteProduct(uintID, hardDelete)

	if !result {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not deleted product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}

func RecoveryProduct(c *gin.Context) {
	id := c.Param("product-id")

	if id == "" {
		log.Print("Missing product id")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id has to be integer"})
		return
	}

	uintID, err := utils.StringToUint(id)
	if err != nil {
		log.Printf("invalid product id: %s", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	result, err := handlers.RecoveryProduct(uintID)

	if result == (models.Product{}) && err == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not recovery product"})
		return
	}

	c.JSON(http.StatusOK, result)
}
