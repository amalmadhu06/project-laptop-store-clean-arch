package domain

import "time"

type Order struct {
	ID                uint `gorm:"primaryKey"`
	UserID            uint
	Users             Users `gorm:"foreignKey:UserID" json:"-"`
	OrderDate         time.Time
	PaymentMethodID   uint
	PaymentMethod     PaymentMethod `gorm:"foreignKey:PaymentMethodID" json:"-"`
	ShippingAddressID uint
	Address           Address `gorm:"foreignKey:ShippingAddressID" json:"-"`
	OrderTotal        float64
	OrderStatusID     uint
	OrderStatus       OrderStatus `gorm:"foreignKey:OrderStatusID" json:"-"`
}

type OrderLine struct {
	ID            uint `gorm:"primaryKey"`
	ProductItemID uint
	ProductItem   ProductItem `gorm:"foreignKey:ProductItemID" json:"-"`
	OrderID       uint
	Order         Order `gorm:"foreignKey:OrderID" json:"-"`
	Quantity      int
	Price         float64
}

type OrderStatus struct {
	ID          uint `gorm:"primaryKey"`
	OrderStatus string
}
