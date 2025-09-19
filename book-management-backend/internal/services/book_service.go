package services

import (
	"book-management/internal/models"
	"book-management/internal/repositories"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type IBookService interface {
	GetBookByID(bookIdStr string) (*models.Book, int, error)
	// GetAllBooks(limit, offset uint) (*[]models.Book, int, error)
}

type BookService struct {
	repo repositories.IBookRepository
	db   *gorm.DB
}

func (s *BookService) GetBookByID(bookIdStr string) (*models.Book, int, error) {
	id, err := strconv.ParseUint(bookIdStr, 10, 32)

	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	book, err := s.repo.GetBookById(s.db, uint(id))

	if err != nil {
		return nil, http.StatusNotFound, err
	}

	return book, http.StatusOK, nil
}

func NewBookService(repo repositories.IBookRepository, db *gorm.DB) IBookService {
	return &BookService{
		repo: repo,
		db:   db,
	}
}
