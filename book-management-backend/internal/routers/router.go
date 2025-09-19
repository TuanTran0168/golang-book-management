package router

import (
	"book-management/internal/handlers"
	// "book-management/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter(authorHandler *handlers.AuthorHandler) *gin.Engine {
	r := gin.Default()
	// r.Use(middlewares.CORSMiddleware())

	api := r.Group("/api")

	RegisterAuthorRoutes(api, authorHandler)

	return r
}
