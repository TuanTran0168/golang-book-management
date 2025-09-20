package models

import "gorm.io/gorm"

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	gorm.Model
	Username     string `gorm:"uniqueIndex;type:varchar(100);not null" json:"username"`
	PasswordHash string `gorm:"type:varchar(255);not null" json:"-"`
	Role         Role   `gorm:"type:varchar(20);not null" json:"role"`
	RefreshToken string `gorm:"type:text" json:"-"` // store refresh token
}
