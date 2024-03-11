package usecase

import (
	"errors"

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

func GetCoupon() ([]models.Coupon, error) {
	coupons, err := repository.GetCoupon()
	if err != nil {
		return []models.Coupon{}, err
	}
	return coupons, nil
}

func ExpireCoupon(couponID int) error {
	couponExist, err := repository.ExpireCoupon(couponID)
	if err != nil {
		return err
	}
	//if it exists expire it ,if already send back relevant message
	if couponExist {
		err = repository.CouponAlreadyExpired(couponID)
		if err != nil{
			return err
		}
		return nil
	}
	return errors.New("coupon does not exist")
}
