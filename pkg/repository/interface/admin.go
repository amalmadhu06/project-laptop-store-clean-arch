package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
)

type AdminRepository interface {
	IsSuperAdmin(ctx context.Context, adminId int) (bool, error)
	CreateAdmin(ctx context.Context, newAdminInfo model.NewAdminInfo) (domain.Admin, error)
	FindAdmin(ctx context.Context, email string) (domain.Admin, error)
	BlockAdmin(ctx context.Context, blockID int) (domain.Admin, error)
	UnblockAdmin(ctx context.Context, unblockID int) (domain.Admin, error)
	AdminDashboard(ctx context.Context) (model.AdminDashboard, error)
	SalesReport(ctx context.Context) ([]model.SalesReport, error)
}
