package db

import (
	"fmt"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/config"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	initDeliveryStatus string = `
INSERT INTO delivery_statuses (status)
	SELECT ds.status
	FROM (
		VALUES ('delivered'),('pending')
	) AS ds(status)
	LEFT JOIN delivery_statuses d 
	    ON d.status = ds.status
WHERE d.status IS NULL;
`

	initOrderStatus string = `
INSERT INTO order_statuses (order_status)
	SELECT status
	FROM (VALUES ('pending'),('cancelled by user'),('cancelled by admin'), ('completed'),('return requested')) AS statuses(status)
	LEFT JOIN order_statuses os ON os.order_status = statuses.status
WHERE os.order_status IS NULL;
`

	initPaymentMethod string = `
INSERT INTO
    payment_methods (payment_method)
	SELECT pm.payment_method FROM
    (VALUES ('cod'), ('online')) AS pm(payment_method)
    LEFT JOIN payment_methods p ON p.payment_method = pm.payment_method
WHERE
    p.payment_method IS NULL;
`

	initPaymentStatus string = `
INSERT INTO
	payment_statuses (payment_status)
	SELECT ps.payment_status FROM
	(VALUES ('pending'), ('completed')) AS ps(payment_status)
	LEFT JOIN payment_statuses p ON p.payment_status = ps.payment_status
`
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	err := db.AutoMigrate(

		//user tables
		&domain.Users{},
		&domain.UserInfo{},
		&domain.Address{},

		//admin tables
		&domain.Admin{},

		//Product tables
		&domain.ProductCategory{},
		&domain.ProductBrand{},
		&domain.Product{},
		&domain.ProductItem{},
		&domain.Coupon{},

		//cart tables
		&domain.Cart{},
		&domain.CartItems{},

		//wishlist tables
		&domain.Wishlist{},
		&domain.WishlistItem{},

		//Order tables
		&domain.Order{},
		&domain.OrderLine{},
		&domain.PaymentMethod{},
		&domain.OrderStatus{},
		&domain.DeliveryStatus{},
		&domain.Return{},

		//	payment details
		&domain.PaymentStatus{},
		&domain.PaymentDetails{},
	)
	if err != nil {
		return nil, err
	}
	// populate status tables with predefined values
	db.Exec(initDeliveryStatus)
	db.Exec(initOrderStatus)
	db.Exec(initPaymentMethod)
	db.Exec(initPaymentStatus)

	return db, dbErr
}
