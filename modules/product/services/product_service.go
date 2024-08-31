package product_service

import (
	"delivery/database"
	product_model "delivery/modules/product/models"
)

func CreateProduct(product *product_model.Product) error {
	return database.DB.Create(product).Error
}

func GetAllProducts() ([]product_model.Product, error) {
	var products []product_model.Product
    err := database.DB.Find(&products).Error
	return products, err
}

func GetProductByID(id uint) (product_model.Product, error) {
	var product product_model.Product
	err := database.DB.First(&product, id).Error
	return product, err
}

func UpdateProduct(product *product_model.Product) error {
	return database.DB.Save(product).Error
}

func DeleteProduct(id uint) error {
	return database.DB.Delete(&product_model.Product{}, id).Error
}