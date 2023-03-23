package handler

import (
	"fmt"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/modelHelper"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/response"
	services "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type PaymentHandler struct {
	paymentUseCase services.PaymentUseCases
}

func NewPaymentHandler(paymentUseCase services.PaymentUseCases) *PaymentHandler {
	return &PaymentHandler{
		paymentUseCase: paymentUseCase,
	}
}

// CreateRazorpayPayment
// @Summary Users can make payment
// @ID create-razorpay-payment
// @Description Users can make payment via Razorpay after placing orders
// @Tags Payment
// @Accept json
// @Produce json
// @Param order_id path string true "Order id"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /payments/razorpay/{order_id} [get]
func (cr *PaymentHandler) CreateRazorpayPayment(c *gin.Context) {
	paramsID := c.Param("order_id")
	orderID, err := strconv.Atoi(paramsID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "failed to read order id", Data: nil, Errors: err.Error()})

	}

	cookie, err := c.Cookie("UserAuth")
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.Response{StatusCode: 401, Message: "unable to fetch cookie", Data: nil, Errors: err.Error()})
		return
	}
	order, razorpayID, err := cr.paymentUseCase.CreateRazorpayPayment(c.Request.Context(), cookie, orderID)
	fmt.Println("order total :", order.OrderTotal)
	fmt.Println(razorpayID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "failed to complete order", Data: nil, Errors: err.Error()})
		return
	}

	c.HTML(200, "app.html", gin.H{
		"UserID":      order.UserID,
		"total_price": order.OrderTotal,
		"total":       order.OrderTotal,
		"orderData":   order.ID,
		"orderid":     razorpayID,
		//"orderid":      order.ID,
		"amount":       order.OrderTotal,
		"Email":        "amalmadhu@gmail.com",
		"Phone_Number": "7902638843",
	})

	//c.JSON(http.StatusAccepted, response.Response{StatusCode: 202, Message: "successfully completed payment using razorpay", Data: nil, Errors: nil})
}

// PaymentSuccess
// @Summary Handling successful payment
// @ID payment-success
// @Description Handler for automatically updating payment details upon successful payment
// @Tags Payment
// @Accept json
// @Produce json
// @Param c query string true "Payment details"
// @Success 202 {object} response.Response "Successfully updated payment details"
// @Failure 500 {object} response.Response "Failed to update payment details"
// @Router /payments/success/ [get]
func (cr *PaymentHandler) PaymentSuccess(c *gin.Context) {
	paymentRef := c.Query("payment_ref")
	fmt.Println("paymentRef from query :", paymentRef)

	idStr := c.Query("order_id")
	fmt.Print("order id from query _:", idStr)

	idStr = strings.ReplaceAll(idStr, " ", "")

	orderID, err := strconv.Atoi(idStr)
	fmt.Println("_converted order  id from query :", orderID)

	uID := c.Query("user_id")
	userID, err := strconv.Atoi(uID)

	t := c.Query("total")
	fmt.Println("total from query :", t)
	total, err := strconv.ParseFloat(t, 32)
	fmt.Println("total from query converted:", total)

	if err != nil {
		//	handle err
		fmt.Println("failed to fetch order id")
	}

	//orderID := strings.Trim("orderid", " ")

	paymentVerifier := modelHelper.PaymentVerification{
		UserID:     userID,
		OrderID:    orderID,
		PaymentRef: paymentRef,
		Total:      total,
	}

	//paymentVerifier.
	err = cr.paymentUseCase.UpdatePaymentDetails(c.Request.Context(), paymentVerifier)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "failed to update payment details", Data: false, Errors: err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, response.Response{StatusCode: 202, Message: "payment success", Data: true, Errors: nil})
}
