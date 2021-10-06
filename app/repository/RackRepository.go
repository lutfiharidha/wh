package repository

import (
	// "github.com/ydhnwb/golang_api/entity"

	"ginWeb/app/model"

	"gorm.io/gorm"
)

type RackRepository interface {
	InsertRack(wh model.Rack) model.Rack
	UpdateRack(wh model.Rack) model.Rack
	DeleteRack(wh model.Rack) model.Rack
	AllRack() []model.Rack
	FindRackByID(rackID string) model.Rack
	DeletedRack() []model.Rack
	RestoreRack(wh model.Rack) model.Rack
	DeletePermanentRack(wh model.Rack)
}

type rackConnection struct {
	connection *gorm.DB
}

func NewRackRepository(dbConn *gorm.DB) RackRepository {
	return &rackConnection{
		connection: dbConn,
	}
}

func (db *rackConnection) InsertRack(rack model.Rack) model.Rack {
	db.connection.Create(&rack)
	db.connection.Preload("Zone").Find(&rack)
	return rack
}

func (db *rackConnection) UpdateRack(rack model.Rack) model.Rack {
	db.connection.Save(&rack)
	db.connection.Preload("Zone").Find(&rack)
	return rack
}

func (db *rackConnection) DeleteRack(rack model.Rack) model.Rack {
	db.connection.Delete(&rack)
	db.connection.Preload("Zone").Find(&rack)
	return rack
}

func (db *rackConnection) AllRack() []model.Rack {
	var racks []model.Rack
	db.connection.Preload("Zone").Find(&racks)
	return racks
}

func (db *rackConnection) FindRackByID(rackID string) model.Rack {
	var rack model.Rack
	db.connection.Where("id =?", rackID).Preload("Zone").First(&rack)
	return rack
}

func (db *rackConnection) DeletedRack() []model.Rack {
	var racks []model.Rack
	db.connection.Unscoped().Where("deleted_at != 0").Preload("Zone").Find(&racks)
	return racks
}

func (db *rackConnection) RestoreRack(rack model.Rack) model.Rack {
	db.connection.Unscoped().First(&rack).Update("deleted_at", nil)
	db.connection.Preload("Zone").Find(&rack)
	return rack
}

func (db *rackConnection) DeletePermanentRack(rack model.Rack) {
	db.connection.Unscoped().Delete(&rack)
}
