package repositories

import (
	"book-management/internal/models"

	"gorm.io/gorm"
)

type IBookRepository interface {
	GetBookById(db *gorm.DB, bookId uint) (*models.Book, error)
	GetAllBooks(db *gorm.DB, limit, offset uint) (*[]models.Book, error)
}

type bookRepository struct{}

func (b *bookRepository) GetBookById(db *gorm.DB, bookId uint) (*models.Book, error) {
	var book models.Book

	result := db.Preload("Author").First(&book, bookId) // result is *gorm.DB

	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func (b *bookRepository) GetAllBooks(db *gorm.DB, limit, offset uint) (*[]models.Book, error) {
	var books []models.Book

	result := db.Limit(int(limit)).
		Offset(int(offset)).
		Preload("Author").
		Find(&books)

	if result.Error != nil {
		return nil, result.Error
	}

	return &books, nil
}

func NewBookRepository() IBookRepository {
	return &bookRepository{}
}
