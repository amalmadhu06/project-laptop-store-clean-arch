package repository

import (
	"context"
	"fmt"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	interfaces "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository/interface"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
	"gorm.io/gorm"
)

type wishlistDatabase struct {
	DB *gorm.DB
}

func NewWishlistRepository(DB *gorm.DB) interfaces.WishlistRepository {
	return &wishlistDatabase{DB}
}

func (c *wishlistDatabase) AddToWishlist(ctx context.Context, userID, productItemID int) error {
	var wishlistID int
	if err := c.DB.Raw("SELECT id FROM wishlists WHERE user_id = $1", userID).Scan(&wishlistID).Error; err != nil {
		return err
	}
	if wishlistID == 0 {
		if err := c.DB.Raw("INSERT INTO wishlists (user_id, updated_at) VALUES($1, NOW()) RETURNING id;", userID).Scan(&wishlistID).Error; err != nil {
			return err
		}
	}
	var wishlistItem domain.Wishlist
	if err := c.DB.Raw("SELECT * FROM wishlist_items WHERE wishlist_id = $1 AND product_item_id = $2", wishlistID, productItemID).Scan(&wishlistItem).Error; err != nil {
		return err
	}

	if wishlistItem.ID == 0 {
		if err := c.DB.Exec("INSERT INTO wishlist_items (wishlist_id, product_item_id) VALUES ($1, $2)", wishlistID, productItemID).Error; err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("product alread in wishlist")
}

func (c *wishlistDatabase) ViewWishlist(ctx context.Context, userID int) (model.ViewWishlist, error) {
	var viewWishlist model.ViewWishlist
	var wishlist domain.Wishlist
	var wishlistItems []model.WishlistItem

	tx := c.DB.Begin()

	if err := tx.Raw("SELECT * FROM wishlists WHERE user_id = $1", userID).Scan(&wishlist).Error; err != nil {
		tx.Rollback()
		return model.ViewWishlist{}, err

	}
	if err := tx.Raw(`
							SELECT w.product_item_id,p.name, pi.model, b.brand, pi.price, pi.product_item_image 
							FROM wishlist_items w 
							INNER JOIN product_items pi ON w.product_item_id = pi.id
							INNER JOIN products p ON p.id = pi.product_id
							INNER JOIN product_brands b ON b.id = p.brand_id
							WHERE w.wishlist_id = $1;`,
		wishlist.ID).Scan(&wishlistItems).Error; err != nil {
		tx.Rollback()
		return model.ViewWishlist{}, err
	}
	viewWishlist.ID = int(wishlist.ID)
	viewWishlist.UserID = wishlist.UserID
	viewWishlist.Items = wishlistItems
	return viewWishlist, nil
}

func (c *wishlistDatabase) RemoveFromWishlist(ctx context.Context, userID, productItemID int) error {
	removeQuery := `DELETE FROM wishlist_items 
					WHERE wishlist_id IN(
										SELECT id FROM wishlists WHERE user_id = $1)
					AND product_item_id = $2`
	if err := c.DB.Exec(removeQuery, userID, productItemID).Error; err != nil {
		return err
	}
	return nil
}

func (c *wishlistDatabase) EmptyWishlist(ctx context.Context, userID int) error {
	emptyQuery := `	DELETE FROM wishlist_items
					WHERE wishlist_id IN(
										SELECT id FROM wishlists WHERE user_id = $1
										);`
	if err := c.DB.Exec(emptyQuery, userID).Error; err != nil {
		return err
	}
	return nil
}
