package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/modelHelper"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user modelHelper.UserDataInput) (modelHelper.UserDataOutput, error)
	FindByEmail(ctx context.Context, email string) (modelHelper.UserLoginVerifier, error)
	FindByPhone(ctx context.Context, phone string) (modelHelper.UserLoginVerifier, error)

	FindAll(ctx context.Context) ([]domain.Users, error)
	FindByID(ctx context.Context, id uint) (domain.Users, error)
	Save(ctx context.Context, user domain.Users) (domain.Users, error)
	Delete(ctx context.Context, user domain.Users) error
}
