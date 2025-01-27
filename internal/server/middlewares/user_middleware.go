package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/herbetyp/go-product-api/internal/models/user"
	"github.com/herbetyp/go-product-api/utils"
)

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user-id")
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

		if !user.Active {
			log.Printf("user %s is not active", sub)
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "Unauthorized"})
			return
		}

		if !user.IsAdmin && sub != userId {
			log.Printf("user id %s is not match sub claim %s", userId, sub)
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "Unauthorized"})
			return
		}
	}
}
