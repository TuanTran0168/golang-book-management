package repositories

import (
	"book-management/internal/models"

	"gorm.io/gorm"
)

type IAuthorRepository interface {
	GetAuthorByID(db *gorm.DB, authorID uint) (*models.Author, error)
	GetAuthors(db *gorm.DB, limit, offset int) ([]models.Author, error) // Pagination
	GetAuthorByEmail(db *gorm.DB, email string) (*models.Author, error)
	CreateAuthor(db *gorm.DB, author *models.Author) error
	UpdateAuthor(db *gorm.DB, author *models.Author) error
	DeleteAuthor(db *gorm.DB, author *models.Author) error
}

type authorRepository struct{}

func (a *authorRepository) GetAuthorByID(db *gorm.DB, authorID uint) (*models.Author, error) {
	var author models.Author
	if err := db.Preload("Books").First(&author, authorID).Error; err != nil {
		return nil, err
	}
	return &author, nil
}

// Pagination
func (a *authorRepository) GetAuthors(db *gorm.DB, limit, offset int) ([]models.Author, error) {
	var authors []models.Author
	if err := db.Preload("Books").
		Limit(limit).
		Offset(offset).
		Find(&authors).Error; err != nil {
		return nil, err
	}
	return authors, nil
}

func (a *authorRepository) GetAuthorByEmail(db *gorm.DB, email string) (*models.Author, error) {
	var author models.Author
	if err := db.Preload("Books").Where("email = ?", email).First(&author).Error; err != nil {
		return nil, err
	}
	return &author, nil
}

func (a *authorRepository) CreateAuthor(db *gorm.DB, author *models.Author) error {
	return db.Create(author).Error
}

func (a *authorRepository) UpdateAuthor(db *gorm.DB, author *models.Author) error {
	return db.Save(author).Error
}

func (a *authorRepository) DeleteAuthor(db *gorm.DB, author *models.Author) error {
	return db.Delete(author).Error
}

func NewAuthorRepository() IAuthorRepository {
	return &authorRepository{}
}
