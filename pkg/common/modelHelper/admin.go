package modelHelper

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
