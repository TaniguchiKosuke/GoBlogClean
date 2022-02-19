package models

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	Title    string    `json:"title"`
	Content  string    `json:"content"`
	UserID   string    `json:"user_id"`
	User     User      `json:"user"`
	Comments []Comment `json:"comment"`
}
