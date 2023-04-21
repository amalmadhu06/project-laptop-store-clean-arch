package repository

import (
	"context"
	interfaces "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository/interface"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
	"gorm.io/gorm"
)

type otpDatabase struct {
	DB *gorm.DB
}

func NewOtpRepository(DB *gorm.DB) interfaces.OtpRepository {
	return &otpDatabase{DB}
}

func (c *otpDatabase) UpdateAsVerified(ctx context.Context, phone string) error {
	updateVerifyQuery := `	UPDATE user_infos 
							SET is_verified = true, verified_at = NOW() 
							WHERE users_id IN ( 
									SELECT id 
									FROM users 
									WHERE phone = $1);`
	err := c.DB.Exec(updateVerifyQuery, phone).Error
	return err
}

func (c *otpDatabase) CheckWithMobile(ctx context.Context, phone string) (bool, error) {
	var isPresent bool
	findQuery := `SELECT EXISTS(
						SELECT 1 
						FROM users 
						WHERE phone = $1);`
	err := c.DB.Raw(findQuery, phone).Scan(&isPresent).Error
	return isPresent, err
}

func (c *otpDatabase) FindByPhone(ctx context.Context, phone string) (model.UserLoginVerifier, error) {
	var userData model.UserLoginVerifier
	findUserQuery := `	SELECT users.id, users.f_name, users.l_name, users.phone, users.phone, users.password, infos.is_blocked, infos.is_verified 
						FROM users as users 
						FULL OUTER JOIN user_infos as infos 
						ON users.id = infos.users_id 	
						WHERE users.phone = $1;`
	//Todo : Context Cancelling
	err := c.DB.Raw(findUserQuery, phone).Scan(&userData).Error
	return userData, err
}
