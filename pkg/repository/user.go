package repository

import (
	"context"
	"fmt"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/common/modelHelper"
	domain "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	interfaces "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

func (c *userDatabase) CreateUser(ctx context.Context, user modelHelper.UserDataInput) (modelHelper.UserDataOutput, error) {
	var userData modelHelper.UserDataOutput
	//query for creating a new entry in users table
	createUserQuery := "INSERT INTO users (f_name,l_name,email,phone,password, created_at)VALUES($1,$2,$3,$4,$5, NOW()) RETURNING id,f_name,l_name,email,phone"
	//Todo : Context Cancelling
	err := c.DB.Raw(createUserQuery, user.FName, user.LName, user.Email, user.Phone, user.Password).Scan(&userData).Error

	if err == nil {
		//query for creating a new entry in the user_infos table.
		insertUserInfoQuery := "INSERT INTO user_infos (is_verified,is_blocked,users_id) VALUES('f','f',$1);"
		err = c.DB.Exec(insertUserInfoQuery, userData.ID).Error
	}
	return userData, err
}

func (c *userDatabase) FindByEmail(ctx context.Context, email string) (modelHelper.UserLoginVerifier, error) {
	var userData modelHelper.UserLoginVerifier
	findUserQuery := "SELECT users.id, users.f_name, users.l_name, users.email, users.phone, users.password, infos.is_blocked, infos.is_verified FROM users as users FULL OUTER JOIN user_infos as infos ON users.id = infos.users_id WHERE users.email = $1;"
	//Todo : Context Cancelling
	err := c.DB.Raw(findUserQuery, email).Scan(&userData).Error
	fmt.Println("printing user data from db", userData)
	return userData, err
}

func (c *userDatabase) FindByPhone(ctx context.Context, phone string) (modelHelper.UserLoginVerifier, error) {
	var userData modelHelper.UserLoginVerifier
	findUserQuery := "SELECT users.id, users.f_name, users.l_name, users.email, users.phone, users.password, infos.is_blocked, infos.is_verified FROM users as users FULL OUTER JOIN user_infos as infos ON users.id = infos.users_id WHERE users.phone = $1;"
	//Todo : Context Cancelling
	err := c.DB.Raw(findUserQuery, phone).Scan(&userData).Error
	fmt.Println("printing user data from db", userData)
	return userData, err
}

func (c *userDatabase) FindAll(ctx context.Context) ([]domain.Users, error) {
	//time.Sleep(60 * time.Second) // sleep for 15 seconds
	var users []domain.Users
	err := c.DB.Find(&users).Error

	return users, err
}

func (c *userDatabase) FindByID(ctx context.Context, id uint) (domain.Users, error) {
	var user domain.Users
	err := c.DB.First(&user, id).Error

	return user, err
}

func (c *userDatabase) Save(ctx context.Context, user domain.Users) (domain.Users, error) {
	err := c.DB.Save(&user).Error

	return user, err
}

func (c *userDatabase) Delete(ctx context.Context, user domain.Users) error {
	err := c.DB.Delete(&user).Error

	return err
}
