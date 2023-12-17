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

type ProductService struct {
	productRepo repositories.ProductRepo
}

func NewProductService(repo repositories.ProductRepo) *ProductService {
	return &ProductService{
		productRepo: repo,
	}
}

func (productService *ProductService) CreateProduct(request params.CreateProduct) *params.Response {
	product := models.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Weight:      request.Weight,
		ImageUrl:    request.ImageUrl,
		Stock:       request.Stock,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	productData, err := productService.productRepo.CreateProduct(product)
	if err != nil {
		return helpers.HandleErrorService(http.StatusBadRequest, err.Error())
	}

	return &params.Response{
		Status:  http.StatusCreated,
		Payload: productData,
	}
}

func (productService *ProductService) UpdateProduct(productId int, request params.UpdateProduct) *params.Response {
	updateProduct := models.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Weight:      request.Weight,
		ImageUrl:    request.ImageUrl,
		Stock:       request.Stock,
		UpdatedAt:   time.Now(),
	}

	updateProductData, err := productService.productRepo.UpdateProduct(productId, updateProduct)
	if err != nil {
		log.Errorln("ERROR:", err)
		return helpers.HandleErrorService(http.StatusInternalServerError, err.Error())
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: updateProductData,
	}
}

func (productService *ProductService) GetProductById(productId int) *params.Response {
	product, err := productService.productRepo.GetProductById(productId)
	if err != nil {
		log.Errorln("ERROR:", err)
		return helpers.HandleErrorService(http.StatusBadRequest, "Can't get product data")
	}

	productRes := params.DetailProduct{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Weight:      product.Weight,
		ImageUrl:    product.ImageUrl,
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt.Format("2006-01-02"),
		UpdatedAt:   product.UpdatedAt.Format("2006-01-02"),
	}

	productResponse := params.ProductResponse{
		Message: "Success Get Product By Id",
		Data:    productRes,
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: productResponse,
	}
}

func (productService *ProductService) GetProductList(queries *params.Query) *params.Response {
	productList, count, err := productService.productRepo.GetProductList(queries)
	if err != nil {
		return helpers.HandleErrorService(http.StatusInternalServerError, fmt.Sprintf("Error on get product list info: [%s]", err.Error()))
	}

	if len(productList) == 0 {
		return &params.Response{
			Status:  http.StatusNoContent,
			Payload: "record not found",
		}
	}

	pagination := helpers.CalculatePagination(count, queries.Page, queries.Size, len(productList))
	result := params.ResponseWithPagination{
		Pagination: pagination,
		Data:       productList,
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: result,
	}
}

func (productService *ProductService) DeleteProduct(productId int) *params.Response {
	_, err := productService.productRepo.GetProductById(productId)
	if err != nil {
		log.Errorln("ERROR:", err)
		return helpers.HandleErrorService(http.StatusInternalServerError, fmt.Sprintf("Error on get product info: [%s]", err.Error()))
	}

	err = productService.productRepo.DeleteProduct(productId)
	if err != nil {
		log.Errorln("ERROR:", err)
		return helpers.HandleErrorService(http.StatusInternalServerError, fmt.Sprintf("Error on delete product : [%s]", err.Error()))

	}

	return &params.Response{
		Status: http.StatusOK,
		Payload: map[string]string{
			"message": fmt.Sprintf("List product with id = %d has been deleted", productId),
		},
	}
}
