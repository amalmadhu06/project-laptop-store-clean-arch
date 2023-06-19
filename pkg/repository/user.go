package repository

import (
	"context"
	"fmt"
	domain "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	interfaces "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository/interface"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
	"gorm.io/gorm"
	"strings"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

func (c *userDatabase) CreateUser(ctx context.Context, user model.UserDataInput) (model.UserDataOutput, error) {
	var userData model.UserDataOutput
	//query for creating a new entry in users table
	createUserQuery := `INSERT INTO users (f_name,l_name,email,phone,password, created_at)	
						VALUES($1,$2,$3,$4,$5, NOW()) 
						RETURNING id,f_name,l_name,email,phone`
	//Todo : Context Cancelling
	err := c.DB.Raw(createUserQuery, user.FName, user.LName, user.Email, user.Phone, user.Password).Scan(&userData).Error

	if err == nil {
		//query for creating a new entry in the user_infos table.
		insertUserInfoQuery := `INSERT INTO user_infos (is_verified,is_blocked,users_id) 
								VALUES('f','f',$1);`
		err = c.DB.Exec(insertUserInfoQuery, userData.ID).Error
	}
	return userData, err
}

func (c *userDatabase) FindByEmail(ctx context.Context, email string) (model.UserLoginVerifier, error) {
	var userData model.UserLoginVerifier
	findUserQuery := `	SELECT users.id, users.f_name, users.l_name, users.email, users.phone, users.password, infos.is_blocked, infos.is_verified 
						FROM users
						FULL OUTER JOIN 
							user_infos as infos
						ON users.id = infos.users_id 
						WHERE users.email = $1;`
	//Todo : Context Cancelling
	err := c.DB.Raw(findUserQuery, email).Scan(&userData).Error
	return userData, err
}

func (c *userDatabase) FindByPhone(ctx context.Context, phone string) (model.UserLoginVerifier, error) {
	var userData model.UserLoginVerifier
	findUserQuery := `SELECT users.id, users.f_name, users.l_name, users.email, users.phone, users.password, infos.is_blocked, infos.is_verified 
					  FROM users
					  FULL OUTER JOIN user_infos as infos 
					  ON users.id = infos.users_id 
					  WHERE users.phone = $1;`
	//Todo : Context Cancelling
	err := c.DB.Raw(findUserQuery, phone).Scan(&userData).Error
	return userData, err
}

func (c *userDatabase) AddAddress(ctx context.Context, userID int, newAddress model.AddressInput) (domain.Address, error) {
	var existingAddress, addedAddress domain.Address
	findAddressQuery := `SELECT * FROM addresses WHERE user_id = $1`
	err := c.DB.Raw(findAddressQuery, userID).Scan(&existingAddress).Error
	if err != nil {
		return domain.Address{}, err
	}
	if existingAddress.ID == 0 || existingAddress.UserID == 0 {
		//no address is found in user table, so insert query
		insertAddressQuery := `	INSERT INTO addresses(
								user_id, house_number, street, city, district, pincode, landmark) 
								VALUES($1,$2,$3,$4,$5,$6, $7) RETURNING *`
		err := c.DB.Raw(insertAddressQuery, userID, newAddress.HouseNumber, newAddress.Street, newAddress.City, newAddress.District, newAddress.Pincode, newAddress.Landmark).Scan(&addedAddress).Error
		return addedAddress, err
	} else {
		//	address is already there, update it
		updateAddressQuery := `	UPDATE addresses SET
								house_number = $1, street = $2, city = $3, district = $4, pincode = $5, landmark = $6
								WHERE user_id = $7
								RETURNING *`
		err := c.DB.Raw(updateAddressQuery, newAddress.HouseNumber, newAddress.Street, newAddress.City, newAddress.District, newAddress.Pincode, newAddress.Landmark, userID).Scan(&addedAddress).Error
		return addedAddress, err
	}
}

func (c *userDatabase) UpdateAddress(ctx context.Context, userID int, address model.AddressInput) (domain.Address, error) {
	var updatedAddress domain.Address
	updateQuery := `UPDATE addresses SET house_number = $1, street = $2, city = $3, district = $4, pincode = $5, landmark = $6 WHERE user_id = $7 RETURNING *`
	err := c.DB.Raw(updateQuery, address.HouseNumber, address.Street, address.City, address.District, address.Pincode, address.Landmark, userID).Scan(&updatedAddress).Error

	if updatedAddress.ID == 0 {
		return domain.Address{}, fmt.Errorf("failed to update address")
	}
	return updatedAddress, err
}

func (c *userDatabase) ViewAddress(ctx context.Context, userID int) (domain.Address, error) {
	var address domain.Address
	fetchAddressQuery := `SELECT * FROM addresses WHERE user_id = $1`
	err := c.DB.Raw(fetchAddressQuery, userID).Scan(&address).Error
	return address, err
}

func (c *userDatabase) ListAllUsers(ctx context.Context, queryParams model.QueryParams) ([]domain.Users, error) {

	findQuery := "SELECT * FROM users"
	if queryParams.Query != "" && queryParams.Filter != "" {
		findQuery = fmt.Sprintf("%s WHERE LOWER(%s) LIKE '%%%s%%'", findQuery, queryParams.Filter, strings.ToLower(queryParams.Query))
	}
	if queryParams.SortBy != "" {
		if queryParams.SortDesc {
			findQuery = fmt.Sprintf("%s ORDER BY %s DESC", findQuery, queryParams.SortBy)
		} else {
			findQuery = fmt.Sprintf("%s ORDER BY %s ASC", findQuery, queryParams.SortBy)
		}
	}
	if queryParams.Limit != 0 && queryParams.Page != 0 {
		findQuery = fmt.Sprintf("%s LIMIT %d OFFSET %d", findQuery, queryParams.Limit, (queryParams.Page-1)*queryParams.Limit)
	}
	if queryParams.Limit == 0 || queryParams.Page == 0 {
		findQuery = fmt.Sprintf("%s LIMIT 10 OFFSET 0", findQuery)
	}
	var users []domain.Users

	err := c.DB.Raw(findQuery).Scan(&users).Error
	if len(users) == 0 {
		return users, fmt.Errorf("no users found")
	}
	return users, err
}

func (c *userDatabase) FindUserByID(ctx context.Context, userID int) (domain.Users, error) {
	var user domain.Users
	findUser := `SELECT * FROM users WHERE id = $1;`
	err := c.DB.Raw(findUser, userID).Scan(&user).Error
	if user.ID == 0 {
		return domain.Users{}, fmt.Errorf("no user found")
	}
	return user, err
}

func (c *userDatabase) BlockUser(ctx context.Context, blockInfo model.BlockUser, adminID int) (domain.UserInfo, error) {
	var userInfo domain.UserInfo
	blockQuery := `UPDATE user_infos SET is_blocked = 'true', blocked_at = NOW(), blocked_by = $1, reason_for_blocking = $2 WHERE users_id = $3 RETURNING *;`
	err := c.DB.Raw(blockQuery, adminID, blockInfo.Reason, blockInfo.UserID).Scan(&userInfo).Error

	if userInfo.UsersID == 0 {
		return domain.UserInfo{}, fmt.Errorf("user not found")
	}

	return userInfo, err
}

func (c *userDatabase) UnblockUser(ctx context.Context, userID int) (domain.UserInfo, error) {
	var userInfo domain.UserInfo
	unblockQuery := `UPDATE user_infos SET is_blocked = 'false', reason_for_blocking = '' WHERE users_id = $1 RETURNING *;`
	err := c.DB.Raw(unblockQuery, userID).Scan(&userInfo).Error
	if userInfo.UsersID == 0 {
		return domain.UserInfo{}, fmt.Errorf("no user found")
	}
	return userInfo, err
}
