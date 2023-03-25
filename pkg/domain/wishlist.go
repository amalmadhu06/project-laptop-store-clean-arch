package domain

import "time"

type Wishlist struct {
	ID        uint `gorm:"primaryKey"`
	UserID    int
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type WishlistItem struct {
	ID            uint `gorm:"primaryKey"`
	WishlistID    int
	Wishlist      Wishlist `gorm:"foreignKey:WishlistID"`
	ProductItemID int
	ProductItem   ProductItem `gorm:"foreignKey:ProductItemID"`
}
