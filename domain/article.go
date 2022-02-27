package domain

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	Title    string
	Content  string
	AuthorID string
	Author   User
	Comments []Comment `gorm:"foreignKey:AuthorID"`
}
