package auth

import (
	"GoBlogClean/pkg/auth/input"
	"GoBlogClean/pkg/auth/output"
)

type UserUsecase interface {
	Signup(*input.SignupRequest) error
	// Login(*User) (*User, error)
	// Logout(*User) (*User, error)
	UpdateUsername(*input.UpdateUsernameRequest) error
	GetUsers() (*output.GetUsersResponse, error)
	Login(*input.LoginRequest) (*output.LoginResponse, error)
}
