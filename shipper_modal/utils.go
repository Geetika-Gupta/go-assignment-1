package shipper_modal

import (
	"github.com/Geetika-Gupta/go-assignment-1/modal"
	"github.com/gin-gonic/gin"
)

//Add new shipper.
func AddShipper() func(context *gin.Context) {
	return func(context *gin.Context) {
		var shipper modal.Shipper
		var err error
		if err = shipper.vAddShipper(context); err == nil {
			err = shipper.pAddShipper()
		}
		buildResponse(nil, err, context)
	}
}

//List Shipper
func GetShipper() (func(context *gin.Context)) {
	return func(context *gin.Context) {
		res, err := pGetShipper(context.Param("name"))
		buildResponse(res, err, context)
	}
}

//Delete shipper
func DeleteShipper() func(context *gin.Context) {
	return func(context *gin.Context) {
		id,_ := strconv.Atoi(context.Param("id"))
		err := pDeleteShipper(id)
		buildResponse(nil, err, context)
	}
}

//Update shipper details.
func UpdateShipperDetails() func(context *gin.Context) {
	return func(context *gin.Context) {
		var shipper UpdateShipper
		var err error
		if err = shipper.vUpdateShipperDetials(context); err == nil {
			err = shipper.pUpdateShipperDetials(context)
		}
		buildResponse(nil, err, context)
	}
}

//Response handler.
func buildResponse(res []modal.Shipper, err error, context *gin.Context) {
	var status int
	if err == nil {
		status = http.StatusOK
	} else {
		status = http.StatusBadRequest
	}
	context.JSON(
		status,
		gin.H{
			"status": status,
			"result": res,
			"error": err,
		},
	)
}