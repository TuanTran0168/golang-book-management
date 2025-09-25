package models

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	Name  string `json:"name"`
	Books []Book `gorm:"many2many:book_genres"`
}

type GenreCreateRequest struct {
	Name  string `json:"name"`
	Books []uint `json:"books"`
}
