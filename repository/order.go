package repository

import "github.com/Adilfarooque/Footgo_Ecommerce/db"

func UpdateHistory(userID, orderID int, amount float64, reason string) error {
	err := db.DB.Exec("INSERT INTO wallet_histories (user_id ,order_id ,description ,amount) VALUES (?,?,?,?)", userID, orderID, reason, amount).Error
	if err != nil {
		return err
	}
	return nil
}
