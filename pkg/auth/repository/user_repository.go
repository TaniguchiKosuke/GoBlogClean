package repository

import (
	"GoBlogClean/config"
	"GoBlogClean/domain"
	"GoBlogClean/pkg/auth"
)

type userRepository struct {
	dbHandler *config.DBHandler
}

func NewUserRepository(dbHandler *config.DBHandler) auth.UserRepository {
	return &userRepository{dbHandler}
}

func (ur *userRepository) CreateUser(user *domain.User) (*domain.User, error) {
	if err := ur.dbHandler.Conn.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (ur *userRepository) GetUsers() ([]*domain.User, error) {
	var users []*domain.User
	if err := ur.dbHandler.Conn.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (ur *userRepository) GetUserByUsername(username string) (*domain.User, error) {
	var user *domain.User
	if err := ur.dbHandler.Conn.Where("username = ?", username).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
