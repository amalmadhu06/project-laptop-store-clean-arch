package modelHelper

type PlaceOrder struct {
	ProductItemID     int `json:"product_item_id,omitempty"`
	PaymentMethodID   int `json:"payment_method_id,omitempty"`
	ShippingAddressID int `json:"shipping_address_id,omitempty"`
}
type PlaceAllOrders struct {
	PaymentMethodID   int `json:"payment_method_id,omitempty"`
	ShippingAddressID int `json:"shipping_address_id,omitempty"`
}
