package repository

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/modelHelper"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	interfaces "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type adminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminDatabase{DB}
}

func (c *adminDatabase) IsSuperAdmin(ctx context.Context, adminId int) (bool, error) {
	var isSuperAdmin bool
	superAdminCheckQuery := `	SELECT is_super_admin
								FROM admins
								WHERE id = $1`
	err := c.DB.Raw(superAdminCheckQuery, adminId).Scan(&isSuperAdmin).Error
	return isSuperAdmin, err
}

func (c *adminDatabase) CreateAdmin(ctx context.Context, newAdminInfo modelHelper.NewAdminInfo) (domain.Admin, error) {
	var newAdmin domain.Admin
	createAdminQuery := `	INSERT INTO admins(user_name, email, password, is_super_admin, is_blocked, created_at, updated_at)
							VALUES($1, $2, $3, false,false, NOW(), NOW())`
	err := c.DB.Raw(createAdminQuery, newAdminInfo.UserName, newAdminInfo.Email, newAdminInfo.Password).Scan(&newAdmin).Error
	return newAdmin, err
}

func (c *adminDatabase) FindAdmin(ctx context.Context, email string) (domain.Admin, error) {
	var adminData domain.Admin
	findAdminQuery := `	SELECT * 
						FROM admins
						WHERE email = $1;`
	//Todo : Context Cancelling
	err := c.DB.Raw(findAdminQuery, email).Scan(&adminData).Error
	return adminData, err
}

//
//UserName     string
//Email        string
//Password     string
//IsSuperAdmin bool
//IsBlocked    bool
//CreatedAt    time.Time
//UpdatedAt    time.Time
