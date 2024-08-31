package return_model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Return struct {
	gorm.Model
	Order_id uint `json:"order_id"`
	User_id uint `json:"user_id"`
	Order_item_id uint `json:"order_item_id"`
	Return_date time.Time `json:"return_date"`
	Reason string `json:"reason"`
}