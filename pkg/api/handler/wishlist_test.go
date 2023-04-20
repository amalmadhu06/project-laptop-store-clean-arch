package handler

import (
	"encoding/json"
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

func TestAddToWishlist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wishlistUseCase := mockUsecase.NewMockWishlistUseCase(ctrl)

	wishlistHandler := NewWishlistHandler(wishlistUseCase)

	r := gin.Default()
	r.POST("/wishlist/:id", wishlistHandler.AddToWishlist)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/wishlist/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Call the AddToWishlist handler function
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Set("userID", 1)
	wishlistUseCase.EXPECT().AddToWishlist(gomock.Any(), 1, 1).Return(
		model.ViewWishlist{
			ID:     1,
			UserID: 1,
			Items:  []model.WishlistItem{}}, nil)
	wishlistHandler.AddToWishlist(c)

	// Check that the response status code is correct
	assert.Equal(t, http.StatusCreated, w.Code)

	// Check that the response body is correct
	var respData gin.H
	err = json.Unmarshal(w.Body.Bytes(), &respData)
	assert.NoError(t, err)
	assert.Equal(t, "successfully added product to wishlist", respData["message"])
}

func TestViewWishlist(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockWishlistUsecase := mockUsecase.NewMockWishlistUseCase(ctrl)

	wishlistHandler := NewWishlistHandler(mockWishlistUsecase)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/wishlist", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Set up a mock WishlistUseCase to return a sample wishlist
	expectedWishlist := model.ViewWishlist{
		ID:     1,
		UserID: 1,
		Items:  []model.WishlistItem{},
	}
	mockWishlistUsecase.EXPECT().ViewWishlist(gomock.Any(), 1).Return(expectedWishlist, nil)

	// Call the ViewWishlist handler function
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Set("userID", 1)
	wishlistHandler.ViewWishlist(c)

	// Check that the response status code is correct
	assert.Equal(t, http.StatusOK, w.Code)

	// Check that the response body is correct
	var respData response.Response
	err = json.Unmarshal(w.Body.Bytes(), &respData)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, respData.StatusCode)
	assert.Equal(t, "successfully fetched wishlist", respData.Message)

	// Type assertion to check that the response data is of type ViewWishlist
	var respWishlist model.ViewWishlist
	respWishlistBytes, err := json.Marshal(respData.Data)
	assert.NoError(t, err)
	err = json.Unmarshal(respWishlistBytes, &respWishlist)
	assert.NoError(t, err)
	assert.Equal(t, expectedWishlist, respWishlist)

	assert.Nil(t, respData.Errors)
}

func TestRemoveFromWishlist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wishlistUseCase := mockUsecase.NewMockWishlistUseCase(ctrl)

	wishlistHandler := NewWishlistHandler(wishlistUseCase)

	r := gin.Default()
	r.DELETE("/wishlist/:id", wishlistHandler.RemoveFromWishlist)

	// Create a new HTTP request
	req, err := http.NewRequest("DELETE", "/wishlist/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Call the RemoveFromWishlist handler function
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Set("userID", 1)
	wishlistUseCase.EXPECT().RemoveFromWishlist(gomock.Any(), 1, 1).Return(nil)
	wishlistHandler.RemoveFromWishlist(c)

	// Check that the response status code is correct
	assert.Equal(t, http.StatusOK, w.Code)

	// Check that the response body is correct
	var respData gin.H
	err = json.Unmarshal(w.Body.Bytes(), &respData)
	assert.NoError(t, err)
	assert.Equal(t, "successfully removed product item from wishlist", respData["message"])
}

func TestEmptyWishlist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	wishlistUseCase := mockUsecase.NewMockWishlistUseCase(ctrl)

	wishlistHandler := NewWishlistHandler(wishlistUseCase)

	r := gin.Default()
	r.DELETE("/wishlist", wishlistHandler.EmptyWishlist)

	// Create a new HTTP request
	req, err := http.NewRequest("DELETE", "/wishlist", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Set the userID value in the request context
	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Set("userID", 1)

	// Set up mock to expect EmptyWishlist call
	wishlistUseCase.EXPECT().EmptyWishlist(gomock.Any(), 1).Return(nil)

	// Call the EmptyWishlist handler function
	wishlistHandler.EmptyWishlist(c)

	// Check that the response status code is correct
	assert.Equal(t, http.StatusOK, w.Code)

	// Check that the response body is correct
	var respData gin.H
	err = json.Unmarshal(w.Body.Bytes(), &respData)
	assert.NoError(t, err)
	assert.Equal(t, "successfully removed everything from wishlist", respData["message"])
}
