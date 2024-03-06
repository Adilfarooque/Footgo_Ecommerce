package handlers

import (
	"net/http"

	"github.com/Adilfarooque/Footgo_Ecommerce/usecase"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/response"
	"github.com/gin-gonic/gin"
)

// @Summary		Get Category
// @Description	Retrieve All Category
// @Tags			Admin Category Management
// @Accept			json
// @Produce		    json
// @Security		Bearer
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/admin/category   [GET]

func GetCategory(c *gin.Context) {
	category, err := usecase.GetCategory()
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Couldn't displayed categories", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Display All Category", category, nil)
	c.JSON(http.StatusOK, success)
}
