package return_pack

import (
	return_crud_router "delivery/modules/return/router"

	"github.com/gin-gonic/gin"
)

func ReturnRouter(router *gin.Engine) {
	return_router := router.Group("/return")
	{
		return_crud_router.ReturnCrudRoutes(return_router)
	}
}