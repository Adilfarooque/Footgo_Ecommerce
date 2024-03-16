package handlers

import (
	"net/http"
	"strconv"

	"github.com/Adilfarooque/Footgo_Ecommerce/usecase"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/response"
	"github.com/gin-gonic/gin"
)

// @Summary Add to Wishlist
// @Description Add To wish List
// @Tags WishList Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param product_id query string true "product id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}

func AddToWishlist(c *gin.Context) {
	userID, _ := c.Get("user_id")
	product_id := c.Query("product_id")
	productID, err := strconv.Atoi(product_id)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "product ID is in wrong format", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	err = usecase.AddToWishlist(productID, userID.(int))
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "failed to item to the wishlist", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully added product to the wishlist", nil, err.Error())
	c.JSON(http.StatusOK, success)
}
