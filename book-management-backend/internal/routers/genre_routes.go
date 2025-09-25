package router

import (
	config "book-management/configs"
	"book-management/internal/handlers"
	"book-management/internal/middlewares"
	"book-management/internal/models"

	"github.com/gin-gonic/gin"
)

func RegisterGenreRoutes(rg *gin.RouterGroup, handler *handlers.GenreHandler, cfg *config.Config) {
	genres := rg.Group("/genres")
	{
		// GET /genres/:id - both admin & user can access
		genres.GET("/:id", middlewares.AuthMiddleware(cfg, models.RoleAdmin, models.RoleUser), handler.GetGenreByID)

		// GET /genres - both admin & user can access
		genres.GET("", middlewares.AuthMiddleware(cfg, models.RoleAdmin, models.RoleUser), handler.GetAllGenres)

		// POST /genres - only admin can create
		genres.POST("", middlewares.AuthMiddleware(cfg, models.RoleAdmin), handler.CreateGenre)

		// PATCH /genres/:id - only admin can update
		genres.PATCH("/:id", middlewares.AuthMiddleware(cfg, models.RoleAdmin), handler.UpdateGenre)

		// DELETE /genres/:id - only admin can delete
		genres.DELETE("/:id", middlewares.AuthMiddleware(cfg, models.RoleAdmin), handler.DeleteGenre)

		// POST /genres/:id/books - only admin can add books to a genre
		genres.POST("/:id/books", middlewares.AuthMiddleware(cfg, models.RoleAdmin), handler.AddBooksToGenre)

		// DELETE /genres/:id/books - only admin can remove books from a genre
		genres.DELETE("/:id/books", middlewares.AuthMiddleware(cfg, models.RoleAdmin), handler.RemoveBooksFromGenre)

		// PUT /genres/:id/books - only admin can replace all books in a genre
		genres.PUT("/:id/books", middlewares.AuthMiddleware(cfg, models.RoleAdmin), handler.ReplaceBooksInGenre)
	}
}
