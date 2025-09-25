package services

import (
	"book-management/internal/models"
	"book-management/internal/repositories"
	"fmt"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type GenreCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

type GenreUpdateRequest struct {
	Name *string `json:"name"`
}

// ManageBooksRequest is used for adding/removing/replacing books in a genre
type ManageBooksRequest struct {
	BookIDs []uint `json:"book_ids" binding:"required"`
}

type IGenreService interface {
	GetGenreByID(genreIdStr string) (*models.Genre, int, error)
	GetAllGenres(limitStr, offsetStr string) (*[]models.Genre, int, error)
	CreateGenre(req GenreCreateRequest) (*models.Genre, error)
	UpdateGenre(genreIdStr string, req GenreUpdateRequest) (*models.Genre, int, error)
	DeleteGenre(genreIdStr string) (int, error)

	// Relationship management
	AddBooksToGenre(genreIdStr string, req ManageBooksRequest) (*models.Genre, int, error)
	RemoveBooksFromGenre(genreIdStr string, req ManageBooksRequest) (*models.Genre, int, error)
	ReplaceBooksInGenre(genreIdStr string, req ManageBooksRequest) (*models.Genre, int, error)
}

type GenreService struct {
	repo repositories.IGenreRepository
	db   *gorm.DB
}

func (s *GenreService) GetGenreByID(genreIdStr string) (*models.Genre, int, error) {
	id, err := strconv.Atoi(genreIdStr)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	genre, err := s.repo.GetGenreByID(s.db, uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, http.StatusNotFound, fmt.Errorf("genre with ID [%d] does not exist", id)
		}
		return nil, http.StatusInternalServerError, err
	}
	return genre, http.StatusOK, nil
}

func (s *GenreService) GetAllGenres(limitStr, offsetStr string) (*[]models.Genre, int, error) {
	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	if limit < 1 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	genres, err := s.repo.GetAllGenres(s.db, uint(limit), uint(offset))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return genres, http.StatusOK, nil
}

func (s *GenreService) CreateGenre(req GenreCreateRequest) (*models.Genre, error) {
	genre := &models.Genre{Name: req.Name}
	return s.repo.CreateGenre(s.db, genre)
}

func (s *GenreService) UpdateGenre(genreIdStr string, req GenreUpdateRequest) (*models.Genre, int, error) {
	id, err := strconv.Atoi(genreIdStr)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	genre, err := s.repo.GetGenreByID(s.db, uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, http.StatusNotFound, fmt.Errorf("genre with ID [%d] does not exist", id)
		}
		return nil, http.StatusInternalServerError, err
	}

	if req.Name != nil {
		genre.Name = *req.Name
	}

	updated, err := s.repo.UpdateGenre(s.db, genre)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return updated, http.StatusOK, nil
}

func (s *GenreService) DeleteGenre(genreIdStr string) (int, error) {
	id, err := strconv.Atoi(genreIdStr)
	if err != nil {
		return http.StatusBadRequest, err
	}

	genre, err := s.repo.GetGenreByID(s.db, uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return http.StatusNotFound, fmt.Errorf("genre with ID [%d] does not exist", id)
		}
		return http.StatusInternalServerError, err
	}

	if err := s.repo.DeleteGenre(s.db, genre); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusNoContent, nil
}

// =============================
// Relationship Management
// =============================

func (s *GenreService) AddBooksToGenre(genreIdStr string, req ManageBooksRequest) (*models.Genre, int, error) {
	id, err := strconv.Atoi(genreIdStr)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	genre, err := s.repo.GetGenreByID(s.db, uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, http.StatusNotFound, fmt.Errorf("genre with ID [%d] does not exist", id)
		}
		return nil, http.StatusInternalServerError, err
	}

	var books []models.Book
	if err := s.db.Where("id IN ?", req.BookIDs).Find(&books).Error; err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err := s.repo.AddBooksToGenre(s.db, genre, books); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return genre, http.StatusOK, nil
}

func (s *GenreService) RemoveBooksFromGenre(genreIdStr string, req ManageBooksRequest) (*models.Genre, int, error) {
	id, err := strconv.Atoi(genreIdStr)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	genre, err := s.repo.GetGenreByID(s.db, uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, http.StatusNotFound, fmt.Errorf("genre with ID [%d] does not exist", id)
		}
		return nil, http.StatusInternalServerError, err
	}

	var books []models.Book
	if err := s.db.Where("id IN ?", req.BookIDs).Find(&books).Error; err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err := s.repo.RemoveBooksFromGenre(s.db, genre, books); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return genre, http.StatusOK, nil
}

func (s *GenreService) ReplaceBooksInGenre(genreIdStr string, req ManageBooksRequest) (*models.Genre, int, error) {
	id, err := strconv.Atoi(genreIdStr)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	genre, err := s.repo.GetGenreByID(s.db, uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, http.StatusNotFound, fmt.Errorf("genre with ID [%d] does not exist", id)
		}
		return nil, http.StatusInternalServerError, err
	}

	var books []models.Book
	if err := s.db.Where("id IN ?", req.BookIDs).Find(&books).Error; err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if err := s.repo.ReplaceBooksInGenre(s.db, genre, books); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return genre, http.StatusOK, nil
}

func NewGenreService(repo repositories.IGenreRepository, db *gorm.DB) IGenreService {
	return &GenreService{repo: repo, db: db}
}
