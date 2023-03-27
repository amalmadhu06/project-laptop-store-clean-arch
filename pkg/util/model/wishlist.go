package model

type ViewWishlist struct {
	ID     int
	UserID int
	Items  []WishlistItem
}

type WishlistItem struct {
	ProductItemID int
	Name          string
	Model         string
	Brand         string
	Price         float64
	Image         string
}
