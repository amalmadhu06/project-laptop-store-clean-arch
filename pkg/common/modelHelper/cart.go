package modelHelper

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
	CartItems []DisplayCart
	CartTotal float64
}
