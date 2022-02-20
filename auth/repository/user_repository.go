package repository

import (
	"GoBlogClean/config"
	"GoBlogClean/models"
)

type userRepository struct {
	dbHandler *config.DBHandler
}

func NewUserRepository(dbHandler *config.DBHandler) models.UserRepository {
	return &userRepository{dbHandler}
}

func (ur *userRepository) Signup(user *models.User) (*models.User, error) {
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