package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
)

type CartRepository interface {
	AddToCart(ctx context.Context, userID int, productItemID int) (domain.CartItems, error)
	RemoveFromCart(ctx context.Context, userID int, productItemID int) error
	ViewCart(ctx context.Context, userID int) (model.ViewCart, error)
	EmptyCart(ctx context.Context, userID int) error
	AddCouponToCart(ctx context.Context, userID, couponID int) (model.ViewCart, error)
}
