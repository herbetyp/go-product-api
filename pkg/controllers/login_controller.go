package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	config "github.com/herbetyp/go-product-api/configs"
	model "github.com/herbetyp/go-product-api/internal/models/login"
	handler "github.com/herbetyp/go-product-api/pkg/handlers"
)

func NewLogin(c *gin.Context) {
	var dto model.LoginDTO

	err := c.BindJSON(&dto)
	if err != nil {
		log.Printf("invalid payload: %s", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	if dto.GranType != "client_credentials" {
		log.Printf("invalid grant type: %s", dto.GranType)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid grant type"})
		return
	}

	token, err := handler.NewLogin(dto)

	if err != nil || token == "" {
		log.Printf("error on login: %s", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	JWTConf := config.GetConfig().JWT
	c.JSON(http.StatusOK, gin.H{"access_token": token, "token_type": "Bearer", "expires_in": JWTConf.ExpiresIn})
}
