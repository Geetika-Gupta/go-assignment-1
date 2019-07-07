package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Geetika-Gupta/go-assignment-1/shipper_modal"
)

//initialize endpoints.
func InitializeRoutes(router *gin.Engine) {
	shipper := router.Group("/shipper")
	{
		shipper.POST("/", shipper_modal.AddShipper())
		shipper.GET("/:name", shipper_modal.GetShipper())
		shipper.PUT("/:id", shipper_modal.UpdateShipperDetails())
		shipper.DELETE("/:id", shipper_modal.DeleteShipper())
	}
}