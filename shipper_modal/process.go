package shipper_modal

import (
	"github.com/Geetika-Gupta/go-assignment-1/modal"
	"github.com/Geetika-Gupta/go-assignment-1/db"
)

//Process to add new shipper to database.
func pAddShipper(shipper *modal.Shipper) (error){
	refDb := db.Connect()
	err := refDb.Insert(shipper)
	db.Close(refDb)
	return err
}

//Process to get list of shipper's from database.
func pGetShipper(name string) ([]modal.Shipper, error) {
	refDb := db.Connect()
	var shipper []Shipper
	err := refDb.Model(&shipper).Where("name LIKE ?", name+"%").Select()
	db.Close(refDb)
	return shipper, err
}

//Process to delete shipper from database.
func pDeleteShipper(id int) (error) {
	refDb := db.Connect()
	var shipper Shipper = Shipper{Id: id}
	_,err := refDb.Model(&shipper).WherePK().Delete()
	return err
}

//Process to update shipper details.
func pUpdateShipperDetials(shipper *modal.UpdateShipper) (error) {
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