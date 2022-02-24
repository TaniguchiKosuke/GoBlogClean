package repository

import (
	"GoBlogClean/auth"
	"GoBlogClean/config"
	"GoBlogClean/models"
)

type userRepository struct {
	dbHandler *config.DBHandler
}

func NewUserRepository(dbHandler *config.DBHandler) auth.UserRepository {
	return &userRepository{dbHandler}
}

func (ur *userRepository) CreateUser(user *models.User) (*models.User, error) {
	if err := ur.dbHandler.Conn.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (ur *userRepository) GetUsers() ([]*models.User, error) {
	var users []*models.User
	if err := ur.dbHandler.Conn.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}
