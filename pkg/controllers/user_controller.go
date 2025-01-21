package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/herbetyp/go-product-api/internal/models"
	"github.com/herbetyp/go-product-api/pkg/handlers"
	"github.com/herbetyp/go-product-api/pkg/services"
)

func NewUser(c *gin.Context) {
	var dto model.UserDTO

	err := c.BindJSON(&dto)
	if err != nil {
		log.Printf("invalid request payload: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	if !services.ValidateEmail(dto.Email) {
		log.Printf("invalid email format: %s", dto.Email)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email format"})
		return
	}

	result, err := handlers.NewUser(dto)

	if err != nil {
		log.Printf("error creating user: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
