package usecase

import "GoBlogClean/models"

type userUsecase struct {
	userRepository models.UserRepository
}

func NewUserUsecase(userRepository models.UserRepository) models.UserUsecase {
	return &userUsecase{userRepository: userRepository}
}