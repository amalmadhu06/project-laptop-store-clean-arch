package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
)

type PaymentUseCases interface {
	CreateRazorpayPayment(ctx context.Context, userID, orderID int) (domain.Order, string, error)
	UpdatePaymentDetails(ctx context.Context, paymentVerifier model.PaymentVerification) error
}
