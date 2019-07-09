package shipper

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Add new shipper.
func addShipper() func(context *gin.Context) {
	return func(context *gin.Context) {
		var shipper Shipper
		var err error
		if err = vAddShipper(context, &shipper); err == nil {
			err = pAddShipper(shipper)
		}
		buildResponse(nil, err, context)
	}
}

//List Shipper.
func getShipper() func(context *gin.Context) {
	return func(context *gin.Context) {
		res, err := pGetShipper(context.Param("name"))
		buildResponse(res, err, context)
	}
}

//Delete shipper.
func deleteShipper() func(context *gin.Context) {
	return func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))
		err := pDeleteShipper(id)
		buildResponse(nil, err, context)
	}
}

//Update shipper details.
func updateShipperDetails() func(context *gin.Context) {
	return func(context *gin.Context) {
		var shipper UpdateShipper
		var err error
		if err = vUpdateShipperDetials(context, &shipper); err == nil {
			err = pUpdateShipperDetials(shipper, context.Param("id"))
		}
		buildResponse(nil, err, context)
	}
}

//Response handler.
func buildResponse(res []Shipper, err error, context *gin.Context) {
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
			"error":  err,
		},
	)
}
