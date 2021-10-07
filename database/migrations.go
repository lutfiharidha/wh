package database

import (
	"ginWeb/app/config"
	"ginWeb/app/model"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
)

func InitialMigration() {
	db.Migrator().DropTable(&model.Warehouse{})
	db.Migrator().CreateTable(&model.Warehouse{})

	db.Migrator().DropTable(&model.Product{})
	db.Migrator().CreateTable(&model.Product{})

	db.Migrator().DropTable(&model.Zone{})
	db.Migrator().CreateTable(&model.Zone{})

	db.Migrator().DropTable(&model.Rack{})
	db.Migrator().CreateTable(&model.Rack{})

	db.Migrator().DropTable(&model.ProductLocation{})
	db.Migrator().CreateTable(&model.ProductLocation{})
}
