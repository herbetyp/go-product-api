package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herbetyp/go-product-api/pkg/services"
)

func getTokenFromHeader(header string) string {
	const BearerScheme = "Bearer "

	if header == "" {
		log.Print("missing authorization header")
		return ""
	}

	if len(header) <= len(BearerScheme) {
		log.Print("invalid AAuthorization header format")
		return ""
	}

	tokenString := header[len(BearerScheme):]
	return tokenString
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		header := c.GetHeader("Authorization")
		tokenString := getTokenFromHeader(header)
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "invalid Authorization header"})
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
		tokenString := getTokenFromHeader(header)

		claims, _ := services.GetJwtClaims(tokenString)
		if userID != "" && claims["sub"] != userID {
			log.Printf("user id %s is not match sub claim %s",
				claims["sub"], userID)
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "Unauthorized"})
			return
		}
	}
}
