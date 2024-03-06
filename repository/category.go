package repository

import (
	"github.com/Adilfarooque/Footgo_Ecommerce/db"
	"github.com/Adilfarooque/Footgo_Ecommerce/domain"
)

func GetCategory() ([]domain.Category, error) {
	var category []domain.Category
	err := db.DB.Raw("SELECT * FROM categories").Scan(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}
