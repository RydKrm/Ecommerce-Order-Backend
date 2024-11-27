package order_crud_router

import (
	"delivery/auth"
	order_crud_controller "delivery/modules/order/controler"

	"github.com/gin-gonic/gin"
)

func OrderCrudRoutes(r *gin.RouterGroup){
	r.POST("/create",auth.Auth([]string{"user"}), order_crud_controller.CreateOrder);
	r.PATCH("/update/:id",auth.Auth([]string{"user"}), order_crud_controller.UpdateOrder);
	r.GET("/single/:id",auth.Auth([]string{"user"}), order_crud_controller.GetSingleOrder);
	r.GET("/all",auth.Auth([]string{"user"}), order_crud_controller.GetAllOrder);
	r.DELETE("/delete/:id",auth.Auth([]string{"user"}), order_crud_controller.DeleteSingleOrder);
}
