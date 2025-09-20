package repositories

import (
	"book-management/internal/models"

	"gorm.io/gorm"
)

type IBookRepository interface {
	GetBookById(db *gorm.DB, bookId uint) (*models.Book, error)
	GetAllBooks(db *gorm.DB, limit, offset uint) (*[]models.Book, error)
	CreateBook(db *gorm.DB, book *models.Book) (*models.Book, error)
	UpdateBook(db *gorm.DB, book *models.Book) (*models.Book, error)
	DeleteBook(db *gorm.DB, book *models.Book) error
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

func (b *bookRepository) CreateBook(db *gorm.DB, book *models.Book) (*models.Book, error) {
	result := db.Create(book)
	if result.Error != nil {
		return nil, result.Error
	}
	return book, nil
}

func (b *bookRepository) UpdateBook(db *gorm.DB, book *models.Book) (*models.Book, error) {
	result := db.Save(book)
	if result.Error != nil {
		return nil, result.Error
	}

	return book, nil
}

func (b *bookRepository) DeleteBook(db *gorm.DB, book *models.Book) error {
	return db.Delete(book).Error
}

func NewBookRepository() IBookRepository {
	return &bookRepository{}
}
