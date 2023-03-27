package model

import "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"

type UserDataInput struct {
	FName    string `json:"f_name"`
	LName    string `json:"l_name"`
	Email    string `json:"email" binding:"required" validate:"required,email"`
	Phone    string `json:"phone" validate:"required,phone"`
	Password string `json:"password" binding:"required" validate:"required,min=8,max=64"`
}

type UserDataOutput struct {
	ID    uint   `json:"user_id"`
	FName string `json:"f_name"`
	LName string `json:"l_name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type UserLoginEmail struct {
	Email    string `json:"email" binding:"required" validate:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserLoginPhone struct {
	Phone    string `json:"phone" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginVerifier struct {
	ID         uint   `json:"user_id"`
	FName      string `json:"f_name"`
	LName      string `json:"l_name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	IsBlocked  bool   `json:"is_blocked"`
	IsVerified bool   `json:"is_verified"`
}

type AddressInput struct {
	HouseNumber string `json:"house_number"`
	Street      string `json:"street"`
	City        string `json:"city"`
	District    string `json:"district"`
	Pincode     string `json:"pincode"`
	Landmark    string `json:"landmark"`
}

type BlockUser struct {
	UserID int    `json:"user_id"`
	Reason string `json:"reason"`
}

type UserProfile struct {
	UserInfo UserDataOutput
	Address  domain.Address
	Orders   []domain.Order
}
