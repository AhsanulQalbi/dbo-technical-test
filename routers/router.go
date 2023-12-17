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

			superadmin := authorized.Group("/")
			superadmin.Use(middlewares.IsSuperAdmin())
			{
				superadmin.POST("/create-user", userController.CreateUser)
			}
		}
	}
	return route
}
