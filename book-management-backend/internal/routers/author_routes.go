package router

import (
	configs "book-management/configs"
	"book-management/internal/handlers"
	"book-management/internal/middlewares"
	"book-management/internal/models"

	"github.com/gin-gonic/gin"
)

// RegisterAuthorRoutes defines routes for authors with optional JWT auth
func RegisterAuthorRoutes(rg *gin.RouterGroup, handler *handlers.AuthorHandler, cfg *configs.Config) {
	authors := rg.Group("/authors")
	{
		// GET /authors - both admin & user can access
		authors.GET("", middlewares.AuthMiddleware(cfg, models.RoleAdmin, models.RoleUser), handler.GetAuthors)
		authors.GET("/:id", middlewares.AuthMiddleware(cfg, models.RoleAdmin, models.RoleUser), handler.GetAuthorByID)

		// POST /authors - only admin can create authors
		authors.POST("", middlewares.AuthMiddleware(cfg, models.RoleAdmin), handler.CreateAuthor)

		// PUT /authors/:id - only admin can update
		authors.PUT("/:id", middlewares.AuthMiddleware(cfg, models.RoleAdmin), handler.UpdateAuthor)

		// DELETE /authors/:id - only admin can delete
		authors.DELETE("/:id", middlewares.AuthMiddleware(cfg, models.RoleAdmin), handler.DeleteAuthor)
	}
}
