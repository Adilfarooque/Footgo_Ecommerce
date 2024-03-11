package repository

import (
	"errors"

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

func GetCoupon() ([]models.Coupon, error) {
	var coupons []models.Coupon

	err := db.DB.Raw("SELECT id,coupon,discount_percentage,minimum_price,validity FROM coupons").Scan(&coupons).Error
	if err != nil {
		return []models.Coupon{}, err
	}
	return coupons, nil
}

func ExpireCoupon(couponID int) (bool, error) {
	var count int
	err := db.DB.Raw("SELECT COUNT(*) FROM coupons WHERE id = ?", couponID).Scan(&count).Error
	if err != nil {
		return false, errors.New("the offer already exists")
	}
	return count > 0, nil
}

func CouponAlreadyExpired(couponID int) error {
	var valid bool
	err := db.DB.Raw("SELECT validity FROM coupons WHERE id = ?", couponID).Scan(&valid).Error
	if err != nil {
		return err
	}
	if valid {
		err := db.DB.Exec("UPDATE coupons SET validity = false WHERE id = ?", couponID).Error
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("already expired")
}
