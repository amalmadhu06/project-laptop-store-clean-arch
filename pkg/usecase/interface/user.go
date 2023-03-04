package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/modelHelper"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, input modelHelper.UserDataInput) (modelHelper.UserDataOutput, error)
	LoginWithEmail(ctx context.Context, input modelHelper.UserLoginEmail) (string, modelHelper.UserDataOutput, error)
	LoginWithPhone(ctx context.Context, input modelHelper.UserLoginPhone) (string, modelHelper.UserDataOutput, error)
	//LoginWithOtp(ctx context.Context, input modelHelper.UserLoginInfo) (string, modelHelper.UserDataOutput, error)
	FindAll(ctx context.Context) ([]domain.Users, error)
	FindByID(ctx context.Context, id uint) (domain.Users, error)
	Save(ctx context.Context, user domain.Users) (domain.Users, error)
	Delete(ctx context.Context, user domain.Users) error
}
