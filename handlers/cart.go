package handlers

import (
	"net/http"
	"strconv"

	"github.com/Adilfarooque/Footgo_Ecommerce/usecase"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/response"
	"github.com/gin-gonic/gin"
)

// @Summary		Add To Cart
// @Description	Add products to carts  for the purchase
// @Tags			User Cart Management
// @Accept			json
// @Produce		    json
// @Param			product_id	query		string	true	"product-id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/user/cart  [post]

func AddToCart(c *gin.Context) {
	id := c.Query("product_id")
	product_id, err := strconv.Atoi(id)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadGateway, "Product id is given in the wrong format", nil, err.Error())
		c.JSON(http.StatusBadGateway, errs)
		return
	}
	user_ID, _ := c.Get("user_id")
	cartResponse, err := usecase.AddToCart(product_id, user_ID.(int))
	if err != nil {
		errs := response.ClientResponse(http.StatusBadGateway, "could not add product to the cart", nil, err.Error())
		c.JSON(http.StatusBadGateway, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Added product Successfully to the cart", cartResponse, nil)
	c.JSON(http.StatusOK, success)

}
