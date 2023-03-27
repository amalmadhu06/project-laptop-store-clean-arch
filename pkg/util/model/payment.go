package model

type PaymentVerification struct {
	UserID     int
	OrderID    int
	PaymentRef string
	Total      float64
}
