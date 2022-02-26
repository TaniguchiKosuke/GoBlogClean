package usecase

import (
	"GoBlogClean/pkg/auth"
	"GoBlogClean/domain"
)

type userUsecase struct {
	userRepository auth.UserRepository
}

func NewUserUsecase(userRepository auth.UserRepository) auth.UserUsecase {
	return &userUsecase{userRepository: userRepository}
}

func (uu *userUsecase) CreateUser(user *domain.User) error {
	_, err := uu.userRepository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (uu *userUsecase) GetUsers() ([]*domain.User, error) {
	users, err := uu.userRepository.GetUsers()
	if err != nil {
		return users, err
	}

	return users, err
}

func (uu *userUsecase) GetUserByUsername(username string) (*domain.User, error) {
	user, err := uu.userRepository.GetUserByUsername(username)
	if err != nil {
		return user, err
	}

	return user, nil
}
