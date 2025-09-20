package handlers

import (
	"book-management/internal/models"
	"book-management/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthHandler handles authentication endpoints
type AuthHandler struct {
	service services.IAuthService
}

// NewAuthHandler constructor
func NewAuthHandler(service services.IAuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// RegisterRequest defines the payload for registration
type RegisterRequest struct {
	Username string      `json:"username" binding:"required"`
	Password string      `json:"password" binding:"required"`
	Role     models.Role `json:"role" binding:"required,oneof=admin user"`
}

// LoginRequest defines the payload for login
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// TokenResponse defines the response with JWTs
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"` // unix timestamp
}

// Register godoc
// @Summary      Register a new user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        payload  body      RegisterRequest  true  "User registration data"
// @Success      201      {object}  models.User
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Register(req.Username, req.Password, req.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Login godoc
// @Summary      Login user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        payload  body      LoginRequest  true  "Login data"
// @Success      200      {object}  TokenResponse
// @Failure      400      {object}  map[string]string
// @Failure      401      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	access, refresh, exp, err := h.service.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, TokenResponse{
		AccessToken:  access,
		RefreshToken: refresh,
		ExpiresAt:    exp.Unix(),
	})
}

// Refresh godoc
// @Summary      Refresh access token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        refresh_token  body      map[string]string  true  "Refresh token"
// @Success      200            {object}  map[string]string
// @Failure      400            {object}  map[string]string
// @Failure      401            {object}  map[string]string
// @Router       /auth/refresh [post]
func (h *AuthHandler) Refresh(c *gin.Context) {
	var body struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	access, err := h.service.Refresh(body.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": access})
}
