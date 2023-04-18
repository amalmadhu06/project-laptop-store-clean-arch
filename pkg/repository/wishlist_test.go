package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

var ErrProductAlreadyInWishlist = errors.New("product already in wishlist")

func TestViewWishlist(t *testing.T) {
	// create a new sqlmock database connection
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
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

	// create a new repository instance using the mockRepo database
	repo := NewWishlistRepository(gormDB)

	// prepare test data
	userID := 1
	wishlistID := 2
	productItemID := 3
	productItemName := "Test Item"
	productItemModel := "Test Model"
	productItemBrand := "Test Brand"
	productItemPrice := 1000
	productItemImage := "test-item-image.jpg"
	mockWishlist := domain.Wishlist{
		ID:     uint(wishlistID),
		UserID: userID,
	}
	fmt.Println(mockWishlist)
	mockWishlistItems := []model.WishlistItem{
		{
			ProductItemID: productItemID,
			Name:          productItemName,
			Model:         productItemModel,
			Brand:         productItemBrand,
			Price:         float64(productItemPrice),
			Image:         "",
		},
	}
	expectedResult := model.ViewWishlist{
		ID:     int(mockWishlist.ID),
		UserID: mockWishlist.UserID,
		Items:  mockWishlistItems,
	}

	// define mockRepo expectations
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT").WithArgs(userID).WillReturnRows(
		sqlmock.NewRows([]string{"id", "user_id"}).AddRow(wishlistID, userID))
	mock.ExpectQuery("SELECT").WithArgs(wishlistID).WillReturnRows(
		sqlmock.NewRows([]string{"product_item_id", "name", "model", "brand", "price", "product_item_image"}).
			AddRow(productItemID, productItemName, productItemModel, productItemBrand, productItemPrice, productItemImage))
	mock.ExpectCommit()

	// invoke the function being tested
	actualResult, err := repo.ViewWishlist(context.Background(), userID)
	require.NoError(t, err)
	require.Equal(t, expectedResult, actualResult)

	// assert that all expectations were met
	//require.NoError(t, mockRepo.ExpectationsWereMet())
}

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
