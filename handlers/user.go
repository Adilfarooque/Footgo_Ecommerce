package handlers

import (
	"net/http"
	"strconv"

	"github.com/Adilfarooque/Footgo_Ecommerce/usecase"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Summary		User Signup
// @Description	user can signup by giving their details
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			signup  body  models.UserSignUp  true	"signup"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/user/signup    [POST]
func UserSignup(c *gin.Context) {
	var SignupDetail models.UserSignUp
	if err := c.ShouldBindJSON(&SignupDetail); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(SignupDetail)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not statisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	user, err := usecase.UsersSignUp(SignupDetail)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "User successfully signed up", user, nil)
	c.JSON(http.StatusCreated, success)
}

// @Summary		User Login
// @Description	user can log in by giving their details
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			login  body  models.LoginDetail  true	"login"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/user/userlogin     [POST]
func Userlogin(c *gin.Context) {
	var UserLoginDetail models.LoginDetail
	if err := c.ShouldBindJSON(&UserLoginDetail); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(UserLoginDetail)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	user, err := usecase.UserLogin(UserLoginDetail)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "field provided are in worng format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "User successfully logged in which password", user, nil)
	c.JSON(http.StatusCreated, success)
}

// @Summary		Get Addresses
// @Description	user can get all their addresses
// @Tags			User Profile
// @Accept          json
// @Produce         json
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router		/user/address       [GET]

func GetAllAddress(c *gin.Context) {
	userID, _ := c.Get("user_id")
	addressInfo, err := usecase.GetAllAddress(userID.(int))
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "filed to retirve details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "User Address", addressInfo, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary		Add Address
// @Description	user can add their addresses
// @Tags			User Profile
// @Accept			json
// @Produce		    json
// @Param			address  body  models.AddressInfo  true	"address"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/user/address    [POST]

func AddAddress(c *gin.Context) {
	user_ID, _ := c.Get("user_id")
	var address models.AddressInfo
	if err := c.ShouldBindJSON(&address); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "field provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(&address)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "constraints does not match", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = usecase.AddAddress(user_ID.(int), address)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "failed adding address", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Address added successfully", nil, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary Update User Address
// @Description Update User address by sending in address id
// @Tags User Profile
// @Accept json
// @Produce json
// @Security Bearer
// @Param address_id query string true "address id"
// @Param address body models.AddressInfo true "User Address Input"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /user/address    [PUT]

func UpdateAddress(c *gin.Context) {
	user_id, _ := c.Get("user_id")
	addressid := c.Query("address_id")
	addressID, _ := strconv.Atoi(addressid)
	var address models.AddressInfo
	if err := c.ShouldBindJSON(&address); err != nil {
		errs := response.ClientResponse(http.StatusBadGateway, "field provided are wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	UpdateAddress, err := usecase.UpdateAddress(address, addressID, user_id.(int))
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "failed to update user address", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Updated user Address", UpdateAddress, nil)
	c.JSON(http.StatusOK, success)

}

// @Summary Delete User Address
// @Description Delete From User Profile
// @Tags User Profile
// @Accept json
// @Produce json
// @Security Bearer
// @Param address_id query string true "address id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /user/address    [DELETE]

func DeleteAddressByID(c *gin.Context) {
	user_id, _ := c.Get("user_id")
	addressid := c.Query("address_id")
	addressID, _ := strconv.Atoi(addressid)
	err := usecase.DeleteAddress(addressID, user_id.(int))
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "failed to delete user address", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Delete User Address", nil, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary User Details
// @Description User Details from User Profile
// @Tags User Profile
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /user/users   [GET]

func UserDetails(c *gin.Context) {
	userID, _ := c.Get("user_id")
	UserDetails, err := usecase.UserDetails(userID.(int))
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "failed to retrieve details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "User Details", UserDetails, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary Update User Details
// @Description Update User Details by sending in user id
// @Tags User Profile
// @Accept json
// @Produce json
// @Security Bearer
// @Param address body models.UsersProfileDetails true "User Details Input"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /user/users [PUT]

func UpdateUserDetails(c *gin.Context) {
	user_id, _ := c.Get("user_id")
	var user models.UsersProfileDetails
	if err := c.ShouldBindJSON(&user); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	UpdateUserDetails, err := usecase.UpdateUserDetails(user, user_id.(int))
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "failed update user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Updated User Details", UpdateUserDetails, err.Error())

	c.JSON(http.StatusOK, success)
}

// @Summary Change User Password
// @Description Change User Password
// @Tags User Profile
// @Accept json
// @Produce json
// @Security Bearer
// @Param body body models.ChangePassword true "User Password Change"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /user/users/changepassword     [PUT]

func ChangePassword(c *gin.Context) {
	user_id, _ := c.Get("user_id")
	var changePassword models.ChangePassword
	if err := c.ShouldBindJSON(&changePassword); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	if err := usecase.ChangePassword(user_id.(int), changePassword.Oldpassword, changePassword.Password, changePassword.Repassword); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Could not change the password", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Password changed successfully", nil, nil)
	c.JSON(http.StatusOK, success)
}
