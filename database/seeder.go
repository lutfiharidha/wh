package database

import (
	"ginWeb/app/model"

	"github.com/google/uuid"
)

func WarehouseSeeder() {
	db.Create(&model.Warehouse{
		ID:            uuid.New().String(),
		WarehouseName: "Gudang Baju A",
		// CompanyId:     "",
		Owner:       "Lutfi Haridha",
		PhoneNumber: "081376867436",
		Status:      "1",
		Address:     "Jalan Kebon Raya No 90, Kel. Duri Kepa, Kec. Kebon Jeruk",
		City:        "Jakarta Barat",
		Country:     "Indonesia",
		Geolocation: "-6.185190, 106.778040",
	})

}

func ProductSeeder() {
	warehouses := model.Warehouse{}
	db.Select("id").Last(&warehouses)
	db.Create(&model.Product{
		ID:          uuid.New().String(),
		ProductName: "Baju Tidur",
		WarehouseID: warehouses.ID,
		SKU:         "pcs",
		Description: "Lorem ipsum",
		Price:       20000.00,
		Image:       "noimage.jpg",
		Status:      1,
	})
}

func ZoneSeeder() {
	warehouse := model.Warehouse{}
	db.Select("id").Last(&warehouse)

	db.Create(&model.Zone{
		ID:          uuid.New().String(),
		ZoneName:    "Food",
		WarehouseID: warehouse.ID,
	})
}

func RackSeeder() {
	zone := model.Zone{}
	db.Select("id").Last(&zone)

	db.Create(&model.Rack{
		ID:       uuid.New().String(),
		ZoneID:   zone.ID,
		Aisle:    "01",
		Rack:     "02",
		Level:    "1",
		Position: "1",
	})
}

func ProductLocationSeeder() {
	location := model.Rack{}
	db.Select("id").Last(&location)
	product := model.Product{}
	db.Select("id").Last(&product)
	db.Create(&model.ProductLocation{
		ID:        uuid.New().String(),
		ProductID: product.ID,
		RackID:    location.ID,
		Stock:     10,
	})
}

func InitialDBSeeder() {
	WarehouseSeeder()
	ProductSeeder()
	ZoneSeeder()
	RackSeeder()
	ProductLocationSeeder()
}
