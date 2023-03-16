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

type UpdateOrder struct {
	OrderID       int `json:"order_id"`
	OrderStatusID int `json:"order_status_id"`
}