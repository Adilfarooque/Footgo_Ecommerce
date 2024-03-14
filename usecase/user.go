package usecase

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Adilfarooque/Footgo_Ecommerce/helper"
	"github.com/Adilfarooque/Footgo_Ecommerce/repository"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

func UsersSignUp(user models.UserSignUp) (*models.TokenUser, error) {
	email, err := repository.CheckUserExistsByEmail(user.Email)
	fmt.Println(email)
	if err != nil {
		return &models.TokenUser{}, errors.New("error with server")
	}
	if email != nil {
		return &models.TokenUser{}, errors.New("user with this email is already exists")
	}

	phone, err := repository.CheckUserExistsByPhone(user.Phone)
	fmt.Println(phone, nil)
	if err != nil {
		return &models.TokenUser{}, errors.New("error with server")
	}
	if phone != nil {
		return &models.TokenUser{}, errors.New("user with this phone is already exists")
	}

	hashPassword, err := helper.PasswordHash(user.Password)
	if err != nil {
		return &models.TokenUser{}, errors.New("error in hashing password")
	}
	user.Password = hashPassword
	userData, err := repository.UserSignUp(user)
	if err != nil {
		return &models.TokenUser{}, errors.New("could not add the user")
	}
	// create referral code for the user and send in details of referred id of user if it exist
	id := uuid.New().ID()
	str := strconv.Itoa(int(id))
	userReferral := str[:8]
	err = repository.CreateReferralEntry(userData, userReferral)
	if err != nil {
		return &models.TokenUser{}, err
	}

	if user.ReferralCode != "" {
		// first check whether if a user with that referralCode exist
		referredUserId, err := repository.GetUserIdFromReferrals(user.ReferralCode)
		if err != nil {
			return &models.TokenUser{}, err
		}
		if referredUserId != 0 {
			referralAmount := 150
			err := repository.UpdateReferralAmount(float64(referralAmount), referredUserId, userData.Id)
			if err != nil {
				return &models.TokenUser{}, err
			}
			referreason := "Amount credited for used referral code"
			err = repository.UpdateHistory(userData.Id, 0, float64(referralAmount), referreason)
			if err != nil {
				return &models.TokenUser{}, err
			}
			amount, err := repository.AmountInrefferals(userData.Id)
			if err != nil {
				return &models.TokenUser{}, err
			}
			wallectExist, err := repository.ExistWallect(userData.Id)
			if err != nil {
				return &models.TokenUser{}, err
			}
			if !wallectExist {
				err = repository.NewWallect(userData.Id, amount)
				if err != nil {
					return &models.TokenUser{}, err
				}
			}
			err = repository.UpdateReferUserWallect(amount, referredUserId)
			if err != nil {
				return &models.TokenUser{}, err
			}
			reason := "Amount credited for refer a new person"
			err = repository.UpdateHistory(referredUserId, 0, amount, reason)
			if err != nil {
				return &models.TokenUser{}, err
			}
		}
	}

	accessToken, err := helper.GenerateAccessToken(userData)
	if err != nil {
		return &models.TokenUser{}, errors.New("couldn't create access token due to error")
	}
	refreshToken, err := helper.GenerateRefreshToken(userData)
	if err != nil {
		return &models.TokenUser{}, errors.New("couldn't create refresh token due to error")
	}
	return &models.TokenUser{
		Users:        userData,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func UserLogin(user models.LoginDetail) (*models.TokenUser, error) {
	email, err := repository.CheckUserExistsByEmail(user.Email)
	if err != nil {
		return &models.TokenUser{}, errors.New("error with server")
	}
	if email == nil {
		return &models.TokenUser{}, errors.New("email doesn't exist")
	}
	userdetails, err := repository.FindUserByEmail(user)
	if err != nil {
		return &models.TokenUser{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userdetails.Password), []byte(user.Password))
	if err != nil {
		return &models.TokenUser{}, err
	}
	var userDetails models.UserDetailsResponse
	err = copier.Copy(&userDetails, &userdetails)
	if err != nil {
		return &models.TokenUser{}, err
	}
	accessToken, err := helper.GenerateAccessToken(userDetails)
	if err != nil {
		return &models.TokenUser{}, errors.New("couldn't create accesstoken due to internal error")
	}
	refreshToken, err := helper.GenerateRefreshToken(userDetails)
	if err != nil {
		return &models.TokenUser{}, errors.New("couldn't create refreshtoken due to internal error")
	}
	return &models.TokenUser{
		Users:        userDetails,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func GetAllAddress(userId int) ([]models.AddressInfoResponse, error) {
	addressinfo, err := repository.GetAllAddress(userId)
	if err != nil {
		return []models.AddressInfoResponse{}, err
	}
	return addressinfo, nil
}

func AddAddress(userID int, address models.AddressInfo) error {
	err := repository.AddAddress(userID, address)
	if err != nil {
		return err
	}
	return nil
}

func UpdateAddress(addressDetails models.AddressInfo, addressID, userID int) (models.AddressInfoResponse, error) {
	addressExist := repository.CheckAddressAvailabilityWithAddressID(addressID, userID)
	if !addressExist {
		return models.AddressInfoResponse{}, errors.New("address doesn't exist")
	}
	if addressDetails.Name != "" {
		repository.UpdateName(addressDetails.Name, addressID)
	}
	if addressDetails.HouseName != "" {
		repository.UpdateHouseName(addressDetails.HouseName, addressID)
	}
	if addressDetails.Street != "" {
		repository.UpdateStreet(addressDetails.Street, addressID)
	}
	if addressDetails.City != "" {
		repository.UpdateCity(addressDetails.City, addressID)
	}
	if addressDetails.State != "" {
		repository.UpdateState(addressDetails.State, addressID)
	}
	if addressDetails.Pin != "" {
		repository.UpdatePin(addressDetails.Pin, addressID)
	}
	return repository.AddressDetails(addressID)

}

func DeleteAddress(addressID, userID int) error {
	addressExist, err := repository.AddressExistInUserProfile(addressID, userID)
	if err != nil {
		return err
	}
	if !addressExist {
		return errors.New("address does not exist in user profile")
	}
	err = repository.RemoveFromUserProfile(userID, addressID)
	if err != nil {
		return err
	}
	return nil
}

func UserDetails(userID int) (models.UsersProfileDetails, error) {
	return repository.UserDetails(userID)
}

func UpdateUserDetails(userDetails models.UsersProfileDetails,userID int)(models.UsersProfileDetails,error){
	userExist := repository.CheckUserAvailabilityWithID(userID)
	if !userExist {
		return models.UsersProfileDetails{},errors.New("user doesn't exist")
	}
	if userDetails.Email != ""{
		 repository.UpdateUserEmail(userDetails.Email,userID)
	}
	if userDetails.Firstname != ""{
		 repository.UpdateUserFirstname(userDetails.Firstname,userID)
	}
	if userDetails.Lastname != ""{
		repository.UpdateUserLastname(userDetails.Lastname,userID)
	}
	if userDetails.Phone != ""{
		repository.UpdateUserPhone(userDetails.Phone,userID)
	}
	return repository.UserDetails(userID)
}

func ChangePassword(id int, old string, password string, repassword string) error {
	userPassword, err := repository.GetPassword(id)
	if err != nil {
		return errors.New("Internal error")
	}
	if err = helper.CompareHashAndPassword(userPassword, old); err != nil {
		return errors.New("Passwod incorrect")
	}
	if password != repassword {
		return errors.New("Passwod doesn't match")
	}
	newPassword, err := helper.PasswordHash(password)
	if err != nil {
		return errors.New("error in hashing password")
	}
	return repository.ChangePassword(id, string(newPassword))

}