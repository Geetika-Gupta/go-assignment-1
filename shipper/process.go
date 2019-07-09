package shipper

import (
	"strconv"

	"github.com/Geetika-Gupta/go-assignment-1/migration"
)

//Process to add new shipper to database.
func pAddShipper(shipper *Shipper) error {
	refDb := migration.Connect()
	err := refDb.Insert(shipper)
	migration.Close(refDb)
	return err
}

//Process to get list of shipper's from database.
func pGetShipper(name string) ([]Shipper, error) {
	refDb := migration.Connect()
	var shipper []Shipper
	err := refDb.Model(&shipper).Where("name LIKE ?", name+"%").Select()
	migration.Close(refDb)
	return shipper, err
}

//Process to delete shipper from database.
func pDeleteShipper(id int) error {
	refDb := migration.Connect()
	var shipper = Shipper{ID: id}
	_, err := refDb.Model(&shipper).WherePK().Delete()
	return err
}

//Process to update shipper details.
func pUpdateShipperDetials(shipper *UpdateShipper, paramID string) error {
	refDb := migration.Connect()
	id, _ := strconv.Atoi(paramID)
	var obj = Shipper{
		ID:                 id,
		ShipperName:        shipper.ShipperName,
		Address:            shipper.Address,
		City:               shipper.City,
		State:              shipper.State,
		GSTIN:              shipper.GSTIN,
		Pincode:            shipper.Pincode,
		TransportationMode: shipper.TransportationMode,
	}
	_, err := refDb.Model(&obj).Where("id = ?id").UpdateNotNull()
	migration.Close(refDb)
	return err
}
