package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
    ID        uint   `gorm:"primaryKey"`
	Title    string `gorm:"not null"`
	Content  string `gorm:"not null"`
	AuthorID uint
	Author   User `gorm:"foreignKey:AuthorID"`
}
