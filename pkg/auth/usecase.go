package auth

import (
	"GoBlogClean/domain"
	"GoBlogClean/pkg/auth/input"
	"GoBlogClean/pkg/auth/output"
)

type UserUsecase interface {
	Signup(*input.UserRequest) error
	// Login(*User) (*User, error)
	// Logout(*User) (*User, error)
	// UpdateUser(*User) (*User, error)
	GetUsers() ([]*domain.User, error)
	Login(*input.LoginRequest) (*output.LoginResponse, error)
}
