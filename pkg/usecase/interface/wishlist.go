package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
)

type WishlistUseCase interface {
	AddToWishlist(ctx context.Context, userID, productItemID int) (model.ViewWishlist, error)
	ViewWishlist(ctx context.Context, userID int) (model.ViewWishlist, error)
	RemoveFromWishlist(ctx context.Context, userID, productItemID int) error
	EmptyWishlist(ctx context.Context, userID int) error
}
