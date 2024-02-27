package handlers

import (
	"net/http"
	"strconv"

	"github.com/Adilfarooque/Footgo_Ecommerce/usecase"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/response"
	"github.com/gin-gonic/gin"
)

// @Summary Get Products Details to users
// @Description Retrieve product images
// @Tags User Product
// @Accept json
// @Produce json
// @Param product_id query string true "product_id"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /user/products/image  [GET]
/*
func ShowImages(c *gin.Context) {
	product_id := c.Query("product_id")
	productID, err := strconv.Atoi(product_id)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "error in stirng conversion", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	image, err := usecase.ShowImages(productID)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadGateway, "couldn't retrive images", nil, err.Error())
		c.JSON(http.StatusBadGateway, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully retrive image", image, err.Error())
	c.JSON(http.StatusOK, success)
}
*/
func ShowImages(c *gin.Context) {
	// Retrieve the product_id query parameter from the request
	product_id := c.Query("product_id")
	if product_id == "" {
		// If product_id is empty, return an error response indicating that it is required
		errs := response.ClientResponse(http.StatusBadRequest, "product_id is required", nil, nil)
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	// Convert the product_id string to an integer
	productID, err := strconv.Atoi(product_id)
	if err != nil {
		// If there's an error in string conversion, return an error response
		errs := response.ClientResponse(http.StatusInternalServerError, "error in string conversion", nil, err)
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	// Call the ShowImages use case function to retrieve images for the specified product ID
	image, err := usecase.ShowImages(productID)
	if err != nil {
		// If there's an error retrieving images, return an error response
		errs := response.ClientResponse(http.StatusBadGateway, "could not retrieve images", nil, err.Error())
		c.JSON(http.StatusBadGateway, errs)
		return
	}
	// Return a success response with the retrieved images
	success := response.ClientResponse(http.StatusOK, "Successfully retrieve images", image, nil)
	c.JSON(http.StatusOK, success)
}
