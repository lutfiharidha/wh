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
	// database.InitialDBSeeder()
	r := gin.New()
	r.Use(config.CORSMiddleware())

	router.WarehouseRoute(r)       //Added all warehouse routes
	router.ProductRoute(r)         //Added all product routes
	router.ZoneRoute(r)            //Added all zone routes
	router.RackRoute(r)            //Added all rack routes
	router.ProductLocationRoute(r) //Added all product location routes
	r.Run()
}
