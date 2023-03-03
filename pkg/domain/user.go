package domain

import (
	"time"
)

type Users struct {
	ID        uint      `gorm:"primaryKey,index" json:"user_id"`
	FName     string    `json:"f_name"`
	LName     string    `json:"l_name"`
	Email     string    `gorm:"unique" json:"email" binding:"required" validate:"required,email"`
	Phone     string    `gorm:"unique" json:"phone" validate:"required,phone"`
	Password  string    `json:"password" binding:"required" validate:"required,min=8,max=64"`
	CreatedAt time.Time `json:"created_at"`
}

type UserInfo struct {
	ID                uint `gorm:"primaryKey"`
	IsVerified        bool
	VerifiedAt        time.Time
	IsBlocked         bool
	BlockedAt         time.Time
	BlockedBy         uint
	ReasonForBlocking string
	UsersID           uint
	Users             Users
}
