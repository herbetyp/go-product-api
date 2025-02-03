package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herbetyp/go-product-api/internal/models"
	userModel "github.com/herbetyp/go-product-api/internal/models/user"
	"github.com/herbetyp/go-product-api/pkg/services"
	"github.com/herbetyp/go-product-api/utils"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		tokenString, _ := utils.GetTokenFromHeader(header)

		claims, _ := utils.GetJwtClaims(tokenString)
		sub := claims["sub"].(string)
		uintSub, _ := utils.StringToUint(sub)

		var user models.User

		cacheKey := utils.USER_AUTHORIZATION_PREFIX + sub
		if services.GetCache(cacheKey, &user) == "" {
			u, err := userModel.Get(uintSub)
			if err != nil {
				log.Printf("error on get user in middleware: %s", err)
				c.AbortWithStatusJSON(http.StatusUnauthorized,
					gin.H{"error": "Unauthorized"})
				return

			}
			user = u
		}

		if !user.Active {
			log.Printf("user is not active")
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

		if !services.CheckAdminPermissions(c.Request.Method, c.Request.URL.Path, sub) {
			log.Printf("Action is not authorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "Unauthorized"})
			return
		}

	}
}
