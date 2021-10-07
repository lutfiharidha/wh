package router

import (
	"ginWeb/app/controller"
	"ginWeb/app/repository"
	"ginWeb/app/service"

	"github.com/gin-gonic/gin"
)

var (
	productLocationRepository repository.ProductLocationRepository = repository.NewProductLocationRepository(db)

	productLocationService service.ProductLocationService = service.NewProductLocationService(productLocationRepository)

	productLocationController controller.ProductLocationController = controller.NewProductLocationController(productLocationService)
)

func ProductLocationRoute(route *gin.Engine) {

	productLocationRoutes := route.Group("api/v1/product-location")
	{
		productLocationRoutes.GET("/", productLocationController.All)
		productLocationRoutes.POST("/", productLocationController.Insert)
		productLocationRoutes.GET("/:id", productLocationController.FindByID)
		productLocationRoutes.PUT("/:id", productLocationController.Update)
		productLocationRoutes.DELETE("/:id", productLocationController.Delete)
		productLocationRoutes.GET("/deleted", productLocationController.Deleted)
		productLocationRoutes.PUT("/restore/:id", productLocationController.Restore)
		productLocationRoutes.DELETE("/delete/:id", productLocationController.DeletePermanent)
	}
}
