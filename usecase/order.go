package usecase

import (
	"github.com/Adilfarooque/Footgo_Ecommerce/repository"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
)

func GetAllOrderDetailsForAdmin(page, pageSize int) ([]models.CombainedOrderDetails, error) {
	orderDetails, err := repository.GetAllOrderDetailsBrief(page, pageSize)
	if err != nil {
		return []models.CombainedOrderDetails{}, err
	}
	return orderDetails, nil
}

func GetOrderDetails(userId int, page int, count int) ([]models.FullOrderDetails, error) {
	fullOrderDetails, err := repository.GetOrderDetails(userId, page, count)
	if err != nil {
		return []models.FullOrderDetails{}, err
	}
	return fullOrderDetails, nil
}

