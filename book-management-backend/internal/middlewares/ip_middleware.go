package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var AllowedIPs = []string{
	"127.0.0.1",
	"172.19.0.1",
}

var AllowAllIPs = true // Set to true to allow all IP addresses

func IPCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		log.Printf("Client IP: %s\n", clientIP)

		if !AllowAllIPs && !isAllowedIP(clientIP) {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Access denied: IP not allowed",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func isAllowedIP(ip string) bool {
	for _, allowed := range AllowedIPs {
		if ip == allowed {
			return true
		}
	}
	return false
}
