package domain

import (
	"time"
)

type Users struct {
	ID        uint   `gorm:"primaryKey,index" json:"user_id"`
	FName     string `json:"f_name"`
	LName     string `json:"l_name"`
	Email     string `gorm:"uniqueIndex" json:"email" binding:"required" validate:"required,email"`
	Phone     string `gorm:"uniqueIndex" json:"phone" validate:"required,phone"`
	Password  string `json:"password" binding:"required" validate:"required,min=8,max=64"`
	CreatedAt time.Time
}

type UserInfo struct {
	ID                uint `gorm:"primaryKey"`
	IsVerified        bool `json:"is_verified"`
	VerifiedAt        time.Time
	IsBlocked         bool `json:"is_blocked"`
	BlockedAt         time.Time
	BlockedBy         uint   `json:"blocked_by"`
	Admin             Admin  `gorm:"foreignKey:BlockedBy" json:"-"`
	ReasonForBlocking string `json:"reason_for_blocking"`
	UsersID           uint   `json:"users_id" json:"-"`
	Users             Users  `gorm:"foreignKey:UsersID" json:"-"`
}

type Address struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	Users       Users  `gorm:"foreignKey:UserID" json:"-"`
	HouseNumber string `json:"house_number"`
	Street      string `json:"street"`
	City        string `json:"city"`
	District    string `json:"district"`
	Pincode     string `json:"pincode"`
	Landmark    string `json:"landmark"`
}
