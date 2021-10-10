package controller

import (
	"ginWeb/app/dto"
	"ginWeb/app/helper"
	"ginWeb/app/model"
	"ginWeb/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WarehouseController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	Deleted(context *gin.Context)
	Restore(context *gin.Context)
	DeletePermanent(context *gin.Context)
}

type warehouseController struct {
	warehouseService service.WarehouseService
}

func NewWarehouseController(warehouseServ service.WarehouseService) WarehouseController {
	return &warehouseController{
		warehouseService: warehouseServ,
	}
}

func (c *warehouseController) All(context *gin.Context) {
	var warehouses []model.Warehouse = c.warehouseService.All()
	res := helper.BuildResponse(true, "OK", warehouses)
	context.Writer.Header().Set("Content-Type", "application/json")
	context.JSON(http.StatusOK, res)
}

func (c *warehouseController) FindByID(context *gin.Context) {
	id := context.Param("id")
	var warehouse model.Warehouse = c.warehouseService.FindByID(id)
	if (warehouse == model.Warehouse{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})

		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", warehouse)
		context.JSON(http.StatusOK, res)
	}
}

func (c *warehouseController) Insert(context *gin.Context) {
	var warehouseCreateDTO dto.WarehouseCreateDTO
	errDTO := context.ShouldBind(&warehouseCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
	result := c.warehouseService.Insert(warehouseCreateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusCreated, response)
}

func (c *warehouseController) Update(context *gin.Context) {
	var warehouseUpdateDTO dto.WarehouseUpdateDTO
	errDTO := context.ShouldBind(&warehouseUpdateDTO)
	id := context.Param("id")
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	var warehouse model.Warehouse = c.warehouseService.FindByID(id)
	if (warehouse == model.Warehouse{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		warehouseUpdateDTO.ID = id
		result := c.warehouseService.Update(warehouseUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *warehouseController) Delete(context *gin.Context) {
	var warehouse model.Warehouse
	id := context.Param("id")
	warehouse = c.warehouseService.FindByID(id)
	if (warehouse == model.Warehouse{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		warehouse.ID = id
		c.warehouseService.Delete(warehouse)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	}
}

func (c *warehouseController) Deleted(context *gin.Context) {
	var warehouses []model.Warehouse = c.warehouseService.Deleted()
	res := helper.BuildResponse(true, "OK", warehouses)
	context.JSON(http.StatusOK, res)
}

func (c *warehouseController) Restore(context *gin.Context) {
	var warehouseRestoreDTO dto.WarehouseRestoreDTO
	id := context.Param("id")
	var warehouse model.Warehouse = c.warehouseService.FindByID(id)
	if (warehouse != model.Warehouse{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		warehouseRestoreDTO.ID = uuid.MustParse(id)
		result := c.warehouseService.Restore(warehouseRestoreDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *warehouseController) DeletePermanent(context *gin.Context) {
	var warehouse model.Warehouse
	id := context.Param("id")
	warehouse = c.warehouseService.FindByID(id)
	if (warehouse != model.Warehouse{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		warehouse.ID = id
		c.warehouseService.DeletePermanent(warehouse)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	}
}
