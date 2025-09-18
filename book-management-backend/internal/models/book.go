package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title    string `json:"title"`
	AuthorID uint   `json:"author_id"`
}
