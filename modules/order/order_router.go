package order

import (
	order_crud_router "delivery/modules/order/router"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine){
	order := router.Group("/order")
	{
		order_crud_router.OrderCrudRoutes(order)
	}
}