package router

import (
	"book-management/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterBookRoutes(rg *gin.RouterGroup, handler *handlers.BookHandler) {
	books := rg.Group("/books")
	{
		books.GET("/:id", handler.GetBookByID)
		books.GET("", handler.GetAllBooks)
	}
}
