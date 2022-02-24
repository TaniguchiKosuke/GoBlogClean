package auth

import "GoBlogClean/models"

type UserRepository interface {
	Signup(*models.User) (*models.User, error)
	// Login(*User) (*User, error)
	// Logout(*User) (*User, error)
	// UpdateUser(*User) (*User, error)
	GetUsers() ([]*models.User, error)
}
