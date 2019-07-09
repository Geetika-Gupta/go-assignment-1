package shipper

import (
	"github.com/gin-gonic/gin"
)

//Validate new shipper before insertion.
func vAddShipper(context *gin.Context, shipper *Shipper) error {
	err := context.ShouldBindJSON(shipper)
	return err
}

//Validate shipper details before update.
func vUpdateShipperDetials(context *gin.Context, shipper *UpdateShipper) error {
	err := context.ShouldBindJSON(shipper)
	return err
}
