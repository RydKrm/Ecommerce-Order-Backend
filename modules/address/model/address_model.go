package address_model

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	User_id uint `json:"user_id"`
	District string `json:"district"`
	Division string `json:"division"`
	Road string `json:"road"`
	House string `json:"house"`
	Description string `json:"description"`
	Status bool `json:"status" gorm:"default:true"`
}