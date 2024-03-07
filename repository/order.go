package repository

import (
	"github.com/Adilfarooque/Footgo_Ecommerce/db"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
)

func UpdateHistory(userID, orderID int, amount float64, reason string) error {
	err := db.DB.Exec("INSERT INTO wallet_histories (user_id ,order_id ,description ,amount) VALUES (?,?,?,?)", userID, orderID, reason, amount).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAllOrderDetailsBrief(page, count int) ([]models.CombainedOrderDetails, error) {
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * count
	var orderDetails []models.CombainedOrderDetails
	err := db.DB.Raw("SELECT orders.id AS order_id , orders.final_price , orders.shipment_status , orders.payment_status , users.firstname,users.email,users.phone,addresses.house_name,addresses.street,addresses.city,addresses.state,addresses.pin FROM orders INNER JOIN users ON orders.user_id = users.id INNER JOIN addresses ON orders.address_id = addresses.id LIMIT ? OFFSET ?", count, offset).Scan(&orderDetails).Error
	if err != nil {
		return []models.CombainedOrderDetails{}, nil
	}
	return orderDetails, nil
}

