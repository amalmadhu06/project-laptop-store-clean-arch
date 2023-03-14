package usecase

import (
	"context"
	"fmt"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/modelHelper"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	interfaces "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository/interface"
	services "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/usecase/interface"
)

type orderUseCase struct {
	orderRepo interfaces.OrderRepository
}

func NewOrderUseCase(repo interfaces.OrderRepository) services.OrderUseCases {
	return &orderUseCase{
		orderRepo: repo,
	}
}

func (c *orderUseCase) BuyProductItem(ctx context.Context, cookie string, orderInfo modelHelper.PlaceOrder) (domain.Order, error) {
	//Find user id
	userID, err := FindUserID(cookie)
	if err != nil {
		return domain.Order{}, fmt.Errorf("failed to fetch user id")
	}
	order, err := c.orderRepo.BuyProductItem(ctx, userID, orderInfo)
	return order, err
}

func (c *orderUseCase) BuyAll(ctx context.Context, cookie string, orderInfo modelHelper.PlaceAllOrders) (domain.Order, error) {
	userID, err := FindUserID(cookie)
	if err != nil {
		return domain.Order{}, fmt.Errorf("failed to fetch user id")
	}
	orders, err := c.orderRepo.BuyAll(ctx, userID, orderInfo)

	return orders, err
}

func (c *orderUseCase) ViewOrderByID(ctx context.Context, orderID int, cookie string) (domain.Order, error) {
	userID, err := FindUserID(cookie)
	if err != nil {
		return domain.Order{}, fmt.Errorf("failed to fetch user id")
	}
	order, err := c.orderRepo.ViewOrderById(ctx, userID, orderID)
	return order, err

}

func (c *orderUseCase) ViewAllOrders(ctx context.Context, cookie string) ([]domain.Order, error) {
	userID, err := FindUserID(cookie)
	if err != nil {
		return []domain.Order{}, fmt.Errorf("failed to fetch user id")
	}
	orders, err := c.orderRepo.ViewAllOrders(ctx, userID)
	return orders, err
}

func (c *orderUseCase) CancelOrder(ctx context.Context, orderID int, cookie string) (domain.Order, error) {
	userID, err := FindUserID(cookie)
	if err != nil {
		return domain.Order{}, fmt.Errorf("failed to fetch user id")
	}
	order, err := c.orderRepo.CancelOrder(ctx, userID, orderID)
	return order, err
}
