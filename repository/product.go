package repository

import (
	"github.com/Adilfarooque/Footgo_Ecommerce/db"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
)

func ShowAllProducts(page int, count int) ([]models.ProductBreif, error) {
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * count

	var ProductBreif []models.ProductBreif
	err := db.DB.Raw("SELECT * FROM products limit ? offset ? ", count, offset).Scan(&ProductBreif).Error
	if err != nil {
		return nil, err
	}
	return ProductBreif, nil
}

func GetImage(productID int) ([]string, error) {
	var url []string
	if err := db.DB.Raw("SELECT url FROM Images WHERE product_id = ?", productID).Scan(&url).Error; err != nil {
		return url, err
	}
	return url, nil
}
