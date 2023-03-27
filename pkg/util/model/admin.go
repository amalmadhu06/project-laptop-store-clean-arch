package model

import "time"

type NewAdminInfo struct {
	UserName string `json:"user_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AdminLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AdminDataOutput struct {
	ID           uint
	UserName     string
	Email        string
	IsSuperAdmin bool
}

type AdminDashboard struct {
	CompletedOrders int     `json:"completed_orders,omitempty"`
	PendingOrders   int     `json:"pending_orders,omitempty"`
	CancelledOrders int     `json:"cancelled_orders,omitempty"`
	TotalOrders     int     `json:"total_orders,omitempty"`
	TotalOrderItems int     `json:"total_order_items,omitempty"`
	OrderValue      float64 `json:"order_value,omitempty"`
	CreditedAmount  float64 `json:"credited_amount,omitempty"`
	PendingAmount   float64 `json:"pending_amount,omitempty"`

	TotalUsers    int `json:"total_users,omitempty"`
	VerifiedUsers int `json:"verified_users,omitempty"`
	OrderedUsers  int `json:"ordered_users,omitempty"`
}

type SalesReport struct {
	OrderID        int
	UserID         int
	Total          float64
	CouponCode     string
	PaymentMethod  string
	OrderStatus    string
	DeliveryStatus string
	OrderDate      time.Time
}
