package controller

import (
	"ginWeb/app/dto"
	"ginWeb/app/helper"
	"ginWeb/app/model"
	"ginWeb/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ZoneController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	Deleted(context *gin.Context)
	Restore(context *gin.Context)
	DeletePermanent(context *gin.Context)
}

type zoneController struct {
	zoneService service.ZoneService
}

func NewZoneController(zoneServ service.ZoneService) ZoneController {
	return &zoneController{
		zoneService: zoneServ,
	}
}

func (c *zoneController) All(context *gin.Context) {
	var zones []model.Zone = c.zoneService.All()
	res := helper.BuildResponse(true, "OK", zones)
	context.JSON(http.StatusOK, res)
}

func (c *zoneController) FindByID(context *gin.Context) {
	id := context.Param("id")
	var zone model.Zone = c.zoneService.FindByID(id)
	if (zone == model.Zone{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", zone)
		context.JSON(http.StatusOK, res)
	}
}

func (c *zoneController) Insert(context *gin.Context) {
	var zoneCreateDTO dto.ZoneCreateDTO
	errDTO := context.ShouldBind(&zoneCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
	result := c.zoneService.Insert(zoneCreateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusCreated, response)
}

func (c *zoneController) Update(context *gin.Context) {
	var zoneUpdateDTO dto.ZoneUpdateDTO
	errDTO := context.ShouldBind(&zoneUpdateDTO)
	id := context.Param("id")
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	var zone model.Zone = c.zoneService.FindByID(id)
	if (zone == model.Zone{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		zoneUpdateDTO.ID = id
		result := c.zoneService.Update(zoneUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *zoneController) Delete(context *gin.Context) {
	var zone model.Zone
	id := context.Param("id")
	zone = c.zoneService.FindByID(id)
	if (zone == model.Zone{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		zone.ID = id
		result := c.zoneService.Delete(zone)
		res := helper.BuildResponse(true, "Deleted", result)
		context.JSON(http.StatusOK, res)
	}
}

func (c *zoneController) Deleted(context *gin.Context) {
	var zones []model.Zone = c.zoneService.Deleted()
	res := helper.BuildResponse(true, "OK", zones)
	context.JSON(http.StatusOK, res)
}

func (c *zoneController) Restore(context *gin.Context) {
	var zoneRestoreDTO dto.ZoneRestoreDTO
	id := context.Param("id")
	var zone model.Zone = c.zoneService.FindByID(id)
	if (zone != model.Zone{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		zoneRestoreDTO.ID = id
		result := c.zoneService.Restore(zoneRestoreDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *zoneController) DeletePermanent(context *gin.Context) {
	var zone model.Zone
	id := context.Param("id")
	zone = c.zoneService.FindByID(id)
	if (zone != model.Zone{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		zone.ID = id
		c.zoneService.DeletePermanent(zone)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	}
}
