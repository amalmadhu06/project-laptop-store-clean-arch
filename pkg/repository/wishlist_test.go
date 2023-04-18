package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

var ErrProductAlreadyInWishlist = errors.New("product already in wishlist")

func TestRemoveFromWishlist(t *testing.T) {
	// create a new mockRepo database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mockRepo database connection: %s", err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	// create a new repository instance using the mockRepo database connection
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to create GORM DB instance: %v", err)
	}
	// set up mockRepo expectations
	mock.ExpectExec("DELETE FROM wishlist_items WHERE wishlist_id IN").WithArgs(1, 2).WillReturnResult(sqlmock.NewResult(0, 1))

	// create a new repository instance using the mockRepo database connection
	repo := &wishlistDatabase{DB: gormDB}

	// call the function being tested
	err = repo.RemoveFromWishlist(context.Background(), 1, 2)
	if err != nil {
		t.Fatalf("Unexpected error while removing from wishlist: %s", err)
	}

	// verify that all of the expected SQL queries were executed
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled SQL mockRepo expectations: %s", err)
	}
}

func TestEmptyWishlist(t *testing.T) {
	// create a new mockRepo database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mockRepo database connection: %s", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to create GORM DB instance: %v", err)
	}

	// create a new repository instance using the mockRepo database connection
	repo := &wishlistDatabase{DB: gormDB}

	// set up mockRepo expectations
	mock.ExpectExec("DELETE FROM wishlist_items WHERE wishlist_id IN").WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))

	// call the function being tested
	err = repo.EmptyWishlist(context.Background(), 1)
	if err != nil {
		t.Fatalf("Unexpected error while emptying wishlist: %s", err)
	}

	// verify that all the expected SQL queries were executed
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled SQL mockRepo expectations: %s", err)
	}
}
