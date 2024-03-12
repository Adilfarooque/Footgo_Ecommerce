package usecase

import (
	"github.com/Adilfarooque/Footgo_Ecommerce/repository"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
)

func AddProductOffer(product models.ProductOfferReciever) error {
	if err := repository.AddProductOffer(product); err != nil {
		return err
	}
	return nil
}
