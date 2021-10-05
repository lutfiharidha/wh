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
}
