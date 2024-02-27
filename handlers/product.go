package handlers

import (
	"net/http"
	"strconv"

	"github.com/Adilfarooque/Footgo_Ecommerce/usecase"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/response"
	"github.com/gin-gonic/gin"
)

// @Summary Get Products Details to users
// @Description Retrieve all product Details with pagination to users
// @Tags User Product
// @Accept json
// @Produce json
// @Param page query string false "Page number"
// @Param count query string false "Page Count"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /user/products     [GET]
func ShowAllProducts(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number no in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	countStr := c.DefaultQuery("count", "10")
	count, err := strconv.Atoi(countStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page count not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	products, err := usecase.ShowAllProducts(page, count)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Couldn't retrieve products", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully Retrieved all products", products, nil)
	c.JSON(http.StatusOK, success)

}

// @Summary Show Products of specified category
// @Description Show all the Products belonging to a specified category
// @Tags User Product
// @Accept json
// @Produce json
// @Param data body map[string]int true "Category IDs and quantities"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /user/products/filter [POST]

func FilterCategory(c *gin.Context) {
	var data map[string]int
	if err := c.ShouldBindJSON(&data); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	productCategory, err := usecase.FilterCategory(data)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Couldn't retrieve products by category", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully filtered the category", productCategory, nil)
	c.JSON(http.StatusOK, success)
}
