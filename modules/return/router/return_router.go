package return_crud_router

import (
	return_controller "delivery/modules/return/controller"

	"github.com/gin-gonic/gin"
)

func ReturnCrudRoutes(r *gin.RouterGroup){
	r.POST("/", return_controller.CreateReturn)
	r.GET("/allByUser/:id", return_controller.GetAllReturn)
	r.GET("/:id", return_controller.GetSingleReturn)
	r.PATCH("/:id", return_controller.UpdateReturn)
	r.DELETE("/:id", return_controller.DeleteSingleReturn)
}
