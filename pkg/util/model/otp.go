package model

type OTPData struct {
	Phone string `json:"phone" validate:"required,phone"`
}

type VerifyData struct {
	Phone *OTPData `json:"phone" validate:"required,phone"`
	Otp   string   `json:"otp" validate:"required"`
}
