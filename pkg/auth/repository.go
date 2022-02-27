package auth

import "GoBlogClean/domain"

type UserRepository interface {
	CreateUser(*domain.User) (*domain.User, error)
	// Login(*User) (*User, error)
	// Logout(*User) (*User, error)
	UpdateUsername(*domain.User) (*domain.User, error)
	GetUsers() ([]*domain.User, error)
	GetUserByUsername(username string) (*domain.User, error)
}
