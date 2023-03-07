package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
)

type CartRepository interface {
	AddToCart(ctx context.Context, userID int, productItemID int) (domain.CartItems, error)
	RemoveFromCart(ctx context.Context, userID int, productItemID int) error
}
