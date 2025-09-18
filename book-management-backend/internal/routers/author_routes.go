package router

import (
	"book-management/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthorRoutes(rg *gin.RouterGroup, handler *handlers.AuthorHandler) {
	authors := rg.Group("/authors")
	{
		authors.GET("", handler.GetAuthors)
		authors.GET("/:id", handler.GetAuthorByID)
		authors.POST("", handler.CreateAuthor)
		authors.PUT("/:id", handler.UpdateAuthor)
		authors.DELETE("/:id", handler.DeleteAuthor)
	}
}
