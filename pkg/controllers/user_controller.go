package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/herbetyp/go-product-api/internal/models"
	"github.com/herbetyp/go-product-api/pkg/handlers"
	"github.com/herbetyp/go-product-api/pkg/services"
	utils "github.com/herbetyp/go-product-api/pkg/services/helpers"
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

func GetUser(c *gin.Context) {
	id := c.Param("user_id")

	if id == "" {
		log.Print("Missing user id")
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing user id"})
		return
	}

	if !utils.UUIDValidate(id) {
		log.Printf("ID has to be uuid: %s", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID has to be uuid"})
		return
	}

	result, err := handlers.GetUser(id)

	if result == (model.User{}) {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not getting user"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetUsers(c *gin.Context) {
	result, err := handlers.GetUsers()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "get users"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("user_id")

	if id == "" {
		log.Print("Missing user id")
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing user ID"})
		return
	}

	if !utils.UUIDValidate(id) {
		log.Printf("ID has to be uuid: %s", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID has to be uuid"})
		return
	}

	var dto model.UserDTO

	err := c.BindJSON(&dto)
	if err != nil {
		log.Printf("invalid request payload: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	result, err := handlers.UpdateUser(id, dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not user updated"})
		return
	} else if result == (model.User{}) {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("user_id")
	hardDelete := c.Query("hard-delete")

	if id == "" {
		log.Print("Missing user id")
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing user ID"})
		return
	}

	if !utils.UUIDValidate(id) {
		log.Printf("ID has to be uuid: %s", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID has to be uuid"})
		return
	}

	result, err := handlers.DeleteUser(id, hardDelete)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not deleted"})
		return
	} else if result == (model.User{}) {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func RecoveryUser(c *gin.Context) {
	id := c.Param("user_id")

	if id == "" {
		log.Print("Missing user id")
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID has to be integer"})
		return
	}

	if !utils.UUIDValidate(id) {
		log.Printf("ID has to be uuid: %s", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID has to be uuid"})
		return
	}

	result, err := handlers.RecoveryUser(id)

	if err != nil {
		log.Printf("Error on recovery user: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "recovery user"})
		return
	} else if result == (model.User{}) {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, result)
}
