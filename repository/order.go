package repository

import (
	"github.com/Adilfarooque/Footgo_Ecommerce/db"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
)

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

func GetShipmentStatus(order_id int) (string, error) {
	var status string
	err := db.DB.Raw("SELECT shipment_status FROM orders WHERE id = ?", order_id).Scan(&status).Error
	if err != nil {
		return "", err
	}
	return status, nil
}

func ApproveOrder(order_id int) error {
	err := db.DB.Exec("UPDATE orders SET shipment_status = 'order placed' , approval = 'true' WHERE id = ?", order_id).Error
	if err != nil {
		return err
	}
	return err
}

func CheckOrderID(orderID int) (bool, error) {
	var count int
	err := db.DB.Raw("SELECT COUNT(*) FROM orders WHERE id = ?", orderID).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func GetProductDetailsFromAdmin(orderID int) ([]models.OrderProducts, error) {
	var OrederProductDetails []models.OrderProducts
	if err := db.DB.Raw("SELECT product_id, quantity AS stock FROM order_items WHERE order_id = ?", orderID).Scan(&OrederProductDetails).Error; err != nil {
		return []models.OrderProducts{}, err
	}
	return OrederProductDetails, nil
}

func CancelOrders(orderId int) error {
	status := "cancelled"
	err := db.DB.Exec("UPDATE orders SET shipment_status = ?, approva = 'false' WHERE id = ?", status, orderId).Error
	if err != nil {
		return err
	}
	var paymentMethod int
	err = db.DB.Raw("SELECT payment_method_id FROM orders WHERE id = ?", orderId).Scan(&paymentMethod).Error
	if err != nil {
		return err
	}
	if paymentMethod == 3 || paymentMethod == 2 {
		err = db.DB.Exec("UPDATE orders SET payment_status = 'refunded' WHERE id = ?", orderId).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateStockOfProduct(orderProducts []models.OrderProducts) error {
	for _, ok := range orderProducts {
		var quantity int
		if err := db.DB.Raw("SELECT stock FROM products WHERE id = ?", ok.ProductId).Scan(&orderProducts).Error; err != nil {
			return err
		}
		ok.Stock += quantity
		if err := db.DB.Exec("UPDATE products SET stock = ? WHERE id = ?", ok.Stock, ok.ProductId).Error; err != nil {
			return err
		}
	}
	return nil
}

func PaymetStatus(orderID int) (string, error) {
	var status string
	err := db.DB.Raw("SELECT payment_status FROM orders WHERE id = ?", orderID).Scan(&status).Error
	if err != nil {
		return "", err
	}
	return status, nil
}

func TotalAmountFromOrder(orderID int) (float64, error) {
	var total float64
	err := db.DB.Raw("SELECT final_price FROM orders WHERE id = ?", orderID).Scan(&total).Error
	if err != nil {
		return 0.0, err
	}
	return total, nil
}

func UserIDFromOrder(orderID int) (int, error) {
	var a int
	err := db.DB.Raw("SELECT user_id FROM orders WHERE id = ?", orderID).Scan(&a).Error
	if err != nil {
		return 0, err
	}
	return a, nil
}

func UpdateAmountToWallet(userID int, amount float64) error {
	err := db.DB.Exec("UPDATE wallets SET amount = amount + ? WHERE user_id = ?", amount, userID).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateHistory(userID, orderID int, amount float64, reason string) error {
	err := db.DB.Exec("INSERT INTO wallet_histories (user_id ,order_id ,description ,amount) VALUES (?,?,?,?)", userID, orderID, reason, amount).Error
	if err != nil {
		return err
	}
	return nil
}
