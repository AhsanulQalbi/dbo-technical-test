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

type CustomerService struct {
	customerRepo repositories.CustomerRepo
}

func NewCustomerService(repo repositories.CustomerRepo) *CustomerService {
	return &CustomerService{
		customerRepo: repo,
	}
}

func (customerService *CustomerService) CreateCustomer(request params.CreateCustomer) *params.Response {
	customerCheck, err := customerService.customerRepo.CheckEmail(request.Email)
	if len(*customerCheck) != 0 {
		return helpers.HandleErrorService(http.StatusBadRequest, fmt.Sprintf("email %s already registered. ", request.Email))
	}

	dateBirth, err := time.Parse("2006-01-02", request.BirthDate)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		// return helpers.HandleErrorService(http.StatusBadRequest, err.Error())
	}

	customer := models.Customer{
		Fullname:  request.Fullname,
		Email:     request.Email,
		BirthDate: dateBirth,
		Phone:     request.Phone,
		Address:   request.Address,
		Gender:    request.Gender,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	customerData, err := customerService.customerRepo.CreateCustomer(customer)
	if err != nil {
		return helpers.HandleErrorService(http.StatusBadRequest, err.Error())
	}

	return &params.Response{
		Status:  http.StatusCreated,
		Payload: customerData,
	}
}

func (customerService *CustomerService) UpdateCustomer(customerId int, request params.UpdateCustomer) *params.Response {
	var updateCustomer models.Customer
	var dateBirth time.Time
	_, err := customerService.customerRepo.GetCustomerById(customerId)
	if err != nil {
		log.Errorln("ERROR:", err)
		return helpers.HandleErrorService(http.StatusBadRequest, "Error / Customer Data not Found")
	}

	if request.BirthDate != "" {
		dateBirth, err = time.Parse("2006-01-02", request.BirthDate)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			// return helpers.HandleErrorService(http.StatusBadRequest, err.Error())
		}
	}

	updateCustomer = models.Customer{
		Fullname:  request.Fullname,
		Phone:     request.Phone,
		Address:   request.Address,
		BirthDate: dateBirth,
		Gender:    request.Gender,
		UpdatedAt: time.Now(),
	}

	updateCustomerData, err := customerService.customerRepo.UpdateCustomer(customerId, updateCustomer)
	if err != nil {
		log.Errorln("ERROR:", err)
		return helpers.HandleErrorService(http.StatusInternalServerError, err.Error())
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: updateCustomerData,
	}
}

func (customerService *CustomerService) GetCustomerById(customerId int) *params.Response {
	customer, err := customerService.customerRepo.GetCustomerById(customerId)
	if err != nil {
		log.Errorln("ERROR:", err)
		return helpers.HandleErrorService(http.StatusBadRequest, "Can't get customer data")
	}

	customerRes := params.DetailCustomer{
		Id:        customer.Id,
		Fullname:  customer.Fullname,
		Email:     customer.Email,
		Phone:     customer.Phone,
		Address:   customer.Address,
		BirthDate: customer.BirthDate.Format("2006-01-02"),
		Gender:    customer.Gender,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
	}

	customerResponse := params.CustomerResponse{
		Message: "Success Get Customer By Id",
		Data:    customerRes,
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: customerResponse,
	}
}

func (customerService *CustomerService) GetCustomerList(queries *params.Query) *params.Response {
	customerList, count, err := customerService.customerRepo.GetCustomerList(queries)
	if err != nil {
		return helpers.HandleErrorService(http.StatusInternalServerError, fmt.Sprintf("Error on get customer info: [%s]", err.Error()))
	}

	if len(customerList) == 0 {
		return &params.Response{
			Status:  http.StatusNoContent,
			Payload: "record not found",
		}
	}

	pagination := helpers.CalculatePagination(count, queries.Page, queries.Size, len(customerList))

	result := params.ResponseWithPagination{
		Pagination: pagination,
		Data:       customerList,
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: result,
	}
}

func (customerService *CustomerService) DeleteCustomer(customerId int) *params.Response {
	_, err := customerService.customerRepo.GetCustomerById(customerId)
	if err != nil {
		log.Errorln("ERROR:", err)
		return helpers.HandleErrorService(http.StatusBadRequest, fmt.Sprintf("Error on get customer info: [%s]", err.Error()))

	}

	err = customerService.customerRepo.DeleteCustomer(customerId)
	if err != nil {
		log.Errorln("ERROR:", err)
		return helpers.HandleErrorService(http.StatusInternalServerError, fmt.Sprintf("Error on delete customer : [%s]", err.Error()))

	}

	return &params.Response{
		Status: http.StatusOK,
		Payload: map[string]string{
			"message": fmt.Sprintf("List customer with id = %d has been deleted", customerId),
		},
	}
}
