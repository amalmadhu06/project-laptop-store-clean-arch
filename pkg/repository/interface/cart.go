package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/modelHelper"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
)

type CartRepository interface {
	AddToCart(ctx context.Context, userID int, productItemID int) (domain.CartItems, error)
	RemoveFromCart(ctx context.Context, userID int, productItemID int) error
	ViewCart(ctx context.Context, userID int) (modelHelper.ViewCart, error)
	EmptyCart(ctx context.Context, userID int) error
	AddCouponToCart(ctx context.Context, userID, couponID int) (modelHelper.ViewCart, error)
}
