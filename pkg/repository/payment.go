package repository

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	interfaces "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type paymentDatabase struct {
	DB *gorm.DB
}

func NewPaymentRepository(DB *gorm.DB) interfaces.PaymentRepository {
	return &paymentDatabase{DB}
}

func (c *paymentDatabase) ViewPaymentDetails(ctx context.Context, orderID int) (domain.PaymentDetails, error) {
	var paymentDetails domain.PaymentDetails
	fetchPaymentDetailsQuery := `SELECT * FROM payment_details WHERE order_id = $1;`
	err := c.DB.Raw(fetchPaymentDetailsQuery, orderID).Scan(&paymentDetails).Error
	return paymentDetails, err
}

func (c *paymentDatabase) UpdatePaymentDetails(ctx context.Context, orderID int, paymentRef string) (domain.PaymentDetails, error) {
	var updatedPayment domain.PaymentDetails
	updatePaymentQuery := `	UPDATE payment_details SET payment_method_id = 2, payment_status_id = 2, payment_ref = $1, updated_at = NOW()
							WHERE order_id = $2 RETURNING *;`
	err := c.DB.Raw(updatePaymentQuery, paymentRef, orderID).Scan(&updatedPayment).Error
	return updatedPayment, err
}
