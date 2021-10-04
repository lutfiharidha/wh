package controller

import (
	"ginWeb/app/config"
	"ginWeb/app/module/warehouse/handler"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type APIEnv struct {
	DB *gorm.DB
}

func (a *APIEnv) GetWarehouses(c *gin.Context) {
	var res config.Response
	warehouses, err := handler.GetWarehouses(a.DB)
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = warehouses
	c.JSON(http.StatusOK, res)
}
