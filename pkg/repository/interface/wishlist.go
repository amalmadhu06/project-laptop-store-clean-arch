package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/modelHelper"
)

type WishlistRepository interface {
	AddToWishlist(ctx context.Context, userID, productItemID int) error
	ViewWishlist(ctx context.Context, userID int) (modelHelper.ViewWishlist, error)
	RemoveFromWishlist(ctx context.Context, userID, productItemID int) error
	EmptyWishlist(ctx context.Context, userID int) error
}
