package repository

import (
	"context"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"
	interfaces "github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/repository/interface"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/util/model"
	"gorm.io/gorm"
)

type adminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminDatabase{DB}
}

func (c *adminDatabase) IsSuperAdmin(ctx context.Context, adminId int) (bool, error) {
	var isSuperAdmin bool
	superAdminCheckQuery := `	SELECT is_super_admin
								FROM admins
								WHERE id = $1`
	err := c.DB.Raw(superAdminCheckQuery, adminId).Scan(&isSuperAdmin).Error
	return isSuperAdmin, err
}

func (c *adminDatabase) CreateAdmin(ctx context.Context, newAdminInfo model.NewAdminInfo) (domain.Admin, error) {
	var newAdmin domain.Admin
	createAdminQuery := `	INSERT INTO admins(user_name, email, password, is_super_admin, is_blocked, created_at, updated_at)
							VALUES($1, $2, $3, false,false, NOW(), NOW()) RETURNING *;`
	err := c.DB.Raw(createAdminQuery, newAdminInfo.UserName, newAdminInfo.Email, newAdminInfo.Password).Scan(&newAdmin).Error
	newAdmin.Password = ""
	return newAdmin, err
}

func (c *adminDatabase) FindAdmin(ctx context.Context, email string) (domain.Admin, error) {
	var adminData domain.Admin
	findAdminQuery := `	SELECT * 
						FROM admins
						WHERE email = $1;`
	//Todo : Context Cancelling
	err := c.DB.Raw(findAdminQuery, email).Scan(&adminData).Error
	return adminData, err
}

func (c *adminDatabase) BlockAdmin(ctx context.Context, blockID int) (domain.Admin, error) {
	var blockedAdmin domain.Admin
	blockQuery := `	UPDATE admins
					SET is_blocked = 'true',
					updated_at = NOW()
					WHERE id = $1
					RETURNING *;`
	err := c.DB.Raw(blockQuery, blockID).Scan(&blockedAdmin).Error
	blockedAdmin.Password = ""
	return blockedAdmin, err
}

func (c *adminDatabase) UnblockAdmin(ctx context.Context, unblockID int) (domain.Admin, error) {
	var unblockedAdmin domain.Admin
	unblockQuery := `	UPDATE admins
					SET is_blocked = 'false',
					updated_at = NOW()
					WHERE id = $1
					RETURNING *;`
	err := c.DB.Raw(unblockQuery, unblockID).Scan(&unblockedAdmin).Error
	unblockedAdmin.Password = ""
	return unblockedAdmin, err
}

func (c *adminDatabase) AdminDashboard(ctx context.Context) (model.AdminDashboard, error) {
	var dashboardData model.AdminDashboard
	fetchOrdersSummaryQuery := `SELECT 
								  COUNT(CASE WHEN order_status_id = 4 THEN id END) AS completed_orders,
								  COUNT(CASE WHEN order_status_id = 1 THEN id END) AS pending_orders,
								  COUNT(CASE WHEN order_status_id = 2 OR order_status_id = 3 THEN id END) AS cancelled_orders,
								  COUNT(id) AS total_orders,
								  SUM(CASE WHEN o.order_status_id != 2 AND o.order_status_id != 3 THEN o.order_total ELSE 0 END) AS order_value,
								  COUNT(DISTINCT o.user_id) AS ordered_users
								FROM orders o;`

	err := c.DB.Raw(fetchOrdersSummaryQuery).Scan(&dashboardData).Error
	if err != nil {
		return dashboardData, err
	}

	totalOrderedItemsQuery := `SELECT count(id) AS total_order_items FROM order_lines;`
	err = c.DB.Raw(totalOrderedItemsQuery).Scan(&dashboardData.TotalOrderItems).Error
	if err != nil {
		return dashboardData, err
	}

	creditedAmountQuery := `SELECT sum(order_total) as credited_amount FROM payment_details WHERE payment_status_id = 2`
	err = c.DB.Raw(creditedAmountQuery).Scan(&dashboardData.CreditedAmount).Error
	if err != nil {
		//return dashboardData, err
	}

	dashboardData.PendingAmount = dashboardData.OrderValue - dashboardData.CreditedAmount

	userCountQuery := `SELECT 
						  COUNT(*) AS total_users, 
						  COUNT(CASE WHEN is_verified = true THEN 1 END) AS verified_users
						FROM user_infos;`

	err = c.DB.Raw(userCountQuery).Scan(&dashboardData).Error

	return dashboardData, err
}

func (c *adminDatabase) SalesReport(ctx context.Context) ([]model.SalesReport, error) {
	var salesData []model.SalesReport
	salesDataQuery := `	
						SELECT
							o.id AS order_id, 
							o.user_id, 
							o.order_total AS total, 	
							c.code AS coupon_code, 
							pm.payment_method, 
							os.order_status, 
							ds.status AS delivery_status,
							o.order_date		
						FROM orders o 
						LEFT JOIN
							payment_methods pm ON o.payment_method_id = pm.id 
						LEFT JOIN
							order_statuses os ON o.order_status_id = os.id  
						LEFT JOIN	
							delivery_statuses ds ON o.delivery_status_id = ds.id
						LEFT JOIN 
							coupons c ON o.coupon_id = c.id;`

	err := c.DB.Raw(salesDataQuery).Scan(&salesData).Error
	return salesData, err
}
