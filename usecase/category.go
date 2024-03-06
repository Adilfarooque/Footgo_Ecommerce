package usecase

import (
	"github.com/Adilfarooque/Footgo_Ecommerce/domain"
	"github.com/Adilfarooque/Footgo_Ecommerce/repository"
)

func GetCategory() ([]domain.Category, error) {
	category, err := repository.GetCategory()
	if err != nil {
		return []domain.Category{}, err
	}
	return category, nil
}
