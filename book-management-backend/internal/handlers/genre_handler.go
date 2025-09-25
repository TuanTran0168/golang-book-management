package handlers

import (
	"book-management/internal/models"
	"book-management/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GenreHandler struct {
	service services.IGenreService
}

func NewGenreHandler(service services.IGenreService) *GenreHandler {
	return &GenreHandler{service: service}
}

type BookResponseForGenre struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type GenreResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Books     []BookResponseForGenre
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func mapGenreResponse(genre *models.Genre) GenreResponse {
	books := make([]BookResponseForGenre, len(genre.Books))
	for i, b := range genre.Books {
		books[i] = BookResponseForGenre{
			ID:   b.ID,
			Name: b.Title,
		}
	}

	return GenreResponse{
		ID:        genre.ID,
		Name:      genre.Name,
		Books:     books,
		CreatedAt: genre.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt: genre.UpdatedAt.Format("02-01-2006 15:04:05"),
	}
}

// GET /genres/:id
// GetGenreByID godoc
// @Summary      Get genre details by ID
// @Description  Get a single genre by genre ID
// @Tags         genres
// @Produce      json
// @Param        id   path      string  true  "Genre ID"
// @Success      200  {object}  GenreResponse
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /genres/{id} [get]
// @Security BearerAuth
func (h *GenreHandler) GetGenreByID(c *gin.Context) {
	idStr := c.Param("id")

	genre, httpStatus, err := h.service.GetGenreByID(idStr)
	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
		return
	}
	c.JSON(httpStatus, mapGenreResponse(genre))
}

// GET /genres?limit=10&offset=0
// GetAllGenres godoc
// @Summary      Get all genres with pagination
// @Description  Retrieve a paginated list of genres
// @Tags         genres
// @Produce      json
// @Param        limit   query     string  false  "Limit number of genres per page"  default(10)
// @Param        offset  query     string  false  "Number of genres to skip"         default(0)
// @Success      200     {object}  map[string]interface{}
// @Failure      500     {object}  map[string]string
// @Router       /genres [get]
// @Security BearerAuth
func (h *GenreHandler) GetAllGenres(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	genres, httpStatus, err := h.service.GetAllGenres(limitStr, offsetStr)
	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
		return
	}

	resp := make([]GenreResponse, len(*genres))
	for i := range *genres {
		resp[i] = mapGenreResponse(&(*genres)[i])
	}

	c.JSON(httpStatus, gin.H{
		"limit":  limitStr,
		"offset": offsetStr,
		"data":   resp,
	})
}

// POST /genres
// CreateGenre godoc
// @Summary      Create a new genre
// @Description  Create a new genre with a name
// @Tags         genres
// @Accept       json
// @Produce      json
// @Param        genre  body      services.GenreCreateRequest  true  "Genre data"
// @Success      201 {object} GenreResponse
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /genres [post]
// @Security BearerAuth
func (h *GenreHandler) CreateGenre(c *gin.Context) {
	var req services.GenreCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdGenre, err := h.service.CreateGenre(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mapGenreResponse(createdGenre))
}

// PATCH /genres/:id
// UpdateGenre godoc
// @Summary      Update a genre partially
// @Description  Update genre fields partially by ID (PATCH)
// @Tags         genres
// @Accept       json
// @Produce      json
// @Param        id     path      string                    true  "Genre ID"
// @Param        genre  body      services.GenreUpdateRequest  true  "Genre fields to update"
// @Success      200 {object} GenreResponse
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /genres/{id} [patch]
// @Security BearerAuth
func (h *GenreHandler) UpdateGenre(c *gin.Context) {
	idStr := c.Param("id")
	var req services.GenreUpdateRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedGenre, httpStatus, err := h.service.UpdateGenre(idStr, req)
	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
		return
	}

	c.JSON(httpStatus, mapGenreResponse(updatedGenre))
}

// DELETE /genres/:id
// DeleteGenre godoc
// @Summary      Delete a genre
// @Description  Delete a genre by its ID
// @Tags         genres
// @Produce      json
// @Param        id   path      string  true  "Genre ID"
// @Success      204  "No Content"
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /genres/{id} [delete]
// @Security BearerAuth
func (h *GenreHandler) DeleteGenre(c *gin.Context) {
	idStr := c.Param("id")
	httpStatus, err := h.service.DeleteGenre(idStr)
	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
		return
	}
	c.JSON(httpStatus, nil)
}

// =============================
// Relationship Management
// =============================

// POST /genres/:id/books
// AddBooksToGenre godoc
// @Summary      Add books to a genre
// @Description  Assign multiple books to a genre by providing their IDs
// @Tags         genres
// @Accept       json
// @Produce      json
// @Param        id    path      string                        true  "Genre ID"
// @Param        body  body      services.ManageBooksRequest    true  "List of book IDs"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /genres/{id}/books [post]
// @Security BearerAuth
func (h *GenreHandler) AddBooksToGenre(c *gin.Context) {
	idStr := c.Param("id")
	var req services.ManageBooksRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	genre, httpStatus, err := h.service.AddBooksToGenre(idStr, req)
	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
		return
	}

	c.JSON(httpStatus, gin.H{
		"message": "Books added successfully",
		"genre":   mapGenreResponse(genre),
	})
}

// DELETE /genres/:id/books
// RemoveBooksFromGenre godoc
// @Summary      Remove books from a genre
// @Description  Detach multiple books from a genre by providing their IDs
// @Tags         genres
// @Accept       json
// @Produce      json
// @Param        id    path      string                        true  "Genre ID"
// @Param        body  body      services.ManageBooksRequest    true  "List of book IDs"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /genres/{id}/books [delete]
// @Security BearerAuth
func (h *GenreHandler) RemoveBooksFromGenre(c *gin.Context) {
	idStr := c.Param("id")
	var req services.ManageBooksRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	genre, httpStatus, err := h.service.RemoveBooksFromGenre(idStr, req)
	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
		return
	}

	c.JSON(httpStatus, gin.H{
		"message": "Books removed successfully",
		"genre":   mapGenreResponse(genre),
	})
}

// PUT /genres/:id/books
// ReplaceBooksInGenre godoc
// @Summary      Replace books in a genre
// @Description  Replace the current list of books with a new set
// @Tags         genres
// @Accept       json
// @Produce      json
// @Param        id    path      string                        true  "Genre ID"
// @Param        body  body      services.ManageBooksRequest    true  "List of book IDs"
// @Success      200   {object}  map[string]interface{}
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /genres/{id}/books [put]
// @Security BearerAuth
func (h *GenreHandler) ReplaceBooksInGenre(c *gin.Context) {
	idStr := c.Param("id")
	var req services.ManageBooksRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	genre, httpStatus, err := h.service.ReplaceBooksInGenre(idStr, req)
	if err != nil {
		c.JSON(httpStatus, gin.H{"error": err.Error()})
		return
	}

	c.JSON(httpStatus, gin.H{
		"message": "Books replaced successfully",
		"genre":   mapGenreResponse(genre),
	})
}
