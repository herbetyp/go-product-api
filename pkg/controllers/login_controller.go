package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herbetyp/go-product-api/configs"
	logger "github.com/herbetyp/go-product-api/configs/logger"
	model "github.com/herbetyp/go-product-api/internal/models/login"
	handler "github.com/herbetyp/go-product-api/pkg/handlers"
	zapLog "go.uber.org/zap"
)

func NewLogin(c *gin.Context) {
	var dto model.LoginDTO

	initLog := logger.InitDefaultLogs(c)

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		initLog.Error("Invalid request payload", zapLog.Error(err))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if dto.GranType != "client_credentials" {
		initLog.Error("Invalid grant type", zapLog.String("grant_type", dto.GranType), zapLog.String("email", dto.Email))
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid grant type"})
		return
	}

	token, jti, userID, err := handler.NewLogin(dto)
	if err != nil || token == "" {
		initLog.Error("Error on login", zapLog.Error(err))
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Access denied"})
		return
	}

	initLog.Info("Login successful",
		zapLog.String("email", dto.Email),
		zapLog.String("jti", jti),
		zapLog.Uint("user_id", userID),
	)

	JWTConf := configs.GetConfig().JWT
	c.JSON(http.StatusOK, gin.H{"access_token": token, "token_type": "Bearer", "expires_in": JWTConf.ExpiresIn})
}
