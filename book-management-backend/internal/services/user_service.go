package services

import (
	"errors"
	"time"

	configs "book-management/configs"
	"book-management/internal/models"
	"book-management/internal/repositories"
	"book-management/pkg/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// IAuthService defines the contract for AuthService
type IAuthService interface {
	Register(username, password string, role models.Role) (*models.User, error)
	Login(username, password string) (accessToken string, refreshToken string, refreshExp time.Time, err error)
	Refresh(refreshToken string) (accessToken string, err error)
}

// AuthService implements IAuthService
type AuthService struct {
	repo repositories.IUserRepository
	cfg  *configs.Config
	db   *gorm.DB
}

// NewAuthService constructor
func NewAuthService(repo repositories.IUserRepository, cfg *configs.Config, db *gorm.DB) IAuthService {
	return &AuthService{
		repo: repo,
		cfg:  cfg,
		db:   db,
	}
}

// Register a new user
func (s *AuthService) Register(username, password string, role models.Role) (*models.User, error) {
	if existing, _ := s.repo.FindByUsername(s.db, username); existing != nil {
		return nil, errors.New("username already taken")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username:     username,
		PasswordHash: string(hash),
		Role:         role,
	}

	if err := s.repo.Create(s.db, user); err != nil {
		return nil, err
	}
	return user, nil
}

// Login authenticates user and returns JWTs
func (s *AuthService) Login(username, password string) (string, string, time.Time, error) {
	user, err := s.repo.FindByUsername(s.db, username)
	if err != nil {
		return "", "", time.Time{}, errors.New("invalid username or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", "", time.Time{}, errors.New("invalid username or password")
	}

	accessToken, err := utils.GenerateAccessToken(user.ID, user.Role, s.cfg)
	if err != nil {
		return "", "", time.Time{}, err
	}

	refreshToken, refreshExp, err := utils.GenerateRefreshToken(user.ID, s.cfg)
	if err != nil {
		return "", "", time.Time{}, err
	}

	user.RefreshToken = refreshToken
	if err := s.repo.Update(s.db, user); err != nil {
		return "", "", time.Time{}, err
	}

	return accessToken, refreshToken, refreshExp, nil
}

// Refresh generates a new access token from refresh token
func (s *AuthService) Refresh(refreshToken string) (string, error) {
	user, err := s.repo.FindByRefreshToken(s.db, refreshToken)
	if err != nil || user == nil {
		return "", errors.New("invalid refresh token")
	}

	// Validate refresh token
	parsed, err := utils.ParseAccessToken(refreshToken, s.cfg)
	if err != nil || time.Now().After(parsed.ExpiresAt.Time) {
		return "", errors.New("expired refresh token")
	}

	accessToken, err := utils.GenerateAccessToken(user.ID, user.Role, s.cfg)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
