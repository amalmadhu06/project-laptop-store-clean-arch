package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
)

type CartUseCases interface {
	AddToCart(ctx context.Context, cookie string, productItemID int) (domain.CartItems, error)
	RemoveFromCart(ctx context.Context, cookie string, productItemID int) error
	ViewCart(ctx context.Context, cookie string) (model.ViewCart, error)
	EmptyCart(ctx context.Context, cookie string) error
	AddCouponToCart(ctx context.Context, userID int, couponID int) (model.ViewCart, error)
}
