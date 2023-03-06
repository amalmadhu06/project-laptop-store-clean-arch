package db

import (
	"fmt"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/config"
	"github.com/amalmadhu06/project-laptop-store-clean-arch/pkg/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

		//admin tables
		&domain.Admin{},

		//Product tables
		&domain.ProductCategory{},
		&domain.ProductBrand{},
		&domain.Product{},
		&domain.ProductItem{},
		//&domain.Variation{},
		//&domain.VariationOption{},
		//&domain.ProductConfiguration{},

		//cart tables
		//Order tables
	)
	if err != nil {
		return nil, err
	}

	return db, dbErr
}
