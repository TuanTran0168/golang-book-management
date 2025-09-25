package repositories

import (
	"book-management/internal/models"

	"gorm.io/gorm"
)

type IGenreRepository interface {
	GetGenreByID(db *gorm.DB, id uint) (*models.Genre, error)
	GetAllGenres(db *gorm.DB, limit, offset uint) (*[]models.Genre, error)
	CreateGenre(db *gorm.DB, genre *models.Genre) (*models.Genre, error)
	UpdateGenre(db *gorm.DB, genre *models.Genre) (*models.Genre, error)
	DeleteGenre(db *gorm.DB, genre *models.Genre) error
	GetGenresByIds(db *gorm.DB, ids []uint) ([]models.Genre, error)

	// Manage relationships between Genres and Books
	AddBooksToGenre(db *gorm.DB, genre *models.Genre, books []models.Book) error
	RemoveBooksFromGenre(db *gorm.DB, genre *models.Genre, books []models.Book) error
	ReplaceBooksInGenre(db *gorm.DB, genre *models.Genre, books []models.Book) error
}

type GenreRepository struct{}

func (r *GenreRepository) GetGenreByID(db *gorm.DB, id uint) (*models.Genre, error) {
	var genre models.Genre
	if err := db.Preload("Books").First(&genre, id).Error; err != nil {
		return nil, err
	}
	return &genre, nil
}

func (r *GenreRepository) GetAllGenres(db *gorm.DB, limit, offset uint) (*[]models.Genre, error) {
	var genres []models.Genre
	if err := db.Preload("Books").Limit(int(limit)).Offset(int(offset)).Find(&genres).Error; err != nil {
		return nil, err
	}
	return &genres, nil
}

func (r *GenreRepository) CreateGenre(db *gorm.DB, genre *models.Genre) (*models.Genre, error) {
	if err := db.Create(genre).Error; err != nil {
		return nil, err
	}
	return genre, nil
}

func (r *GenreRepository) UpdateGenre(db *gorm.DB, genre *models.Genre) (*models.Genre, error) {
	if err := db.Save(genre).Error; err != nil {
		return nil, err
	}
	return genre, nil
}

func (r *GenreRepository) DeleteGenre(db *gorm.DB, genre *models.Genre) error {
	return db.Delete(genre).Error
}

func (r *GenreRepository) GetGenresByIds(db *gorm.DB, ids []uint) ([]models.Genre, error) {
	var genres []models.Genre
	if err := db.Where("id IN ?", ids).Find(&genres).Error; err != nil {
		return nil, err
	}
	return genres, nil
}

// =============================
// Manage Books inside a Genre
// =============================

// AddBooksToGenre appends new books to a genre without removing existing ones
func (r *GenreRepository) AddBooksToGenre(db *gorm.DB, genre *models.Genre, books []models.Book) error {
	return db.Model(genre).Association("Books").Append(books)
}

// RemoveBooksFromGenre deletes specific books from a genre
func (r *GenreRepository) RemoveBooksFromGenre(db *gorm.DB, genre *models.Genre, books []models.Book) error {
	return db.Model(genre).Association("Books").Delete(books)
}

// ReplaceBooksInGenre replaces all existing books of a genre with a new set
func (r *GenreRepository) ReplaceBooksInGenre(db *gorm.DB, genre *models.Genre, books []models.Book) error {
	return db.Model(genre).Association("Books").Replace(books)
}

// NewGenreRepository creates a new instance of GenreRepository
func NewGenreRepository() IGenreRepository {
	return &GenreRepository{}
}
