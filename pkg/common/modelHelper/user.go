package modelHelper

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

type UserLoginInfo struct {
	Email    string `json:"email" binding:"required" validate:"required,email"`
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
