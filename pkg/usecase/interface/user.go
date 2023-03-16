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

	AddAddress(ctx context.Context, newAddress modelHelper.AddressInput, cookie string) (domain.Address, error)
	UpdateAddress(ctx context.Context, addressInfo modelHelper.AddressInput, cookie string) (domain.Address, error)

	ListAllUsers(ctx context.Context) ([]domain.Users, error)
	FindUserByID(ctx context.Context, userID int) (domain.Users, error)
	BlockUser(ctx context.Context, blockInfo modelHelper.BlockUser, cookie string) (domain.UserInfo, error)
	UnblockUser(ctx context.Context, userID int) (domain.UserInfo, error)
}
