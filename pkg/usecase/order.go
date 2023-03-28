package usecase

import (
	"context"
	"fmt"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	interfaces "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository/interface"
	services "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/usecase/interface"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
	"time"
)

type orderUseCase struct {
	orderRepo   interfaces.OrderRepository
	userRepo    interfaces.UserRepository
	productRepo interfaces.ProductRepository
}

func NewOrderUseCase(orderRepo interfaces.OrderRepository, userRepo interfaces.UserRepository, productRepo interfaces.ProductRepository) services.OrderUseCases {
	return &orderUseCase{
		orderRepo:   orderRepo,
		userRepo:    userRepo,
		productRepo: productRepo,
	}
}

func (c *orderUseCase) BuyProductItem(ctx context.Context, userID int, orderInfo model.PlaceOrder) (domain.Order, error) {

	//check if user has added address. If not, return error
	address, err := c.userRepo.ViewAddress(ctx, userID)
	if err != nil {
		return domain.Order{}, err
	}
	if address.ID == 0 {
		return domain.Order{}, fmt.Errorf("cannot place order without adding address")
	}

	//check if coupon is valid, applicable
	var appliedCoupon domain.Coupon

	//if user applied a coupon
	if orderInfo.CouponID != 0 {

		//check if coupon is already used
		isUsed, err := c.productRepo.CouponUsed(ctx, userID, orderInfo.CouponID)
		if err != nil {
			return domain.Order{}, err
		}
		if isUsed {
			return domain.Order{}, fmt.Errorf("coupon already used")
		}

		appliedCoupon, err = c.productRepo.ViewCouponByID(ctx, orderInfo.CouponID)
		if err != nil {
			return domain.Order{}, fmt.Errorf("failed to fetch coupon details")
		}
		currentTime := time.Now()
		if appliedCoupon.ValidTill.Before(currentTime) {
			return domain.Order{}, fmt.Errorf("expired coupon")
		}

		productInfo, err := c.productRepo.FindProductItemByID(ctx, orderInfo.ProductItemID)
		if err != nil {
			return domain.Order{}, err
		}
		if productInfo.ID == 0 {
			return domain.Order{}, fmt.Errorf("failed to fetch product item details")
		}

		//	check is product is eligible for coupon discount
		if productInfo.Price < appliedCoupon.MinOrderValue {
			return domain.Order{}, fmt.Errorf("cannot apply coupon as order total is less than required")
		}
	}

	order, err := c.orderRepo.BuyProductItem(ctx, userID, orderInfo)
	return order, err
}

func (c *orderUseCase) BuyAll(ctx context.Context, userID int, orderInfo model.PlaceAllOrders) (domain.Order, error) {

	//check if user has added address. If not, return error
	address, err := c.userRepo.ViewAddress(ctx, userID)
	if err != nil {
		return domain.Order{}, err
	}
	if address.ID == 0 {
		return domain.Order{}, fmt.Errorf("cannot place order without adding address")
	}

	orders, err := c.orderRepo.BuyAll(ctx, userID, orderInfo)

	return orders, err
}

func (c *orderUseCase) ViewOrderByID(ctx context.Context, orderID int, userID int) (domain.Order, error) {
	order, err := c.orderRepo.ViewOrderById(ctx, userID, orderID)
	return order, err

}

func (c *orderUseCase) ViewAllOrders(ctx context.Context, userID int) ([]domain.Order, error) {
	orders, err := c.orderRepo.ViewAllOrders(ctx, userID)
	return orders, err
}

func (c *orderUseCase) CancelOrder(ctx context.Context, orderID, userID int) (domain.Order, error) {
	order, err := c.orderRepo.CancelOrder(ctx, userID, orderID)
	return order, err
}

func (c *orderUseCase) UpdateOrder(ctx context.Context, orderInfo model.UpdateOrder) (domain.Order, error) {
	order, err := c.orderRepo.UpdateOrder(ctx, orderInfo)
	return order, err
}

func (c *orderUseCase) ReturnRequest(ctx context.Context, userID int, returnRequest model.ReturnRequest) (domain.Order, error) {
	//	check if order is eligible to be returned
	//	users can request for return only if order status is completed(4), delivery status is delivered(5) and is within 15 days of order delivery

	orderDetails, err := c.orderRepo.ViewOrderById(ctx, userID, returnRequest.OrderID)
	fmt.Println(orderDetails)
	if err != nil {
		return domain.Order{}, err
	}
	if orderDetails.ID == 0 {
		return domain.Order{}, fmt.Errorf("no such order found")
	}
	if orderDetails.DeliveryUpdatedAt.Sub(time.Now()) > time.Hour*24*15 {
		return domain.Order{}, fmt.Errorf("failed to place return request as it is more than 15 days")
	}
	if orderDetails.OrderStatusID != 4 || orderDetails.DeliveryStatusID != 5 {
		return domain.Order{}, fmt.Errorf("cannot return as order status is %v and delivery status is %v", orderDetails.OrderStatusID, orderDetails.DeliveryStatusID)
	}
	order, err := c.orderRepo.ReturnRequest(ctx, returnRequest)
	if err != nil {
		return domain.Order{}, fmt.Errorf("failed to place return request")
	}
	return order, nil
}
