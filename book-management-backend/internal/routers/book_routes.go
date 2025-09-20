package router

import (
	config "book-management/configs"
	"book-management/internal/handlers"
	"book-management/internal/middlewares"
	"book-management/internal/models"

	"github.com/gin-gonic/gin"
)

// func RegisterBookRoutes(rg *gin.RouterGroup, handler *handlers.BookHandler) {
// 	books := rg.Group("/books")
// 	{
// 		books.GET("/:id", handler.GetBookByID)
// 		books.GET("", handler.GetAllBooks)
// 		books.POST("", handler.CreateBook)
// 		books.PATCH("/:id", handler.UpdateBook)
// 		books.DELETE("/:id", handler.DeleteBook)
// 	}
// }

func RegisterBookRoutes(rg *gin.RouterGroup, handler *handlers.BookHandler, cfg *config.Config) {
	books := rg.Group("/books")
	{
		// GET /books/:id - both admin & user can access
		books.GET("/:id", middlewares.AuthMiddleware(cfg, models.RoleAdmin, models.RoleUser), handler.GetBookByID)

		// GET /books - both admin & user can access
		books.GET("", middlewares.AuthMiddleware(cfg, models.RoleAdmin, models.RoleUser), handler.GetAllBooks)

		// POST /books - both admin & user can create
		books.POST("", middlewares.AuthMiddleware(cfg, models.RoleAdmin, models.RoleUser), handler.CreateBook)

		// PATCH /books/:id - only admin can update
		books.PATCH("/:id", middlewares.AuthMiddleware(cfg, models.RoleAdmin), handler.UpdateBook)

		// DELETE /books/:id - only admin can delete
		books.DELETE("/:id", middlewares.AuthMiddleware(cfg, models.RoleAdmin), handler.DeleteBook)
	}
}
