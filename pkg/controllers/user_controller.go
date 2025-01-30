package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	config "github.com/herbetyp/go-product-api/configs"
	model "github.com/herbetyp/go-product-api/internal/models"
	"github.com/herbetyp/go-product-api/pkg/handlers"
	"github.com/herbetyp/go-product-api/pkg/services"
	"github.com/herbetyp/go-product-api/utils"
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
		initLog.Error("Invalid email format", zapLog.String("email", dto.Email))
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

	initLog := config.InitDefaultLogs(c)

	if err != nil {
		initLog.Error("Error on get users", zapLog.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not get users"})
		return
	}

	initLog.Info("Get users successfully")
	c.JSON(http.StatusOK, result)
}

func GetUser(c *gin.Context) {
	id := c.Param("user-id")

	initLog := config.InitDefaultLogs(c)

	uintID, err := utils.StringToUint(id)
	if err != nil {
		initLog.Error("Invalid type user id", zapLog.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	result, err := handlers.GetUser(uintID)

	if result == (model.User{}) && err == nil {
		initLog.Error("User not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	} else if err != nil {
		initLog.Error("Error on get user", zapLog.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not getting user"})
		return
	}

	initLog.Info("Get user successfully", zapLog.Uint("user_id", result.ID),
		zapLog.String("email", result.Email), zapLog.String("username", result.Username))
	c.JSON(http.StatusOK, result)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("user-id")

	initLog := config.InitDefaultLogs(c)

	uintID, err := utils.StringToUint(id)
	if err != nil {
		initLog.Error("Invalid type user id", zapLog.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var dto model.UserDTO

	err = c.BindJSON(&dto)
	if err != nil {
		initLog.Error("invalid request payload", zapLog.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	result, err := handlers.UpdateUser(uintID, dto)

	if result == (model.User{}) && err == nil {
		initLog.Error("User not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	} else if err != nil {
		initLog.Error("Error on update user", zapLog.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "not updated user"})
		return
	}

	initLog.Info("User updated successfully", zapLog.Uint("user_id", result.ID),
		zapLog.String("email", result.Email), zapLog.String("username", result.Username))
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

	initLog := config.InitDefaultLogs(c)

	uintID, err := utils.StringToUint(id)
	if err != nil {
		initLog.Error("Invalid type user id", zapLog.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	result, err := handlers.RecoveryUser(uintID)

	if result == (model.User{}) && err == nil {
		initLog.Error("User not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	} else if err != nil {
		initLog.Error("Error on recovery user", zapLog.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "not recovery user"})
		return
	}

	initLog.Info("User recovery successfully", zapLog.Uint("user_id", result.ID),
		zapLog.String("email", result.Email), zapLog.String("username", result.Username))
	c.JSON(http.StatusOK, result)
}

func UpdateUserStatus(c *gin.Context) {
	id := c.Param("user-id")
	status := c.Query("active")

	initLog := config.InitDefaultLogs(c)

	uintID, err := utils.StringToUint(id)
	if err != nil {
		initLog.Error("Invalid user id", zapLog.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	active, err := utils.StringToBoolean(status)
	if err != nil {
		initLog.Error("Invalid status", zapLog.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		return
	}

	result, err := handlers.UpdateUserStatus(uintID, active)
	if !result {
		initLog.Error("User not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	} else if err != nil {
		initLog.Error("Error on update user status", zapLog.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not updated user status"})
		return
	}

	initLog.Info("User status updated successfully", zapLog.String("user_id", id),
		zapLog.Bool("active", active))
	c.JSON(http.StatusOK, gin.H{"message": "User status updated"})
}
