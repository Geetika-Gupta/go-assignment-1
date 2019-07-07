package shipper_modal

import (
	"github.com/gin-gonic/gin"
	"github.com/Geetika-Gupta/go-assignment-1/modal"
)

//Validate new shipper before insertion.
func vAddShipper(context *gin.Context, shipper *modal.Shipper) (error) {
	err := context.ShouldBindJSON(shipper)
	return err
}

//Validate shipper details before update.
func vUpdateShipperDetials(context *gin.Context, shipper *modal.UpdateShipper) (error) {
	err := context.ShouldBindJSON(shipper)
	return err
}