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

func RemoveFromCart(product_id, user_id int) (models.CartResponse, error) {
	// Check if the product exists in the user's cart
	ok, err := repository.ProductExist(user_id, product_id)
	if err != nil {
		// If there's an error, return an empty CartResponse struct and the error
		return models.CartResponse{}, err
	}
	if !ok {
		// If the product does not exist, return an error
		return models.CartResponse{}, errors.New("product doesn't exist in the cart")
	}

	// Struct to store the quantity and total price of the product
	var cartDetails struct {
		Quantity   int
		TotalPrice float64
	}

	// Get the quantity and total price of the product from the cart
	cartDetails, err = repository.GetQuantityAndProductFromID(user_id, product_id, cartDetails)
	if err != nil {
		// If there's an error, return an empty CartResponse struct and the error
		return models.CartResponse{}, err
	}

	// Remove the product from the cart
	if err := repository.RemoveProductFromCart(user_id, product_id); err != nil {
		// If there's an error, return an empty CartResponse struct and the error
		return models.CartResponse{}, err
	}

	// If the product quantity is not zero, update the cart details
	if cartDetails.Quantity != 0 {
		product_price, err := repository.GetPriceOfProductFromID(product_id)
		if err != nil {
			// If there's an error, return an empty CartResponse struct and the error
			return models.CartResponse{}, err
		}
		// Calculate the new total price
		cartDetails.TotalPrice = float64(cartDetails.Quantity) * product_price
		// Update the cart with the new total price
		err = repository.UpdateCartDetails(cartDetails, user_id, product_id)
		if err != nil {
			// If there's an error, return an empty CartResponse struct and the error
			return models.CartResponse{}, err
		}
	}

	// Get the updated cart details after the product has been removed
	updateCart, err := repository.CartAfterRemovalOfProduct(user_id)
	if err != nil {
		// If there's an error, return an empty CartResponse struct and the error
		return models.CartResponse{}, err
	}

	// Calculate the new total price for the cart
	cartTotal, err := repository.GetTotalPrice(user_id)
	if err != nil {
		// If there's an error, return an empty CartResponse struct and the error
		return models.CartResponse{}, err
	}

	// Return the updated CartResponse with the user's name, new total price, and updated cart details
	return models.CartResponse{
		UserName:   cartTotal.UserName,
		TotalPrice: cartTotal.TotalPrice,
		Cart:       updateCart,
	}, nil
}
