package interfaces

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, input model.UserDataInput) (model.UserDataOutput, error)
	LoginWithEmail(ctx context.Context, input model.UserLoginEmail) (string, model.UserDataOutput, error)
	LoginWithPhone(ctx context.Context, input model.UserLoginPhone) (string, model.UserDataOutput, error)

	AddAddress(ctx context.Context, newAddress model.AddressInput, userID int) (domain.Address, error)
	UpdateAddress(ctx context.Context, addressInfo model.AddressInput, userID int) (domain.Address, error)

	ListAllUsers(ctx context.Context, viewUserInfo model.QueryParams) ([]domain.Users, error)
	FindUserByID(ctx context.Context, userID int) (domain.Users, error)
	BlockUser(ctx context.Context, blockInfo model.BlockUser, adminID int) (domain.UserInfo, error)
	UnblockUser(ctx context.Context, userID int) (domain.UserInfo, error)

	UserProfile(ctx context.Context, userID int) (model.UserProfile, error)
}
