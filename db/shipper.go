package db

import (
	"github.com/Geetika-Gupta/go-assignment-1/migration"
	"github.com/Geetika-Gupta/go-assignment-1/shipper"
	"github.com/go-pg/pg/orm"
)

//CreateShipper creates a table if not exists.
func CreateShipper() {
	dbRef := migration.Connect()
	ormOptions := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	dbRef.CreateTable(&shipper.Shipper{}, ormOptions)
	migration.Close(dbRef)
}
