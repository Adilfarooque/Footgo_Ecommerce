package repository

import (
	"github.com/Adilfarooque/Footgo_Ecommerce/db"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
)

func CouponExist(couponName string) (bool, error) {
	var count int
	err := db.DB.Raw("SELECT COUNT(*) FROM coupons WHERE coupon = ?", couponName).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func CouponRevalidateIfExpired(couponName string) (bool, error) {
	var isValid bool
	err := db.DB.Raw("SELECT validty FROM coupons WHERE coupon = ?", couponName).Scan(&isValid).Error
	if err != nil {
		return false, err
	}

	if isValid {
		return true, nil
	}

	err = db.DB.Exec("UPDATE coupons SET validity = true WHERE coupon = ?", couponName).Error
	if err != nil {
		return false, err
	}
	return false, nil

}

func AddCoupon(coupon models.AddCoupon) error {
	err := db.DB.Exec("INSERT INTO coupons (coupon,discount_percentage,minimum_price,validity) VALUES (?, ?, ?, ?)", coupon.Coupon, coupon.DiscountPercentage, coupon.MinimumPrice, true).Error
	if err != nil {
		return nil
	}
	return nil
}
