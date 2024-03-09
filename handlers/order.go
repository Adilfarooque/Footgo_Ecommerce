package handlers

import (
	"net/http"
	"strconv"

	"github.com/Adilfarooque/Footgo_Ecommerce/usecase"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/response"
	"github.com/gin-gonic/gin"
)

// @Summary Get All order details for admin
// @Description Get all order details to the admin side
// @Tags Admin Order Management
// @Accept   json
// @Produce  json
// @Security Bearer
// @Param page query string false "Page number"
// @Param count query string false "Page Count"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/order   [GET]

func GetAllOrderDetailsForAdmin(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	countStr := c.DefaultQuery("count", "20")
	pageSize, err := strconv.Atoi(countStr)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "page not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	allOrderDetails, err := usecase.GetAllOrderDetailsForAdmin(page, pageSize)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Could not retrive order details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Order Details Retrived Successfully", allOrderDetails, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary Approve Order
// @Description Approve Order from admin side which is in processing state
// @Tags Admin Order Management
// @Accept   json
// @Produce  json
// @Security Bearer
// @Param    order_id   query   string   true    "Order ID"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/order/approve [GET]

func ApproveOrder(c *gin.Context) {
	orderId, err := strconv.Atoi(c.Query("order_id"))
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "error from orderID", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	err = usecase.ApproveOrder(orderId)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Couldn't approved the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Order Approved Successfully", nil, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary Cancel Order Admin
// @Description Cancel Order from admin side
// @Tags Admin Order Management
// @Accept   json
// @Produce  json
// @Security Bearer
// @Param order_id query string true "Order ID"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/order/cancel   [GET]
func CancelOrderFromAdmin(c *gin.Context) {
	order_id, err := strconv.Atoi(c.Query("order_id"))
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "error from orderID", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	err = usecase.CancelOrderFromAdmin(order_id)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Couldn't cancel the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Order Cancel Successfully", nil, nil)
	c.JSON(http.StatusOK, success)
}
