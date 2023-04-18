package repository

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mockRepo: %v", err)
	}
	defer db.Close()
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	c := NewUserRepository(gormDB)

	// Mock the insert query for users table
	rows := sqlmock.NewRows([]string{"id", "f_name", "l_name", "email", "phone", "password"}).
		AddRow(1, "John", "Doe", "johndoe@example.com", "1234567890", "password")
	mock.ExpectQuery("^INSERT INTO users (.+)$").WithArgs("John", "Doe", "johndoe@example.com", "1234567890", "password").WillReturnRows(rows)

	// Mock the insert query for user_infos table
	mock.ExpectExec("^INSERT INTO user_infos (.+)$").WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))

	// Call the function being tested
	input := model.UserDataInput{
		FName:    "John",
		LName:    "Doe",
		Email:    "johndoe@example.com",
		Phone:    "1234567890",
		Password: "password",
	}
	expectedOutput := model.UserDataOutput{
		ID:    1,
		FName: "John",
		LName: "Doe",
		Email: "johndoe@example.com",
		Phone: "1234567890",
	}
	actualOutput, err := c.CreateUser(ctx, input)

	// Check the output and error
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(expectedOutput, actualOutput) {
		t.Errorf("Unexpected output. Expected %v, but got %v", expectedOutput, actualOutput)
	}

	// Check that all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestFindByEmail(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mockRepo: %v", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	c := NewUserRepository(gormDB)

	// Test Case 1: Successful execution of query
	rows := sqlmock.NewRows([]string{"id", "f_name", "l_name", "email", "phone", "password", "is_blocked", "is_verified"}).
		AddRow(1, "John", "Doe", "john.doe@example.com", "1234567890", "password", false, true)

	mock.ExpectQuery("SELECT users.id, users.f_name, users.l_name, users.email, users.phone, users.password, infos.is_blocked, infos.is_verified FROM users as users FULL OUTER JOIN user_infos as infos ON users.id = infos.users_id WHERE users.email = (.+)").
		WithArgs("john.doe@example.com").
		WillReturnRows(rows)

	result, err := c.FindByEmail(ctx, "john.doe@example.com")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if result.ID != 1 || result.FName != "John" || result.LName != "Doe" || result.Email != "john.doe@example.com" ||
		result.Phone != "1234567890" || result.Password != "password" || result.IsBlocked || !result.IsVerified {
		t.Errorf("unexpected result: %v", result)
	}

	// Test Case 2: Query returning no rows
	rows = sqlmock.NewRows([]string{})

	mock.ExpectQuery("SELECT users.id, users.f_name, users.l_name, users.email, users.phone, users.password, infos.is_blocked, infos.is_verified FROM users as users FULL OUTER JOIN user_infos as infos ON users.id = infos.users_id WHERE users.email = (.+)").
		WithArgs("invalid@example.com").
		WillReturnRows(rows)

	result, err = c.FindByEmail(ctx, "invalid@example.com")
	if result.ID != 0 {
		t.Errorf("unexpected result: %v", result)
	}
}
