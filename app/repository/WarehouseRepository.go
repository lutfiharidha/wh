package repository

import (
	// "github.com/ydhnwb/golang_api/entity"

	"ginWeb/app/model"

	"gorm.io/gorm"
)

type WarehouseRepository interface {
	InsertWarehouse(wh model.Warehouse) model.Warehouse
	UpdateWarehouse(wh model.Warehouse) model.Warehouse
	DeleteWarehouse(wh model.Warehouse)
	AllWarehouse() []model.Warehouse
	FindWarehouseByID(warehouseID string) model.Warehouse
	DeletedWarehouse() []model.Warehouse
	RestoreWarehouse(wh model.Warehouse) model.Warehouse
	DeletePermanentWarehouse(wh model.Warehouse)
}

type warehouseConnection struct {
	connection *gorm.DB
}

func NewWarehouseRepository(dbConn *gorm.DB) WarehouseRepository {
	return &warehouseConnection{
		connection: dbConn,
	}
}

func (db *warehouseConnection) InsertWarehouse(wh model.Warehouse) model.Warehouse {
	db.connection.Create(&wh)
	db.connection.Find(&wh)
	return wh
}

func (db *warehouseConnection) UpdateWarehouse(wh model.Warehouse) model.Warehouse {
	db.connection.Save(&wh)
	db.connection.Find(&wh)
	return wh
}

func (db *warehouseConnection) DeleteWarehouse(wh model.Warehouse) {
	db.connection.Delete(&wh)
}

func (db *warehouseConnection) FindWarehouseByID(warehouseID string) model.Warehouse {
	var warehouse model.Warehouse
	db.connection.Where("id =?", warehouseID).First(&warehouse)
	return warehouse
}

func (db *warehouseConnection) AllWarehouse() []model.Warehouse {
	var warehouses []model.Warehouse
	db.connection.Find(&warehouses)
	return warehouses
}

func (db *warehouseConnection) DeletedWarehouse() []model.Warehouse {
	var warehouses []model.Warehouse
	db.connection.Unscoped().Where("deleted_at != 0").Find(&warehouses)
	return warehouses
}

func (db *warehouseConnection) RestoreWarehouse(wh model.Warehouse) model.Warehouse {
	db.connection.Unscoped().First(&wh).Update("deleted_at", nil)
	db.connection.Find(&wh)
	return wh
}

func (db *warehouseConnection) DeletePermanentWarehouse(wh model.Warehouse) {
	db.connection.Unscoped().Delete(&wh)
}
