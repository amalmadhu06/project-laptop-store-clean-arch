package domain

type Cart struct {
	ID     uint    `json:"id"`
	UserID uint    `json:"user_id"`
	Users  Users   `gorm:"foreignKey:UserID" json:"_"`
	Total  float64 `json:"total"`
}

type CartItems struct {
	ID            uint        `json:"id"`
	CartID        uint        `json:"cart_id"`
	Cart          Cart        `gorm:"foreignKey:CartID" json:"-"`
	ProductItemID uint        `json:"product_item_id"`
	ProductItem   ProductItem `json:"-"`
	Quantity      uint        `json:"quantity"`
}
