package router

import (
	"ginWeb/app/config"
	"ginWeb/app/module/product/controller"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	api := controller.APIEnv{
		DB: config.InitDb(),
	}

	r.GET("/products", api.GetBooks)
	// r.GET("/:id", api.GetBook)
	// r.POST("", api.CreateBook)
	// r.PUT("/:id", api.UpdateBook)
	// r.DELETE("/:id", api.DeleteBook)

	return r
}
