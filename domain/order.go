package domain

import "gorm.io/gorm"

type Orders struct {
	gorm.Model
	UserID          int           `json:"user_id" gorm:"not null"`
	User            User          `json:"-" gorm:"foreignkey:UserID"`
	AddressID       uint          `json:"address_id" gorm:"not null"`
	Address         Address       `json:"address" gorm:"foreignkey:AddressID"`
	PaymentMethodID uint          `json:"paymentmethod_id"`
	PaymentMethod   PaymentMethod `json:"-" gorm:"foreignkey:PaymentMethodID"`
	ShipmentStatus  string        `json:"shipment_status" gorm:"default:'pending'"`
	PaymentStatus   string        `json:"payment_status" gorm:"default:'not paid'"`
	FinalPrice      float64       `json:"final_price"`
	Approval        bool          `json:"approval" gorm:"default:false"`
}

type OrderItem struct {
	ID         uint    `json:"id" gorm:"PrimaryKey:not null"`
	OrderID    uint    `json:"order_id"`
	Order      Orders  `json:"-" gorm:"foreignkey:OrderID;constraint:OneDelete:CASCADE"`
	ProductID  uint    `json:"product_id"`
	Products   Product `json:"-" gorm:"foreignkey:ProductID"`
	Quantity   float64 `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

type OrderSuccessResponse struct {
	OrderID        uint   `json:"order_id"`
	ShipmentStatus string `json:"shipment_status"`
}
