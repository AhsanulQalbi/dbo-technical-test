package controllers

import (
	"net/http"
	"strconv"

	"dbo-technical-test/helpers"
	"dbo-technical-test/params"
	"dbo-technical-test/services"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	customerService   services.CustomerService
	validationService *helpers.ValidatorHelpers
}

func NewCustomerController(customerService services.CustomerService, validationService *helpers.ValidatorHelpers) *CustomerController {
	return &CustomerController{customerService, validationService}
}

func (customerController *CustomerController) CreateCustomer(c *gin.Context) {
	var request params.CreateCustomer
	err := c.ShouldBind(&request)
	if err != nil {
		helpers.HandleErrorController(c, http.StatusBadRequest, err.Error())
		return
	}

	err = customerController.validationService.ValidateIncomingRequest(request)
	if err != nil {
		validationMessage := customerController.validationService.BuildAndGetValidationMessage(err)

		helpers.HandleErrorController(c, http.StatusBadRequest, validationMessage)
		return
	}

	result := customerController.customerService.CreateCustomer(request)
	c.JSON(result.Status, result.Payload)
}

func (customerController *CustomerController) GetCustomerList(c *gin.Context) {
	var queries params.Query
	err := c.ShouldBindQuery(&queries)
	if err != nil {
		helpers.HandleErrorController(c, http.StatusInternalServerError, "Failed Bind Params")
		return
	}

	result := customerController.customerService.GetCustomerList(&queries)
	c.JSON(result.Status, result.Payload)
}

func (customerController *CustomerController) GetCustomerById(c *gin.Context) {
	customerStr := c.Param("customerId")
	customerId, err := strconv.Atoi(customerStr)
	if err != nil {
		helpers.HandleErrorController(c, http.StatusBadRequest, err.Error())
		return
	}
	result := customerController.customerService.GetCustomerById(customerId)
	c.JSON(result.Status, result.Payload)
}

func (customerController *CustomerController) UpdateCustomer(c *gin.Context) {
	var request params.UpdateCustomer
	customerStr := c.Param("customerId")
	customerId, err := strconv.Atoi(customerStr)
	if err != nil {
		helpers.HandleErrorController(c, http.StatusBadRequest, err.Error())
		return
	}

	err = c.ShouldBind(&request)
	if err != nil {
		helpers.HandleErrorController(c, http.StatusBadRequest, err.Error())
		return
	}

	err = customerController.validationService.ValidateIncomingRequest(request)
	if err != nil {
		validationMessage := customerController.validationService.BuildAndGetValidationMessage(err)

		helpers.HandleErrorController(c, http.StatusBadRequest, validationMessage)
		return
	}

	result := customerController.customerService.UpdateCustomer(customerId, request)
	c.JSON(result.Status, result.Payload)
}

func (customerController *CustomerController) DeleteCustomer(c *gin.Context) {
	customerStr := c.Param("customerId")
	customerId, err := strconv.Atoi(customerStr)
	if err != nil {
		helpers.HandleErrorController(c, http.StatusBadRequest, err.Error())
		return
	}

	result := customerController.customerService.DeleteCustomer(customerId)
	c.JSON(result.Status, result.Payload)
}
