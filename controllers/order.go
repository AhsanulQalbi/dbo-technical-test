package controllers

import (
	"net/http"
	"strconv"

	"dbo-technical-test/helpers"
	"dbo-technical-test/params"
	"dbo-technical-test/services"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService      services.OrderService
	validationService *helpers.ValidatorHelpers
}

func NewOrderController(orderService services.OrderService, validationService *helpers.ValidatorHelpers) *OrderController {
	return &OrderController{orderService, validationService}
}

func (orderController *OrderController) CreateOrder(c *gin.Context) {
	var request params.CreateOrder
	err := c.ShouldBind(&request)
	if err != nil {
		helpers.HandleErrorController(c, http.StatusBadRequest, err.Error())
		return
	}

	err = orderController.validationService.ValidateIncomingRequest(request)
	if err != nil {
		validationMessage := orderController.validationService.BuildAndGetValidationMessage(err)

		helpers.HandleErrorController(c, http.StatusBadRequest, validationMessage)
		return
	}

	result := orderController.orderService.CreateOrder(request)
	c.JSON(result.Status, result.Payload)
}

func (orderController *OrderController) GetOrderList(c *gin.Context) {
	var queries params.Query
	err := c.ShouldBindQuery(&queries)
	if err != nil {
		helpers.HandleErrorController(c, http.StatusInternalServerError, "Failed Bind Params")
		return
	}

	result := orderController.orderService.GetOrderList(&queries)
	c.JSON(result.Status, result.Payload)
}

func (orderController *OrderController) GetOrderById(c *gin.Context) {
	orderStr := c.Param("orderId")
	orderId, err := strconv.Atoi(orderStr)
	if err != nil {
		helpers.HandleErrorController(c, http.StatusBadRequest, err.Error())
		return
	}
	result := orderController.orderService.GetOrderById(orderId)
	c.JSON(result.Status, result.Payload)
}

func (orderController *OrderController) UpdateOrder(c *gin.Context) {
	var request params.UpdateOrder
	orderStr := c.Param("orderId")
	orderId, err := strconv.Atoi(orderStr)
	if err != nil {
		helpers.HandleErrorController(c, http.StatusBadRequest, err.Error())
		return
	}

	err = c.ShouldBind(&request)
	if err != nil {
		helpers.HandleErrorController(c, http.StatusBadRequest, err.Error())
		return
	}

	err = orderController.validationService.ValidateIncomingRequest(request)
	if err != nil {
		validationMessage := orderController.validationService.BuildAndGetValidationMessage(err)

		helpers.HandleErrorController(c, http.StatusBadRequest, validationMessage)
		return
	}

	result := orderController.orderService.UpdateOrder(orderId, request)
	c.JSON(result.Status, result.Payload)
}

func (orderController *OrderController) DeleteOrder(c *gin.Context) {
	orderStr := c.Param("orderId")
	orderId, err := strconv.Atoi(orderStr)
	if err != nil {
		helpers.HandleErrorController(c, http.StatusBadRequest, err.Error())
		return
	}

	result := orderController.orderService.DeleteOrder(orderId)
	c.JSON(result.Status, result.Payload)
}
