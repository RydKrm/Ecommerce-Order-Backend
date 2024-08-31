package address

import (
	address_crud_router "delivery/modules/address/router"

	"github.com/gin-gonic/gin"
)

func AddressRouter(router *gin.Engine){
	address := router.Group("/address")
	{
		address_crud_router.AddressCrudRoutes(address)
	}
}
