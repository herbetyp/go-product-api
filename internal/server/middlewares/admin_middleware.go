package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/herbetyp/go-product-api/internal/models/user"
	"github.com/herbetyp/go-product-api/utils"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		tokenString, _ := utils.GetTokenFromHeader(header)

		claims, _ := utils.GetJwtClaims(tokenString)
		sub := claims["sub"].(string)
		uintSub, _ := utils.StringToUint(sub)

		user, err := model.Get(uintSub)
		if err != nil {
			log.Printf("error on get user in middleware: %s", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "Unauthorized"})
			return
		}
		if !user.IsAdmin {
			log.Printf("user is not admin")
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "Unauthorized"})
			return
		}
	}
}
