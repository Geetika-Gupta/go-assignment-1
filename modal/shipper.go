package modal

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