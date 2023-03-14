package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/modelHelper"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
)

type OrderUseCases interface {
	//	BuyProductItem
	BuyProductItem(ctx context.Context, cookie string, orderInfo modelHelper.PlaceOrder) (domain.Order, error)
	BuyAll(ctx context.Context, cookie string, orderInfo modelHelper.PlaceAllOrders) (domain.Order, error)
	ViewOrderByID(ctx context.Context, orderID int, cookie string) (domain.Order, error)
	ViewAllOrders(ctx context.Context, cookie string) ([]domain.Order, error)
	CancelOrder(ctx context.Context, orderID int, cookie string) (domain.Order, error)
	//	BuyAll
	//	ViewOrderByID
	//	ViewAllOrders
	//	CancelOrder

}
