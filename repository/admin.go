package repository

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Adilfarooque/Footgo_Ecommerce/db"
	"github.com/Adilfarooque/Footgo_Ecommerce/domain"
	"github.com/Adilfarooque/Footgo_Ecommerce/helper"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
)

// func LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error) {
// 	var details domain.Admin
// 	if err := db.DB.Raw("SELECT * FROM users WHERE email = ? AND isadmin = ture", adminDetails).Scan(&details).Error; err != nil {
// 		return domain.Admin{}, err
// 	}
// 	return details, nil
// }

func LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error) {
	var details domain.Admin
	if err := db.DB.Raw("SELECT * FROM users WHERE email = ? AND isadmin = true", adminDetails.Email).Scan(&details).Error; err != nil {
		return domain.Admin{}, err
	}
	return details, nil
}

func DashBoardUserDetails() (models.DashBoardUser, error) {
	var userDetails models.DashBoardUser
	err := db.DB.Raw("SELECT COUNT(*) FROM users WHERE isadmin = false").Scan(&userDetails.TotalUsers).Error
	if err != nil {
		return models.DashBoardUser{}, nil
	}
	err = db.DB.Raw("SELECT COUNT(*) FROM users WHERE blocked = true").Scan(&userDetails.BlockedUser).Error
	if err != nil {
		return models.DashBoardUser{}, nil
	}
	return userDetails, nil
}

func DashBoardProductDetails() (models.DashBoardProduct, error) {
	var productDetails models.DashBoardProduct
	err := db.DB.Raw("SELECT COUNT(*) FROM products").Scan(&productDetails.TotalProduct).Error
	if err != nil {
		return models.DashBoardProduct{}, nil
	}
	err = db.DB.Raw("SELECT COUNT(*) FROM products WHERE stock <= 0").Scan(&productDetails.OutofStockProduct).Error
	if err != nil {
		return models.DashBoardProduct{}, nil
	}
	return productDetails, nil

}

func DashBoardOrder() (models.DashBoardOrder, error) {
	var orderDetails models.DashBoardOrder
	err := db.DB.Raw("SELECT COUNT(*) FROM orders WHERE payment_status = 'paid' AND approval = true").Scan(&orderDetails.CompletedOrder).Error
	if err != nil {
		return models.DashBoardOrder{}, err
	}

	err = db.DB.Raw("SELECT COUNT(*) FROM orders WHERE shipment_status = 'pending' OR shipment_status = 'processing'").Scan(&orderDetails.PendingOrder).Error
	if err != nil {
		return models.DashBoardOrder{}, err
	}

	err = db.DB.Raw("SELECT COUNT(*) FROM orders WHERE shipment_status = 'cancelled'").Scan(&orderDetails.CancelledOrder).Error
	if err != nil {
		return models.DashBoardOrder{}, nil
	}

	err = db.DB.Raw("SELECT COUNT(*) FROM orders").Scan(&orderDetails.TotalOrder).Error
	if err != nil {
		return models.DashBoardOrder{}, nil
	}

	err = db.DB.Raw("SELECT COALESCE(SUM(quantity),0) FROM carts").Scan(&orderDetails.TotalOrderItem).Error
	if err != nil {
		return models.DashBoardOrder{}, nil
	}
	return orderDetails, nil
}

func TotalRevenue() (models.DashBoardRevenue, error) {
	var revenueDetails models.DashBoardRevenue
	startTime := time.Now().AddDate(0, 0, -1)
	endTime := time.Now()
	err := db.DB.Raw("SELECT COALESCE(SUM(final_price),0) FROM orders WHERE payment_status = 'paid' AND approval = true AND created_at >= ? AND created_at <= ?", startTime, endTime).Scan(revenueDetails.TodayRevenue).Error
	if err != nil {
		return models.DashBoardRevenue{}, nil
	}
	startTime, endTime = helper.GetTimeFromPeriod("month")
	err = db.DB.Raw("SELECT COALESCE(SUM(final_price),0) FROM orders WHERE payment_status = 'paid' AND approval = true AND created_at >= ? AND created_at <= ?", startTime, endTime).Scan(&revenueDetails.MonthRevenue).Error
	if err != nil {
		return models.DashBoardRevenue{}, nil
	}

	startTime, endTime = helper.GetTimeFromPeriod("year")
	err = db.DB.Raw("SELECT COALESCE(SUM(final_price),0) FROM orders WHERE payment_status = 'paid' AND approval = true AND created_at >= ? AND created_at <= ?", startTime, endTime).Scan(&revenueDetails.YearRevenue).Error
	if err != nil {
		return models.DashBoardRevenue{}, nil
	}

	return revenueDetails, nil
}

func AmountDetails() (models.DashBoardAmount, error) {
	var amountDetails models.DashBoardAmount
	err := db.DB.Raw("SELECT COALESCE(SUM(final_price),0) FROM orders WHERE payment_status = 'paid' AND approval = true").Scan(&amountDetails.CreditedAmount).Error
	if err != nil {
		return models.DashBoardAmount{}, nil
	}

	err = db.DB.Raw("SELECT COALESCE(SUM(final_price),0) FROM orders WHERE payment_status = 'pending' OR shipment_status = 'order placed'").Scan(&amountDetails.PendingAmount).Error
	if err != nil {
		return models.DashBoardAmount{}, nil
	}
	return models.DashBoardAmount{}, nil
}

func FilteredSalesReport(startTime time.Time, endTime time.Time) (models.SalesReport, error) {
	var salesReport models.SalesReport
	result := db.DB.Raw("SELECT COALESCE(SUM(final_price),0) FROM orders WHERE payment_status='paid' AND approval = true AND created_at >= ? AND created_at <= ?", startTime, endTime).Scan(&salesReport.TotalSales)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}
	result = db.DB.Raw("SELECT COUNT(*) FROM orders").Scan(&salesReport.TotalOrders)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}
	result = db.DB.Raw("SELECT COUNT(*) FROM orders WHERE payment_status = 'paid' and approval = true and created_at >= ? AND created_at <= ?", startTime, endTime).Scan(&salesReport.CompletedOrders)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}
	result = db.DB.Raw("SELECT COUNT(*) FROM orders WHERE shipment_status = 'processing' AND approval = false AND created_at >= ? AND created_at<=?", startTime, endTime).Scan(&salesReport.PendingOrders)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}
	var productID int
	result = db.DB.Raw("SELECT product_id FROM order_items GROUP BY product_id order by SUM(quantity) DESC LIMIT 1").Scan(&productID)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}
	result = db.DB.Raw("SELECT name FROM products WHERE id = ?", productID).Scan(&salesReport.TrendingProduct)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}
	return salesReport, nil
}

func ShowAllUsers(page, count int) ([]models.UserDetailsAtAdmin, error) {
	var users []models.UserDetailsAtAdmin
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * count
	err := db.DB.Raw("SELECT id, firstname, lastname, email, phone, blocked FROM users WHERE isadmin = 'false' LIMIT ? OFFSET ?", count, offset).Scan(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByID(id string) (domain.User, error) {
	user_id, err := strconv.Atoi(id)
	if err != nil {
		return domain.User{}, err
	}

	var count int
	if err := db.DB.Raw("SELECT COUNT(*) FROM users WHERE id = ?", user_id).Scan(&count).Error; err != nil {
		return domain.User{}, err
	}

	if count < 1 {
		return domain.User{}, errors.New("user for the given id does not exist")
	}

	var userDetails domain.User
	if err := db.DB.Raw("SELECT * FROM users WHERE id = ?", user_id).Scan(&userDetails).Error; err != nil {
		return domain.User{}, err
	}
	return userDetails, nil
}

/*
func UpdateBlockedUserByID(user domain.User) error {
	err := db.DB.Raw("UPDATE users SET blocked = ? WHERE id = ?", user.Blocked, user.ID).Error
	if err != nil {
		fmt.Println("Error updating user :", user)
		return err
	}
	return nil
}
*/

func UpdateBlockedUserByID(user domain.User) error {
	query := "UPDATE users SET blocked = ? WHERE id = ?"
	err := db.DB.Exec(query, user.Blocked, user.ID).Error
	if err != nil {
		fmt.Printf("Error updating user with ID %d: %v\n", user.ID, err)
		return err
	}
	fmt.Printf("User with ID %d successfully updated\n", user.ID)
	return nil
}

func ShowAllProductsFromAdmin(page, count int) ([]models.ProductBreif, error) {
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * count
	var ProductBreif []models.ProductBreif
	err := db.DB.Raw("SELECT * FROM products LIMIT ? OFFSET ?", count, offset).Scan(&ProductBreif).Error
	if err != nil {
		return nil, err
	}
	return ProductBreif, nil
}

