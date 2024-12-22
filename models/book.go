package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Price       string `gorm:"not null"`
	Image       string `gorm:"not null"`
	AuthorID    uint
	UserID      uint
	Genres      []Genre `gorm:"many2many:book_genre;"`
}
