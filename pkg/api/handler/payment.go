package handler

import (
	"fmt"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/modelHelper"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/response"
	services "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PaymentHandler struct {
	paymentUseCase services.PaymentUseCases
}

func NewPaymentHandler(paymentUseCase services.PaymentUseCases) *PaymentHandler {
	return &PaymentHandler{
		paymentUseCase: paymentUseCase,
	}
}

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

// url: `/payment-handler?user_id=${userid}&payment_ref=${res.razorpay_payment_id}&order_id=${orderData}
//
//	&signature=${res.razorpay_signature}&id=${orderid}&total=${total}`,
func (cr *PaymentHandler) PaymentSuccess(c *gin.Context) {
	paymentRef := c.Query("paymentid")

	id := c.Query("order_id")
	orderID, err := strconv.Atoi(id)

	uID := c.Query("user_id")
	userID, err := strconv.Atoi(uID)

	t := c.Query("total")
	total, err := strconv.ParseFloat(t, 32)

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

	paymentVerifier.
		err = cr.paymentUseCase.UpdatePaymentDetails(c.Request.Context(), paymentVerifier)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Response{StatusCode: 500, Message: "failed to update payment details", Data: false, Errors: err.Error()})
	}

	c.JSON(http.StatusAccepted, response.Response{StatusCode: 202, Message: "payment success", Data: true, Errors: nil})
}
