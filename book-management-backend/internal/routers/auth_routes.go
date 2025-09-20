// internal/router/auth_router.go
package router

import (
	"book-management/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, handler *handlers.AuthHandler) {
	auth := rg.Group("/auth")
	{
		auth.POST("/register", handler.Register) // Register new user
		auth.POST("/login", handler.Login)       // Login, return JWT tokens
		auth.POST("/refresh", handler.Refresh)   // Refresh access token
	}
}
