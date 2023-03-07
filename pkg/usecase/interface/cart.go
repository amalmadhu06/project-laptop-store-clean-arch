package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
)

type CartUseCases interface {
	AddToCart(ctx context.Context, cookie string, productItemID int) (domain.CartItems, error)
	RemoveFromCart(ctx context.Context, cookie string, productItemID int) error
	//ViewCart(ctx context.Context, cookie string) (modelHelper.ViewCart, error)
}
