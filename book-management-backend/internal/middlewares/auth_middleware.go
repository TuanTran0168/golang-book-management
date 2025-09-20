package middlewares

import (
	configs "book-management/configs"
	"book-management/internal/models"
	"book-management/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates JWT access token and attaches user claims to context
func AuthMiddleware(cfg *configs.Config, requiredRoles ...models.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from "Authorization" header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing Authorization header"})
			c.Abort()
			return
		}

		// Expected format: "Bearer <token>"
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid Authorization header format"})
			c.Abort()
			return
		}
		tokenStr := parts[1]

		// Parse and validate token
		claims, err := utils.ParseAccessToken(tokenStr, cfg)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		// Role-based access control
		if len(requiredRoles) > 0 {
			allowed := false
			for _, role := range requiredRoles {
				if claims.Role == role {
					allowed = true
					break
				}
			}
			if !allowed {
				c.JSON(http.StatusForbidden, gin.H{"error": "forbidden: insufficient permissions"})
				c.Abort()
				return
			}
		}

		// Save user info into context
		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}
