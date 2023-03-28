package handler

import (
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/api/handlerUtil"
	services "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/usecase/interface"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	orderUseCase services.OrderUseCases
}

func NewOrderHandler(usecase services.OrderUseCases) *OrderHandler {
	return &OrderHandler{
		orderUseCase: usecase,
	}
}

// BuyProductItem handles an HTTP POST request to buy a specific product item.
// @Summary Buy product item
// @ID buy-product-item
// @Description Buy a product item using product item ID.
// @Tags Order
// @Accept json
// @Produce json
// @Param order_details body model.PlaceOrder true "Order Details"
// @Success 201 {object} response.Response "Successfully ordered product item"
// @Failure 400 {object} response.Response "Failed to order the product item"
// @Failure 401 {object} response.Response "Unable to fetch authentication cookie"
// @Failure 422 {object} response.Response "Unable to process the request"
// @Router /orders/ [post]
func (cr *OrderHandler) BuyProductItem(c *gin.Context) {
	var body model.PlaceOrder
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "unable to read request body", Data: nil, Errors: err.Error()})
		return
	}

	userID, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Response{StatusCode: 400, Message: "unable to fetch user id from context", Data: nil, Errors: err.Error()})
		return
	}

	order, err := cr.orderUseCase.BuyProductItem(c.Request.Context(), userID, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to order the product item", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, response.Response{StatusCode: 201, Message: "Successfully ordered product item", Data: order, Errors: nil})
}

// BuyAll
// @Summary Buy all items from the user's cart
// @ID buyAll
// @Description This endpoint allows a user to purchase all items in their cart
// @Tags Order
// @Accept json
// @Produce json
// @Param order_details body model.PlaceAllOrders true "Order Details"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /orders/buy-all [post]
func (cr *OrderHandler) BuyAll(c *gin.Context) {
	var body model.PlaceAllOrders
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "unable to read request body", Data: nil, Errors: err.Error()})
		return
	}

	userID, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Response{StatusCode: 400, Message: "unable to fetch user id from context", Data: nil, Errors: err.Error()})
		return
	}

	order, err := cr.orderUseCase.BuyAll(c.Request.Context(), userID, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to create order", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, response.Response{StatusCode: 201, Message: "Successfully ordered all items from cart", Data: order, Errors: nil})
}

// ViewOrderByID function retrieves order details for a given order ID, if authorized.
// @Summary Retrieves order details for a given order ID, if authorized.
// @ID view-order-by-id
// @Description This function handles requests for retrieving the details of a specific order identified by its order ID. The user must be authorized with a valid cookie to view the order details.
// @Tags Order
// @Accept json
// @Produce json
// @Param order_id path int true "Order ID"
// @Success 200 {object} response.Response "Successfully fetched order details"
// @Failure 400 {object} response.Response "Failed to fetch order details"
// @Failure 401 {object} response.Response "Failed to authorize user"
// @Failure 422 {object} response.Response "Failed to read order ID from path"
// @Router /orders/{order_id} [get]
func (cr *OrderHandler) ViewOrderByID(c *gin.Context) {
	paramsId := c.Param("order_id")
	orderID, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to read order id from path", Data: nil, Errors: err.Error()})
		return
	}

	userID, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Response{StatusCode: 400, Message: "unable to fetch user id from context", Data: nil, Errors: err.Error()})
		return
	}
	order, err := cr.orderUseCase.ViewOrderByID(c.Request.Context(), orderID, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to fetch order details", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{StatusCode: 200, Message: "successfully fetched order details", Data: order, Errors: nil})
}

// ViewAllOrders
// @Summary Retrieves all orders of currently logged in user
// @ID view-all-orders
// @Description Endpoint for getting all orders associated with a user
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router  /orders/ [get]
func (cr *OrderHandler) ViewAllOrders(c *gin.Context) {
	userID, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Response{StatusCode: 400, Message: "unable to fetch user id from context", Data: nil, Errors: err.Error()})
		return
	}
	orders, err := cr.orderUseCase.ViewAllOrders(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to fetch orders", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{StatusCode: 200, Message: "successfully fetched orders", Data: orders, Errors: nil})
}

// CancelOrder
// @Summary Cancels a specific order for the currently logged in user
// @ID cancel-order
// @Description Endpoint for cancelling an order associated with a user
// @Tags Order
// @Accept json
// @Produce json
// @Param order_id path int true "ID of the order to be cancelled"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /orders/cancel/{order_id} [put]
func (cr *OrderHandler) CancelOrder(c *gin.Context) {
	paramsId := c.Param("order_id")
	orderID, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to read order id from path", Data: nil, Errors: err.Error()})
		return
	}
	userID, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Response{StatusCode: 400, Message: "unable to fetch user id from context", Data: nil, Errors: err.Error()})
		return
	}
	order, err := cr.orderUseCase.CancelOrder(c.Request.Context(), orderID, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to cancel order", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{StatusCode: 200, Message: "successfully cancelled order", Data: order, Errors: nil})
}

// UpdateOrder allows admins to update order status
// @Summary Admin can update order status of any order using order_id
// @ID update-order
// @Description Endpoint for updating order status
// @Tags Order
// @Accept json
// @Produce json
// @Param order_info body model.UpdateOrder true "Details of the order to be updated"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /admin/orders [put]
func (cr *OrderHandler) UpdateOrder(c *gin.Context) {
	var body model.UpdateOrder
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to read request bod", Data: nil, Errors: err.Error()})
		return
	}
	order, err := cr.orderUseCase.UpdateOrder(c.Request.Context(), body)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to update order", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.Response{StatusCode: 200, Message: "successfully updated order", Data: order, Errors: nil})
}

// ReturnRequest
// @Summary User can request for returning products within 15 days of order delivery
// @ID return-request
// @Description User can request for returning products withing 15 days of order delivery
// @Tags Order
// @Accept json
// @Produce json
// @Param return_request body model.ReturnRequest true "Return details"
// @Success 202 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /orders/return [post]
func (cr *OrderHandler) ReturnRequest(c *gin.Context) {
	var body model.ReturnRequest
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to read request body", Data: nil, Errors: err.Error()})
		return
	}
	userID, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Response{StatusCode: 400, Message: "unable to fetch user id from context", Data: nil, Errors: err.Error()})
		return
	}
	order, err := cr.orderUseCase.ReturnRequest(c.Request.Context(), userID, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 400, Message: "failed to place return request", Data: nil, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, response.Response{StatusCode: 202, Message: "successfully placed return request", Data: order, Errors: nil})
}
