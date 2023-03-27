package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user model.UserDataInput) (model.UserDataOutput, error)
	FindByEmail(ctx context.Context, email string) (model.UserLoginVerifier, error)
	FindByPhone(ctx context.Context, phone string) (model.UserLoginVerifier, error)

	AddAddress(ctx context.Context, userID int, newAddress model.AddressInput) (domain.Address, error)
	UpdateAddress(ctx context.Context, userID int, address model.AddressInput) (domain.Address, error)
	ViewAddress(ctx context.Context, userID int) (domain.Address, error)

	ListAllUsers(ctx context.Context, queryParams model.QueryParams) ([]domain.Users, error)
	FindUserByID(ctx context.Context, userID int) (domain.Users, error)
	BlockUser(ctx context.Context, blockInfo model.BlockUser, adminID int) (domain.UserInfo, error)
	UnblockUser(ctx context.Context, userID int) (domain.UserInfo, error)
}
