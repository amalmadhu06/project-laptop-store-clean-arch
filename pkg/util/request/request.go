package request

type UserData struct {
	FName string `json:"f_name"`
	LName string `json:"l_name"`
	Email string `gorm:"unique" json:"email" binding:"required" validate:"required,email"`
	Phone string `gorm:"unique" json:"phone" validate:"required,phone"`
}
