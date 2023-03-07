package handler

import (
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/response"
	services "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CartHandler struct {
	cartUseCase services.CartUseCases
}

func NewCartHandler(usecase services.CartUseCases) *CartHandler {
	return &CartHandler{
		cartUseCase: usecase,
	}
}

// AddToCart
// @Summary Add a product item to the cart
// @ID add-to-cart
// @Description User's can add product to cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param
// @Success
// @Failure
// @Failure
// @Router /user/add-to-cart [post]
func (cr *CartHandler) AddToCart(c *gin.Context) {
	paramsID := c.Param("product_item_id")
	productItemID, err := strconv.Atoi(paramsID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "unable to process the request",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	cookie, err := c.Cookie("UserAuth")
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Response{
			StatusCode: 401,
			Message:    "unable to fetch cookie",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	cartItem, err := cr.cartUseCase.AddToCart(c.Request.Context(), cookie, productItemID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "failed to add product to the cart",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, response.Response{
		StatusCode: 201,
		Message:    "Successfully added product item to the cart",
		Data:       cartItem,
		Errors:     nil,
	})
}

// RemoveFromCart
// @Summary Remove a product from the cart
// @ID remove-from-cart
// @Description User can remove product from cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param
// @Success
// @Failure
// @Failure
// @Router /user/remove-from-cart [delete]
func (cr *CartHandler) RemoveFromCart(c *gin.Context) {
	paramsID := c.Param("product_item_id")
	productItemID, err := strconv.Atoi(paramsID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{
			StatusCode: 422,
			Message:    "unable to process the request",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	cookie, err := c.Cookie("UserAuth")
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Response{
			StatusCode: 401,
			Message:    "unable to fetch cookie",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	err = cr.cartUseCase.RemoveFromCart(c.Request.Context(), cookie, productItemID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{
			StatusCode: 400,
			Message:    "failed to remove product from the cart",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusNoContent, response.Response{
		StatusCode: 204,
		Message:    "Successfully removed product from the cart",
		Data:       nil,
		Errors:     nil,
	})
}

// ViewCart
// @Summary
// @ID
// @Description
// @Tags
// @Accept json
// @Produce json
// @Param
// @Success
// @Failure
// @Failure
//func (cr *CartHandler) ViewCart(c *gin.Context) {
//	cookie, err := c.Cookie("UserAuth")
//	if err != nil {
//		c.JSON(http.StatusUnauthorized, response.Response{
//			StatusCode: 401,
//			Message:    "unable to fetch cookie",
//			Data:       nil,
//			Errors:     err.Error(),
//		})
//		return
//	}
//
//	cart, err := cr.cartUseCase.ViewCart(c.Request.Context(), cookie)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, response.Response{
//			StatusCode: 500,
//			Message:    "failed to cart details",
//			Data:       nil,
//			Errors:     err.Error(),
//		})
//		return
//	}
//	c.JSON(http.StatusOK, response.Response{
//		StatusCode: 200,
//		Message:    "Successfully fetched cart",
//		Data:       cart,
//		Errors:     nil,
//	})
//}
