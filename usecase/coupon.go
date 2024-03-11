package usecase

import (
	"github.com/Adilfarooque/Footgo_Ecommerce/repository"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
)

func AddCoupon(coupon models.AddCoupon) (string, error) {

	couponExist, err := repository.CouponExist(coupon.Coupon)
	if err != nil {
		return "", err
	}

	if couponExist {
		alreadyValid, err := repository.CouponRevalidateIfExpired(coupon.Coupon)
		if err != nil {
			return "", err
		}
		if alreadyValid {
			return "The coupon which is already exists", nil
		}
		return "Made the coupon valid", nil
	}
	err = repository.AddCoupon(coupon)
	if err != nil {
		return "", err
	}
	return "Successfully added the coupon", nil
}
