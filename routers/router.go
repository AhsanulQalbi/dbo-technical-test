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
	mainRouter := route.Group("/v1")
	{
		mainRouter.POST("/login", userController.Login)
		authorized := mainRouter.Group("/")
		authorized.Use(middlewares.Auth())
		{
			superadmin := authorized.Group("/")
			superadmin.Use(middlewares.IsSuperAdmin())
			{
				superadmin.POST("/create-user", userController.CreateUser)
			}
		}
	}
	return route
}
