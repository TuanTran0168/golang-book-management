package handlers

import (
	"book-management/internal/models"
	"book-management/internal/services"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	service services.IBookService
}

func NewBookHandler(service services.IBookService) *BookHandler {
	return &BookHandler{service: service}
}

type AuthorResponseForBook struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type BookResponse struct {
	ID     uint                  `json:"id"`
	Title  string                `json:"title"`
	Author AuthorResponseForBook `json:"author"`
}

func mapBookResponse(book *models.Book) BookResponse {

	AuthorResponseForBook := AuthorResponseForBook{
		ID:   book.Author.ID,
		Name: book.Author.Name,
	}

	return BookResponse{
		ID:     book.ID,
		Title:  book.Title,
		Author: AuthorResponseForBook,
	}
}

// GET /books/:id
// GetBookByID godoc
// @Summary      Get book details by ID
// @Description  Get a single book along with its author by book ID
// @Tags         books
// @Produce      json
// @Param        id   path      string  true  "Book ID"
// @Success      200  {object}  BookResponse
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /books/{id} [get]
func (h *BookHandler) GetBookByID(c *gin.Context) {
	idStr := c.Param("id")

	book, httpStatus, err := h.service.GetBookByID(idStr)
	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
	}

	c.JSON(httpStatus, mapBookResponse(book))
}
