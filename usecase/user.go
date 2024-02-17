package usecase

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Adilfarooque/Footgo_Ecommerce/helper"
	"github.com/Adilfarooque/Footgo_Ecommerce/repository"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
	"github.com/google/uuid"
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
