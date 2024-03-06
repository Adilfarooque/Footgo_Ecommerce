package repository

import (
	"github.com/Adilfarooque/Footgo_Ecommerce/db"
	"github.com/Adilfarooque/Footgo_Ecommerce/domain"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
)

func GetCategory() ([]domain.Category, error) {
	var category []domain.Category
	err := db.DB.Raw("SELECT * FROM categories").Scan(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}

func CheckIfCategoryAlreadyExists(category string) (bool, error) {
	var count int64
	err := db.DB.Raw("SELECT COUNT(*) FROM categories WHERE category = $1", category).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func AddCategory(category models.Category) (domain.Category, error) {
	var cate string
	err := db.DB.Raw("INSERT INTO categories (category) VALUES (?) RETURNING category", category.Category).Scan(&cate).Error
	if err != nil {
		return domain.Category{}, err
	}

	var categoriesResponse domain.Category
	err = db.DB.Raw("SELECT id, category FROM categories WHERE category = ?", cate).Scan(&categoriesResponse).Error
	if err != nil {
		return domain.Category{}, err
	}

	return categoriesResponse, nil
}
