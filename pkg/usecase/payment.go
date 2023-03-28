package usecase

import (
	"context"
	"fmt"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	interfaces "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository/interface"
	services "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/usecase/interface"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
	"github.com/razorpay/razorpay-go"
)

const (
	razorpayID     = "rzp_test_W8zjL7kK4HLV61"
	razorpaySecret = "D5LlGN9gnynqZRacjS6qNB86"
)

type paymentUseCase struct {
	paymentRepo interfaces.PaymentRepository
	orderRepo   interfaces.OrderRepository
}

func NewPaymentUseCase(orderRepo interfaces.OrderRepository, paymentRepo interfaces.PaymentRepository) services.PaymentUseCases {
	return &paymentUseCase{
		paymentRepo: paymentRepo,
		orderRepo:   orderRepo,
	}
}

func (cr *paymentUseCase) CreateRazorpayPayment(ctx context.Context, userID, orderID int) (domain.Order, string, error) {
	//check payment status. if already paid, no need to proceed with payment. If not paid yet, proceed with transaction.
	paymentDetails, err := cr.paymentRepo.ViewPaymentDetails(ctx, orderID)
	if paymentDetails.PaymentStatusID == 2 {
		return domain.Order{}, "", fmt.Errorf("payment already completed")
	}
	//fetch order details from the db
	order, err := cr.orderRepo.ViewOrderById(ctx, userID, orderID)
	if order.ID == 0 {
		return domain.Order{}, "", fmt.Errorf("no such order found")
	}
	client := razorpay.NewClient(razorpayID, razorpaySecret)

	data := map[string]interface{}{
		"amount":   order.OrderTotal * 100,
		"currency": "INR",
		"receipt":  "test_receipt_id",
	}

	body, err := client.Order.Create(data, nil)
	if err != nil {
		return domain.Order{}, "", err
	}

	value := body["id"]
	razorpayID := value.(string)
	return order, razorpayID, err
}

func (cr *paymentUseCase) UpdatePaymentDetails(ctx context.Context, paymentVerifier model.PaymentVerification) error {
	//	fetch payment details

	paymentDetails, err := cr.paymentRepo.ViewPaymentDetails(ctx, paymentVerifier.OrderID)
	if err != nil {
		return err
	}

	if paymentDetails.ID == 0 {
		return fmt.Errorf("no order found")
	}

	if paymentDetails.OrderTotal != paymentVerifier.Total {
		return fmt.Errorf("payment amount and order amount does not match")
	}

	updatedPayment, err := cr.paymentRepo.UpdatePaymentDetails(ctx, paymentVerifier.OrderID, paymentVerifier.PaymentRef)
	if err != nil {
		return err
	}

	if updatedPayment.ID == 0 {
		return fmt.Errorf("failed to update payment details")
	}
	return nil
}
