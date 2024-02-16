package handlers

import (
	"net/http"

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
	// Declare a variable to hold the user sign-up details
	var SignupDetails models.UserSignUp
	// Bind JSON data from request body into signupDetail struct
	if err := c.ShouldBindJSON(&SignupDetails); err != nil {
		// Return a client response with status code 400 (Bad Request) and an error message
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	// Validate signupDetail struct
	err := validator.New().Struct(SignupDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	//Call usecase to sign up the user
	user, err := usecase.UsersSignUp(SignupDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	//Return a client response with status code 201 (Created) and a success message
	success := response.ClientResponse(http.StatusCreated, "User successfully signed up", user, err.Error())
	c.JSON(http.StatusCreated, success)
}
