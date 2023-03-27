package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
)

type WishlistRepository interface {
	AddToWishlist(ctx context.Context, userID, productItemID int) error
	ViewWishlist(ctx context.Context, userID int) (model.ViewWishlist, error)
	RemoveFromWishlist(ctx context.Context, userID, productItemID int) error
	EmptyWishlist(ctx context.Context, userID int) error
}
