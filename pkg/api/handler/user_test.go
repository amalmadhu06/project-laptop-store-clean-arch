package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/usecase/mockUsecase"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/response"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateUser(t *testing.T) {
	//NewController creates a mock controller for testing
	ctrl := gomock.NewController(t)
	//NewMockUserUseCase creates a new mock instance of the user use case
	userUseCase := mockUsecase.NewMockUserUseCase(ctrl)
	//NewUserHandler create a new user user handler
	userHandler := NewUserHandler(userUseCase)

	//testData is a slice of anonymous structs which are initialized
	testData := []struct {
		name             string // name of the test case
		userData         model.UserDataInput
		buildStub        func(userUsecase mockUsecase.MockUserUseCase) // function to generate the mock user use case
		expectedCode     int                                           // expected status code
		expectedResponse response.Response                             // expected response for the request
		expectedError    error                                         // expected error in the response
	}{
		{ // test case for checking user sign up for a non-registered users
			name: "create non existing user",
			userData: model.UserDataInput{
				FName:    "Amal",
				LName:    "Madhu",
				Email:    "amalmadhu@gmail.com",
				Phone:    "7902631234",
				Password: "password@123",
			},
			buildStub: func(userUsecase mockUsecase.MockUserUseCase) {
				userUsecase.EXPECT(). //setting the expected behaviour of the usecase method
							CreateUser(gomock.Any(), model.UserDataInput{ //CreateUser usecase receives two arguments, one is context.so we can use gomock.Any(). next one is user signup information
						FName:    "Amal",
						LName:    "Madhu",
						Email:    "amalmadhu@gmail.com",
						Phone:    "7902631234",
						Password: "password@123",
					}).
					Times(1).                    //how many times the CreateUser usecase should be called
					Return(model.UserDataOutput{ //what should CreateUser usecase return. Here it should return user info and nil (error)
						ID:    1,
						FName: "Amal",
						LName: "Madhu",
						Email: "amalmadhu@gmail.com",
						Phone: "7902631234",
					}, nil)
			},
			expectedCode: 201,
			expectedResponse: response.Response{
				StatusCode: 201,
				Message:    "User created successfully",
				Data: model.UserDataOutput{
					ID:    1,
					FName: "Amal",
					LName: "Madhu",
					Email: "amalmadhu@gmail.com",
					Phone: "7902631234",
				},
				Errors: nil,
			},
		},
	}

	// looping through the test cases and running them individually.
	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			//passing mock use case to buildStub function which is
			tt.buildStub(*userUseCase)
			// gin.Default will create a new engine instance with logger middleware by default
			engine := gin.Default()
			//NewRecorder from httptest package will create a recorder which records the response
			recorder := httptest.NewRecorder()
			//create new route for testing
			engine.POST("/signup", userHandler.CreateUser)
			//body will hold a slice of bytes. It is used for Marshaling json data and passing to the request body
			var body []byte
			//marshaling userdata in testcase to body
			body, err := json.Marshal(tt.userData)
			//validating no error occurred while marshaling userdata to body
			assert.NoError(t, err)
			//url for the test
			url := "/signup"
			// req is a pointer to http.Request . With httptest.NewRequest we are mentioning the http method, endpoint and body
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			//recorder will record the response, and req is the mock req that we created with test data
			engine.ServeHTTP(recorder, req)

			var actual response.Response
			err = json.Unmarshal(recorder.Body.Bytes(), &actual)

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedCode, recorder.Code)
			assert.Equal(t, tt.expectedResponse.Message, actual.Message)
			// Todo : compare actual response body and actual response body
		})
	}

}

func TestFindUserByID(t *testing.T) {

	ctrl := gomock.NewController(t)
	userUseCase := mockUsecase.NewMockUserUseCase(ctrl)
	userHandler := NewUserHandler(userUseCase)

	testData := []struct {
		name             string
		user             domain.Users
		pathParam        string
		buildStub        func(userUseCase mockUsecase.MockUserUseCase)
		expectedCode     int
		expectedResponse response.Response
		expectedError    error
	}{
		{
			name: "existing user",
			user: domain.Users{
				ID:       1,
				FName:    "Amal",
				LName:    "Madhu",
				Email:    "amalmadhu@gmail.com",
				Phone:    "7901212345",
				Password: "amal@123",
			},
			pathParam: "1",
			buildStub: func(userUsecase mockUsecase.MockUserUseCase) {
				userUsecase.EXPECT().
					FindUserByID(gomock.Any(), 1).
					Times(1).
					Return(domain.Users{
						ID:       1,
						FName:    "Amal",
						LName:    "Madhu",
						Email:    "amalmadhu@gmail.com",
						Phone:    "7901212345",
						Password: "amal@123",
					}, nil)
			},
			expectedCode: 200,
			expectedResponse: response.Response{
				StatusCode: 200,
				Message:    "Successfully fetched user details",
				Data: domain.Users{
					ID:       1,
					FName:    "Amal",
					LName:    "Madhu",
					Email:    "amalmadhu@gmail.com",
					Phone:    "7901212345",
					Password: "amal@123",
				},
				Errors: "",
			},
			expectedError: nil,
		},
		{
			name:      "no user found",
			user:      domain.Users{},
			pathParam: "2",
			buildStub: func(userUsecase mockUsecase.MockUserUseCase) {
				userUsecase.EXPECT().
					FindUserByID(gomock.Any(), 2).
					Times(1).
					Return(domain.Users{}, errors.New("no such user"))
			},
			expectedCode: 500,
			expectedResponse: response.Response{
				StatusCode: 500,
				Message:    "failed fetch user",
				Data:       nil,
				Errors:     "failed to fetch user",
			},
			expectedError: errors.New("failed to fetch user"),
		}, {
			name:      "invalid param",
			user:      domain.Users{},
			pathParam: "invalid",
			buildStub: func(userUsecase mockUsecase.MockUserUseCase) {
				userUsecase.EXPECT().
					FindUserByID(nil, nil).
					Times(0).
					Return(domain.Users{}, nil)
			},
			expectedCode: 422,
			expectedResponse: response.Response{
				StatusCode: 422,
				Message:    "failed to parse user id",
				Data:       nil,
				Errors:     "failed to parse user id",
			},
			expectedError: errors.New("failed to parse user id"),
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.buildStub(*userUseCase)

			engine := gin.Default()
			recorder := httptest.NewRecorder()

			engine.GET("/admin/users/:id", userHandler.FindUserByID)

			var body []byte
			body, err := json.Marshal(tt.user)
			assert.NoError(t, err)
			url := fmt.Sprintf("/admin/users/%s", tt.pathParam)
			req := httptest.NewRequest(http.MethodGet, url, bytes.NewBuffer(body))
			engine.ServeHTTP(recorder, req)

			var actual response.Response
			err = json.Unmarshal(recorder.Body.Bytes(), &actual)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedCode, recorder.Code)
			assert.Equal(t, tt.expectedResponse.Message, actual.Message)

		})
	}
}
