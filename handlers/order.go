package handlers

import (
	"fmt"
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

// Get All order details for admin
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

// Approve Order
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

// @Summary Get Order Details to user side
// @Description Get all order details done by user to user side
// @Tags User Order Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param page query string false "Page"
// @Param count query string false "Count"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /user/order   [GET]

// Get Order Details to user side
func GetOrderDetails(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "page number is not correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("count", "10"))
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "page count not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	id, _ := c.Get("user_id")
	UserID := id.(int)
	OrderDetails, err := usecase.GetOrderDetails(UserID, page, pageSize)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Could not do the order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Full Order Details", OrderDetails, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary Cancel order
// @Description Cancel order by the user using order ID
// @Tags User Order Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param id query string true "Order ID"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /user/order   [PUT]

// Cancel order
func CancelOrder(c *gin.Context) {
	orderID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "error from orderID", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	id, _ := c.Get("user_id")
	userID := id.(int)
	if err := usecase.CancelOrders(orderID, userID); err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Could not cancel the order", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Cancel Successfull", nil, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary		Checkout section
// @Description	Add products to carts  for the purchase
// @Tags			User Order Management
// @Accept			json
// @Produce		    json
// @Param    order_id    query    int    true    "address id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/user/order/place-order     [GET]

// Checkout section
func PlaceOrderCOD(c *gin.Context) {
	order_id, err := strconv.Atoi(c.Query("order_id"))
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "error from orderID", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	paymentMethod, err := usecase.PaymentMethodID(order_id)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "error from paymentId", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	if paymentMethod == 1 {
		if err := usecase.ExecutePurchaseCOD(order_id); err != nil {
			errs := response.ClientResponse(http.StatusInternalServerError, "error in cash on delivery", nil, err.Error())
			c.JSON(http.StatusBadRequest, errs)
			return
		}
		success := response.ClientResponse(http.StatusOK, "Placed Order with cash on delivery", nil, nil)
		c.JSON(http.StatusOK, success)
	}
	if paymentMethod == 2 {
		link := fmt.Sprintf("http://localhost:8000/user/razorpay?order_id=%d", order_id)
		success := response.ClientResponse(http.StatusOK, "Placed order with razor pay following link", link, nil)
		c.JSON(http.StatusOK, success)
	}
}

// @Summary		Checkout section
// @Description	Add products to carts  for the purchase
// @Tags			User Order Management
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/user/checkout    [GET]

func CheckOut(c *gin.Context) {
	userID, _ := c.Get("user_id")
	checkoutDetails, err := usecase.CheckOut(userID.(int))
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "failed to retrive details", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Checkout page loaded successfully", checkoutDetails, nil)
	c.JSON(http.StatusOK, success)
}
