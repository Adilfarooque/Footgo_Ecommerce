package usecase

import (
	"errors"

	"github.com/Adilfarooque/Footgo_Ecommerce/domain"
	"github.com/Adilfarooque/Footgo_Ecommerce/repository"
	"github.com/Adilfarooque/Footgo_Ecommerce/utils/models"
)

func ShowAllProducts(page int, count int) ([]models.ProductBreif, error) {
	//Retrieve product details from the repository
	productDetails, err := repository.ShowAllProducts(page, count)
	if err != nil {
		return []models.ProductBreif{}, err
	}
	//Loop through each product detail
	for i := range productDetails {
		p := &productDetails[i]
		//Prd status based on stock availability
		if p.Stock <= 0 {
			p.ProductStatus = "out of stock"
		} else {
			p.ProductStatus = "in stock"
		}
	}
	//Loop through each product details to calculate and apply discounts
	for j := range productDetails {
		//Find the discount percentage for the product
		discount_percentage, err := repository.FindDiscountPercentageForProduct(int(productDetails[j].ID))
		if err != nil {
			//If there's an error finding discount percentage, return an empty slice and error
			return []models.ProductBreif{}, errors.New("there was some error in finding the discounted prices")
		}
		//Calculte discount if applicable
		var discount float64
		if discount_percentage > 0 {
			discount = (productDetails[j].Price * float64(discount_percentage)) / 100
		}
		//Apply discount to calculate discount price
		productDetails[j].DiscountedPrice = productDetails[j].Price - discount
		//Find discount percentage for the product's category
		discount_percentageCategory, err := repository.FindDiscountPercentageForCategory(int(productDetails[j].CatergoryID))
		if err != nil {
			return []models.ProductBreif{}, errors.New("there was some error in finding the discounted prices")
		}
		//Calculte category-level discount if applicable
		var categorydiscount float64
		if discount_percentageCategory > 0 {
			categorydiscount = (productDetails[j].Price * float64(discount_percentageCategory)) / 100
		}
		//Apply category-level discount to the discounted price
		productDetails[j].DiscountedPrice = productDetails[j].DiscountedPrice - categorydiscount
	}
	//Initialize a slice to hold updated prd details
	var updatedproductDetails []models.ProductBreif
	//Loop through each prd detail to fetch and update image details
	for _, p := range productDetails {
		img, err := repository.GetImage(int(p.ID))
		if err != nil {
			return nil, err
		}
		//Updat product image details
		p.Image = img
		updatedproductDetails = append(updatedproductDetails, p)
	}
	//Return updated prd details and nil error
	return updatedproductDetails, nil
}
func FilterCategory(data map[string]int) ([]models.ProductBreif, error) {
	//Check if the provided category IDs are valid
	err := repository.CheckValidateCategory(data)
	if err != nil {
		//If validation fails , return an empty slice of ProductBrief and the error
		return []models.ProductBreif{}, err
	}
	//Initialize a slice to hold products from the specified categories
	var ProductFromCategory []models.ProductBreif
	//Iterate over each category ID
	for _, id := range data {
		//Retrieve products for the current category ID
		product, err := repository.GetProductFromCategory(id)
		if err != nil {
			//If there's an error fetching products, return an empty slice and error
			return []models.ProductBreif{}, err
		}
		//Iterate over fetched products
		for _, products := range product {
			//Retrieve stock quantity for the current product
			stock, err := repository.GetQuantityFromProductID(int(products.ID))
			if err != nil {
				//If there's and error fetching quantity
				return []models.ProductBreif{}, err
			}
			//Set Product status based on stock availability
			if stock <= 0 {
				products.ProductStatus = "out of stock"
			} else {
				products.ProductStatus = "in stock"
			}
			//Append the prd to the slice if its ID is not zero
			if products.ID != 0 {
				ProductFromCategory = append(ProductFromCategory, products)
			}
		}
	}
	//Iterate over products to calculate and apply discounts
	for j := range ProductFromCategory {
		//Find discount percentage for the current product
		discount_percentage, err := repository.FindDiscountPercentageForCategory(int(ProductFromCategory[j].ID))
		if err != nil {
			//If there's an error finding discount percentage , return and empty slice and error
			return []models.ProductBreif{}, errors.New("there was some error in finding the discounted prices")
		}
		var discount float64
		//Calculate discount if applicable
		if discount_percentage > 0 {
			discount = (ProductFromCategory[j].Price * float64(discount_percentage)) / 100
		}
		//Apply discount to calculate discounted price
		ProductFromCategory[j].DiscountedPrice = ProductFromCategory[j].Price - discount

		//Find category-level discount percentage for the current product
		discount_percentageCategory, err := repository.FindDiscountPercentageForCategory(ProductFromCategory[j].CatergoryID)
		if err != nil {
			return []models.ProductBreif{}, errors.New("there was some error in finding the discount prices")
		}

		var categorydiscount float64
		// Calculate category-level discount if applicable
		if discount_percentageCategory > 0 {
			categorydiscount = (ProductFromCategory[j].Price * float64(discount_percentageCategory)) / 100
		}
		// Apply category-level discount to the discounted price
		ProductFromCategory[j].DiscountedPrice = ProductFromCategory[j].DiscountedPrice - categorydiscount
	}
	//Initialize a slice to hold update product details
	updatedproductDetails := make([]models.ProductBreif, 0)
	//Iterate over products to fetch and update image details
	for _, p := range ProductFromCategory {
		img, err := repository.GetImage(int(p.ID))
		if err != nil {
			return nil, err
		}
		// Update product image details
		p.Image = img
		updatedproductDetails = append(updatedproductDetails, p)
	}
	// Return updated product details and nil error
	return updatedproductDetails, nil
}

func ShowAllProductsFromAdmin(page, count int) ([]models.ProductBreif, error) {
	productDetails, err := repository.ShowAllProductsFromAdmin(page, count)
	if err != nil {
		return []models.ProductBreif{}, err
	}
	for i := range productDetails {
		p := &productDetails[i]
		if p.Stock <= 0 {
			p.ProductStatus = "out of stock"
		} else {
			p.ProductStatus = "in stock"
		}
	}

	for j := range productDetails {
		discount_percentage, err := repository.FindDiscountPercentageForProduct(int(productDetails[j].ID))
		if err != nil {
			return []models.ProductBreif{}, errors.New("there was some error in finding the discounted prices")
		}
		var discount float64
		if discount_percentage > 0 {
			discount = (productDetails[j].Price * float64(discount_percentage)) / 100
		}
		productDetails[j].DiscountedPrice = productDetails[j].Price - discount

		discount_percentageCategory, err := repository.FindDiscountPercentageForCategory(int(productDetails[j].CatergoryID))
		if err != nil {
			return []models.ProductBreif{}, errors.New("there was some error in finding the discounted prices")
		}
		var Catergorydiscount float64
		if discount_percentageCategory > 0 {
			Catergorydiscount = (productDetails[j].Price * float64(discount_percentageCategory)) / 100
		}
		productDetails[j].DiscountedPrice = productDetails[j].DiscountedPrice - Catergorydiscount
	}
	var updatedPorductDetails []models.ProductBreif
	for _, p := range productDetails {
		img, err := repository.GetImage(int(p.ID))
		if err != nil {
			return nil, err
		}
		p.Image = img
		updatedPorductDetails = append(updatedPorductDetails, p)
	}
	return updatedPorductDetails, nil
}

func AddProducts(product models.Product) (domain.Product, error) {
	exist := repository.ProductAlreadyExist(product.Name)
	if exist {
		return domain.Product{}, errors.New("product already exist")
	}
	productResponse, err := repository.AddProducts(product)
	if err != nil {
		return domain.Product{}, err
	}
	stock := repository.StockInvalid(productResponse.Name)
	if !stock {
		return domain.Product{}, errors.New("stock is invalid input")
	}
	return productResponse, nil
}
