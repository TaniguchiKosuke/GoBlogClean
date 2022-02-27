package domain

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        int `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	Content   string
	AuthorID  string
	Author    User
	ArticleID int
}
