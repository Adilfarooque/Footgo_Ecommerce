package handlers

import (
	"net/http"

	"github.com/Adilfarooque/Footgo_Ecommerce/usecase"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Summary Add  a new coupon by Admin
// @Description Add A new Coupon which can be used by the users from the checkout section
// @Tags Admin Coupon Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param coupon body models.AddCoupon true "Add new Coupon"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/coupons [POST]

func AddCoupon(c *gin.Context) {
	var coupon models.AddCoupon
	if err := c.ShouldBindJSON(&coupon); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not bind the coupon details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
	}
	err := validator.New().Struct(coupon)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	message, err := usecase.AddCoupon(coupon)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "Could not add coupon", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusCreated, "Coupon Added", message, nil)
	c.JSON(http.StatusCreated, successRes)
}
