package middlewares

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/herbetyp/go-product-api/configs"
	"github.com/herbetyp/go-product-api/pkg/services"
	"github.com/herbetyp/go-product-api/utils"
	"golang.org/x/time/rate"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		tokenString, err := utils.GetTokenFromHeader(header)
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

func RateLimitByIPMiddleware() gin.HandlerFunc {
	var mu sync.Mutex
	apiConfig := configs.GetConfig().API

	clients := make(map[string]*rate.Limiter)
	return func(c *gin.Context) {
		mu.Lock()

		ip := c.ClientIP()
		limiter, exists := clients[ip]
		if !exists {
			limiter = rate.NewLimiter(1, apiConfig.RateLimit)
			clients[ip] = limiter
		}

		defer mu.Unlock()
		if !limiter.Allow() {
			log.Printf("Limit excede was block IP: %s\n", ip)
			c.AbortWithStatusJSON(http.StatusTooManyRequests,
				gin.H{"error": "Too many requests"})
			return
		}
	}
}
