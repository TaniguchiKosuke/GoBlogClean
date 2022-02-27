package usecase

import (
	"time"

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

func (uu *userUsecase) Signup(signupRequest *input.SignupRequest) error {
	hashedBytePassword, err := bcrypt.GenerateFromPassword([]byte(signupRequest.Password), bcrypt.DefaultCost)
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
		Username: signupRequest.Username,
		Password: string(hashedBytePassword),
	}

	_, err = uu.userRepository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (uu *userUsecase) GetUsers() (*output.GetUsersResponse, error) {
	users, err := uu.userRepository.GetUsers()
	if err != nil {
		return &output.GetUsersResponse{}, err
	}

	usersResponse := make([]*output.UserResponse, 0, len(users))

	for _, user := range users {
		userResponse := &output.UserResponse{
			ID:        user.ID,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
			DeletedAt: user.DeletedAt.Time.Format(time.RFC3339),
			Username:  user.Username,
		}

		usersResponse = append(usersResponse, userResponse)
	}

	getUsersResponse := &output.GetUsersResponse{
		Users: usersResponse,
	}

	return getUsersResponse, err
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

func (uu *userUsecase) UpdateUsername(userRequest *input.UpdateUsernameRequest) error {
	user := &domain.User{
		ID:       userRequest.ID,
		Username: userRequest.Username,
	}

	_, err := uu.userRepository.UpdateUsername(user)
	if err != nil {
		return err
	}

	return nil
}
