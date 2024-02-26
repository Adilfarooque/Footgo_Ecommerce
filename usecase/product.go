package usecase

import (
	"errors"

	"github.com/Adilfarooque/Footgo_Ecommerce/repository"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
)

func ShowAllProducts(page int, count int) ([]models.ProductBreif, error) {
	productDetails, err := repository.ShowAllProducts(page, count)
	if err != nil {
		return []models.ProductBreif{}, err
	}

	for i := range productDetails {
		p := &productDetails[i]
		if p.Stock <= 0 {
			p.ProductStatus = "out of stock"
		} else {
			p.ProductStatus = "in stock"
		}
	}
	//Loop inside products and then calculate discounted price of each then return
	for j := range productDetails {
		discount_percentage, err := repository.FindDiscountPercentageForProduct(int(productDetails[j].ID))
		if err != nil {
			return []models.ProductBreif{}, errors.New("there was some error in finding the discounted prices")
		}
		var discount float64
		if discount_percentage > 0 {
			discount = (productDetails[j].Price * float64(discount_percentage)) / 100
		}
		productDetails[j].DiscountedPrice = productDetails[j].Price - discount

		discount_percentageCategory, err := repository.FindDiscountPercentageForCategory(int(productDetails[j].CatergoryID))
		if err != nil {
			return []models.ProductBreif{}, errors.New("there was some error in finding the discounted prices")
		}
		var categorydiscount float64
		if discount_percentageCategory > 0 {
			categorydiscount = (productDetails[j].Price * float64(discount_percentageCategory)) / 100
		}

		productDetails[j].DiscountedPrice = productDetails[j].DiscountedPrice - categorydiscount
	}
	var updatedproductDetails []models.ProductBreif
	for _, p := range productDetails {
		img, err := repository.GetImage(int(p.ID))
		if err != nil {
			return nil, err
		}
		p.Image = img
		updatedproductDetails = append(updatedproductDetails, p)
	}
	return updatedproductDetails, nil
}
