package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/herbetyp/go-product-api/internal/models"
	"github.com/herbetyp/go-product-api/pkg/handlers"
	"github.com/herbetyp/go-product-api/pkg/services"
	"github.com/herbetyp/go-product-api/pkg/services/helpers"
)

func CreateUser(c *gin.Context) {
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

	result, err := handlers.CreateUser(dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not created user"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetUsers(c *gin.Context) {
	result, err := handlers.GetUsers()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not get users"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetUser(c *gin.Context) {
	id := c.Param("user-id")

	if id == "" {
		log.Print("Missing user id")
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing user id"})
		return
	}

	uintID, err := helpers.StringToUint(id)
	if err != nil {
		log.Printf("invalid user id: %s", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	result, err := handlers.GetUser(uintID)

	if result == (model.User{}) && err == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not getting user"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("user-id")

	if id == "" {
		log.Print("Missing user id")
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing user ID"})
		return
	}

	uintID, err := helpers.StringToUint(id)
	if err != nil {
		log.Printf("invalid user id: %s", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var dto model.UserDTO

	err = c.BindJSON(&dto)
	if err != nil {
		log.Printf("invalid request payload: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	result, err := handlers.UpdateUser(uintID, dto)

	if result == (model.User{}) && err == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not updated user"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("user-id")
	hardDelete := c.Query("hard-delete")

	if id == "" {
		log.Print("Missing user id")
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing user ID"})
		return
	}

	uintID, err := helpers.StringToUint(id)
	if err != nil {
		log.Printf("invalid user id: %s", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	result, err := handlers.DeleteUser(uintID, hardDelete)

	if !result {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not deleted user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

func RecoveryUser(c *gin.Context) {
	id := c.Param("user-id")

	if id == "" {
		log.Print("Missing user id")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id has to be integer"})
		return
	}

	uintID, err := helpers.StringToUint(id)
	if err != nil {
		log.Printf("invalid user id: %s", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	result, err := handlers.RecoveryUser(uintID)

	if result == (model.User{}) && err == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not recovery user"})
		return
	}

	c.JSON(http.StatusOK, result)
}
