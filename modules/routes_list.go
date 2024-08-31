package modules

import (
	"delivery/modules/address"
	"delivery/modules/order"

	"github.com/gin-gonic/gin"
)

func RouteList(router *gin.Engine){

	// Order routes
	order.OrderRoutes(router)

	// Address routes
	address.AddressRouter(router)

}