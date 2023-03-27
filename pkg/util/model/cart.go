package model

type DisplayCart struct {
	ProductItemID    uint
	Brand            string
	Name             string
	Model            string
	Quantity         uint
	ProductItemImage string
	Price            float64
	Total            float64
}

type ViewCart struct {
	CartItems []DisplayCart `json:"cart_items,omitempty"`
	CouponID  int           `json:"coupon_id,omitempty"`
	SubTotal  float64       `json:"sub_total"`
	Discount  float64       `json:"discount"`
	CartTotal float64       `json:"cart_total,omitempty"`
}
