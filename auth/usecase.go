package auth

import "GoBlogClean/models"

type UserUsecase interface {
	Signup(*models.User) error
	// Login(*User) (*User, error)
	// Logout(*User) (*User, error)
	// UpdateUser(*User) (*User, error)
	GetUsers() ([]*models.User, error)
}
