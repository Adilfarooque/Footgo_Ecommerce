package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/Adilfarooque/Footgo_Ecommerce/domain"
	"github.com/Adilfarooque/Footgo_Ecommerce/helper"
	"github.com/Adilfarooque/Footgo_Ecommerce/repository"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(adminDetails models.AdminLogin) (domain.TokenAdmin, error) {
	adminCompareDetails, err := repository.LoginHandler(adminDetails)
	if err != nil {
		return domain.TokenAdmin{}, err
	}
	// compare password from database and that provided from admins
	err = bcrypt.CompareHashAndPassword([]byte(adminCompareDetails.Password), []byte(adminDetails.Password))
	if err != nil {
		return domain.TokenAdmin{}, err
	}
	var adminDetailsResponse models.AdminDetailsResponse
	//  copy all details except password
	err = copier.Copy(&adminDetailsResponse, &adminCompareDetails)
	if err != nil {
		return domain.TokenAdmin{}, err
	}

	tokenString, err := helper.GenerateTokenAdmin(adminDetailsResponse)

	if err != nil {
		return domain.TokenAdmin{}, err
	}

	return domain.TokenAdmin{
		Admin: adminDetailsResponse,
		Token: tokenString,
	}, nil
}

func DashBoard() (models.CompleteAdminDashboard, error) {
	userDetails, err := repository.DashBoardUserDetails()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	productDetails, err := repository.DashBoardProductDetails()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	orderDetails, err := repository.DashBoardOrder()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	totalRevenue, err := repository.TotalRevenue()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}
	amountDetails, err := repository.AmountDetails()
	if err != nil {
		return models.CompleteAdminDashboard{}, err
	}

	return models.CompleteAdminDashboard{
		DashboardUser:    userDetails,
		DashBoardProduct: productDetails,
		DashBoardOrder:   orderDetails,
		DashBoardRevenue: totalRevenue,
		DashBoardAmount:  amountDetails,
	}, nil

}

func FilteredSalesReport(timePeriod string) (models.SalesReport, error) {
	starTime, endTime := helper.GetTimeFromPeriod(timePeriod)
	salesReport, err := repository.FilteredSalesReport(starTime, endTime)
	if err != nil {
		return models.SalesReport{}, err
	}
	return salesReport, nil
}

func ExicuteSalesReportByDate(startDate, endDate time.Time) (models.SalesReport, error) {
	orders, err := repository.FilteredSalesReport(startDate, endDate)
	if err != nil {
		return models.SalesReport{}, errors.New("report fetching failed")
	}
	return orders, nil
}

func ShowAllUsers(page, count int) ([]models.UserDetailsAtAdmin, error) {
	users, err := repository.ShowAllUsers(page, count)
	if err != nil {
		return []models.UserDetailsAtAdmin{}, err
	}
	return users, nil
}

func BlockedUser(id string) error {
	user, err := repository.GetUserByID(id)
	if err != nil {
		return err
	}
	if user.Blocked {
		return errors.New("already blocked")
	}

	user.Blocked = true

	err = repository.UpdateBlockedUserByID(user)
	if err != nil {
		return err
	}
	return nil
}

func UnBlockUser(id string) error {
	user, err := repository.GetUserByID(id)
	if err != nil {
		return err
	}
	if user.Blocked {
		user.Blocked = false
	} else {
		return errors.New("already unblocked")
	}

	err = repository.UpdateBlockedUserByID(user)
	if err != nil {
		return err
	}
	return nil
}

func ApproveOrder(order_id int) error {
	ShipmentStatus, err := repository.GetShipmentStatus(order_id)
	if err != nil {
		return err
	}
	if ShipmentStatus == "cancelled" {
		return errors.New("the order is cancelled, cannot approved it")
	}
	if ShipmentStatus == "pending" {
		return errors.New("the order is pending,cannot approve it")
	}
	if ShipmentStatus == "Processing" {
		err := repository.ApproveOrder(order_id)
		if err != nil {
			return err
		}
		return nil
	}
	//if the shipment status is not processing or cancelled . Then it is defenetely cancelled.
	return nil
}

func CancelOrderFromAdmin(orderID int) error {
	ok, err := repository.CheckOrderID(orderID)
	fmt.Println(err)
	if !ok {
		return err
	}
	orderProduct, err := repository.GetProductDetailsFromAdmin(orderID)
	if err != nil {
		return err
	}
	err = repository.CancelOrders(orderID)
	if err != nil {
		return err
	}
	err = repository.UpdateStockOfProduct(orderProduct)
	if err != nil {
		return err
	}
	payment_status, err := repository.PaymetStatus(orderID)
	if err != nil {
		return err
	}
	amount, err := repository.TotalAmountFromOrder(orderID)
	if err != nil {
		return err
	}
	userID, err := repository.UserIDFromOrder(orderID)
	if err != nil {
		return err
	}
	if payment_status == "refunded" {
		err := repository.UpdateAmountToWallet(userID, amount)
		if err != nil {
			return err
		}
		reason := "Amount credited for cancellation of order by admin"
		err = repository.UpdateHistory(userID, orderID, amount, reason)
		if err != nil {
			return err
		}
	}
	return nil
}

func AdddPaymentMehod(payment models.NewPaymentMethod) (domain.PaymentMethod, error) {
	exists, err := repository.CheckifPaymentMethodAlreadyExists(payment.PaymentName)
	if err != nil {
		return domain.PaymentMethod{}, err
	}
	if exists {
		return domain.PaymentMethod{}, errors.New("payment method already exists")
	}
	paymentAdd, err := repository.AdddPaymentMehod(payment)
	if err != nil {
		return domain.PaymentMethod{}, err
	}
	return paymentAdd, nil
}
