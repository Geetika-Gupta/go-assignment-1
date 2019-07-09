package shipper

import (
	"github.com/gin-gonic/gin"
)

//InitializeRoutes defines endpoints.
func InitializeRoutes(router *gin.Engine) {
	shipper := router.Group("/shipper")
	{
		shipper.POST("/", addShipper())
		shipper.GET("/:name", getShipper())
		shipper.PUT("/:id", updateShipperDetails())
		shipper.DELETE("/:id", deleteShipper())
	}
}
