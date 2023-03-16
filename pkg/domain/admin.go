package domain

import "time"

type Admin struct {
	ID           uint   `gorm:"primaryKey,index" json:"id"`
	UserName     string `gorm:"uniqueIndex" json:"user_name"`
	Email        string `gorm:"uniqueIndex" json:"email"`
	Password     string `json:"password,omitempty"`
	IsSuperAdmin bool   `json:"is_super_admin"`
	IsBlocked    bool   `json:"is_blocked"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
