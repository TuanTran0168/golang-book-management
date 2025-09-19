package handlers

import (
	"net/http"
	"strconv"

	"book-management/internal/models"
	"book-management/internal/services"

	"github.com/gin-gonic/gin"
)

type BookSimple struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

type AuthorResponse struct {
	ID        uint         `json:"id"`
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	CreatedAt string       `json:"created_at"`
	UpdatedAt string       `json:"updated_at"`
	Books     []BookSimple `json:"books"`
}

type CreateAuthorRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type UpdateAuthorRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type AuthorHandler struct {
	service services.IAuthorService
}

func NewAuthorHandler(service services.IAuthorService) *AuthorHandler {
	return &AuthorHandler{service: service}
}

// Helper: map GORM Author to AuthorResponse
func mapAuthorResponse(author *models.Author) AuthorResponse {
	books := make([]BookSimple, len(author.Books))
	for i, b := range author.Books {
		books[i] = BookSimple{
			ID:    b.ID,
			Title: b.Title,
		}
	}

	return AuthorResponse{
		ID:        author.ID,
		Name:      author.Name,
		Email:     author.Email,
		CreatedAt: author.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt: author.UpdatedAt.Format("02-01-2006 15:04:05"),
		Books:     books,
	}
}

// GET /authors?limit=10&offset=0
// GetAuthors godoc
// @Summary      Get list of authors
// @Description  Retrieve authors with pagination
// @Tags         authors
// @Accept       json
// @Produce      json
// @Param        limit   query     int  false  "Number of authors per page" default(10)
// @Param        offset  query     int  false  "Starting offset" default(0)
// @Success      200     {object}  map[string]interface{}
// @Failure      500     {object}  map[string]string
// @Router       /authors [get]
func (h *AuthorHandler) GetAuthors(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if limit < 1 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	authors, err := h.service.GetAuthors(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := make([]AuthorResponse, len(authors))
	for i := range authors {
		resp[i] = mapAuthorResponse(&authors[i])
	}

	c.JSON(http.StatusOK, gin.H{
		"limit":  limit,
		"offset": offset,
		"data":   resp,
	})
}

// GET /authors/:id
// GetAuthorByID godoc
// @Summary      Get author details by ID
// @Tags         authors
// @Produce      json
// @Param        id   path      int  true  "Author ID"
// @Success      200  {object}  AuthorResponse
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /authors/{id} [get]
func (h *AuthorHandler) GetAuthorByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid author id"})
		return
	}

	author, err := h.service.GetAuthorByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "author not found"})
		return
	}

	c.JSON(http.StatusOK, mapAuthorResponse(author))
}

// POST /authors
// CreateAuthor godoc
// @Summary      Create a new author
// @Tags         authors
// @Accept       json
// @Produce      json
// @Param        author  body      CreateAuthorRequest  true  "Author data"
// @Success      201     {object}  AuthorResponse
// @Failure      400     {object}  map[string]string
// @Failure      500     {object}  map[string]string
// @Router       /authors [post]
func (h *AuthorHandler) CreateAuthor(c *gin.Context) {
	var req CreateAuthorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author := &models.Author{
		Name:  req.Name,
		Email: req.Email,
	}

	if err := h.service.CreateAuthor(author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, mapAuthorResponse(author))
}

// PUT /authors/:id
// UpdateAuthor godoc
// @Summary      Update an author
// @Tags         authors
// @Accept       json
// @Produce      json
// @Param        id      path      int                 true  "Author ID"
// @Param        author  body      UpdateAuthorRequest true  "Updated author data"
// @Success      200     {object}  AuthorResponse
// @Failure      400     {object}  map[string]string
// @Failure      404     {object}  map[string]string
// @Failure      500     {object}  map[string]string
// @Router       /authors/{id} [put]
func (h *AuthorHandler) UpdateAuthor(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid author id"})
		return
	}

	var req UpdateAuthorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author, err := h.service.GetAuthorByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "author not found"})
		return
	}

	author.Name = req.Name
	author.Email = req.Email

	if err := h.service.UpdateAuthor(author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mapAuthorResponse(author))
}

// DELETE /authors/:id
// DeleteAuthor godoc
// @Summary      Delete an author
// @Tags         authors
// @Param        id   path  int  true  "Author ID"
// @Success      204  "No Content"
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /authors/{id} [delete]
func (h *AuthorHandler) DeleteAuthor(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid author id"})
		return
	}

	author, err := h.service.GetAuthorByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "author not found"})
		return
	}

	if err := h.service.DeleteAuthor(author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
