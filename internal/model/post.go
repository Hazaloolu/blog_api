package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model

	Title    string `gorm:"not null"`
	Content  string `gorm:"not null"`
	AuthorID uint
	Author   User `gorm:"foreignKey:AuthorID"`
}
