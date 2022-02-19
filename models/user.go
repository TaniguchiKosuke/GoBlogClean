package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRepository interface {
	// SignUp(*User) (*User, error)
	// Login(*User) (*User, error)
	// Logout(*User) (*User, error)
	// UpdateUser(*User) (*User, error)
}

type UserUsecase interface {
	// SignUp(*User) (*User, error)
	// Login(*User) (*User, error)
	// Logout(*User) (*User, error)
	// UpdateUser(*User) (*User, error)
}