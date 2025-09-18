package services

import (
	"book-management/internal/models"
	"book-management/internal/repositories"

	"gorm.io/gorm"
)

type IAuthorService interface {
	GetAuthorByID(authorID uint) (*models.Author, error)
	GetAuthors(limit, offset int) ([]models.Author, error) // Pagination
	GetAuthorByEmail(email string) (*models.Author, error)
	CreateAuthor(author *models.Author) error
	UpdateAuthor(author *models.Author) error
	DeleteAuthor(author *models.Author) error
}

type AuthorService struct {
	repo repositories.IAuthorRepository
	db   *gorm.DB
}

func NewAuthorService(repo repositories.IAuthorRepository, db *gorm.DB) IAuthorService {
	return &AuthorService{
		repo: repo,
		db:   db,
	}
}

func (s *AuthorService) GetAuthorByID(authorID uint) (*models.Author, error) {
	return s.repo.GetAuthorByID(s.db, authorID)
}

// Pagination
func (s *AuthorService) GetAuthors(limit, offset int) ([]models.Author, error) {
	return s.repo.GetAuthors(s.db, limit, offset)
}

func (s *AuthorService) GetAuthorByEmail(email string) (*models.Author, error) {
	return s.repo.GetAuthorByEmail(s.db, email)
}

func (s *AuthorService) CreateAuthor(author *models.Author) error {
	return s.repo.CreateAuthor(s.db, author)
}

func (s *AuthorService) UpdateAuthor(author *models.Author) error {
	return s.repo.UpdateAuthor(s.db, author)
}

func (s *AuthorService) DeleteAuthor(author *models.Author) error {
	return s.repo.DeleteAuthor(s.db, author)
}
