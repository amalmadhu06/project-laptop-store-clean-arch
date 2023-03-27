package handler

import (
	"fmt"
	services "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/usecase/interface"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type WishlistHandler struct {
	wishlistUsecase services.WishlistUseCase
}

func NewWishlistHandler(usecase services.WishlistUseCase) *WishlistHandler {
	return &WishlistHandler{
		wishlistUsecase: usecase,
	}
}

// AddToWishlist
// @Summary User can add product item to wishlist
// @ID add-to-wishlist
// @Description User can add product item to wishlist
// @Tags Wishlist
// @Accept json
// @Produce json
// @Param id path string true "ID of the product item to be added to wishlist"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /wishlist/{id} [post]
func (cr *WishlistHandler) AddToWishlist(c *gin.Context) {
	paramsID := c.Param("id")
	productItemID, err := strconv.Atoi(paramsID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to fetch product item id from request", Data: nil, Errors: err.Error()})
		return
	}
	uID := c.Value("userID")
	userID, err := strconv.Atoi(fmt.Sprintf("%v", uID))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to fetch user id", Data: nil, Errors: err.Error()})
		return
	}
	wishlist, err := cr.wishlistUsecase.AddToWishlist(c.Request.Context(), userID, productItemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "failed to add product item to wishlist", Data: nil, Errors: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response.Response{StatusCode: 201, Message: "successfully added product to wishlist", Data: wishlist, Errors: nil})
}

// ViewWishlist
// @Summary User can view items in wishlist
// @ID view-wishlist
// @Description User view product items in wishlist
// @Tags Wishlist
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /wishlist/ [get]
func (cr *WishlistHandler) ViewWishlist(c *gin.Context) {
	uID := c.Value("userID")
	userID, err := strconv.Atoi(fmt.Sprintf("%v", uID))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to fetch user id", Data: nil, Errors: err.Error()})
		return
	}
	wishlist, err := cr.wishlistUsecase.ViewWishlist(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "failed to fetch wishlist info", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{StatusCode: 200, Message: "successfully fetched wishlist", Data: wishlist, Errors: nil})

}

// RemoveFromWishlist
// @Summary User can remove product item from wishlist
// @ID remove-from-wishlist
// @Description User can remove product item from wishlist
// @Tags Wishlist
// @Accept json
// @Produce json
// @Param id path string true "ID of the product item to be removed from wishlist"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /wishlist/{id} [delete]
func (cr *WishlistHandler) RemoveFromWishlist(c *gin.Context) {
	paramsID := c.Param("id")
	productItemID, err := strconv.Atoi(paramsID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to fetch product item id", Data: nil, Errors: err.Error()})
	}
	uID := c.Value("userID")
	userID, err := strconv.Atoi(fmt.Sprintf("%v", uID))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to fetch user ID from context", Data: nil, Errors: err.Error()})
		return
	}
	if err := cr.wishlistUsecase.RemoveFromWishlist(c.Request.Context(), userID, productItemID); err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "failed to remove product item from wishlist", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{StatusCode: 200, Message: "successfully removed product item from wishlist", Data: nil, Errors: nil})
}

// EmptyWishlist
// @Summary User can remove all product items from wishlist
// @ID empty-wishlist
// @Description User can remove all product items from wishlist
// @Tags Wishlist
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /wishlist/ [delete]
func (cr *WishlistHandler) EmptyWishlist(c *gin.Context) {
	uID := c.Value("userID")
	userID, err := strconv.Atoi(fmt.Sprintf("%v", uID))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to fetch user ID from context", Data: nil, Errors: err.Error()})
		return
	}
	if err := cr.wishlistUsecase.EmptyWishlist(c.Request.Context(), userID); err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "failed to empty wishlist", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{StatusCode: 200, Message: "successfully removed everything from wishlist", Data: nil, Errors: nil})
}
