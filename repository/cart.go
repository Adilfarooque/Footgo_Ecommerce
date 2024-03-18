package repository

import (
	"errors"

	"github.com/Adilfarooque/Footgo_Ecommerce/db"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
	"gorm.io/gorm"
)

func CheckProduct(product_id int) (bool, string, error) {
	var count int
	err := db.DB.Raw("SELECT COUNT(*) FROM products WHERE id = ?", product_id).Scan(&count).Error
	if err != nil {
		return false, "", err
	}
	if count > 0 {
		var category string
		err := db.DB.Raw("SELECT categories.category FROM categories INNER JOIN products ON products.category_id = categories.id WHERE products.id = ?", product_id).Scan(&category).Error
		if err != nil {
			return false, "", err
		}
		return true, category, nil
	}
	return false, "", nil
}

func QuantityOfProductsInCart(userId int, productId int) (int, error) {
	var productQty int
	if err := db.DB.Raw("SELECT quantity FROM carts WHERE user_id = ? AND product_id = ?", userId, productId).Scan(&productQty).Error; err != nil {
		return 0, err
	}
	return productQty, nil
}

func AddItemInToCart(userId, productId, Quantity int, productprice float64) error {
	if err := db.DB.Exec("INSERT INTO carts (user_id,product_id,quantity,total_price) VALUES(?,?,?,?)", userId, productId, Quantity, productprice).Error; err != nil {
		return err
	}
	return nil
}

func TotalPriceForProductInCart(userID, productID int) (float64, error) {
	var totalprice float64
	if err := db.DB.Raw("SELECT SUM(total_price) AS total_price FROM carts WHERE user_id = ? AND product_id = ?", userID, productID).Scan(&totalprice).Error; err != nil {
		return 0.0, err
	}
	return totalprice, nil
}

func UpdateCart(quantity int, price float64, userID, productID int) error {
	if err := db.DB.Exec("UPDATE carts SET quantity = ?, total_price = ? WHERE user_id = ? AND product_id = ?", quantity, price, productID, userID).Error; err != nil {
		return err
	}
	return nil
}

func DisplayCart(userID int) ([]models.Cart, error) {
	var count int                  // To store the count of cart items
	var cartResponse []models.Cart // To store the cart details

	// Count the number of cart items for the user
	if err := db.DB.Raw("SELECT COUNT(*) FROM carts WHERE user_id = ?", userID).Scan(&count).Error; err != nil {
		// If there's an error, return an empty cart slice and the error
		return []models.Cart{}, err
	}

	// If there are no cart items, return an empty cart slice and no error
	if count == 0 {
		return []models.Cart{}, nil
	}

	// Fetch the cart details
	if err := db.DB.Raw("SELECT carts.user_id, users.firstname AS user_name, carts.product_id, products.name AS product_name, carts.quantity, carts.total_price FROM carts INNER JOIN users ON carts.user_id = users.id INNER JOIN products ON carts.product_id = products.id WHERE carts.user_id = ?", userID).Scan(&cartResponse).Error; err != nil {
		// If there's an error fetching the cart details, check if it's because no records were found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// If no records are found, return an empty cart slice and the error
			return []models.Cart{}, err
		}
		// For any other error, return an empty cart slice and the error
		return []models.Cart{}, err
	}
	// If successful, return the cart details and no error
	return cartResponse, nil
}

func GetTotalPrice(userID int) (models.CartTotal, error) {
	var cartTotal models.CartTotal // Struct to store the total price and user's name

	// Query to calculate the total price of the user's cart
	err := db.DB.Raw("SELECT COALESCE(SUM(total_price), 0) FROM carts WHERE user_id = ?", userID).Scan(&cartTotal.TotalPrice).Error
	if err != nil {
		// If there's an error, return an empty CartTotal struct and the error
		return models.CartTotal{}, err
	}

	// Query to get the user's name based on the userID
	err = db.DB.Raw("SELECT firstname as user_name FROM users WHERE id = ?", userID).Scan(&cartTotal.UserName).Error
	if err != nil {
		// If there's an error, return an empty CartTotal struct and the error
		return models.CartTotal{}, err
	}

	// If successful, return the cartTotal struct with the total price and user's name, and no error
	return cartTotal, nil
}

func ProductStockMinus(productID, stock int) error {
	if err := db.DB.Exec("UPDATE products SET stock = stock - ? WHERE id = ?", stock, productID).Error; err != nil {
		return err
	}
	return nil
}
