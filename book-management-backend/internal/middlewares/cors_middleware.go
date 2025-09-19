package middlewares

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Whitelist = []string{
	"https://golang-book-management-h5pt.onrender.com",
	"http://localhost:3000",
}

func CORSMiddleware() gin.HandlerFunc {
	config := cors.Config{
		AllowOrigins:     Whitelist,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	return cors.New(config)
}
