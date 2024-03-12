package repository

import (
	"errors"
	"time"

	"github.com/Adilfarooque/Footgo_Ecommerce/db"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
)

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

func AddProductOffer(productOffer models.ProductOfferReciever) error {
	var count int
	err := db.DB.Raw("SELECT COUNT(*) FROM product_offers WHERE offer_name = ? AND product_id = ?", productOffer.OfferName, productOffer.ProductID).Scan(&count).Error
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("the offer already exists")
	}

	//if there is any other offer for this product delete that before adding this one
	count = 0
	err = db.DB.Raw("SELECT COUNT(*) FROM product_offers WHERE product_id = ?", productOffer.ProductID).Scan(&count).Error
	if err != nil {
		return err
	}

	if count > 0 {
		err = db.DB.Exec("DELETE FROM product_offers WHERE product_id = ?", productOffer.ProductID).Scan(&count).Error
		if err != nil {
			return err
		}
	}
	startDate := time.Now()
	endDate := time.Now().Add(time.Hour * 24 * 5)
	err = db.DB.Exec("INSERT INTO product_offers (product_id , offer_name , discount_percentage , start_date , end_date) VALUES(?,?,?,?,?)", productOffer.ProductID, productOffer.OfferName, productOffer.DiscountPercentage, startDate, endDate).Error
	if err != nil {
		return err
	}
	return nil
}
