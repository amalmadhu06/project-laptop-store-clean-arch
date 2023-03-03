package repository

import (
	interfaces "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type otpDatabase struct {
	DB *gorm.DB
}

func NewOtpRepository(DB *gorm.DB) interfaces.OtpRepository {
	return &otpDatabase{DB}
}

//
//func (c *otpDatabase) UpdateAsVerified(ctx context.Context, phone string) error {
//	insertQuery := "INSERT INTO user_info s "
//	return nil
//}
