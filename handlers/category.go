package handlers

import (
	"net/http"

	"github.com/Adilfarooque/Footgo_Ecommerce/usecase"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
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

// admin
// @Summary		Add Category
// @Description	Admin can add new categories for products
// @Tags			Admin Category Management
// @Accept			json
// @Produce		    json
// @Param			category	body	models.Category	true	"category"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/category [POST]

func AddCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "fields provided are in worng format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	cate, err := usecase.AddCategory(category)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Could not the Category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully added Category", cate, nil)
	c.JSON(http.StatusOK, success)
}

// @Summary		Update Category
// @Description	Admin can update name of a category into new name
// @Tags			Admin Category Management
// @Accept			json
// @Produce		    json
// @Param			set_new_name	body	models.SetNewName	true	"set new name"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/category     [PUT]

func UpdateCategory(c *gin.Context) {
	var categoryUpdate models.SetNewName
	if err := c.ShouldBindJSON(&categoryUpdate); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	ok, err := usecase.UpdateCategory(categoryUpdate.Current, categoryUpdate.New)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "coudn't update the product", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully updated category", ok, nil)
	c.JSON(http.StatusOK, success)
}

