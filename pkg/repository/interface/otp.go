package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/modelHelper"
)

type OtpRepository interface {
	UpdateAsVerified(ctx context.Context, phone string) error
	CheckWithMobile(ctx context.Context, phone string) (bool, error)
	FindByPhone(ctx context.Context, phone string) (modelHelper.UserLoginVerifier, error)
}
