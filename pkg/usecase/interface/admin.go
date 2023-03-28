package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
)

type AdminUseCase interface {
	CreateAdmin(ctx context.Context, newAdmin model.NewAdminInfo, adminID int) (domain.Admin, error)
	FindAdminID(ctx context.Context, cookie string) (int, error)
	AdminLogin(ctx context.Context, input model.AdminLogin) (string, model.AdminDataOutput, error)
	BlockAdmin(ctx context.Context, blockID int, superAdminID int) (domain.Admin, error)
	UnblockAdmin(ctx context.Context, unblockID int, superAdminID int) (domain.Admin, error)
	AdminDashboard(ctx context.Context) (model.AdminDashboard, error)
	SalesReport(ctx context.Context) ([]model.SalesReport, error)
}
