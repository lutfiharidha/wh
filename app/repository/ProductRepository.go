package repository

import (
	// "github.com/ydhnwb/golang_api/entity"

	"ginWeb/app/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	InsertProduct(wh model.Product) model.Product
	UpdateProduct(wh model.Product) model.Product
	DeleteProduct(wh model.Product) model.Product
	AllProduct() []model.Product
	FindProductByID(productID string) model.Product
	DeletedProduct() []model.Product
	RestoreProduct(wh model.Product) model.Product
	DeletePermanentProduct(wh model.Product)
}

type productConnection struct {
	connection *gorm.DB
}

func NewProductRepository(dbConn *gorm.DB) ProductRepository {
	return &productConnection{
		connection: dbConn,
	}
}

func (db *productConnection) InsertProduct(prod model.Product) model.Product {
	db.connection.Create(&prod)
	db.connection.Preload("Warehouse").Find(&prod)
	return prod
}

func (db *productConnection) UpdateProduct(prod model.Product) model.Product {
	db.connection.Save(&prod)
	db.connection.Preload("Warehouse").Find(&prod)
	return prod
}

func (db *productConnection) DeleteProduct(prod model.Product) model.Product {
	db.connection.Delete(&prod)
	db.connection.Preload("Warehouse").Find(&prod)
	return prod
}

func (db *productConnection) AllProduct() []model.Product {
	var products []model.Product
	db.connection.Preload("Warehouse").Find(&products)
	return products
}

func (db *productConnection) FindProductByID(productID string) model.Product {
	var product model.Product
	db.connection.Where("id =?", productID).Preload("Warehouse").First(&product)
	return product
}

func (db *productConnection) DeletedProduct() []model.Product {
	var products []model.Product
	db.connection.Unscoped().Where("deleted_at != 0").Preload("Warehouse").Find(&products)
	return products
}

func (db *productConnection) RestoreProduct(prod model.Product) model.Product {
	db.connection.Unscoped().First(&prod).Update("deleted_at", nil)
	db.connection.Preload("Warehouse").Find(&prod)
	return prod
}

func (db *productConnection) DeletePermanentProduct(prod model.Product) {
	db.connection.Unscoped().Delete(&prod)
}
