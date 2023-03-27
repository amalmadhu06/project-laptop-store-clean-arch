package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
)

type OrderUseCases interface {
	//	BuyProductItem
	BuyProductItem(ctx context.Context, cookie string, orderInfo model.PlaceOrder) (domain.Order, error)
	BuyAll(ctx context.Context, cookie string, orderInfo model.PlaceAllOrders) (domain.Order, error)
	ViewOrderByID(ctx context.Context, orderID int, cookie string) (domain.Order, error)
	ViewAllOrders(ctx context.Context, cookie string) ([]domain.Order, error)
	CancelOrder(ctx context.Context, orderID int, cookie string) (domain.Order, error)
	UpdateOrder(ctx context.Context, orderInfo model.UpdateOrder) (domain.Order, error)
	ReturnRequest(ctx context.Context, userID int, returnRequest model.ReturnRequest) (domain.Order, error)
}
