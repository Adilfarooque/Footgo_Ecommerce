package db

import (
	"fmt"

	"github.com/Adilfarooque/Footgo_Ecommerce/config"
	"github.com/Adilfarooque/Footgo_Ecommerce/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(confg config.Config) (*gorm.DB, error) {
	connectTo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", confg.DBHost, confg.DBUser, confg.DBName, confg.DBPort, confg.DBPassword)
	db, err := gorm.Open(postgres.Open(connectTo), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database:%w", err)
	}
	DB = db
	db.AutoMigrate(&domain.Admin{})
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Product{})
	db.AutoMigrate(&domain.Image{})
	db.AutoMigrate(&domain.Orders{})
	db.AutoMigrate(&domain.PaymentMethod{})
	db.AutoMigrate(&domain.Coupons{})
	db.AutoMigrate(&domain.ProductOffer{})
	db.AutoMigrate(&domain.CategoryOffer{})
	//db.AutoMigrate(&domain.Referral{})
	db.AutoMigrate(&domain.Address{})
	db.AutoMigrate(&domain.WishList{})
	db.AutoMigrate(&domain.Cart{})
	db.AutoMigrate(&domain.Orders{})
	db.AutoMigrate(&domain.WalletHistory{})
	db.AutoMigrate(&domain.Wallet{})

	return DB, err
}
