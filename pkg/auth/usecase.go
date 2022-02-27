package auth

import (
	"GoBlogClean/domain"
	"GoBlogClean/pkg/auth/input"
)

type UserUsecase interface {
	CreateUser(*input.UserRequest) error
	// Login(*User) (*User, error)
	// Logout(*User) (*User, error)
	// UpdateUser(*User) (*User, error)
	GetUsers() ([]*domain.User, error)
	GetUserByUsername(username string) (*domain.User, error)
}
