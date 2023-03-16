package usecase

import (
	"context"
	"fmt"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/modelHelper"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	interfaces "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository/interface"
	services "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/usecase/interface"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (c *userUseCase) CreateUser(ctx context.Context, input modelHelper.UserDataInput) (modelHelper.UserDataOutput, error) {
	//Hashing user password
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		return modelHelper.UserDataOutput{}, err
	}
	input.Password = string(hash)
	userData, err := c.userRepo.CreateUser(ctx, input)

	return userData, err
}

func (c *userUseCase) LoginWithEmail(ctx context.Context, input modelHelper.UserLoginEmail) (string, modelHelper.UserDataOutput, error) {

	var userData modelHelper.UserDataOutput

	// 1. Find the userData with given email
	user, err := c.userRepo.FindByEmail(ctx, input.Email)
	if err != nil {
		return "", userData, fmt.Errorf("error finding userData")
	}
	if user.Email == "" {
		return "", userData, fmt.Errorf("no such userData found")
	}

	// 2. Compare and hash the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return "", userData, err
	}

	// 3. Check whether the userData is blocked by admin
	if user.IsBlocked {
		return "", userData, fmt.Errorf("userData account is blocked")
	}

	// 4. Create JWT Token
	// creating jwt token and sending it in cookie
	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// singed string
	ss, err := token.SignedString([]byte("secret"))

	// 4. Send back the created token

	//user data for sending back as response
	userData.ID, userData.FName, userData.LName, userData.Email, userData.Phone = user.ID, user.FName, user.LName, user.Email, user.Phone

	return ss, userData, err
}

func (c *userUseCase) LoginWithPhone(ctx context.Context, input modelHelper.UserLoginPhone) (string, modelHelper.UserDataOutput, error) {

	var userData modelHelper.UserDataOutput

	// 1. Find the userData with given email
	user, err := c.userRepo.FindByPhone(ctx, input.Phone)
	if err != nil {
		return "", userData, fmt.Errorf("error finding userData")
	}
	if user.Email == "" {
		return "", userData, fmt.Errorf("no such userData found")
	}

	// 2. Compare and hash the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return "", userData, err
	}

	// 3. Check whether the userData is blocked by admin
	if user.IsBlocked {
		return "", userData, fmt.Errorf("userData account is blocked")
	}

	// 4. Create JWT Token
	// creating jwt token and sending it in cookie
	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// singed string
	ss, err := token.SignedString([]byte("secret"))

	// 4. Send back the created token

	//user data for sending back as response
	userData.ID, userData.FName, userData.LName, userData.Email, userData.Phone = user.ID, user.FName, user.LName, user.Email, user.Phone

	return ss, userData, err
}

func (c *userUseCase) AddAddress(ctx context.Context, newAddress modelHelper.AddressInput, cookie string) (domain.Address, error) {

	userID, err := FindUserID(cookie)
	if err != nil {
		return domain.Address{}, err
	}
	address, err := c.userRepo.AddAddress(ctx, userID, newAddress)
	return address, err
}

func (c *userUseCase) UpdateAddress(ctx context.Context, addressInfo modelHelper.AddressInput, cookie string) (domain.Address, error) {
	userID, err := FindUserID(cookie)
	if err != nil {
		return domain.Address{}, err
	}
	updatedAddress, err := c.userRepo.UpdateAddress(ctx, userID, addressInfo)
	return updatedAddress, err
}

func (c *userUseCase) ListAllUsers(ctx context.Context) ([]domain.Users, error) {
	users, err := c.userRepo.ListAllUsers(ctx)
	return users, err
}

func (c *userUseCase) FindUserByID(ctx context.Context, userID int) (domain.Users, error) {
	user, err := c.userRepo.FindUserByID(ctx, userID)
	return user, err

}

func (c *userUseCase) BlockUser(ctx context.Context, blockInfo modelHelper.BlockUser, cookie string) (domain.UserInfo, error) {
	adminID, err := FindAdminID(cookie)
	if err != nil {
		return domain.UserInfo{}, err
	}
	blockedUser, err := c.userRepo.BlockUser(ctx, blockInfo, adminID)
	return blockedUser, err
}

func (c *userUseCase) UnblockUser(ctx context.Context, userID int) (domain.UserInfo, error) {
	unblockedUser, err := c.userRepo.UnblockUser(ctx, userID)
	return unblockedUser, err
}
