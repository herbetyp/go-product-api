package middlewares

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herbetyp/go-product-api/pkg/services"
)

func getTokenFromHeader(header string) (string, error) {
	const BearerScheme = "Bearer "

	if header == "" {
		msg := "missing Authorization header"
		log.Print(msg)
		return "", fmt.Errorf("%s", msg)
	}

	if len(header) <= len(BearerScheme) {
		msg := "invalid Authorization header format"
		log.Printf("%s: %s", msg, header)
		return "", fmt.Errorf("%s", msg)
	}

	tokenString := header[len(BearerScheme):]
	return tokenString, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		header := c.GetHeader("Authorization")
		tokenString, err := getTokenFromHeader(header)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": err.Error()})
			return
		}

		ok, _, err := services.ValidateToken(tokenString)

		if !ok || err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "Unauthorized"})
			return
		}
	}
}

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("user_id")

		header := c.GetHeader("Authorization")
		tokenString, _ := getTokenFromHeader(header)

		claims, _ := services.GetJwtClaims(tokenString)
		sub := claims["sub"]

		if claims["active"] != true {
			log.Printf("user is not active")
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "Unauthorized"})
			return
		}

		if userID == sub && c.Request.Method == "DELETE" {
			log.Printf("is not authorized self delete user %s", sub)
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "Unauthorized"})
			return
		}

		if userID != "" && sub != userID {
			log.Printf("user id %s is not match sub claim %s", sub, userID)
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "Unauthorized"})
			return
		}
	}
}
