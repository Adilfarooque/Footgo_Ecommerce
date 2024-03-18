package usecase

import (
	"errors"

	"github.com/Adilfarooque/Footgo_Ecommerce/repository"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
)

func AddToCart(product_id int, user_id int) (models.CartResponse, error) {
	ok, _, err := repository.CheckProduct(product_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	if !ok {
		return models.CartResponse{}, errors.New("product doesn't exists")
	}
	QuantityOfProductsInCart, err := repository.QuantityOfProductsInCart(user_id, product_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	quantityOfProduct, err := repository.GetQuantityFromProductID(product_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	if quantityOfProduct <= 0 {
		return models.CartResponse{}, errors.New("out of stock")
	}
	if quantityOfProduct == QuantityOfProductsInCart {
		return models.CartResponse{}, errors.New("stock limit exceeded")
	}
	ProductPrice, err := repository.GetPriceOfProductFromID(product_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	discount_percentage, err := repository.FindDiscountPercentageForProduct(product_id)
	if err != nil {
		return models.CartResponse{}, errors.New("there was some error in finding the discounted price")
	}
	var discount float64
	if discount_percentage > 0 {
		discount = (ProductPrice * float64(discount_percentage)) / 100
	}

	Price := ProductPrice - discount
	categoryID, err := repository.FindCategoryID(product_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	discount_percentageCategory, err := repository.FindDiscountPercentageForCategory(categoryID)
	if err != nil {
		return models.CartResponse{}, errors.New("there was some error in finding the discounted prices")
	}
	var discountcategory float64
	if discount_percentageCategory > 0 {
		discountcategory = (ProductPrice * float64(discount_percentageCategory)) / 100
	}

	FinalPrice := Price - discountcategory

	if QuantityOfProductsInCart == 0 {
		err := repository.AddItemInToCart(user_id, product_id, 1, FinalPrice)
		if err != nil {
			return models.CartResponse{}, err
		}
	} else {
		currentTotal, err := repository.TotalPriceForProductInCart(user_id, product_id)
		if err != nil {
			return models.CartResponse{}, err
		}
		err = repository.UpdateCart(QuantityOfProductsInCart+1, currentTotal+ProductPrice, user_id, product_id)
		if err != nil {
			return models.CartResponse{}, err
		}
	}
	cartDetails, err := repository.DisplayCart(user_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	cartTotal, err := repository.GetTotalPrice(user_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	err = repository.ProductStockMinus(product_id, quantityOfProduct)
	if err != nil {
		return models.CartResponse{}, err
	}
	return models.CartResponse{
		UserName:   cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		Cart:       cartDetails,
	}, nil
}
