package repository

import (
	"github.com/Adilfarooque/Footgo_Ecommerce/db"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
)

func ShowImages(productID int) ([]models.Image, error) {
	var image []models.Image
	err := db.DB.Raw("SELECT url FROM images WHERE images.product_id = $1", productID).Scan(&image).Error
	if err != nil {
		return nil, err
	}
	return image, nil
}
