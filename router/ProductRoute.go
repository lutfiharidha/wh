package router

import (
	"ginWeb/app/controller"
	"ginWeb/app/repository"
	"ginWeb/app/service"

	"github.com/gin-gonic/gin"
)

var (
	productRepository repository.ProductRepository = repository.NewProductRepository(db)

	productService service.ProductService = service.NewProductService(productRepository)

	productController controller.ProductController = controller.NewProductController(productService)
)

func ProductRoute(route *gin.Engine) {

	productRoutes := route.Group("api/v1/product")
	{
		productRoutes.GET("/", productController.All)
		productRoutes.POST("/", productController.Insert)
		productRoutes.GET("/:id", productController.FindByID)
		productRoutes.PUT("/:id", productController.Update)
		productRoutes.DELETE("/:id", productController.Delete)
		productRoutes.GET("/deleted", productController.Deleted)
		productRoutes.PUT("/restore/:id", productController.Restore)
		productRoutes.DELETE("/delete/:id", productController.DeletePermanent)
	}
}
