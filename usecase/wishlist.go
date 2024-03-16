package usecase

import (
	"errors"

	"github.com/Adilfarooque/Footgo_Ecommerce/repository"
)

func AddToWishlist(userID, productID int) error {
	// Check if the product exists
	productExist, err := repository.DoesProductExist(productID)
	if err != nil {
		return err
	}
	if !productExist {
		return errors.New("product does not exist")
	}

	// Check if the product already exists in the user's wishlist
	productExistInWishList, err := repository.ProductExistInWishList(productID, userID)
	if err != nil {
		return err
	}
	if productExistInWishList {
		return errors.New("product already exists in wishlist")
	}

	// Add the product to the wishlist
	err = repository.AddToWishlist(userID, productID)
	if err != nil {
		return err
	}
	return nil
}
