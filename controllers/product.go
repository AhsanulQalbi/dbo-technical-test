package controllers

import (
	"net/http"
	"strconv"

	"dbo-technical-test/helpers"
	"dbo-technical-test/params"
	"dbo-technical-test/services"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService    services.ProductService
	validationService *helpers.ValidatorHelpers
}

func NewProductController(productService services.ProductService, validationService *helpers.ValidatorHelpers) *ProductController {
	return &ProductController{productService, validationService}
}

func (productController *ProductController) CreateProduct(c *gin.Context) {
	var request params.CreateProduct
	err := c.ShouldBind(&request)
	if err != nil {
		helpers.HandleErrorController(c, http.StatusBadRequest, err.Error())
		return
	}

	err = productController.validationService.ValidateIncomingRequest(request)
	if err != nil {
		validationMessage := productController.validationService.BuildAndGetValidationMessage(err)

		helpers.HandleErrorController(c, http.StatusBadRequest, validationMessage)
		return
	}

	result := productController.ProductService.CreateProduct(request)
	c.JSON(result.Status, result.Payload)
}

func (productController *ProductController) GetProductList(c *gin.Context) {
	var queries params.Query
	err := c.ShouldBindQuery(&queries)
	if err != nil {
		helpers.HandleErrorController(c, http.StatusInternalServerError, "Failed Bind Params")
		return
	}

	result := productController.ProductService.GetProductList(&queries)
	c.JSON(result.Status, result.Payload)
}

func (productController *ProductController) GetProductById(c *gin.Context) {
	productStr := c.Param("productId")
	productId, err := strconv.Atoi(productStr)
	if err != nil {
		helpers.HandleErrorController(c, http.StatusBadRequest, err.Error())
		return
	}
	result := productController.ProductService.GetProductById(productId)
	c.JSON(result.Status, result.Payload)
}

func (productController *ProductController) UpdateProduct(c *gin.Context) {
	var request params.UpdateProduct
	productStr := c.Param("productId")
	productId, err := strconv.Atoi(productStr)
	if err != nil {
		helpers.HandleErrorController(c, http.StatusBadRequest, err.Error())
		return
	}

	err = c.ShouldBind(&request)
	if err != nil {
		helpers.HandleErrorController(c, http.StatusBadRequest, err.Error())
		return
	}

	err = productController.validationService.ValidateIncomingRequest(request)
	if err != nil {
		validationMessage := productController.validationService.BuildAndGetValidationMessage(err)

		helpers.HandleErrorController(c, http.StatusBadRequest, validationMessage)
		return
	}

	result := productController.ProductService.UpdateProduct(productId, request)
	c.JSON(result.Status, result.Payload)
}

func (productController *ProductController) DeleteProduct(c *gin.Context) {
	productStr := c.Param("productId")
	productId, err := strconv.Atoi(productStr)
	if err != nil {
		helpers.HandleErrorController(c, http.StatusBadRequest, err.Error())
		return
	}

	result := productController.ProductService.DeleteProduct(productId)
	c.JSON(result.Status, result.Payload)
}
