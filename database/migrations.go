package database

import (
	"ginWeb/app/config"
	"ginWeb/app/module/product/model"
)

func InitialMigration() {
	config.Db.Migrator().CreateTable(&model.Product{})
}
