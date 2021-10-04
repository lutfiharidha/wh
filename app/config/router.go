package config

import (
	"ginWeb/app/module/warehouse/router"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	router.V1Warehouse(r)
	return r
}

func Run() *gin.Engine {
	r := SetupRouter()
	r.Run()

	return r
}
