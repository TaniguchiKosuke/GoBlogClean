package usecase

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"GoBlogClean/domain"
	"GoBlogClean/pkg/auth"
	"GoBlogClean/pkg/auth/input"
	"GoBlogClean/pkg/auth/output"
	"GoBlogClean/pkg/util"
)

type userUsecase struct {
	userRepository auth.UserRepository
}

func NewUserUsecase(userRepository auth.UserRepository) auth.UserUsecase {
	return &userUsecase{userRepository: userRepository}
}

func (uu *userUsecase) Signup(userRequest *input.UserRequest) error {
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

func (uu *userUsecase) Login(loginRequest *input.LoginRequest) (*output.LoginResponse, error) {
	user, err := uu.userRepository.GetUserByUsername(loginRequest.Username)
	if err != nil {
		return &output.LoginResponse{}, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		return &output.LoginResponse{}, err
	}

	jwtToken, err := util.CreateJWTToken(user.ID, user.Username)
	if err != nil {
		return &output.LoginResponse{}, err
	}

	loginResponse := &output.LoginResponse{
		Token: jwtToken,
	}

	return loginResponse, nil
}
