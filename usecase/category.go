package usecase

import (
	"errors"

	"github.com/Adilfarooque/Footgo_Ecommerce/domain"
	"github.com/Adilfarooque/Footgo_Ecommerce/repository"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
)

func GetCategory() ([]domain.Category, error) {
	category, err := repository.GetCategory()
	if err != nil {
		return []domain.Category{}, err
	}
	return category, nil
}

func AddCategory(category models.Category) (domain.Category, error) {
	exists, err := repository.CheckIfCategoryAlreadyExists(category.Category)
	if err != nil {
		return domain.Category{}, err
	}
	if exists {
		return domain.Category{}, errors.New("category already exists")
	}
	categories, err := repository.AddCategory(category)
	if err != nil {
		return domain.Category{}, err
	}
	return categories, nil
}
