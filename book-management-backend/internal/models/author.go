package models

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name  string `gorm:"type:varchar(100);not null"`
	Email string `gorm:"type:varchar(100);uniqueIndex"`
	Books []Book `json:"books"`
}
