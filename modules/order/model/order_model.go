package order_model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type OrderStatus string

const (
	StatusPending OrderStatus = "pending"
	StatusDelivered OrderStatus = "delivered"
	StatusWarehouse OrderStatus = "warehouse"
)

type Order struct {
	gorm.Model
	User_id uint `json:"user_id"`
	Address_id uint `json:"address_id"`
	Order_date time.Time `json:"order_date"`
	Total_amount uint64 `json:"total_amount"`
	Status OrderStatus `json:"status"`
	OrderItems []OrderItem `gorm:"foreignKey:Order_id" json:"order_items"`
}

type OrderItem struct{
	gorm.Model
	Order_id uint `json:"order_id"`
	Product_name string `json:"product_name"`
	Product_id string `json:"product_id"`
	Shop_id string `json:"shop_id"`
	Quantity uint `json:"quantity"`
	Price uint `json:"price"`
}