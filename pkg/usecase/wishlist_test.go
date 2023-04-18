package usecase

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository/mockRepo"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddToWishlist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWishlistRepo := mockRepo.NewMockWishlistRepository(ctrl)
	mockWishlist := model.ViewWishlist{ //expected output of ViewWishlist method
		UserID: 1,
		Items:  []model.WishlistItem{},
	}

	userID := 1
	productItemID := 2

	// expect AddToWishlist to be called with userID and productItemID
	mockWishlistRepo.EXPECT().AddToWishlist(gomock.Any(), userID, productItemID).Return(nil)

	// expect ViewWishlist to be called with userID and return mockWishlist
	mockWishlistRepo.EXPECT().ViewWishlist(gomock.Any(), userID).Return(mockWishlist, nil)

	wishlistUC := NewWishlistUsecase(mockWishlistRepo)

	// call AddToWishlist method
	result, err := wishlistUC.AddToWishlist(context.Background(), userID, productItemID)

	// check that AddToWishlist and ViewWishlist were called as expected
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	assert.Equal(t, result, mockWishlist)
}

func TestViewWishlist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWishlistRepo := mockRepo.NewMockWishlistRepository(ctrl)
	mockWishlist := model.ViewWishlist{
		UserID: 1,
		Items:  []model.WishlistItem{},
	}

	userID := 1

	// expect ViewWishlist to be called with userID and return mockWishlist
	mockWishlistRepo.EXPECT().ViewWishlist(gomock.Any(), userID).Return(mockWishlist, nil)

	wishlistUC := NewWishlistUsecase(mockWishlistRepo)

	// call ViewWishlist method
	result, err := wishlistUC.ViewWishlist(context.Background(), userID)

	// check that ViewWishlist was called as expected and result is equal to mockWishlist
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	assert.Equal(t, result, mockWishlist)
}

func TestRemoveFromWishlist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWishlistRepo := mockRepo.NewMockWishlistRepository(ctrl)

	userID := 1
	productItemID := 2

	// expect RemoveFromWishlist to be called with userID and productItemID
	mockWishlistRepo.EXPECT().RemoveFromWishlist(gomock.Any(), userID, productItemID).Return(nil)

	wishlistUC := NewWishlistUsecase(mockWishlistRepo)

	// call RemoveFromWishlist method
	err := wishlistUC.RemoveFromWishlist(context.Background(), userID, productItemID)

	// check that RemoveFromWishlist was called as expected
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestEmptyWishlist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWishlistRepo := mockRepo.NewMockWishlistRepository(ctrl)

	userID := 1

	// expect EmptyWishlist to be called with userID
	mockWishlistRepo.EXPECT().EmptyWishlist(gomock.Any(), userID).Return(nil)

	wishlistUC := NewWishlistUsecase(mockWishlistRepo)

	// call EmptyWishlist method
	err := wishlistUC.EmptyWishlist(context.Background(), userID)

	// check that EmptyWishlist was called as expected
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
