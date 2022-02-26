package usecase

import (
	"GoBlogClean/domain"
	"GoBlogClean/pkg/auth"
	"GoBlogClean/pkg/auth/input"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepository auth.UserRepository
}

func NewUserUsecase(userRepository auth.UserRepository) auth.UserUsecase {
	return &userUsecase{userRepository: userRepository}
}

func (uu *userUsecase) CreateUser(userRequest *input.UserRequest) error {
	hashedBytePassword, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	uuidStr := u.String()

	user := &domain.User{
		ID:       uuidStr,
		Username: userRequest.Username,
		Password: string(hashedBytePassword),
	}

	_, err = uu.userRepository.CreateUser(user)
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
