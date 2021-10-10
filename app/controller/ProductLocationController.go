package controller

import (
	"ginWeb/app/dto"
	"ginWeb/app/helper"
	"ginWeb/app/model"
	"ginWeb/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductLocationController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	Deleted(context *gin.Context)
	Restore(context *gin.Context)
	DeletePermanent(context *gin.Context)
}

type productLocationController struct {
	productLocationService service.ProductLocationService
}

func NewProductLocationController(productLocationServ service.ProductLocationService) ProductLocationController {
	return &productLocationController{
		productLocationService: productLocationServ,
	}
}

func (c *productLocationController) All(context *gin.Context) {
	var pl []model.ProductLocation = c.productLocationService.All()
	res := helper.BuildResponse(true, "OK", pl)
	context.JSON(http.StatusOK, res)
}

func (c *productLocationController) FindByID(context *gin.Context) {
	id := context.Param("id")
	var pl model.ProductLocation = c.productLocationService.FindByID(id)
	if (pl == model.ProductLocation{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", pl)
		context.JSON(http.StatusOK, res)
	}
}

func (c *productLocationController) Insert(context *gin.Context) {
	var productLocationCreateDTO dto.ProductLocationCreateDTO
	errDTO := context.ShouldBind(&productLocationCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
	result := c.productLocationService.Insert(productLocationCreateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusCreated, response)
}

func (c *productLocationController) Update(context *gin.Context) {
	var productLocationUpdateDTO dto.ProductLocationUpdateDTO
	errDTO := context.ShouldBind(&productLocationUpdateDTO)
	id := context.Param("id")
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	var product model.ProductLocation = c.productLocationService.FindByID(id)
	if (product == model.ProductLocation{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		productLocationUpdateDTO.ID = id
		result := c.productLocationService.Update(productLocationUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *productLocationController) Delete(context *gin.Context) {
	var pl model.ProductLocation
	id := context.Param("id")
	pl = c.productLocationService.FindByID(id)
	if (pl == model.ProductLocation{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		pl.ID = id
		result := c.productLocationService.Delete(pl)
		res := helper.BuildResponse(true, "Deleted", result)
		context.JSON(http.StatusOK, res)
	}
}

func (c *productLocationController) Deleted(context *gin.Context) {
	var pl []model.ProductLocation = c.productLocationService.Deleted()
	res := helper.BuildResponse(true, "OK", pl)
	context.JSON(http.StatusOK, res)
}

func (c *productLocationController) Restore(context *gin.Context) {
	var productLocationRestoreDTO dto.ProductLocationRestoreDTO
	id := context.Param("id")
	var pl model.ProductLocation = c.productLocationService.FindByID(id)
	if (pl != model.ProductLocation{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		productLocationRestoreDTO.ID = id
		result := c.productLocationService.Restore(productLocationRestoreDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *productLocationController) DeletePermanent(context *gin.Context) {
	var pl model.ProductLocation
	id := context.Param("id")
	pl = c.productLocationService.FindByID(id)
	if (pl != model.ProductLocation{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		pl.ID = id
		c.productLocationService.DeletePermanent(pl)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	}
}
