package repository

import (
	"context"
	interfaces "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type otpDatabase struct {
	DB *gorm.DB
}

func NewOtpRepository(DB *gorm.DB) interfaces.OtpRepository {
	return &otpDatabase{DB}
}

func (c *otpDatabase) UpdateAsVerified(ctx context.Context, phone string) error {
	updateVerifyQuery := "UPDATE user_infos SET is_verified = true, verified_at = NOW() WHERE users_id IN ( SELECT id FROM users WHERE phone = $1);"
	err := c.DB.Exec(updateVerifyQuery, phone).Error
	return err
}
