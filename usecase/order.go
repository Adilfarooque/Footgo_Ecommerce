package usecase

import (
	"errors"
	"fmt"

	"github.com/Adilfarooque/Footgo_Ecommerce/repository"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
)

func GetAllOrderDetailsForAdmin(page, pageSize int) ([]models.CombainedOrderDetails, error) {
	orderDetails, err := repository.GetAllOrderDetailsBrief(page, pageSize)
	if err != nil {
		return []models.CombainedOrderDetails{}, err
	}
	return orderDetails, nil
}

func GetOrderDetails(userId int, page int, count int) ([]models.FullOrderDetails, error) {
	fullOrderDetails, err := repository.GetOrderDetails(userId, page, count)
	if err != nil {
		return []models.FullOrderDetails{}, err
	}
	return fullOrderDetails, nil
}

func CancelOrders(orderID, userId int) error {
	userTest, err := repository.UserOrderRelationship(orderID, userId)
	if err != nil {
		return err
	}
	if userTest != userId {
		return errors.New("the order is not done by this user")
	}
	orderProductDetails, err := repository.GetProductDetailsFromOrders(orderID)
	if err != nil {
		return err
	}
	shipmentStatus, err := repository.GetShipmentStatus(orderID)
	if err != nil {
		return err
	}
	if shipmentStatus == "delivered" {
		return errors.New("item already delivered , cannot cancel")
	}

	if shipmentStatus == "pending" || shipmentStatus == "returned" || shipmentStatus == "return" {
		message := fmt.Sprint(shipmentStatus)
		return errors.New("the order is in " + message + ", so no point in cancelling")
	}

	if shipmentStatus == "cancelled" {
		return errors.New("the order is already cancelled, so no point in cancelling")
	}
	if err := repository.CancelOrders(orderID); err != nil {
		return err
	}
	payment_status, err := repository.PaymetStatus(orderID)
	if err != nil {
		return err
	}
	if err := repository.UpdateQuantityOfProduct(orderProductDetails); err != nil {
		return err
	}
	amount, err := repository.TotalAmountFromOrder(orderID)
	if err != nil {
		return err
	}
	if payment_status == "refunded" {
		if err = repository.UpdateAmountToWallet(userId, amount); err != nil {
			return err
		}
		reason := "Amount credited for cancellation of order by user"
		if err := repository.UpdateHistory(userId, orderID, amount, reason); err != nil {
			return err
		}
	}
	return nil
}

func PaymentMethodID(order_id int) (int, error) {
	id, err := repository.PaymentMethodID(order_id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func ExecutePurchaseCOD(orderID int) error {
	if err := repository.OrderExist(orderID); err != nil {
		return err
	}
	shipmentStatus, err := repository.GetShipmentStatus(orderID)
	if err != nil {
		return err
	}
	if shipmentStatus == "delivered" {
		return errors.New("item delivered, cannot pay")
	}
	if shipmentStatus == "order placed" {
		return errors.New("item placed, cannot pay")
	}
	if shipmentStatus == "cancelled" || shipmentStatus == "returned" || shipmentStatus == "return" {
		message := fmt.Sprint(shipmentStatus)
		return errors.New("the order is in " + message + "so can't paid")
	}
	if shipmentStatus == "processing" {
		return errors.New("the order is already paid")
	}
	if err := repository.UpdateOrder(orderID); err != nil {
		return err
	}
	return nil
}

func CheckOut(userID int) (models.CheckoutDetails, error) {
	allUserAddress, err := repository.GetAllAddress(userID)
	if err != nil {
		return models.CheckoutDetails{}, err
	}
	paymentDetails, err := repository.GetAllPaymentOption(userID)
	if err != nil {
		return models.CheckoutDetails{}, err
	}
	cartItems, err := repository.DisplayCart(userID)
	if err != nil {
		return models.CheckoutDetails{}, err
	}
	grandTotal, err := repository.GetTotalPrice(userID)
	if err != nil {
		return models.CheckoutDetails{}, err
	}
	return models.CheckoutDetails{
		AddressInfoResponse: allUserAddress,
		Payment_Method:      paymentDetails,
		Cart:                cartItems,
		Total_Price:         grandTotal.FinalPrice,
	}, nil

}

