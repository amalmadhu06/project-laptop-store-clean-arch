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

func (c *userUseCase) UserLogin(ctx context.Context, input modelHelper.UserLoginInfo) (string, modelHelper.UserDataOutput, error) {

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

func (c *userUseCase) FindAll(ctx context.Context) ([]domain.Users, error) {
	users, err := c.userRepo.FindAll(ctx)
	return users, err
}

func (c *userUseCase) FindByID(ctx context.Context, id uint) (domain.Users, error) {
	user, err := c.userRepo.FindByID(ctx, id)
	return user, err
}

func (c *userUseCase) Save(ctx context.Context, user domain.Users) (domain.Users, error) {
	user, err := c.userRepo.Save(ctx, user)

	return user, err
}

func (c *userUseCase) Delete(ctx context.Context, user domain.Users) error {
	err := c.userRepo.Delete(ctx, user)

	return err
}
