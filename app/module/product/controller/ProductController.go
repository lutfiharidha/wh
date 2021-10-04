package controller

import (
	"fmt"
	"ginWeb/app/module/product/handler"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type APIEnv struct {
	DB *gorm.DB
}

func (a *APIEnv) GetBooks(c *gin.Context) {
	products, err := handler.GetBooks(a.DB)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, products)
}
