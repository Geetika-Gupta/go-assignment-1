package routes

import (
	"github.com/gin-gonic/gin"
	shipperModel "project/model/shipper"
)

//initialize endpoints.
func InitializeRoutes(router *gin.Engine) {
	shipper := router.Group("/shipper")
	{
		shipper.POST("/", shipperModel.AddShipper())
		shipper.GET("/:name", shipperModel.GetShipper())
		shipper.PUT("/:id", shipperModel.UpdateShipperDetails())
		shipper.DELETE("/:id", shipperModel.DeleteShipper())
	}
}