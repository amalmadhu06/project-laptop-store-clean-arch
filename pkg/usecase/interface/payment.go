package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/modelHelper"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
)

type PaymentUseCases interface {
	CreateRazorpayPayment(ctx context.Context, cookie string, orderID int) (domain.Order, string, error)
	UpdatePaymentDetails(ctx context.Context, paymentVerifier modelHelper.PaymentVerification) error
}
