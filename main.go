package main

import (
	"ginWeb/app/config"
	"ginWeb/router"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = config.SetupDatabaseConnection()

func main() {
	defer config.CloseDatabaseConnection(db)

	// database.InitialMigration()
	// database.InitDbialSeeder()

	r := gin.Default()
	router.WarehouseRoute(r) //Added all auth routes
	router.ProductRoute(r)   //Added all user routes
	r.Run()
}
