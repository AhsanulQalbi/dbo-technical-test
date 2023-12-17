package routers

import (
	"dbo-technical-test/controllers"
	"dbo-technical-test/helpers"
	"dbo-technical-test/middlewares"
	"dbo-technical-test/repositories"
	"dbo-technical-test/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func RouterConfig(db *gorm.DB) *gin.Engine {
	route := gin.Default()
	logger := log.New()
	route.Use(middlewares.ErrorHandler(logger))

	// helpers
	validationService := helpers.NewValidatorService()
	repoHelpers := repositories.NewRepoHelpers()

	//User Repo
	userRepo := repositories.NewUserRepo(db, *repoHelpers)
	userService := services.NewUserService(*userRepo)
	userController := controllers.NewUserController(userService, validationService)

	//Customer Repo
	customerRepo := repositories.NewCustomerRepo(db, *repoHelpers)
	customerService := services.NewCustomerService(*customerRepo)
	customerController := controllers.NewCustomerController(*customerService, validationService)

	//Product Repo
	productRepo := repositories.NewProductRepo(db, *repoHelpers)
	productService := services.NewProductService(*productRepo)
	productController := controllers.NewProductController(*productService, validationService)

	//Order Repo
	orderRepo := repositories.NewOrderRepo(db, *repoHelpers)
	orderService := services.NewOrderService(*orderRepo, *productRepo, *customerRepo)
	orderController := controllers.NewOrderController(*orderService, validationService)

	mainRouter := route.Group("/v1")
	{
		mainRouter.POST("/login", userController.Login)
		authorized := mainRouter.Group("/")
		authorized.Use(middlewares.Auth())
		{
			authorized.GET("/customer", customerController.GetCustomerList)
			authorized.GET("/customer/:customerId", customerController.GetCustomerById)
			authorized.POST("/customer", customerController.CreateCustomer)
			authorized.PUT("/customer/:customerId", customerController.UpdateCustomer)
			authorized.DELETE("/customer/:customerId", customerController.DeleteCustomer)

			authorized.GET("/product", productController.GetProductList)
			authorized.GET("/product/:productId", productController.GetProductById)
			authorized.POST("/product", productController.CreateProduct)
			authorized.PUT("/product/:productId", productController.UpdateProduct)
			authorized.DELETE("/product/:productId", productController.DeleteProduct)

			authorized.GET("/order", orderController.GetOrderList)
			authorized.GET("/order/:orderId", orderController.GetOrderById)
			authorized.POST("/order", orderController.CreateOrder)
			authorized.PUT("/order/:orderId", orderController.UpdateOrder)
			authorized.DELETE("/order/:orderId", orderController.DeleteOrder)

			superadmin := authorized.Group("/")
			superadmin.Use(middlewares.IsSuperAdmin())
			{
				superadmin.POST("/create-user", userController.CreateUser)
			}
		}
	}
	return route
}
