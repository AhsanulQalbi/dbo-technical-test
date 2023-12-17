package services

import (
	"dbo-technical-test/helpers"
	"dbo-technical-test/models"
	"dbo-technical-test/params"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"dbo-technical-test/repositories"
	"net/http"
)

type OrderService struct {
	orderRepo    repositories.OrderRepo
	productRepo  repositories.ProductRepo
	customerRepo repositories.CustomerRepo
}

func NewOrderService(
	repo repositories.OrderRepo,
	productRepo repositories.ProductRepo,
	customerRepo repositories.CustomerRepo) *OrderService {
	return &OrderService{
		orderRepo:    repo,
		productRepo:  productRepo,
		customerRepo: customerRepo,
	}
}

func (orderService *OrderService) CreateOrder(request params.CreateOrder) *params.Response {
	productInfo, err := orderService.productRepo.GetProductById(request.ProductID)
	if err != nil {
		log.Errorln("ERROR:", err)
		return helpers.HandleErrorService(http.StatusBadRequest, fmt.Sprintf("Error on get product info: [%s]", err.Error()))
	}

	_, err = orderService.customerRepo.GetCustomerById(request.CustomerID)
	if err != nil {
		log.Errorln("ERROR:", err)
		return helpers.HandleErrorService(http.StatusBadRequest, fmt.Sprintf("Error on get customer info: [%s]", err.Error()))
	}

	orderDate, err := time.Parse("2006-01-02", request.OrderDate)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		// return helpers.HandleErrorService(http.StatusBadRequest, err.Error())
	}
	shippedDate, err := time.Parse("2006-01-02", request.ShippedDate)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		// return helpers.HandleErrorService(http.StatusBadRequest, err.Error())
	}
	arrivedDate, err := time.Parse("2006-01-02", request.ArrivedDate)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		// return helpers.HandleErrorService(http.StatusBadRequest, err.Error())
	}

	order := models.Order{
		OrderName:       request.OrderName,
		Quantity:        request.Quantity,
		TotalAmount:     productInfo.Price * request.Quantity,
		OrderDate:       orderDate,
		ShippedDate:     shippedDate,
		ArrivedDate:     arrivedDate,
		ShipAddress:     request.ShipAddress,
		ShipProvince:    request.ShipProvince,
		ShipCity:        request.ShipCity,
		ShipDistrict:    request.ShipDistrict,
		ShipSubDistrict: request.ShipSubDistrict,
		OrderStatus:     request.OrderStatus,
		CustomerID:      request.CustomerID,
		ProductID:       request.ProductID,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	orderData, err := orderService.orderRepo.CreateOrder(order)
	if err != nil {
		return helpers.HandleErrorService(http.StatusBadRequest, err.Error())
	}
	orderData.Product = nil
	orderData.Customer = nil

	return &params.Response{
		Status:  http.StatusCreated,
		Payload: orderData,
	}
}

func (orderService *OrderService) UpdateOrder(orderId int, request params.UpdateOrder) *params.Response {
	var shippedDate time.Time
	var arrivedDate time.Time

	_, err := orderService.orderRepo.GetOrderById(orderId)
	if err != nil {
		log.Errorln("ERROR:", err)
		return helpers.HandleErrorService(http.StatusBadRequest, "Error / Order Data not Found")
	}

	if request.ShippedDate != "" {
		shippedDate, err = time.Parse("2006-01-02", request.ShippedDate)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			// return helpers.HandleErrorService(http.StatusBadRequest, err.Error())
		}
	}
	if request.ShippedDate != "" {
		arrivedDate, err = time.Parse("2006-01-02", request.ArrivedDate)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			// return helpers.HandleErrorService(http.StatusBadRequest, err.Error())
		}
	}

	updateOrder := models.Order{
		ShippedDate:     shippedDate,
		ArrivedDate:     arrivedDate,
		ShipAddress:     request.ShipAddress,
		ShipProvince:    request.ShipProvince,
		ShipCity:        request.ShipCity,
		ShipDistrict:    request.ShipDistrict,
		ShipSubDistrict: request.ShipSubDistrict,
		OrderStatus:     request.OrderStatus,
		UpdatedAt:       time.Now(),
	}

	updateOrderData, err := orderService.orderRepo.UpdateOrder(orderId, updateOrder)
	if err != nil {
		log.Errorln("ERROR:", err)
		return helpers.HandleErrorService(http.StatusInternalServerError, err.Error())
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: updateOrderData,
	}
}

func (orderService *OrderService) GetOrderById(orderId int) *params.Response {
	order, err := orderService.orderRepo.GetOrderById(orderId)
	if err != nil {
		log.Errorln("ERROR:", err)
		return helpers.HandleErrorService(http.StatusBadRequest, "Can't get order data")
	}

	orderResponse := params.OrderResponse{
		Message: "Success Get Order By Id",
		Data:    order,
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: orderResponse,
	}
}

func (orderService *OrderService) GetOrderList(queries *params.Query) *params.Response {
	orderList, count, err := orderService.orderRepo.GetOrderList(queries)
	if err != nil {
		return helpers.HandleErrorService(http.StatusInternalServerError, fmt.Sprintf("Error on get order info: [%s]", err.Error()))
	}

	if len(orderList) == 0 {
		return &params.Response{
			Status:  http.StatusNoContent,
			Payload: "record not found",
		}
	}

	pagination := helpers.CalculatePagination(count, queries.Page, queries.Size, len(orderList))
	result := params.ResponseWithPagination{
		Pagination: pagination,
		Data:       orderList,
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: result,
	}
}

func (orderService *OrderService) DeleteOrder(orderId int) *params.Response {
	_, err := orderService.orderRepo.GetOrderById(orderId)
	if err != nil {
		log.Errorln("ERROR:", err)
		return helpers.HandleErrorService(http.StatusInternalServerError, fmt.Sprintf("Error on get order info: [%s]", err.Error()))

	}

	err = orderService.orderRepo.DeleteOrder(orderId)
	if err != nil {
		log.Errorln("ERROR:", err)
		return helpers.HandleErrorService(http.StatusInternalServerError, fmt.Sprintf("Error on delete order : [%s]", err.Error()))

	}

	return &params.Response{
		Status: http.StatusOK,
		Payload: map[string]string{
			"message": fmt.Sprintf("List order with id = %d has been deleted", orderId),
		},
	}
}
