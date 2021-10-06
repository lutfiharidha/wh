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

type RackController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	Deleted(context *gin.Context)
	Restore(context *gin.Context)
	DeletePermanent(context *gin.Context)
}

type rackController struct {
	rackService service.RackService
}

func NewRackController(rackServ service.RackService) RackController {
	return &rackController{
		rackService: rackServ,
	}
}

func (c *rackController) All(context *gin.Context) {
	var racks []model.Rack = c.rackService.All()
	res := helper.BuildResponse(true, "OK", racks)
	context.JSON(http.StatusOK, res)
}

func (c *rackController) FindByID(context *gin.Context) {
	id := context.Param("id")
	var rack model.Rack = c.rackService.FindByID(id)
	if (rack == model.Rack{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", rack)
		context.JSON(http.StatusOK, res)
	}
}

func (c *rackController) Insert(context *gin.Context) {
	var rackCreateDTO dto.RackCreateDTO
	errDTO := context.ShouldBind(&rackCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
	result := c.rackService.Insert(rackCreateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusCreated, response)
}

func (c *rackController) Update(context *gin.Context) {
	var rackUpdateDTO dto.RackUpdateDTO
	errDTO := context.ShouldBind(&rackUpdateDTO)
	id := context.Param("id")
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	var rack model.Rack = c.rackService.FindByID(id)
	if (rack == model.Rack{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		rackUpdateDTO.ID = uuid.MustParse(id)
		result := c.rackService.Update(rackUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *rackController) Delete(context *gin.Context) {
	var rack model.Rack
	id := context.Param("id")
	rack = c.rackService.FindByID(id)
	if (rack == model.Rack{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		rack.ID = uuid.MustParse(id)
		result := c.rackService.Delete(rack)
		res := helper.BuildResponse(true, "Deleted", result)
		context.JSON(http.StatusOK, res)
	}
}

func (c *rackController) Deleted(context *gin.Context) {
	var racks []model.Rack = c.rackService.Deleted()
	res := helper.BuildResponse(true, "OK", racks)
	context.JSON(http.StatusOK, res)
}

func (c *rackController) Restore(context *gin.Context) {
	var rackRestoreDTO dto.RackRestoreDTO
	id := context.Param("id")
	var rack model.Rack = c.rackService.FindByID(id)
	if (rack != model.Rack{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		rackRestoreDTO.ID = uuid.MustParse(id)
		result := c.rackService.Restore(rackRestoreDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *rackController) DeletePermanent(context *gin.Context) {
	var rack model.Rack
	id := context.Param("id")
	rack = c.rackService.FindByID(id)
	if (rack != model.Rack{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		rack.ID = uuid.MustParse(id)
		c.rackService.DeletePermanent(rack)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	}
}
