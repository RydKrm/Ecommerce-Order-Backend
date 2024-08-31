package product_model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name string `json:"name"`
	Price float64 `json:"price"`
	Description string `json:"description"`
}