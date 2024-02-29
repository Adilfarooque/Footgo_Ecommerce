package handlers

import (
	"net/http"
	"strconv"
	"time"

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

// @Summary		Admin Dashboard
// @Description	Retrieve admin dashboard
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/admin/dashboard [GET]
func DashBoard(c *gin.Context) {
	adminDashboard, err := usecase.DashBoard()
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Dashboard could not be displayed", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Admin dashboard displayed", adminDashboard, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary Filtered Sales Report
// @Description Get Filtered sales report by week, month and year
// @Tags Admin
// @Accept json
// @Produce json
// @Security Bearer
// @Param period query string true "sales report"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/sales-report    [GET]

func FilteredSalesReport(c *gin.Context) {
	timePeriod := c.Query("period")
	salesreport, err := usecase.FilteredSalesReport(timePeriod)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "sales report could not be retrieved", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	success := response.ClientResponse(http.StatusOK, "sales report retrieved successfully", salesreport, nil)
	c.JSON(http.StatusOK, success)
}

//	@Summary		Sales report by date
//	@Description	Showing the sales report with respect to the given date
//	@Tags			Admin
//	@Accept			json
//	@Produce		json
// @Security        Bearer
//	@Param			start	query	string		true	"start date DD-MM-YYYY"
//	@Param			end		query	string		true	"end   date DD-MM-YYYY"
//	@Success		200		body	entity.SalesReport	"report"
//	@Router			/admin/sales-report-date   [GET]

func SalesReportByDate(c *gin.Context) {
	startDateStr := c.Query("start")
	endDateStr := c.Query("end")
	if startDateStr == "" || endDateStr == "" {
		err := response.ClientResponse(http.StatusBadRequest, "start or end date is empty", nil, "Empty date string")
		c.JSON(http.StatusBadRequest, err)
		return
	}
	startDate, err := time.Parse("2-1-2006", startDateStr)
	if err != nil {
		err := response.ClientResponse(http.StatusBadRequest, "start date conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	endDate, err := time.Parse("2-1-2006", endDateStr)
	if err != nil {
		err := response.ClientResponse(http.StatusBadRequest, "end date conversion failed", nil, err.Error())
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if startDate.After(endDate) {
		err := response.ClientResponse(http.StatusBadRequest, "start date is after end date", nil, "Invalid date range")
		c.JSON(http.StatusBadRequest, err)
		return
	}

	report, err := usecase.ExicuteSalesReportByDate(startDate, endDate)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "sales report could not be retrieved", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	success := response.ClientResponse(http.StatusOK, "sales report retrieved successfully", report, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary		Get Users
// @Description	Retrieve users with pagination
// @Tags			Admin User Management
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Param page query string false "Page number"
// @Param count query string false "Page size"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/admin/users   [GET]

func GetUsers(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "page number in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	countStr := c.DefaultQuery("count", "10")
	pageSize, err := strconv.Atoi(countStr)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "user count in a page not in right format", nil, err.Error())
		c.JSON(http.StatusBadGateway, errRes)
		return
	}

	users, err := usecase.ShowAllUsers(page, pageSize)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "couldn't retrieve users", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully Retrieved all Users", users, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary		Block User
// @Description	using this handler admins can block an user
// @Tags			Admin User Management
// @Accept			json
// @Produce		json
// @Security		Bearer
// @Param			id	query		string	true	"user-id"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/users/block   [PUT]

func BlockedUser(c *gin.Context) {
	id := c.Query("id")
	err := usecase.BlockedUser(id)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "user could not be blocked", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully blocked the user", nil, nil)
	c.JSON(http.StatusOK, success)
}
