package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
)

type PaymentRepository interface {
	ViewPaymentDetails(ctx context.Context, orderID int) (domain.PaymentDetails, error)
	UpdatePaymentDetails(ctx context.Context, orderID int, paymentRef string) (domain.PaymentDetails, error)
}
