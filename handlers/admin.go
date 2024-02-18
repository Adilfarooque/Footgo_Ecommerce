package handlers

import (
	"net/http"

	"github.com/Adilfarooque/Footgo_Ecommerce/usecase"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/response"
	"github.com/gin-gonic/gin"
)

// @Summary		Admin Login
// @Description	Login handler for jerseyhub admins
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			admin	body		models.AdminLogin	true	"Admin login details"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/admin/adminlogin [POST]
func LoginHandler(c *gin.Context) {
	var adminDetails models.AdminLogin
	if err := c.ShouldBindJSON(&adminDetails); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Detials not correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	admin, err := usecase.LoginHandler(adminDetails)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Cannot authenicate user", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Admin authenticate successfully", admin, nil)
	c.JSON(http.StatusOK, success)
}
