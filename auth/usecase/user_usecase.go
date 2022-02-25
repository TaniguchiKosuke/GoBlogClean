package usecase

import (
	"GoBlogClean/auth"
	"GoBlogClean/models"
)

type userUsecase struct {
	userRepository auth.UserRepository
}

func NewUserUsecase(userRepository auth.UserRepository) auth.UserUsecase {
	return &userUsecase{userRepository: userRepository}
}

func (uu *userUsecase) CreateUser(user *models.User) error {
	_, err := uu.userRepository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (uu *userUsecase) GetUsers() ([]*models.User, error) {
	users, err := uu.userRepository.GetUsers()
	if err != nil {
		return users, err
	}

	return users, err
}

func (uu *userUsecase) GetUserByUsername(username string) (*models.User, error) {
	user, err := uu.userRepository.GetUserByUsername(username)
	if err != nil {
		return user, err
	}

	return user, nil
}
