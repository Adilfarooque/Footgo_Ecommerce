package repository

import (
	"errors"

	"github.com/Adilfarooque/Footgo_Ecommerce/db"
	"github.com/Adilfarooque/Footgo_Ecommerce/domain"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
	"gorm.io/gorm"
)

func CheckUserExistByEmail(email string) (*domain.User, error) {
	//Declare a variable to hold the user data
	var user domain.User
	//To find a user with the given email
	res := db.DB.Where(&domain.User{Email: email})
	if res.Error != nil {
		//If the error indicates that no records were found, return nil, nil
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		// If there's an error other than "record not found", return the error
		return nil, res.Error
	}
	//If the query is successfull , return a pointer to the user
	return &user, nil
}

func CheckUserExistByPhone(phone string) (*domain.User, error) {
	var user domain.User
	res := db.DB.Where(&domain.User{Phone: phone})
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
	//insert user details and retrieve specific columns
	err := db.DB.Raw("INSERT INTO users(firstname,lastname,email,password,phone)VALUES(?,?,?,?,?)RETURNING id,firstname,lastname,password,phone", user.Firstname, user.Lastname, user.Email, user.Password, user.Phone).Scan(&SignupDetail).Error
	if err != nil {
		//Return an empty UserDetailsResponse and the error if there's an error during execution
		return models.UserDetailsResponse{}, err
	}
	// Return the UserDetailsResponse containing the user's details and nil error if successful
	return SignupDetail, nil
}
