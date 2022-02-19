package repository

import (
	"GoBlogClean/config"
	"GoBlogClean/models"
)

type userRepository struct {
	dbHandler *config.DBHandler
}

func NewUserRepository(dbHandler *config.DBHandler) models.UserRepository {
	return &userRepository{*&dbHandler}
}