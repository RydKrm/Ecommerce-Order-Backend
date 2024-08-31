package order_crud_router

import (
	order_crud_controller "delivery/modules/order/controler"

	"github.com/gin-gonic/gin"
)

func OrderCrudRoutes(r *gin.RouterGroup){
	r.POST("/create", order_crud_controller.CreateOrder);
	r.PATCH("/update/:id", order_crud_controller.UpdateOrder);
	r.GET("/single/:id", order_crud_controller.GetSingleOrder);
	r.GET("/all", order_crud_controller.GetAllOrder);
	r.DELETE("/delete/:id", order_crud_controller.DeleteSingleOrder);
}
