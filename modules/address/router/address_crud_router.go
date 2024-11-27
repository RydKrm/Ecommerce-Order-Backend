package address_crud_router

import (
	"delivery/auth"
	address_controller "delivery/modules/address/controler"

	"github.com/gin-gonic/gin"
)

func AddressCrudRoutes(r *gin.RouterGroup){
	r.POST("/",auth.Auth([]string{"user"}), address_controller.CreateAddress)
	r.GET("/",auth.Auth([]string{"user","seller","admin"}), address_controller.GetAllAddress)
	r.GET("/:id",auth.Auth([]string{"user","seller","admin"}), address_controller.DeleteSingleAddress)
	r.PATCH("/:id",auth.Auth([]string{"user"}), address_controller.UpdateAddress)
	r.DELETE("/:id",auth.Auth([]string{"user"}), address_controller.DeleteSingleAddress)
}
