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

	AddAddress(ctx context.Context, userID int, newAddress modelHelper.AddressInput) (domain.Address, error)
	UpdateAddress(ctx context.Context, userID int, address modelHelper.AddressInput) (domain.Address, error)
	ViewAddress(ctx context.Context, userID int) (domain.Address, error)

	ListAllUsers(ctx context.Context) ([]domain.Users, error)
	FindUserByID(ctx context.Context, userID int) (domain.Users, error)
	BlockUser(ctx context.Context, blockInfo modelHelper.BlockUser, adminID int) (domain.UserInfo, error)
	UnblockUser(ctx context.Context, userID int) (domain.UserInfo, error)
}
