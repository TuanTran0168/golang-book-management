package handlers

import (
	"book-management/internal/models"
	"book-management/internal/services"
	"net/http"

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
	ID        uint                  `json:"id"`
	Title     string                `json:"title"`
	Author    AuthorResponseForBook `json:"author"`
	CreatedAt string                `json:"created_at"`
	UpdatedAt string                `json:"updated_at"`
}

func mapBookResponse(book *models.Book) BookResponse {

	AuthorResponseForBook := AuthorResponseForBook{
		ID:   book.Author.ID,
		Name: book.Author.Name,
	}

	return BookResponse{
		ID:        book.ID,
		Title:     book.Title,
		Author:    AuthorResponseForBook,
		CreatedAt: book.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt: book.UpdatedAt.Format("02-01-2006 15:04:05"),
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
		return
	}

	c.JSON(httpStatus, mapBookResponse(book))
}

// GET /books?limit=10&offset=0
// GetAllBooks godoc
// @Summary      Get all books with pagination
// @Description  Retrieve a paginated list of books, each including its author
// @Tags         books
// @Produce      json
// @Param        limit   query     string  false  "Limit number of books per page"  default(10)
// @Param        offset  query     string  false  "Number of books to skip"         default(0)
// @Success      200     {object}  map[string]interface{}
// @Failure      400     {object}  map[string]string
// @Failure      500     {object}  map[string]string
// @Router       /books [get]
func (h *BookHandler) GetAllBooks(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	books, httpStatus, err := h.service.GetAllBooks(limitStr, offsetStr)

	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
		return
	}

	resp := make([]BookResponse, len(*books))
	for i := range *books {
		resp[i] = mapBookResponse(&(*books)[i])
	}

	c.JSON(httpStatus, gin.H{
		"limit":  limitStr,
		"offset": offsetStr,
		"data":   resp,
	})
}

// POST /books
// CreateBook godoc
// @Summary      Create a new book
// @Description  Create a new book with a title and an existing author
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        book  body      services.BookCreateRequest  true  "Book Create Request"
// @Success      201   {object}  handlers.BookResponse
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /books [post]
func (h *BookHandler) CreateBook(c *gin.Context) {
	var bookCreateRequest services.BookCreateRequest
	if err := c.ShouldBindJSON(&bookCreateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdBook, err := h.service.CreateBook(bookCreateRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mapBookResponse(createdBook))
}

// PATCH /books/:id
// UpdateBook godoc
// @Summary      Update a book partially
// @Description  Update book fields partially by ID (PATCH)
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id    path      string                  true  "Book ID"
// @Param        book  body      services.BookUpdateRequest  true  "Book fields to update"
// @Success      200   {object}  handlers.BookResponse
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /books/{id} [patch]
func (h *BookHandler) UpdateBook(c *gin.Context) {
	var bookUpdateRequest services.BookUpdateRequest
	idStr := c.Param("id")

	if err := c.ShouldBindJSON(&bookUpdateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedBook, httpStatus, err := h.service.UpdateBook(idStr, bookUpdateRequest)

	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
		return
	}

	c.JSON(httpStatus, mapBookResponse(updatedBook))
}
