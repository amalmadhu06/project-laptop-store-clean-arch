package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/modelHelper"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
)

type AdminUseCase interface {
	CreateAdmin(ctx context.Context, newAdmin modelHelper.NewAdminInfo, adminID int) (domain.Admin, error)
	FindAdminID(ctx context.Context, cookie string) (int, error)
	AdminLogin(ctx context.Context, input modelHelper.AdminLogin) (string, modelHelper.AdminDataOutput, error)
	BlockAdmin(ctx context.Context, blockID int, cookie string) (domain.Admin, error)
	UnblockAdmin(ctx context.Context, unblockID int, cookie string) (domain.Admin, error)
}
