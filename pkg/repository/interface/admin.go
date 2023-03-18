package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/modelHelper"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
)

type AdminRepository interface {
	IsSuperAdmin(ctx context.Context, adminId int) (bool, error)
	CreateAdmin(ctx context.Context, newAdminInfo modelHelper.NewAdminInfo) (domain.Admin, error)
	FindAdmin(ctx context.Context, email string) (domain.Admin, error)
	BlockAdmin(ctx context.Context, blockID int) (domain.Admin, error)
	UnblockAdmin(ctx context.Context, unblockID int) (domain.Admin, error)
	AdminDashboard(ctx context.Context) (modelHelper.AdminDashboard, error)
}
