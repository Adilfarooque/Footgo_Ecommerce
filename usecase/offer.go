package usecase

import (
	"github.com/Adilfarooque/Footgo_Ecommerce/domain"
	"github.com/Adilfarooque/Footgo_Ecommerce/repository"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
)

func AddProductOffer(product models.ProductOfferReciever) error {
	if err := repository.AddProductOffer(product); err != nil {
		return err
	}
	return nil
}

func Getoffers() ([]domain.ProductOffer, error) {
	offers, err := repository.Getoffers()
	if err != nil {
		return []domain.ProductOffer{}, err
	}
	return offers, nil
}

func AddCategoryOffer(model models.CategoryOfferReceiver) error {
	if err := repository.AddCategoryOffer(model); err != nil {
		return err
	}
	return nil
}
