package modules

import (
	"delivery/modules/address"
	"delivery/modules/order"
	return_pack "delivery/modules/return"

	"github.com/gin-gonic/gin"
)

func RouteList(router *gin.Engine){

	// Order routes
	order.OrderRoutes(router)

	// Address routes
	address.AddressRouter(router)

	// Return routes
	return_pack.ReturnRouter(router)

}