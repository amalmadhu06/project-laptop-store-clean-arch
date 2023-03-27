package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
)

type OrderRepository interface {
	BuyProductItem(ctx context.Context, userID int, orderInfo model.PlaceOrder) (domain.Order, error)
	BuyAll(ctx context.Context, userID int, orderInfo model.PlaceAllOrders) (domain.Order, error)
	ViewOrderById(ctx context.Context, userID int, orderID int) (domain.Order, error)
	ViewAllOrders(ctx context.Context, userID int) ([]domain.Order, error)
	CancelOrder(ctx context.Context, userID int, orderID int) (domain.Order, error)
	UpdateOrder(ctx context.Context, orderInfo model.UpdateOrder) (domain.Order, error)
	ReturnRequest(ctx context.Context, returnRequest model.ReturnRequest) (domain.Order, error)
}
