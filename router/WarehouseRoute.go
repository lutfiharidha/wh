package router

import (
	"ginWeb/app/config"
	"ginWeb/app/controller"
	"ginWeb/app/repository"
	"ginWeb/app/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()

	warehouseRepository repository.WarehouseRepository = repository.NewWarehouseRepository(db)

	warehouseService service.WarehouseService = service.NewWarehouseService(warehouseRepository)

	warehouseController controller.WarehouseController = controller.NewWarehouseController(warehouseService)
)

func WarehouseRoute(route *gin.Engine) {
	warehouseRoutes := route.Group("api/v1/warehouse")
	{
		warehouseRoutes.GET("/", warehouseController.All)
		warehouseRoutes.POST("/", warehouseController.Insert)
		warehouseRoutes.GET("/:id", warehouseController.FindByID)
		warehouseRoutes.PUT("/:id", warehouseController.Update)
		warehouseRoutes.DELETE("/:id", warehouseController.Delete)
		warehouseRoutes.GET("/deleted", warehouseController.Deleted)
		warehouseRoutes.PUT("/restore/:id", warehouseController.Restore)
		warehouseRoutes.DELETE("/delete/:id", warehouseController.DeletePermanent)
	}
}
