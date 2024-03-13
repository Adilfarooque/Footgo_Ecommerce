package repository

import (
	"errors"

	"github.com/Adilfarooque/Footgo_Ecommerce/db"
	"github.com/Adilfarooque/Footgo_Ecommerce/domain"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
	"gorm.io/gorm"
)

func CheckUserExistsByEmail(email string) (*domain.User, error) {
	var user domain.User
	res := db.DB.Where(&domain.User{Email: email}).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	return &user, nil
}
func CheckUserExistsByPhone(phone string) (*domain.User, error) {
	var user domain.User
	res := db.DB.Where(&domain.User{Phone: phone}).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	return &user, nil
}
func UserSignUp(user models.UserSignUp) (models.UserDetailsResponse, error) {
	var SignupDetail models.UserDetailsResponse
	err := db.DB.Raw("INSERT INTO users(firstname,lastname,email,password,phone)VALUES(?,?,?,?,?)RETURNING id,firstname,lastname,email,password,phone", user.Firstname, user.Lastname, user.Email, user.Password, user.Phone).Scan(&SignupDetail).Error
	if err != nil {
		return models.UserDetailsResponse{}, err
	}
	return SignupDetail, nil
}

func CreateReferralEntry(userDetails models.UserDetailsResponse, userReferral string) error {

	err := db.DB.Exec("INSERT INTO referrals (user_id,referral_code,referral_amount) VALUES (?,?,?)", userDetails.Id, userReferral, 0).Error
	if err != nil {
		return err
	}

	return nil

}
func GetUserIdFromReferrals(ReferralCode string) (int, error) {

	var referredUserId int
	err := db.DB.Raw("SELECT user_id FROM referrals WHERE referral_code = ?", ReferralCode).Scan(&referredUserId).Error
	if err != nil {
		return 0, nil
	}

	return referredUserId, nil
}

func UpdateReferralAmount(referralAmount float64, referredUserId int, currentUserID int) error {

	err := db.DB.Exec("UPDATE referrals SET referral_amount = ? , referred_user_id = ? WHERE user_id = ? ", referralAmount, referredUserId, currentUserID).Error
	if err != nil {
		return err
	}

	// find the current amount in referred users referral table and add 100 with that
	err = db.DB.Exec("UPDATE referrals SET referral_amount = referral_amount + ? WHERE user_id = ? ", referralAmount, referredUserId).Error
	if err != nil {
		return err
	}

	return nil

}
func AmountInrefferals(userID int) (float64, error) {
	var a float64
	err := db.DB.Raw("SELECT referral_amount FROM referrals WHERE user_id = ?", userID).Scan(&a).Error
	if err != nil {
		return 0.0, err
	}
	return a, nil
}
func ExistWallect(userID int) (bool, error) {
	var count int
	err := db.DB.Raw("SELECT COUNT(*) FROM wallets WHERE user_id = ?", userID).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil

}
func UpdateWallect(amount float64, userID int) error {
	err := db.DB.Exec("UPDATE wallets SET amount = amount + ?  WHERE user_id = ? ", amount, userID).Error
	if err != nil {
		return err
	}

	return nil
}
func UpdateReferUserWallect(amount float64, userID int) error {
	err := db.DB.Exec("UPDATE wallets SET amount = amount + ?  WHERE user_id = ? ", amount, userID).Error
	if err != nil {
		return err
	}

	return nil
}
func NewWallect(userID int, amount float64) error {
	err := db.DB.Exec("INSERT INTO wallets (user_id,amount) VALUES(?,?) ", userID, amount).Error
	if err != nil {
		return err
	}

	return nil
}

func FindUserByEmail(user models.LoginDetail) (models.UserLoginResponse, error) {
	var userDetails models.UserLoginResponse
	err := db.DB.Raw("SELECT * FROM users WHERE email = ? AND blocked=false and isadmin=false", user.Email).Scan(&userDetails).Error
	if err != nil {
		return models.UserLoginResponse{}, err
	}
	return userDetails, nil
}

func GetAllAddress(userId int) ([]models.AddressInfoResponse, error) {
	var AddressInfoResponse []models.AddressInfoResponse
	if err := db.DB.Raw("SELECT * FROM addresses WHERE user_id = ?", userId).Scan(&AddressInfoResponse).Error; err != nil {
		return []models.AddressInfoResponse{}, err
	}
	return AddressInfoResponse, nil
}

func AddAddress(userID int, address models.AddressInfo) error {
	err := db.DB.Exec("INSERT INTO addresses(user_id,name,house_name,street,city,state,pin)VALUES(?,?,?,?,?,?,?)", userID, address.Name, address.HouseName, address.Street, address.City, address.State, address.Pin).Error
	if err != nil {
		return errors.New("could not add address")
	}
	return nil
}

func CheckAddressAvailabilityWithAddressID(addressID, userID int) bool {
	var count int
	if err := db.DB.Raw("SELECT COUNT(*) FROM addresses WHERE id = ? AND user_id = ?", addressID, userID).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func UpdateName(name string, addressID int) error {
	err := db.DB.Exec("UPDATE addresses SET name = ? WHERE id = ?", name, addressID).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateHouseName(HouseName string, addressID int) error {
	err := db.DB.Exec("UPDATE addresses SET  house_name = ? WHERE id = ?", HouseName, addressID).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateStreet(Street string, addressID int) error {
	err := db.DB.Exec("UPDATE addresses SET street = ? WHERE id = ?", Street, addressID).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateCity(city string, addressID int) error {
	if err := db.DB.Exec("UPDATE addresses SET city = ? WHERE id = ?", city, addressID).Error; err != nil {
		return err
	}
	return nil
}

func UpdateState(state string, addressID int) error {
	if err := db.DB.Exec("UPDATE addresses SET state = ? WHERE id = ?", state, addressID).Error; err != nil {
		return err
	}
	return nil
}

func UpdatePin(pin string, addressID int) error {
	if err := db.DB.Exec("UPDATE addresses SET pin = ? WHERE id = ?", pin, addressID).Error; err != nil {
		return err
	}
	return nil
}

func AddressDetails(addressID int) (models.AddressInfoResponse, error) {
	var addressDetails models.AddressInfoResponse
	err := db.DB.Raw("SELECT a.id, a.name, a.house_name, a.street, a.city, a.state, a.pin FROM addresses a WHERE a.id = ?", addressID).Row().Scan(&addressDetails.ID, &addressDetails.Name, &addressDetails.HouseName, &addressDetails.Street, &addressDetails.City, &addressDetails.State, &addressDetails.Pin)
	if err != nil {
		return models.AddressInfoResponse{}, err
	}
	return addressDetails, nil
}

func AddressExistInUserProfile(addressID, userID int) (bool, error) {
	var count int
	err := db.DB.Raw("SELECT COUNT(*) FROM addresses WHERE user_id = $1 AND id = $2", userID, addressID).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func RemoveFromUserProfile(userID, addressID int) error {
	err := db.DB.Exec("DELETE FROM addresses WHERE user_id = ? AND id = ?", userID, addressID).Error
	if err != nil {
		return err
	}
	return nil
}
