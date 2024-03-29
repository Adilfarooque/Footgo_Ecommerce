package usecase

import (
	"errors"
	"fmt"

	"github.com/Adilfarooque/Footgo_Ecommerce/domain"
	"github.com/Adilfarooque/Footgo_Ecommerce/repository"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
	"github.com/jinzhu/copier"
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

func OrderItemsFromCart(orderFromCart models.OrderFromCart, userID int) (domain.OrderSuccessResponse, error) {
	var orderBody models.OrederIncoming
	if err := copier.Copy(&orderBody, &orderFromCart); err != nil {
		return domain.OrderSuccessResponse{}, err
	}
	orderBody.UserID = userID
	cartExist, err := repository.DoesCartExist(userID)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}
	if !cartExist {
		return domain.OrderSuccessResponse{}, errors.New("cart empty can't order")
	}

	addressExist, err := repository.AddressExist(orderBody)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}
	if !addressExist {
		return domain.OrderSuccessResponse{}, errors.New("address does not exist")
	}

	paymentExist, err := repository.PaymentExist(orderBody)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}
	if !paymentExist {
		return domain.OrderSuccessResponse{}, errors.New("payment method doesn't exist")
	}

	cartItems, err := repository.GetAllItemsFromCart(orderBody.UserID)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	total, err := repository.TotalAmountInCart(orderBody.UserID)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	discount_price, err := repository.GetCouponDiscountPrice(int(orderBody.UserID), total)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	if err := repository.UpdateCouponDetails(discount_price, orderBody.UserID); err != nil {
		return domain.OrderSuccessResponse{}, err
	}

	FinalPrice := total - discount_price
	if orderBody.PaymentID == 3 {
		walletAmount, err := repository.WalletAmount(userID)
		if err != nil {
			return domain.OrderSuccessResponse{}, err
		}
		if FinalPrice >= walletAmount {
			return domain.OrderSuccessResponse{}, errors.New("this much of amount not available in wallet")
		}
	}

	order_id, err := repository.OrderItems(orderBody, FinalPrice)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}
	if orderBody.PaymentID == 3 {
		if err := repository.UpdateWalletAfterOrder(userID, FinalPrice); err != nil {
			return domain.OrderSuccessResponse{}, err
		}
		reason := "Amount debited for purchasing products"
		if err := repository.UpdateHistoryForDebit(userID, order_id, FinalPrice, reason); err != nil {
			return domain.OrderSuccessResponse{}, err
		}
	}
	if err := repository.AddOrderProducts(order_id, cartItems); err != nil {
		return domain.OrderSuccessResponse{}, err
	}
	orderSuccessResponse, err := repository.GetBriefOrderDetails(order_id, orderBody.PaymentID)
	if err != nil {
		return domain.OrderSuccessResponse{}, err
	}
	var orderItemDetails domain.OrderItem
	for _, c := range cartItems {
		orderItemDetails.ProductID = c.ProductID
		orderItemDetails.Quantity = c.Quantity
		err := repository.UpdateCartAfterOrder(userID, int(orderItemDetails.ProductID), orderItemDetails.Quantity)
		if err != nil {
			return domain.OrderSuccessResponse{}, err
		}
	}
	return orderSuccessResponse, nil

}
