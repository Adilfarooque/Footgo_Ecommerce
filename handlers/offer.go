package handlers

import (
	"net/http"

	"github.com/Adilfarooque/Footgo_Ecommerce/usecase"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Summary Add  Product Offer
// @Description Add a new Offer for a product by specifying a limit
// @Tags Admin Offer Management
// @Accept json
// @Produce json
// @Security Bearer
// @Param coupon body models.ProductOfferReceiver true "Add new Product Offer"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/productoffer [POST]

func AddProductOffer(c *gin.Context) {
	var productOffer models.ProductOfferReciever
	if err := c.ShouldBindJSON(&productOffer); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "request fields in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(productOffer)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err = usecase.AddProductOffer(productOffer)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "could not add offer", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "Successfully added offer", nil, nil)
	c.JSON(http.StatusCreated, success)

}

// @Summary Show  Product Offer
// @Description Add a new Offer for a product by specifying a limit
// @Tags Admin Offer Management
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /admin/productoffer [GET]

func GetProductOffer(c *gin.Context) {
	categories, err := usecase.Getoffers()
	if err != nil {
		errsRes := response.ClientResponse(http.StatusBadRequest, "fields provided are wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errsRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all offers", categories, nil)
	c.JSON(http.StatusOK, successRes)
}
