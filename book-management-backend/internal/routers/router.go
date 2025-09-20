package router

import (
	config "book-management/configs"
	"book-management/internal/handlers"
	"book-management/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	authorHandler *handlers.AuthorHandler,
	bookHandler *handlers.BookHandler,
	authHandler *handlers.AuthHandler,
	cfg *config.Config,
) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())
	r.Use(middlewares.IPCheckMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello Book Management API By Tuan Tran!"})
	})

	api := r.Group("/api")

	RegisterAuthRoutes(api, authHandler)
	RegisterAuthorRoutes(api, authorHandler, cfg)
	RegisterBookRoutes(api, bookHandler, cfg)

	return r
}
