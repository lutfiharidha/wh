package router

import (
	"ginWeb/app/controller"
	"ginWeb/app/repository"
	"ginWeb/app/service"

	"github.com/gin-gonic/gin"
)

var (
	rackRepository repository.RackRepository = repository.NewRackRepository(db)

	rackService service.RackService = service.NewRackService(rackRepository)

	rackController controller.RackController = controller.NewRackController(rackService)
)

func RackRoute(route *gin.Engine) {

	rackRoutes := route.Group("api/v1/rack")
	{
		rackRoutes.GET("/", rackController.All)
		rackRoutes.POST("/", rackController.Insert)
		rackRoutes.GET("/:id", rackController.FindByID)
		rackRoutes.PUT("/:id", rackController.Update)
		rackRoutes.DELETE("/:id", rackController.Delete)
		rackRoutes.GET("/deleted", rackController.Deleted)
		rackRoutes.PUT("/restore/:id", rackController.Restore)
		rackRoutes.DELETE("/delete/:id", rackController.DeletePermanent)
	}
}
