package handler

import (
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/api/handlerUtil"
	services "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/usecase/interface"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/response"
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
// @Summary User can add a product item to the cart
// @ID add-to-cart
// @Description User can add product item to the cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param product_item_id path string true "product_item_id"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /cart/add/{product_item_id} [post]
func (cr *CartHandler) AddToCart(c *gin.Context) {
	paramsID := c.Param("product_item_id")
	productItemID, err := strconv.Atoi(paramsID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "unable to process the request", Data: nil, Errors: err.Error()})
		return
	}

	userID, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Response{StatusCode: 400, Message: "unable to fetch user id from context", Data: nil, Errors: err.Error()})
		return
	}

	cartItem, err := cr.cartUseCase.AddToCart(c.Request.Context(), userID, productItemID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to add product to the cart", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, response.Response{StatusCode: 201, Message: "Successfully added product item to the cart", Data: cartItem, Errors: nil})
}

// RemoveFromCart
// @Summary Remove a product from the cart
// @ID remove-from-cart
// @Description User can remove product from cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param product_item_id path string true "product_item_id"
// @Success 204 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /cart/remove/{product_item_id} [delete]
func (cr *CartHandler) RemoveFromCart(c *gin.Context) {
	paramsID := c.Param("product_item_id")
	productItemID, err := strconv.Atoi(paramsID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "unable to process the request", Data: nil, Errors: err.Error()})
		return
	}

	userID, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Response{StatusCode: 400, Message: "unable to fetch user id from context", Data: nil, Errors: err.Error()})
		return
	}

	err = cr.cartUseCase.RemoveFromCart(c.Request.Context(), userID, productItemID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to remove product from the cart", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, response.Response{StatusCode: 204, Message: "Successfully removed product from the cart", Data: nil, Errors: nil})
}

// ViewCart
// @Summary User can view cart items and total
// @ID view-cart
// @Description User can view cart and cart items
// @Tags Cart
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /cart [get]
func (cr *CartHandler) ViewCart(c *gin.Context) {
	userID, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Response{StatusCode: 400, Message: "unable to fetch user id from context", Data: nil, Errors: err.Error()})
		return
	}

	cart, err := cr.cartUseCase.ViewCart(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "failed to cart details", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{StatusCode: 200, Message: "Successfully fetched cart", Data: cart, Errors: nil})
}

// EmptyCart
// @Summary Remove everything from cart
// @ID empty-cart
// @Description User can remove everything from cart at once
// @Tags Cart
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /cart [delete]
func (cr *CartHandler) EmptyCart(c *gin.Context) {
	userID, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Response{StatusCode: 400, Message: "unable to fetch user id from context", Data: nil, Errors: err.Error()})
		return
	}

	err = cr.cartUseCase.EmptyCart(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "failed to empty cart", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{StatusCode: 200, Message: "Successfully emptied cart", Data: nil, Errors: nil})
}

// AddCouponToCart
// @Summary User can add a coupon to the cart
// @ID add-coupon-to-cart
// @Description User can add coupon to the cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param coupon_id path string true "coupon_id"
// @Success 202 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /cart/coupon/{coupon_id} [post]
func (cr *CartHandler) AddCouponToCart(c *gin.Context) {
	userID, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Response{StatusCode: 400, Message: "unable to fetch user id from context", Data: nil, Errors: err.Error()})
		return
	}

	paramsID := c.Param("coupon_id")
	couponID, err := strconv.Atoi(paramsID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to fetch coupon id", Data: nil, Errors: "failed to fetch coupon id from request"})
		return
	}

	cart, err := cr.cartUseCase.AddCouponToCart(c.Request.Context(), userID, couponID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to add coupon to the cart", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, response.Response{StatusCode: 202, Message: "successfully added coupon to the cart", Data: cart, Errors: nil})
}
