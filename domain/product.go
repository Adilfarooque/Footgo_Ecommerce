package domain

type Product struct {
	ID          uint    `json:"id" gorm:"unique:not null"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	CatergoryID uint    `json:"category_id"`
	Size        int     `json:"size"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}

type ProductImage struct {
	ID              uint   `json:"id" gorm:"unique:not null"`
	ProductImageUrl string `json:"product_image_url"`
}

type Image struct {
	ID        uint   `json:"id" gorm:"unique:not null"`
	ProductId uint   `json:"product_id"`
	Url       string `json:"url"`
}

type Category struct {
	ID       uint   `json:"id" gorm:"unique:not null"`
	Category string `json:"category" gorm:"unique:not null"`
}
