package models

type OrderDetails struct {
	OrderId        int
	FinalPrice     float64
	ShipmentStatus string
	PaymentStatus  string
}

type OrederProductDetails struct {
	ProductID   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	TotalPrice  float64 `json:"total_price"`
}

type FullOrderDetails struct {
	OrderDetails         OrderDetails
	OrederProductDetails []OrederProductDetails
}

type OrderProducts struct {
	ProductId string `json:"id"`
	Stock     int    `json:"stock"`
}
type CombainedOrderDetails struct {
	OrederId       string  `json:"order_id"`
	FinalPrice     float64 `json:"final_price"`
	ShipmentStatus string  `json:"shipment_status"`
	PaymentStatus  string  `json:"payment_status"`
	Firstname      string  `json:"firstname"`
	Email          string  `json:"email"`
	Phone          string  `json:"phone"`
	HouseName      string  `json:"house_name" validate:"required"`
	Street         string  `json:"street"`
	City           string  `json:"city"`
	State          string  `json:"state" validate:"required"`
	Pin            string  `json:"pin" validate:"required"`
}

type OrederPaymentDetails struct {
	UserID     int     `json:"user_id"`
	Username   string  `json:"username"`
	Razor_id   string  `json:"razor_id"`
	OrderID    int     `json:"order_id"`
	FinalPrice float64 `json:"final_price"`
}

type AddOrderProductDetails struct {
	UserID       int `json:"user_id"`
	AddressID    int `json:"address_id"`
	PaymetMethod int `json:"payment_id"`
}

type OrderResponse struct {
	AddOrderProductDetails AddOrderProductDetails
	OrderDetails           OrderDetails
}

type OrderFromCart struct {
	PaymentID uint `json:"payment_id"`
	AddressID uint `json:"address_id"`
}

type OrederIncoming struct {
	UserID    int `json:"user_id"`
	PaymentID int `json:"payment_id"`
	AddressID int `json:"address_id"`
}

type Invoice struct {
	ProductID     uint    `json:"product_id"`
	ProductName   string  `json:"product_name"`
	Quantity      float64 `json:"quantity"`
	DiscountPrice float64 `json:"discount_price"`
	TotalPrice    float64 `json:"total_price"`
}
