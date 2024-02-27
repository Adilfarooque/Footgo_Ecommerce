package repository

import (
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
