package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	config "github.com/herbetyp/go-product-api/configs"
	model "github.com/herbetyp/go-product-api/internal/models"
	"github.com/herbetyp/go-product-api/pkg/handlers"
	"github.com/herbetyp/go-product-api/pkg/services"
	"github.com/herbetyp/go-product-api/utils"
	"go.uber.org/zap"
	zapLog "go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	var dto model.UserDTO

	initLog := config.InitDefaultLogs(c)

	err := c.BindJSON(&dto)
	if err != nil {
		initLog.Error("Invalid request payload", zapLog.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if !services.ValidateEmail(dto.Email) {
		initLog.Error("Invalid email format", zap.String("email", dto.Email))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	result, err := handlers.CreateUser(dto)
	if err != nil {
		initLog.Error("Error on create user", zapLog.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not created user"})
		return
	}

	initLog.Info("User created successfully",
		zapLog.String("email", result.Email),
		zapLog.String("username", result.Username),
		zapLog.Uint("user_id", result.ID),
		zapLog.Bool("active", result.Active),
	)

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

	uintID, err := utils.StringToUint(id)
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

	uintID, err := utils.StringToUint(id)
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

	initLog := config.InitDefaultLogs(c)

	uintID, err := utils.StringToUint(id)
	if err != nil {
		initLog.Error("Invalid type user id", zapLog.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type user id"})
		return
	}

	result, err := handlers.DeleteUser(uintID, hardDelete)
	if !result {
		initLog.Error("User not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	} else if err != nil {
		initLog.Error("Error on deleted user", zapLog.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not deleted user"})
		return
	}

	initLog.Info("User deleted successfully", zapLog.String("user_id", id))
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func RecoveryUser(c *gin.Context) {
	id := c.Param("user-id")

	if id == "" {
		log.Print("Missing user id")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id has to be integer"})
		return
	}

	uintID, err := utils.StringToUint(id)
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

func UpdateUserStatus(c *gin.Context) {
	id := c.Param("user-id")
	status := c.Query("active")

	if id == "" {
		log.Print("Missing user ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing user ID"})
		return
	}

	uintID, err := utils.StringToUint(id)
	if err != nil {
		log.Printf("Invalid user id: %s", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	active, err := utils.StringToBoolean(status)
	if err != nil {
		log.Printf("Invalid status: %s", status)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	result, err := handlers.UpdateUserStatus(uintID, active)
	if !result {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not updated user status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User status updated"})
}
