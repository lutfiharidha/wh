package router

import (
	"ginWeb/app/controller"
	"ginWeb/app/repository"
	"ginWeb/app/service"

	"github.com/gin-gonic/gin"
)

var (
	zoneRepository repository.ZoneRepository = repository.NewZoneRepository(db)

	zoneService service.ZoneService = service.NewZoneService(zoneRepository)

	zoneController controller.ZoneController = controller.NewZoneController(zoneService)
)

func ZoneRoute(route *gin.Engine) {

	zoneRoutes := route.Group("api/v1/zone")
	{
		zoneRoutes.GET("/", zoneController.All)
		zoneRoutes.POST("/", zoneController.Insert)
		zoneRoutes.GET("/:id", zoneController.FindByID)
		zoneRoutes.PUT("/:id", zoneController.Update)
		zoneRoutes.DELETE("/:id", zoneController.Delete)
		zoneRoutes.GET("/deleted", zoneController.Deleted)
		zoneRoutes.PUT("/restore/:id", zoneController.Restore)
		zoneRoutes.DELETE("/delete/:id", zoneController.DeletePermanent)
	}
}
