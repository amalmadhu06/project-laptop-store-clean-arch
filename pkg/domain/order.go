package domain

import "time"

type Order struct {
	ID                uint           `gorm:"primaryKey"`
	UserID            uint           `json:"user_id"`
	Users             Users          `gorm:"foreignKey:UserID" json:"-"`
	OrderDate         time.Time      `json:"order_date"`
	PaymentMethodID   uint           `json:"payment_method_id"`
	PaymentMethod     PaymentMethod  `gorm:"foreignKey:PaymentMethodID" json:"-"`
	ShippingAddressID uint           `json:"shipping_address_id"`
	Address           Address        `gorm:"foreignKey:ShippingAddressID" json:"-"`
	OrderTotal        float64        `json:"order_total"`
	OrderStatusID     uint           `json:"order_status_id"`
	OrderStatus       OrderStatus    `gorm:"foreignKey:OrderStatusID" json:"-"`
	CouponID          uint           `json:"coupon_id"`
	DeliveryStatusID  int            `json:"delivery_status_id"`
	DeliveryStatus    DeliveryStatus `gorm:"primaryKey" json:"-"`
	DeliveryUpdatedAt time.Time      `json:"delivery_time"`
}

type OrderLine struct {
	ID            uint        `gorm:"primaryKey"`
	ProductItemID uint        `json:"product_item_id"`
	ProductItem   ProductItem `gorm:"foreignKey:ProductItemID" json:"-"`
	OrderID       uint        `json:"order_Id"`
	Order         Order       `gorm:"foreignKey:OrderID" json:"-"`
	Quantity      int         `json:"quantity"`
	Price         float64     `json:"price"`
}

type OrderStatus struct {
	ID          uint `gorm:"primaryKey"`
	OrderStatus string
}

type DeliveryStatus struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Status string `json:"status"`
}

type Return struct {
	ID       uint   `gorm:"primaryKey"`
	OrderID  int    `json:"order_id"`
	Order    Order  `gorm:"foreignKey:OrderID"`
	Reason   string `json:"string"`
	Approved bool   `json:"approved"`
}
