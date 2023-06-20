package repository

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name           string
		input          model.UserDataInput
		expectedOutput model.UserDataOutput
		buildStub      func(mock sqlmock.Sqlmock)
		expectedErr    error
	}{
		{ //test case for creating a new user
			name: "successful create user",
			input: model.UserDataInput{
				FName:    "Sujith",
				LName:    "S",
				Email:    "sujith@gmail.com",
				Phone:    "7902638845",
				Password: "sujith@123",
			},
			expectedOutput: model.UserDataOutput{
				ID:    3,
				FName: "Sujith",
				LName: "S",
				Email: "sujith@gmail.com",
				Phone: "7902638845",
			},
			buildStub: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "f_name", "l_name", "email", "phone", "password"}).
					AddRow(3, "Sujith", "S", "sujith@gmail.com", "7902638845", "sujith@123")

				mock.ExpectQuery("^INSERT INTO users (.+)$").
					WithArgs("Sujith", "S", "sujith@gmail.com", "7902638845", "sujith@123").
					WillReturnRows(rows)

				mock.ExpectExec("^INSERT INTO user_infos (.+)$").
					WithArgs(3).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			expectedErr: nil,
		},
		{ //test case for trying to insert a user with duplicate phone id
			name: "duplicate email",
			input: model.UserDataInput{
				FName:    "Sujith",
				LName:    "S",
				Email:    "sujith@gmail.com",
				Phone:    "7902638845",
				Password: "sujith@123",
			},
			expectedOutput: model.UserDataOutput{},
			buildStub: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("^INSERT INTO users (.+)$").
					WithArgs("Sujith", "S", "sujith@gmail.com", "7902638845", "sujith@123").
					WillReturnError(errors.New("duplicate key value violates unique constraint 'email'"))
			},
			expectedErr: errors.New("duplicate key value violates unique constraint 'email'"),
		},
		{ //test case for inserting a user with duplicate phone number
			name: "duplicate phone",
			input: model.UserDataInput{
				FName:    "Sujith",
				LName:    "S",
				Email:    "sujith@gmail.com",
				Phone:    "7902638845",
				Password: "sujith@123",
			},
			expectedOutput: model.UserDataOutput{},
			buildStub: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("^INSERT INTO users (.+)$").
					WithArgs("Sujith", "S", "sujith@gmail.com", "7902638845", "sujith@123").
					WillReturnError(errors.New("duplicate key value violates unique constraint 'phone'"))
			},
			expectedErr: errors.New("duplicate key value violates unique constraint 'phone'"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//New() method from sqlmock package create sqlmock database connection and a mock to manage expectations.
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			//close the mock db connection after testing.
			defer db.Close()

			//initialize a mock db session
			gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
			if err != nil {
				t.Fatalf("an error '%s' was not expected when initializing a mock db session", err)
			}

			//create NewUserRepository mock by passing a pointer to gorm.DB
			userRepository := NewUserRepository(gormDB)

			// before we actually execute our function, we need to expect required DB actions
			tt.buildStub(mock)

			//call the actual method
			actualOutput, actualErr := userRepository.CreateUser(context.TODO(), tt.input)
			// validate err is nil if we are not expecting to receive an error
			if tt.expectedErr == nil {
				assert.NoError(t, actualErr)
			} else { //validate whether expected and actual errors are same
				assert.Equal(t, tt.expectedErr, actualErr)
			}

			if !reflect.DeepEqual(tt.expectedOutput, actualOutput) {
				t.Errorf("got %v, but want %v", actualOutput, tt.expectedOutput)
			}

			// Check that all expectations were met
			err = mock.ExpectationsWereMet()
			if err != nil {
				t.Errorf("Unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestFindByEmail(t *testing.T) {
	testData := []struct {
		name           string
		email          string
		expectedOutput model.UserLoginVerifier
		buildStub      func(mock sqlmock.Sqlmock)
		expectedErr    error
	}{
		{ //test case for finding a valid user
			name:  "valid email",
			email: "sujith@gmail.com",
			expectedOutput: model.UserLoginVerifier{
				ID:         3,
				FName:      "Sujith",
				LName:      "S",
				Email:      "sujith@gmail.com",
				Phone:      "7902638845",
				Password:   "sujith@123",
				IsBlocked:  false,
				IsVerified: true,
			},
			buildStub: func(mock sqlmock.Sqlmock) {
				columns := []string{"id", "f_name", "l_name", "email", "phone", "password", "is_blocked", "is_verified"}
				mock.ExpectQuery("SELECT users.id, users.f_name, users.l_name, users.email, users.phone, users.password, infos.is_blocked, infos.is_verified FROM users FULL OUTER JOIN user_infos as infos ON users.id = infos.users_id WHERE users.email (.+)$").
					WithArgs("sujith@gmail.com").
					WillReturnRows(sqlmock.NewRows(columns).FromCSVString("3,Sujith,S,sujith@gmail.com,7902638845,sujith@123,f,t"))
			},
			expectedErr: nil,
		},
		{ //test case when user does not exists for the given mail
			name:           "invalid email",
			email:          "nonexistinguser@gmail.com",
			expectedOutput: model.UserLoginVerifier{},
			buildStub: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT users.id, users.f_name, users.l_name, users.email, users.phone, users.password, infos.is_blocked, infos.is_verified FROM users FULL OUTER JOIN user_infos as infos ON users.id = infos.users_id WHERE users.email (.+)$").
					WithArgs("nonexistinguser@gmail.com").
					WillReturnError(errors.New("no rows found"))
			},
			expectedErr: errors.New("no rows found"),
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			// create an sqlmock database connection and mock to manage expectations
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			//	initialize a mock db session
			gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})

			if err != nil {
				t.Fatalf("an error '%s' was not expected when initializing a mock db session", err)
			}

			//create NewUserRepository mock by passing a pointer to gorm.DB
			userRepository := NewUserRepository(gormDB)

			// before we actually execute our function, we need to expect required DB actions
			tt.buildStub(mock)

			//call the actual method
			actualOutput, actualErr := userRepository.FindByEmail(context.TODO(), tt.email)

			// validate err is nil if we are not expecting to receive an error
			if tt.expectedErr == nil {
				assert.NoError(t, actualErr)
			} else { //validate whether expected and actual errors are same
				assert.Equal(t, tt.expectedErr, actualErr)
			}

			if !reflect.DeepEqual(tt.expectedOutput, actualOutput) {
				t.Errorf("got %v, but want %v", actualOutput, tt.expectedOutput)
			}

			// Check that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestFindByPhone(t *testing.T) {
	testData := []struct {
		name           string
		phone          string
		expectedOutput model.UserLoginVerifier
		buildStub      func(mock sqlmock.Sqlmock)
		expectedErr    error
	}{
		{ //test case for finding a valid user
			name:  "valid phone",
			phone: "7902638875",
			expectedOutput: model.UserLoginVerifier{
				ID:         3,
				FName:      "Sujith",
				LName:      "S",
				Email:      "sujith@gmail.com",
				Phone:      "7902638875",
				Password:   "sujith@123",
				IsBlocked:  false,
				IsVerified: true,
			},
			buildStub: func(mock sqlmock.Sqlmock) {
				columns := []string{"id", "f_name", "l_name", "email", "phone", "password", "is_blocked", "is_verified"}
				mock.ExpectQuery("SELECT users.id, users.f_name, users.l_name, users.email, users.phone, users.password, infos.is_blocked, infos.is_verified FROM users FULL OUTER JOIN user_infos as infos ON users.id = infos.users_id WHERE users.phone (.+)$").
					WithArgs("7902638875").
					WillReturnRows(sqlmock.NewRows(columns).FromCSVString("3,Sujith,S,sujith@gmail.com,7902638875,sujith@123,f,t"))
			},
			expectedErr: nil,
		},
		{ //test case when user does not exists for the given mail
			name:           "invalid phone",
			phone:          "7902638875",
			expectedOutput: model.UserLoginVerifier{},
			buildStub: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT users.id, users.f_name, users.l_name, users.email, users.phone, users.password, infos.is_blocked, infos.is_verified FROM users FULL OUTER JOIN user_infos as infos ON users.id = infos.users_id WHERE users.phone (.+)$").
					WithArgs("7902638875").
					WillReturnError(errors.New("no rows found"))
			},
			expectedErr: errors.New("no rows found"),
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			// create an sqlmock database connection and mock to manage expectations
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			//	initialize a mock db session
			gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})

			if err != nil {
				t.Fatalf("an error '%s' was not expected when initializing a mock db session", err)
			}

			//create NewUserRepository mock by passing a pointer to gorm.DB
			userRepository := NewUserRepository(gormDB)

			// before we actually execute our function, we need to expect required DB actions
			tt.buildStub(mock)

			//call the actual method
			actualOutput, actualErr := userRepository.FindByPhone(context.TODO(), tt.phone)

			// validate err is nil if we are not expecting to receive an error
			if tt.expectedErr == nil {
				assert.NoError(t, actualErr)
			} else { //validate whether expected and actual errors are same
				assert.Equal(t, tt.expectedErr, actualErr)
			}

			if !reflect.DeepEqual(tt.expectedOutput, actualOutput) {
				t.Errorf("got %v, but want %v", actualOutput, tt.expectedOutput)
			}

			// Check that all expectations were met
			if err = mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Unfulfilled expectations: %s", err)
			}
		})
	}
}
