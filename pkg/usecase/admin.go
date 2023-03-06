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

type adminUseCase struct {
	adminRepo interfaces.AdminRepository
}

func NewAdminUseCase(repo interfaces.AdminRepository) services.AdminUseCase {
	return &adminUseCase{
		adminRepo: repo,
	}
}

func (c *adminUseCase) CreateAdmin(ctx context.Context, newAdmin modelHelper.NewAdminInfo, adminID int) (domain.Admin, error) {

	isSuperAdmin, err := c.adminRepo.IsSuperAdmin(ctx, adminID)
	if err != nil {
		return domain.Admin{}, err
	}
	if !isSuperAdmin {
		return domain.Admin{}, fmt.Errorf("only super admin can create new admins")
	}

	//Hashing admin password
	hash, err := bcrypt.GenerateFromPassword([]byte(newAdmin.Password), 10)
	if err != nil {
		return domain.Admin{}, err
	}
	newAdmin.Password = string(hash)
	newAdminOutput, err := c.adminRepo.CreateAdmin(ctx, newAdmin)
	return newAdminOutput, err
}

func (c *adminUseCase) AdminLogin(ctx context.Context, input modelHelper.AdminLogin) (string, modelHelper.AdminDataOutput, error) {

	var adminData modelHelper.AdminDataOutput

	// 1. Find the adminData with given email
	adminInfo, err := c.adminRepo.FindAdmin(ctx, input.Email)
	if err != nil {
		return "", adminData, fmt.Errorf("error finding admin")
	}
	if adminInfo.Email == "" {
		return "", adminData, fmt.Errorf("no such admin found")
	}

	// 2. Compare and hash the password
	//if err := bcrypt.CompareHashAndPassword([]byte(adminInfo.Password), []byte(input.Password)); err != nil {
	//	return "", adminData, err
	//}

	// 3. Check whether the adminData is blocked by admin
	if adminInfo.IsBlocked {
		return "", adminData, fmt.Errorf(" admin account is blocked")
	}

	// 4. Create JWT Token
	// creating jwt token and sending it in cookie
	claims := jwt.MapClaims{
		"id":  adminInfo.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// singed string
	ss, err := token.SignedString([]byte("secret"))

	// 4. Send back the created token

	//adminInfo data for sending back as response
	adminData.ID, adminData.UserName, adminData.Email = adminInfo.ID, adminInfo.UserName, adminInfo.Email
	return ss, adminData, err
}

func (c *adminUseCase) FindAdminID(ctx context.Context, cookie string) (int, error) {
	//parses, validates, verifies the signature and returns the parsed token
	token, err := jwt.Parse(cookie, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		//secret was used for signing the string
		return []byte("secret"), nil
	})
	if err != nil {
		return 0, err
	}
	var parsedID interface{}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		parsedID = claims["id"]
	}
	//type assertion
	value, ok := parsedID.(float64)
	if !ok {
		return 0, fmt.Errorf("expected an int value, but got %T", parsedID)
	}

	id := int(value)

	return id, nil
}
