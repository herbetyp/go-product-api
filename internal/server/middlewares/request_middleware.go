package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herbetyp/go-product-api/configs"
	"github.com/herbetyp/go-product-api/utils"
	"golang.org/x/time/rate"
)

func RateLimitByIPMiddleware() gin.HandlerFunc {
	apiConfig := configs.GetConfig().API
	clients := make(map[string]*rate.Limiter)

	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter, exists := clients[ip]
		if !exists {
			// Create a new limiter for the IP
			limiter = rate.NewLimiter(rate.Limit(apiConfig.RateLimit), apiConfig.RateLimitBurst)
			clients[ip] = limiter
		}

		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests,
				gin.H{"error": "Too many requests"})
			return
		}
	}
}

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
