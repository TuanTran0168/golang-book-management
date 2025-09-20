package repositories

import (
	"book-management/internal/models"

	"gorm.io/gorm"
)

// IUserRepository defines all user-related DB operations
type IUserRepository interface {
	Create(db *gorm.DB, user *models.User) error
	Update(db *gorm.DB, user *models.User) error
	FindByUsername(db *gorm.DB, username string) (*models.User, error)
	FindByID(db *gorm.DB, id uint) (*models.User, error)
	FindByRefreshToken(db *gorm.DB, token string) (*models.User, error)
}

// userRepository is the concrete implementation of IUserRepository
type userRepository struct{}

// NewUserRepository creates a new IUserRepository
func NewUserRepository() IUserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(db *gorm.DB, user *models.User) error {
	return db.Create(user).Error
}

func (r *userRepository) Update(db *gorm.DB, user *models.User) error {
	return db.Save(user).Error
}

func (r *userRepository) FindByUsername(db *gorm.DB, username string) (*models.User, error) {
	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByID(db *gorm.DB, id uint) (*models.User, error) {
	var user models.User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByRefreshToken(db *gorm.DB, token string) (*models.User, error) {
	var user models.User
	if err := db.Where("refresh_token = ?", token).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
