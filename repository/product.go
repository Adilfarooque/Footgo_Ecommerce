package repository

import (
	"errors"
	"log"
	"strconv"

	"github.com/Adilfarooque/Footgo_Ecommerce/db"
	"github.com/Adilfarooque/Footgo_Ecommerce/domain"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
)

func ShowAllProducts(page int, count int) ([]models.ProductBreif, error) {
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * count

	var ProductBreif []models.ProductBreif
	err := db.DB.Raw("SELECT * FROM products limit ? offset ? ", count, offset).Scan(&ProductBreif).Error
	if err != nil {
		return nil, err
	}
	return ProductBreif, nil
}

func GetImage(productID int) ([]string, error) {
	var url []string
	if err := db.DB.Raw("SELECT url FROM Images WHERE product_id = ?", productID).Scan(&url).Error; err != nil {
		return url, err
	}
	return url, nil
}

func CheckValidateCategory(data map[string]int) error {
	for _, id := range data {
		var count int
		err := db.DB.Raw("SELECT COUNT(*) FROM categories WHERE id = ? ", id).Scan(&count).Error
		if err != nil {
			return err
		}
		if count < 1 {
			return errors.New("doesn't exist")
		}
	}
	return nil
}

func GetProductFromCategory(id int) ([]models.ProductBreif, error) {
	var product []models.ProductBreif
	err := db.DB.Raw("SELECT * FROM products JOIN categories ON products.category_id = categories.id WHERE categories.id = ?", id).Scan(&product).Error
	if err != nil {
		return []models.ProductBreif{}, err
	}
	return product, nil
}

func GetQuantityFromProductID(id int) (int, error) {
	var quantity int
	err := db.DB.Raw("SELECT stock FROM products WHERE id = ?", id).Scan(&quantity).Error
	if err != nil {
		return 0.0, err
	}
	return quantity, nil
}
func GetPriceOfProductFromID(productId int) (float64, error) {
	var productPrice float64
	if err := db.DB.Raw("SELECT price FROM products WHERE id = ?", productId).Scan(&productPrice).Error; err != nil {
		return 0.0, err
	}
	return productPrice, nil
}

func ProductAlreadyExist(Name string) bool {
	var count int
	if err := db.DB.Raw("SELECT COUNT(*) FROM products WHERE name = ?", Name).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0

}

func AddProduct(product models.Product) (domain.Product, error) {
	var p domain.Product
	query := "INSERT INTO products (name, description, category_id, size, stock, price) VALUES ($1, $2, $3, $4, $5, $6) RETURNING name, description, category_id, size, stock, price"
	err := db.DB.Raw(query, product.Name, product.Description, product.CategoryID, product.Size, product.Stock, product.Price).Scan(&p).Error
	if err != nil {
		log.Println(err.Error())
		return domain.Product{}, err
	}
	var prodctResponse domain.Product
	err = db.DB.Raw("SELECT * FROM products WHERE name = ?", p.Name).Scan(&prodctResponse).Error
	if err != nil {
		log.Println(err.Error())
		return domain.Product{}, err
	}
	return prodctResponse, nil
}

func StockValid(Name string) bool {
	var count int
	if err := db.DB.Raw("SELECT SUM(stock) FROM products WHERE name = ? AND stock >= 0", Name).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func CheckProductExist(productID int) (bool, error) {
	var prd int
	err := db.DB.Raw("SELECT COUNT(*) FROM products WHERE id = ?", productID).Scan(&prd).Error
	if err != nil {
		return false, err // Return false and the error
	}
	// If the product count is greater than 0, the product exists
	return prd > 0, nil // Return true if prd > 0, otherwise false
}

/*
	func UpdateProduct(productID int, stock int) (models.ProductUpdateReciever, error) {
		if stock <= 0 {
			return models.ProductUpdateReciever{}, errors.New("stock doesnot update invalid input")
		}
		if db.DB == nil {
			return models.ProductUpdateReciever{}, errors.New("database connection is nil")
		}
		if err := db.DB.Exec("UPDATE products SET stock = stock + $1 WHERE id = $2", stock, productID).Error; err != nil {
			return models.ProductUpdateReciever{}, err
		}
		var newdetails models.ProductUpdateReciever
		var newQuantity int
		if err := db.DB.Raw("SELECT stock FROM products WHERE id =?", productID).Scan(&newQuantity).Error; err != nil {
			return models.ProductUpdateReciever{}, err
		}
		newdetails.ProductID = productID
		newdetails.Stock = newQuantity
		return newdetails, nil
	}
*/
func UpdateProduct(productID int, stock int) (models.ProductUpdateReciever, error) {
	if stock <= 0 {
		return models.ProductUpdateReciever{}, errors.New("stock does not update invalid input")
	}
	if db.DB == nil {
		return models.ProductUpdateReciever{}, errors.New("database connection is nil")
	}
	if err := db.DB.Exec("UPDATE products SET stock = stock + $1 WHERE id = $2", stock, productID).Error; err != nil {
		return models.ProductUpdateReciever{}, err
	}
	var newdetails models.ProductUpdateReciever
	var newQuantity int
	if err := db.DB.Raw("SELECT stock FROM products WHERE id =?", productID).Scan(&newQuantity).Error; err != nil {
		return models.ProductUpdateReciever{}, err
	}
	newdetails.ProductID = productID
	newdetails.Stock = newQuantity
	return newdetails, nil
}

func DeleteProduct(id string) error {
	product_id, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	var count int
	if err := db.DB.Raw("SELECT COUNT(*) FROM products WHERE id = ?", product_id).Scan(&count).Error; err != nil {
		return err
	}
	if count < 1 {
		return errors.New("product for given id does not exist")
	}
	if err := db.DB.Exec("DELETE FROM products WHERE id = ?", product_id).Error; err != nil {
		return err
	}
	// if err := db.DB.Exec("DELETE FROM images WHERE product_id = ?", product_id).Error; err != nil {
	// 	return err
	// }
	return nil
}

func GetInventory(prefix string) ([]models.ProductBreif, error) {
	var productDetails []models.ProductBreif
	query := `
	SELECT i.*
	FROM products i
	LEFT JOIN categories c ON i.category_id = c.id
	WHERE i.name ILIKE '%' || $1 || '%'
    OR c.category ILIKE '%' || $1 || '%';`
	if err := db.DB.Raw(query, prefix).Scan(&productDetails).Error; err != nil {
		return []models.ProductBreif{}, err
	}
	return productDetails, nil
}

func UpdateProductImage(productID int, url string) error {
	err := db.DB.Exec("INSERT INTO images (product_id,url) VALUES($1,$2) RETURNING *", productID, url)
	if err != nil {
		return errors.New("error while insert to database")
	}
	return nil
}

func DoesProductExist(productID int) (bool, error) {
	var count int
	if err := db.DB.Raw("SELECT COUNT(*) FROM products WHERE id = ?", productID).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func FindCategoryID(categoryID int) (int, error) {
	var catid int
	if err := db.DB.Raw("SELECT category_id FROM products WHERE id = ?", categoryID).Scan(&catid).Error; err != nil {
		return 0, err
	}
	return catid, nil
}
