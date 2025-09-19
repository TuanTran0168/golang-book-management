package services

import (
	"book-management/internal/models"
	"book-management/internal/repositories"
	"fmt"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type BookCreateRequest struct {
	Title    string `json:"title" binding:"required"`
	AuthorId uint   `json:"authorId" binding:"required"`
}

// Use pointer to check nil or empty for PATCH update api
type BookUpdateRequest struct {
	Title    *string `json:"title"`
	AuthorId *uint   `json:"authorId"`
}

func mapBook(bookCreateRequest BookCreateRequest) *models.Book {
	return &models.Book{
		Title:    bookCreateRequest.Title,
		AuthorID: bookCreateRequest.AuthorId,
	}
}

type IBookService interface {
	GetBookByID(bookIdStr string) (*models.Book, int, error)
	GetAllBooks(limitStr, offsetStr string) (*[]models.Book, int, error)
	CreateBook(book BookCreateRequest) (*models.Book, error)
	UpdateBook(bookIdStr string, book BookUpdateRequest) (*models.Book, int, error)
}

type BookService struct {
	repo       repositories.IBookRepository
	authorRepo repositories.IAuthorRepository
	db         *gorm.DB
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
func (s *BookService) GetAllBooks(limitStr, offsetStr string) (*[]models.Book, int, error) {
	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	if limit < 1 {
		limit = 10
	}
	if offset < 0 {
		offset = 0
	}

	books, err := s.repo.GetAllBooks(s.db, uint(limit), uint(offset))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return books, http.StatusOK, nil
}

func (s *BookService) CreateBook(book BookCreateRequest) (*models.Book, error) {
	_, err := s.authorRepo.GetAuthorByID(s.db, book.AuthorId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("author with ID [%d] does not exist", book.AuthorId)
		}
		return nil, err
	}

	bookModel := mapBook(book)
	createdBook, err := s.repo.CreateBook(s.db, bookModel)
	if err != nil {
		return nil, err
	}
	s.db.Preload("Author").First(createdBook, createdBook.ID)

	return createdBook, nil
}

func (s *BookService) UpdateBook(bookIdStr string, book BookUpdateRequest) (*models.Book, int, error) {
	bookId, err := strconv.Atoi(bookIdStr)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	bookObj, err := s.repo.GetBookById(s.db, uint(bookId))

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, http.StatusNotFound, fmt.Errorf("book with ID [%d] does not exist", bookId)
		}
		return nil, http.StatusInternalServerError, err
	}

	if book.Title != nil {
		bookObj.Title = *book.Title
	}

	if book.AuthorId != nil {
		author, err := s.authorRepo.GetAuthorByID(s.db, *book.AuthorId)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, http.StatusNotFound, fmt.Errorf("author with ID [%d] does not exist", *book.AuthorId)
			}
			return nil, http.StatusInternalServerError, err
		}
		bookObj.AuthorID = author.ID
		bookObj.Author = *author
	}

	savedBook, err := s.repo.UpdateBook(s.db, bookObj)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return savedBook, http.StatusOK, nil
}

func NewBookService(repo repositories.IBookRepository, authorRepo repositories.IAuthorRepository, db *gorm.DB) IBookService {
	return &BookService{
		repo:       repo,
		authorRepo: authorRepo,
		db:         db,
	}
}
