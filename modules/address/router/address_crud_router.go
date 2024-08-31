package address_crud_router

import (
	address_controller "delivery/modules/address/controler"

	"github.com/gin-gonic/gin"
)

func AddressCrudRoutes(r *gin.RouterGroup){
	r.POST("/", address_controller.CreateAddress)
	r.GET("/", address_controller.GetAllAddress)
	r.GET("/:id", address_controller.DeleteSingleAddress)
	r.PATCH("/:id", address_controller.UpdateAddress)
	r.DELETE("/:id", address_controller.DeleteSingleAddress)
}
