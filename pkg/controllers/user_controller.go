package controllers

import (
	"log"
	"net/http"
	"regexp"

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

func GetUser(c *gin.Context) {
	id := c.Param("user_id")

	if id == "" {
		log.Print("Missing user id")
		c.JSON(400, gin.H{
			"error": "Missing user id",
		})
		return
	}

	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	if !r.MatchString(id) {
		log.Printf("ID has to be uuid: %s", id)
		c.JSON(400, gin.H{
			"error": "ID has to be uuid",
		})
		return
	}

	result, err := handlers.GetUser(id)

	if err != nil {
		log.Printf("error getting user: %s", err)
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
