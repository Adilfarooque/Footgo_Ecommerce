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

// @Summary		Remove From Cart
// @Description	Remove products to carts  for the purchase
// @Tags			User Cart Management
// @Accept			json
// @Produce		    json
// @Param			id	query		string	true	"product-id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/user/cart    [DELETE]

func RemoveFromCart(c *gin.Context) {
	id := c.Query("id")
	product_id, err := strconv.Atoi(id)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "product not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	user_ID, _ := c.Get("user_id")
	updateCart, err := usecase.RemoveFromCart(product_id, user_ID.(int))
	if err != nil {
		errs := response.ClientResponse(http.StatusBadGateway, "can't remove product from cart", nil, err.Error())
		c.JSON(http.StatusBadGateway, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "product removed successfully", updateCart, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary		Display Cart
// @Description	Display products to carts
// @Tags			User Cart Management
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/user/cart  [GET]

func DisplayCart(c *gin.Context) {
	userID, _ := c.Get("user_id")
	cart, err := usecase.DisplayCart(userID.(int))
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "cannot display cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Cart items display successfully", cart, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary		Empty Cart
// @Description	Empty products to carts  for the purchase
// @Tags			User Cart Management
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/user/cart/empty   [DELETE]

func EmptyCart(c *gin.Context) {
	userID, _ := c.Get("user_id")
	cart, err := usecase.EmptyCart(userID.(int))
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "cannot empty the cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Cart emptied Successfully", cart, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary		Add quantity in cart by one
// @Description	user can add 1 quantity of product to their cart
// @Tags			User Cart Management
// @Accept			json
// @Produce		    json
// @Param			product_id	query	string	true	"product_id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/user/cart/updatequantityadd   [PUT]

func UpdateQuantityAdd(c *gin.Context) {
	id, _ := c.Get("user_id")
	productID, err := strconv.Atoi(c.Query("product_id"))
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "check product id parameter properly", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	if err := usecase.UpdateQuantityAdd(id.(int), productID); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "could not add quantity", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully added quantity", nil, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary		Subtract quantity in cart by one
// @Description	user can subtract 1 quantity of product from their cart
// @Tags			User Cart Management
// @Accept			json
// @Produce		    json
// @Param			product_id	query	string	true	"product_id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/user/cart/updatequantityless     [PUT]

//Decreases the quantity of a product in a user’s cart
func UpdateQuantityLess(c *gin.Context){
	id , _ := c.Get("user_id")
	productID,err := strconv.Atoi(c.Query("product_id"))
	if err != nil{
		errs := response.ClientResponse(http.StatusBadRequest,"check parameter properly",nil,err.Error())
		c.JSON(http.StatusBadRequest,errs)
		return
	}
	if err := usecase.UpdateQuantityLess(id.(int),productID); err != nil{
		errs := response.ClientResponse(http.StatusBadRequest,"could not less the quantity",nil,err.Error())
		c.JSON(http.StatusBadRequest,errs)
		return
	}
	success := response.ClientResponse(http.StatusOK,"Successfully less quantity",nil,nil)
	c.JSON(http.StatusOK,success)
}
