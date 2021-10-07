package repository

import (
	// "github.com/ydhnwb/golang_api/entity"

	"ginWeb/app/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductLocationRepository interface {
	InsertProductLocation(pl model.ProductLocation) model.ProductLocation
	UpdateProductLocation(pl model.ProductLocation) model.ProductLocation
	DeleteProductLocation(pl model.ProductLocation) model.ProductLocation
	AllProductLocation() []model.ProductLocation
	FindProductLocationByID(productID string) model.ProductLocation
	DeletedProductLocation() []model.ProductLocation
	RestoreProductLocation(pl model.ProductLocation) model.ProductLocation
	DeletePermanentProductLocation(pl model.ProductLocation)
}

type productLocationConnection struct {
	connection *gorm.DB
}

func NewProductLocationRepository(dbConn *gorm.DB) ProductLocationRepository {
	return &productLocationConnection{
		connection: dbConn,
	}
}

func (db *productLocationConnection) InsertProductLocation(pl model.ProductLocation) model.ProductLocation {
	db.connection.Create(&pl)
	db.connection.Joins("Product").Preload("Rack").Find(&pl)
	return pl
}

func (db *productLocationConnection) UpdateProductLocation(pl model.ProductLocation) model.ProductLocation {
	db.connection.Save(&pl)
	db.connection.Preload("Product").Preload("Rack").Find(&pl)
	return pl
}

func (db *productLocationConnection) DeleteProductLocation(pl model.ProductLocation) model.ProductLocation {
	db.connection.Delete(&pl)
	db.connection.Preload("Product").Preload("Rack").Find(&pl)
	return pl
}

func (db *productLocationConnection) AllProductLocation() []model.ProductLocation {
	var pl []model.ProductLocation
	db.connection.Preload("Product").Preload("Rack").Find(&pl)
	return pl
}

func (db *productLocationConnection) FindProductLocationByID(productLocationID string) model.ProductLocation {
	var pl model.ProductLocation
	db.connection.Where("id =?", productLocationID).Preload(clause.Associations).First(&pl)
	return pl
}

func (db *productLocationConnection) DeletedProductLocation() []model.ProductLocation {
	var pl []model.ProductLocation
	db.connection.Unscoped().Where("deleted_at != 0").Preload("Product").Preload("Rack").Find(&pl)
	return pl
}

func (db *productLocationConnection) RestoreProductLocation(pl model.ProductLocation) model.ProductLocation {
	db.connection.Unscoped().First(&pl).Update("deleted_at", nil)
	db.connection.Preload("Product").Preload("Rack").Find(&pl)
	return pl
}

func (db *productLocationConnection) DeletePermanentProductLocation(pl model.ProductLocation) {
	db.connection.Unscoped().Delete(&pl)
}
