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

type ProductController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	Deleted(context *gin.Context)
	Restore(context *gin.Context)
	DeletePermanent(context *gin.Context)
}

type productController struct {
	productService service.ProductService
}

func NewProductController(productServ service.ProductService) ProductController {
	return &productController{
		productService: productServ,
	}
}

func (c *productController) All(context *gin.Context) {
	var products []model.Product = c.productService.All()
	res := helper.BuildResponse(true, "OK", products)
	context.JSON(http.StatusOK, res)
}

func (c *productController) FindByID(context *gin.Context) {
	id := context.Param("id")
	var product model.Product = c.productService.FindByID(id)
	if (product == model.Product{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		res := helper.BuildResponse(true, "OK", product)
		context.JSON(http.StatusOK, res)
	}
}

func (c *productController) Insert(context *gin.Context) {
	var productCreateDTO dto.ProductCreateDTO
	errDTO := context.ShouldBind(&productCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	}
	result := c.productService.Insert(productCreateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusCreated, response)
}

func (c *productController) Update(context *gin.Context) {
	var productUpdateDTO dto.ProductUpdateDTO
	errDTO := context.ShouldBind(&productUpdateDTO)
	id := context.Param("id")
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	var product model.Product = c.productService.FindByID(id)
	if (product == model.Product{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		productUpdateDTO.ID = uuid.MustParse(id)
		result := c.productService.Update(productUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *productController) Delete(context *gin.Context) {
	var product model.Product
	id := context.Param("id")
	product = c.productService.FindByID(id)
	if (product == model.Product{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		product.ID = uuid.MustParse(id)
		result := c.productService.Delete(product)
		res := helper.BuildResponse(true, "Deleted", result)
		context.JSON(http.StatusOK, res)
	}
}

func (c *productController) Deleted(context *gin.Context) {
	var products []model.Product = c.productService.Deleted()
	res := helper.BuildResponse(true, "OK", products)
	context.JSON(http.StatusOK, res)
}

func (c *productController) Restore(context *gin.Context) {
	var productRestoreDTO dto.ProductRestoreDTO
	id := context.Param("id")
	var product model.Product = c.productService.FindByID(id)
	if (product != model.Product{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		productRestoreDTO.ID = uuid.MustParse(id)
		result := c.productService.Restore(productRestoreDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)
	}
}

func (c *productController) DeletePermanent(context *gin.Context) {
	var product model.Product
	id := context.Param("id")
	product = c.productService.FindByID(id)
	if (product != model.Product{}) {
		res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)
	} else {
		product.ID = uuid.MustParse(id)
		c.productService.DeletePermanent(product)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	}
}
