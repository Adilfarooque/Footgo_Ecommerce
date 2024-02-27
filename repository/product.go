package repository

import (
	"errors"

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

func CheckValidateCategory(data map[string]int) error {
	for _, id := range data {
		var count int
		err := db.DB.Raw("SELECT COUNT(*) FROM categories WHERE id = ? ", id).Scan(&count).Error
		if err != nil {
			return err
		}
		if count < 1 {
			return errors.New("doesn't exist..")
		}
	}
	return nil
}

func GetProductFromCategory(id int) ([]models.ProductBreif, error) {
	var product []models.ProductBreif
	err := db.DB.Raw("SELECT * FROM products JOIN categories ON products.category_id = categories.id WHERE categories.id = ?", id).Scan(&product).Error
	if err != nil {
		return []models.ProductBreif{}, err
	}
	return product, nil
}

func GetQuantityFromProductID(id int) (int, error) {
	var quantity int
	err := db.DB.Raw("SELECT stock FROM products WHERE id = ?", id).Scan(&quantity).Error
	if err != nil {
		return 0.0, err
	}
	return quantity, nil
}
