package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/herbetyp/go-product-api/utils"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = utils.NewUUID()
			c.Request.Header.Set("X-Request-ID", requestID)
		}

		c.Header("X-Request-ID", requestID)
		c.Next()
	}
}
