package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository/mockRepo"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"reflect"

	"github.com/golang/mock/gomock"
	"testing"
)

type eqCreateUserParamsMatcher struct {
	arg      model.UserDataInput
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(model.UserDataInput)
	if !ok {
		return false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(arg.Password), []byte(e.password)); err != nil {
		return false
	}
	e.arg.Password = arg.Password
	return reflect.DeepEqual(e.arg, arg)
}
func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserParams(arg model.UserDataInput, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}

// Todo : fix password hashing []
func TestCreateUser(t *testing.T) {
	// NewController from gomock package returns a new controller for testing
	ctrl := gomock.NewController(t)
	//NewMockUserRepository creates a new mock instance of the user repo
	userRepo := mockRepo.NewMockUserRepository(ctrl)
	//NewMockOrderRepository creates a new mock instance of the order repo
	orderRepo := mockRepo.NewMockOrderRepository(ctrl)

	userUseCase := NewUserUseCase(userRepo, orderRepo)
	mockPassword := "password@123"

	testData := []struct {
		name           string
		input          model.UserDataInput
		buildStub      func(userRepo mockRepo.MockUserRepository)
		expectedOutput model.UserDataOutput
		expectedError  error
	}{
		{
			name: "new user",
			input: model.UserDataInput{
				FName:    "Amal",
				LName:    "Madhu",
				Email:    "amalmadhu@gmail.com",
				Phone:    "7902638843",
				Password: mockPassword,
			},
			buildStub: func(userRepo mockRepo.MockUserRepository) {

				//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(mockPassword), 10)
				//if err != nil {
				//	t.Errorf("failed to hash password for testing : %q", err)
				//}

				userRepo.EXPECT().
					CreateUser(
						gomock.Any(),
						EqCreateUserParams(
							model.UserDataInput{
								FName:    "Amal",
								LName:    "Madhu",
								Email:    "amalmadhu@gmail.com",
								Phone:    "7902638843",
								Password: mockPassword},
							mockPassword),
					).
					Times(1).
					Return(model.UserDataOutput{
						ID:    1,
						FName: "Amal",
						LName: "Madhu",
						Email: "amalmadhu@gmail.com",
						Phone: "7902638843",
					}, nil)
			},
			expectedOutput: model.UserDataOutput{
				ID:    1,
				FName: "Amal",
				LName: "Madhu",
				Email: "amalmadhu@gmail.com",
				Phone: "7902638843",
			},
			expectedError: nil,
		},
		{
			name: "duplicate user",
			input: model.UserDataInput{
				FName:    "Amal",
				LName:    "Madhu",
				Email:    "amalmadhu@gmail.com",
				Phone:    "7902638843",
				Password: mockPassword,
			},
			buildStub: func(userRepo mockRepo.MockUserRepository) {
				userRepo.EXPECT().
					CreateUser(
						gomock.Any(),
						EqCreateUserParams(
							model.UserDataInput{
								FName:    "Amal",
								LName:    "Madhu",
								Email:    "amalmadhu@gmail.com",
								Phone:    "7902638843",
								Password: mockPassword},
							mockPassword),
					).
					Times(1).
					Return(model.UserDataOutput{}, errors.New("user already exists"))
			},
			expectedOutput: model.UserDataOutput{},
			expectedError:  errors.New("user already exists"),
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.buildStub(*userRepo)
			actualUser, err := userUseCase.CreateUser(context.TODO(), tt.input)
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, actualUser, tt.expectedOutput)
		})
	}

}
