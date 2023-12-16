package controllers

import (
	"dbo-technical-test/helpers"
	"dbo-technical-test/params"
	"dbo-technical-test/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService      services.UserService
	validatorHelpers helpers.ValidatorHelpers
}

func NewUserController(service *services.UserService, validatorHelpers *helpers.ValidatorHelpers) *UserController {
	return &UserController{
		userService:      *service,
		validatorHelpers: *validatorHelpers,
	}
}

func (userController *UserController) CreateUser(c *gin.Context) {
	var req params.CreateUser

	err := c.ShouldBind(&req)

	if err != nil {
		helpers.HandleErrorController(c, http.StatusBadRequest, err.Error())
		return
	}

	err = userController.validatorHelpers.Validate.Struct(req)

	if err != nil {
		validationMessage := userController.validatorHelpers.BuildAndGetValidationMessage(err)

		helpers.HandleErrorController(c, http.StatusBadRequest, validationMessage)
		return
	}

	result := userController.userService.CreateUser(req)

	c.JSON(result.Status, result.Payload)
}
