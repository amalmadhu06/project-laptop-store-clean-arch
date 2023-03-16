package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/modelHelper"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
)

type OrderRepository interface {
	BuyProductItem(ctx context.Context, userID int, orderInfo modelHelper.PlaceOrder) (domain.Order, error)
	BuyAll(ctx context.Context, userID int, orderInfo modelHelper.PlaceAllOrders) (domain.Order, error)
	ViewOrderById(ctx context.Context, userID int, orderID int) (domain.Order, error)
	ViewAllOrders(ctx context.Context, userID int) ([]domain.Order, error)
	CancelOrder(ctx context.Context, userID int, orderID int) (domain.Order, error)
	UpdateOrder(ctx context.Context, orderInfo modelHelper.UpdateOrder) (domain.Order, error)
}
