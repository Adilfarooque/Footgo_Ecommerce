package repository

import (
	"errors"

	"github.com/Adilfarooque/Footgo_Ecommerce/db"
)

func ProductExistInWishList(productID, userID int) (bool, error) {
	var count int
	err := db.DB.Raw("SELECT COUNT(*) FROM wish_lists WHERE user_id = ? AND product_id = ?", userID, productID).Scan(&count).Error
	if err != nil {
		return false, errors.New("error checking user product already present")
	}
	return count > 0, nil
}


func AddToWishlist(userID, productID int) error {
	err := db.DB.Exec("INSERT INTO wish_lists (user_id,product_id) VALUES (?,?)", userID, productID).Error
	if err != nil {
		return err
	}
	return nil
}