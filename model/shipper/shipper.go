package shipper

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "project/db"
	"strconv"
)

type Shipper struct {
	Id int `sql:"id,pk" json:"id"`
	ShipperName string `sql:"name" json:"name" binding:"required,validname,max=255"`
	Address string `sql:"address" json:"address" binding:"required"`
	TransportationMode string `sql:"transportation_mode" json:"transportation_mode" binding:"required,validtransportmode"`
	Pincode string `sql:"pincode" json:"pincode" binding:"required,validpin"`
	City string `sql:"city" json:"city" binding:"required,max=255"`
	State string `sql:"state" json:"state" binding:"required,max=255"`
	GSTIN string `sql:"gstin" json:"gstin" binding:"required,max=255"`
}

type UpdateShipper struct {
	ShipperName string `json:"name" binding:"validname,max=255"`
	Address string `json:"address"`
	TransportationMode string `json:"transportation_mode" binding:"validtransportmode"`
	Pincode string `json:"pincode" binding:"validpin"`
	City string `json:"city" binding:"max=255"`
	State string `json:"state" binding:"max=255"`
	GSTIN string `json:"gstin" binding:"max=255"`
}

//Add new shipper.
func AddShipper() func(context *gin.Context) {
	return func(context *gin.Context) {
		var shipper Shipper
		var err error
		if err = shipper.vAddShipper(context); err == nil {
			err = shipper.pAddShipper()
		}
		buildResponse(nil, err, context)
	}
}
//Validate new shipper before insertion.
func (shipper *Shipper)vAddShipper(context *gin.Context) (error) {
	err := context.ShouldBindJSON(shipper)
	return err
}
//Process to add new shipper to database.
func (shipper *Shipper) pAddShipper() (error){
	refDb := db.Connect()
	err := refDb.Insert(shipper)
	db.Close(refDb)
	return err
}

//List Shipper
func GetShipper() (func(context *gin.Context)) {
	return func(context *gin.Context) {
		res, err := pGetShipper(context.Param("name"))
		buildResponse(res, err, context)
	}
}

//Process to get list of shipper's from database.
func pGetShipper(name string) ([]Shipper, error) {
	refDb := db.Connect()
	var shipper []Shipper
	err := refDb.Model(&shipper).Where("name LIKE ?", name+"%").Select()
	db.Close(refDb)
	return shipper, err
}

//Delete shipper
func DeleteShipper() func(context *gin.Context) {
	return func(context *gin.Context) {
		id,_ := strconv.Atoi(context.Param("id"))
		err := pDeleteShipper(id)
		buildResponse(nil, err, context)
	}
}

//Process to delete shipper from database.
func pDeleteShipper(id int) (error) {
	refDb := db.Connect()
	var shipper Shipper = Shipper{Id: id}
	_,err := refDb.Model(&shipper).WherePK().Delete()
	return err
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

//Validate shipper details before update.
func (shipper *UpdateShipper) vUpdateShipperDetials(context *gin.Context) (error) {
	err := context.ShouldBindJSON(shipper)
	return err
}

//Process to update shipper details.
func (shipper *UpdateShipper) pUpdateShipperDetials(context *gin.Context) (error) {
	refDb := db.Connect()
	id,_ := strconv.Atoi(context.Param("id"))
	var obj Shipper = Shipper{
		Id: id,
		ShipperName: shipper.ShipperName,
		Address: shipper.Address,
		City: shipper.City,
		State: shipper.State,
		GSTIN: shipper.GSTIN,
		Pincode: shipper.Pincode,
		TransportationMode: shipper.TransportationMode,
	}
	_, err := refDb.Model(&obj).Where("id = ?id").UpdateNotNull()
	db.Close(refDb)
	return err
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
			"error": err,
		},
	)
}