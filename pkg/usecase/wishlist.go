package usecase

import (
	"context"
	"fmt"
	interfaces "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository/interface"
	services "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/usecase/interface"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
)

type wishlistUsecase struct {
	wishlistRepo interfaces.WishlistRepository
}

func NewWishlistUsecase(wishlistRepo interfaces.WishlistRepository) services.WishlistUseCase {
	return &wishlistUsecase{
		wishlistRepo: wishlistRepo,
	}
}

func (c *wishlistUsecase) AddToWishlist(ctx context.Context, userID, productItemID int) (model.ViewWishlist, error) {
	if err := c.wishlistRepo.AddToWishlist(ctx, userID, productItemID); err != nil {
		return model.ViewWishlist{}, fmt.Errorf("failed to add product item to the cart : %v", err)
	}
	wishlist, err := c.wishlistRepo.ViewWishlist(ctx, userID)
	if err != nil {
		return model.ViewWishlist{}, fmt.Errorf("failed to retrieve wishlist info : %w", err)
	}
	return wishlist, nil
}

func (c *wishlistUsecase) ViewWishlist(ctx context.Context, userID int) (model.ViewWishlist, error) {
	wishlist, err := c.wishlistRepo.ViewWishlist(ctx, userID)
	if err != nil {
		return model.ViewWishlist{}, fmt.Errorf("failed to retrive wishlist info : %v", err)
	}
	return wishlist, nil

}

func (c *wishlistUsecase) RemoveFromWishlist(ctx context.Context, userID, productItemID int) error {
	err := c.wishlistRepo.RemoveFromWishlist(ctx, userID, productItemID)
	return err
}

func (c *wishlistUsecase) EmptyWishlist(ctx context.Context, userID int) error {
	err := c.wishlistRepo.EmptyWishlist(ctx, userID)
	return err
}
