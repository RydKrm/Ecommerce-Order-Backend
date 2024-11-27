package return_crud_router

import (
	"delivery/auth"
	return_controller "delivery/modules/return/controller"

	"github.com/gin-gonic/gin"
)

func ReturnCrudRoutes(r *gin.RouterGroup){
	r.POST("/", auth.Auth([]string{"user"}), return_controller.CreateReturn)
	r.GET("/allByUser/:id",auth.Auth([]string{"user","seller","admin"}), return_controller.GetAllReturn)
	r.GET("/:id",auth.Auth([]string{"user","seller","admin"}), return_controller.GetSingleReturn)
	r.PATCH("/:id",auth.Auth([]string{"user"}), return_controller.UpdateReturn)
	r.DELETE("/:id",auth.Auth([]string{"user"}), return_controller.DeleteSingleReturn)
}
