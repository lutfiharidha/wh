package repository

import (
	// "github.com/ydhnwb/golang_api/entity"

	"ginWeb/app/model"

	"gorm.io/gorm"
)

type ZoneRepository interface {
	InsertZone(wh model.Zone) model.Zone
	UpdateZone(wh model.Zone) model.Zone
	DeleteZone(wh model.Zone) model.Zone
	AllZone() []model.Zone
	FindZoneByID(zoneID string) model.Zone
	DeletedZone() []model.Zone
	RestoreZone(wh model.Zone) model.Zone
	DeletePermanentZone(wh model.Zone)
}

type zoneConnection struct {
	connection *gorm.DB
}

func NewZoneRepository(dbConn *gorm.DB) ZoneRepository {
	return &zoneConnection{
		connection: dbConn,
	}
}

func (db *zoneConnection) InsertZone(zone model.Zone) model.Zone {
	db.connection.Create(&zone)
	db.connection.Preload("Warehouse").Find(&zone)
	return zone
}

func (db *zoneConnection) UpdateZone(zone model.Zone) model.Zone {
	db.connection.Save(&zone)
	db.connection.Preload("Warehouse").Find(&zone)
	return zone
}

func (db *zoneConnection) DeleteZone(zone model.Zone) model.Zone {
	db.connection.Delete(&zone)
	db.connection.Preload("Warehouse").Find(&zone)
	return zone
}

func (db *zoneConnection) AllZone() []model.Zone {
	var zones []model.Zone
	db.connection.Preload("Warehouse").Find(&zones)
	return zones
}

func (db *zoneConnection) FindZoneByID(zoneID string) model.Zone {
	var zone model.Zone
	db.connection.Where("id =?", zoneID).Preload("Warehouse").First(&zone)
	return zone
}

func (db *zoneConnection) DeletedZone() []model.Zone {
	var zones []model.Zone
	db.connection.Unscoped().Where("deleted_at != 0").Preload("Warehouse").Find(&zones)
	return zones
}

func (db *zoneConnection) RestoreZone(zone model.Zone) model.Zone {
	db.connection.Unscoped().First(&zone).Update("deleted_at", nil)
	db.connection.Preload("Warehouse").Find(&zone)
	return zone
}

func (db *zoneConnection) DeletePermanentZone(zone model.Zone) {
	db.connection.Unscoped().Delete(&zone)
}
