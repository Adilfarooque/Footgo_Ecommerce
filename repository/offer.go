package repository

import "github.com/Adilfarooque/Footgo_Ecommerce/db"

func FindDiscountPercentageForProduct(id int) (int, error) {
	var percentage int
	err := db.DB.Raw("SELECT discount_percentage FROM product_offers WHERE product_id $1 ", id).Scan(&percentage).Error
	if err != nil {
		return 0, err
	}
	return percentage, nil
}

func FindDiscountPercentageForCategory(id int) (int, error) {
	var percetage int
	err := db.DB.Raw("SELECT discount_percentage FROM category_offers WHERE category_id = $1", id).Scan(&percetage).Error
	if err != nil {
		return 0, err
	}
	return percetage, nil
}
