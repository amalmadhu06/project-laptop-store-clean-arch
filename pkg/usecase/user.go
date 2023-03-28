package usecase

import (
	"context"
	"fmt"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	interfaces "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository/interface"
	services "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/usecase/interface"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type userUseCase struct {
	userRepo  interfaces.UserRepository
	orderRepo interfaces.OrderRepository
}

func NewUserUseCase(userRepo interfaces.UserRepository, orderRepo interfaces.OrderRepository) services.UserUseCase {
	return &userUseCase{
		userRepo:  userRepo,
		orderRepo: orderRepo,
	}
}

func (c *userUseCase) CreateUser(ctx context.Context, input model.UserDataInput) (model.UserDataOutput, error) {
	//Hashing user password
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		return model.UserDataOutput{}, err
	}
	input.Password = string(hash)
	userData, err := c.userRepo.CreateUser(ctx, input)

	return userData, err
}

func (c *userUseCase) LoginWithEmail(ctx context.Context, input model.UserLoginEmail) (string, model.UserDataOutput, error) {

	var userData model.UserDataOutput

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

func (c *userUseCase) LoginWithPhone(ctx context.Context, input model.UserLoginPhone) (string, model.UserDataOutput, error) {

	var userData model.UserDataOutput

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

func (c *userUseCase) AddAddress(ctx context.Context, newAddress model.AddressInput, userID int) (domain.Address, error) {
	address, err := c.userRepo.AddAddress(ctx, userID, newAddress)
	return address, err
}

func (c *userUseCase) UpdateAddress(ctx context.Context, addressInfo model.AddressInput, userID int) (domain.Address, error) {
	updatedAddress, err := c.userRepo.UpdateAddress(ctx, userID, addressInfo)
	return updatedAddress, err
}

func (c *userUseCase) ListAllUsers(ctx context.Context, viewUserInfo model.QueryParams) ([]domain.Users, error) {
	users, err := c.userRepo.ListAllUsers(ctx, viewUserInfo)
	return users, err
}

func (c *userUseCase) FindUserByID(ctx context.Context, userID int) (domain.Users, error) {
	user, err := c.userRepo.FindUserByID(ctx, userID)
	return user, err

}

func (c *userUseCase) BlockUser(ctx context.Context, blockInfo model.BlockUser, adminID int) (domain.UserInfo, error) {
	blockedUser, err := c.userRepo.BlockUser(ctx, blockInfo, adminID)
	return blockedUser, err
}

func (c *userUseCase) UnblockUser(ctx context.Context, userID int) (domain.UserInfo, error) {
	unblockedUser, err := c.userRepo.UnblockUser(ctx, userID)
	return unblockedUser, err
}

func (c *userUseCase) UserProfile(ctx context.Context, userID int) (model.UserProfile, error) {

	// fetch user details from users_table
	userInfoRetrieved, err := c.userRepo.FindUserByID(ctx, userID)
	if err != nil {
		return model.UserProfile{}, err
	}
	userInfo := model.UserDataOutput{
		ID:    userInfoRetrieved.ID,
		FName: userInfoRetrieved.FName,
		LName: userInfoRetrieved.LName,
		Email: userInfoRetrieved.Email,
		Phone: userInfoRetrieved.Phone,
	}

	// fetch address from address table
	address, err := c.userRepo.ViewAddress(ctx, userID)
	if err != nil {
		return model.UserProfile{}, err
	}

	// fetch orders from orders table
	orders, err := c.orderRepo.ViewAllOrders(ctx, userID)
	if err != nil {
		return model.UserProfile{}, err
	}

	var userProfile model.UserProfile
	userProfile.UserInfo = userInfo
	userProfile.Address = address
	userProfile.Orders = orders

	return userProfile, nil
}
