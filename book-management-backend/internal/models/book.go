package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title    string  `json:"title"`
	AuthorID uint    `json:"author_id"`
	Author   Author  `gorm:"foreignKey:AuthorID" json:"author"`
	Image    string  `json:"image"`
	Genres   []Genre `gorm:"many2many:book_genres"`
}
