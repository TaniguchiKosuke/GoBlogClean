package usecase

import "GoBlogClean/models"

type userUsecase struct {
	userRepository models.UserRepository
}

func NewUserUsecase(userRepository models.UserRepository) models.UserUsecase {
	return &userUsecase{userRepository: userRepository}
}

func (uu *userUsecase) Signup(user *models.User) error {
	_, err := uu.userRepository.Signup(user)
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