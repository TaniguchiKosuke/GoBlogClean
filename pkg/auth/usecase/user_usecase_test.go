package usecase

import (
	"GoBlogClean/domain"
	mock_auth "GoBlogClean/pkg/auth/repository/mock"
	"log"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestSignup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testCase := []struct{
		arg *domain.User
		expected *domain.User
	}{
		{
			arg: &domain.User{
				ID: "1",
				Username: "test1",
				Password: "0628",
			},
			expected: &domain.User{
				ID: "1",
				Username: "test1",
				Password: "0628",
			},
		},
		{
			arg: &domain.User{
				ID: "2",
				Username: "test2",
				Password: "0628",
			},
			expected: &domain.User{
				ID: "2",
				Username: "test2",
				Password: "0628",
			},
		},
	}

	for _, tt := range testCase {
		mockUserRepository := mock_auth.NewMockUserRepository(ctrl)
		mockUserRepository.EXPECT().CreateUser(tt.arg).Return(tt.arg, nil)

		mockUserUsecase := &userUsecase{userRepository: mockUserRepository}
		resultUser, err := mockUserUsecase.userRepository.CreateUser(tt.arg)
		if err != nil {
			log.Printf("failed to Signup in test=%s", err.Error())
		}

		if resultUser != tt.expected {
			log.Printf("the result is not what is expected=%v", resultUser)
		}
	}

	log.Println("all test cases are succeed")
}