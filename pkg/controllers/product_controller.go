package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	logger "github.com/herbetyp/go-product-api/configs/logger"
	"github.com/herbetyp/go-product-api/internal/models"
	model "github.com/herbetyp/go-product-api/internal/models/product"
	"github.com/herbetyp/go-product-api/pkg/handlers"
	"github.com/herbetyp/go-product-api/utils"
	zapLog "go.uber.org/zap"
)

func CreateProduct(c *gin.Context) {
	var dto model.ProductDTO

	initLog := logger.InitDefaultLogs(c)

	err := c.BindJSON(&dto)
	if err != nil {
		initLog.Error("Invalid request payload", zapLog.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	result, err := handlers.CreateProduct(dto)
	if err != nil {
		initLog.Error("Error on create product", zapLog.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not created product"})
		return
	}

	initLog.Info("Product created successfully", zapLog.Uint("product_id", result.ID),
		zapLog.String("code", result.Code), zapLog.String("name", result.Name))
	c.JSON(http.StatusOK, result)
}

func GetProduct(c *gin.Context) {
	id := c.Param("product-id")

	initLog := logger.InitDefaultLogs(c)

	if id == "" {
		initLog.Error("Missing product ID")
		c.JSON(http.StatusBadRequest, "Missing product ID")
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		initLog.Error("ID has to be integer", zapLog.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID has to be integer"})
		return
	}

	result, err := handlers.GetProduct(uint(productId))
	if result == (models.Product{}) && err == nil {
		initLog.Error("Not found product")
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found product"})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not get product"})
		return
	}

	initLog.Info("Get product successfully", zapLog.Uint("product_id", result.ID),
		zapLog.String("code", result.Code), zapLog.String("name", result.Name))
	c.JSON(http.StatusOK, result)
}

func GetProducts(c *gin.Context) {
	result, err := handlers.GetProducts()

	initLog := logger.InitDefaultLogs(c)

	if len(result) == 0 && err == nil {
		initLog.Error("Not found products")
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found products"})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"error": "Not get products"})
		return
	}

	initLog.Info("Get products successfully")
	c.JSON(http.StatusOK, result)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("product-id")

	initLog := logger.InitDefaultLogs(c)

	if id == "" {
		initLog.Error("Missing product ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing product ID"})
		return
	}

	prodID, err := strconv.Atoi(id)
	if err != nil {
		initLog.Error("ID has to be integer", zapLog.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID has to be integer"})
		return
	}

	var dto model.ProductDTO

	err = c.BindJSON(&dto)
	if err != nil {
		initLog.Error("Invalid request payload", zapLog.Error(err))
		c.JSON(http.StatusBadRequest, "invalid request payload")
		return
	}

	result, err := handlers.UpdateProduct(uint(prodID), dto)
	if result == (models.Product{}) && err == nil {
		initLog.Error("Not found product")
		c.JSON(http.StatusNotFound, gin.H{"error": "not found product"})
		return
	} else if err != nil {
		initLog.Error("Error on update product", zapLog.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "not updated product"})
		return
	}

	initLog.Info("Product updated successfully", zapLog.Uint("product_id", result.ID),
		zapLog.String("code", result.Code), zapLog.String("name", result.Name))
	c.JSON(http.StatusOK, result)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("product-id")
	hardDelete := c.Query("hard-delete")

	initLog := logger.InitDefaultLogs(c)

	if id == "" {
		initLog.Error("Missing product ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing product ID"})
		return
	}

	uintID, err := utils.StringToUint(id)
	if err != nil {
		initLog.Error("Invalid product id", zapLog.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product id"})
		return
	}

	result, err := handlers.DeleteProduct(uintID, hardDelete)

	if !result {
		initLog.Error("Product not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	} else if err != nil {
		initLog.Error("Error on delete product", zapLog.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not deleted product"})
		return
	}

	initLog.Info("Product deleted successfully", zapLog.Uint("product_id", uintID))
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

func RecoveryProduct(c *gin.Context) {
	id := c.Param("product-id")

	initLog := logger.InitDefaultLogs(c)

	if id == "" {
		initLog.Error("Missing product ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID has to be integer"})
		return
	}

	uintID, err := utils.StringToUint(id)
	if err != nil {
		initLog.Error("Invalid product id", zapLog.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product id"})
		return
	}

	result, err := handlers.RecoveryProduct(uintID)
	if result == (models.Product{}) && err == nil {
		initLog.Error("Product not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	} else if err != nil {
		initLog.Error("Error on recovery product", zapLog.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not recovery product"})
		return
	}

	initLog.Info("Product recovery successfully", zapLog.Uint("product_id", result.ID))
	c.JSON(http.StatusOK, result)
}
