package services

import (
	"dbo-technical-test/helpers"
	"dbo-technical-test/models"
	"dbo-technical-test/params"
	"time"

	"dbo-technical-test/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	userRepo repositories.UserRepo
}

func NewUserService(repo repositories.UserRepo) *UserService {
	return &UserService{
		userRepo: repo,
	}
}

func (userService *UserService) Login(request params.UserLogin) *params.Response {
	if request.Email == "" {
		return helpers.HandleErrorService(http.StatusBadRequest, "Email cannot be null")
	}

	if request.Password == "" {
		return helpers.HandleErrorService(http.StatusBadRequest, "Password cannot be null")
	}

	userDB, err := userService.userRepo.FindByEmail(request.Email)
	if err != nil {
		return helpers.HandleErrorService(http.StatusBadRequest, err.Error())
	}

	passwordMatch := helpers.ComparePassword([]byte(userDB.Password), []byte(request.Password))
	if !passwordMatch {
		return helpers.HandleErrorService(http.StatusBadRequest, "Password doesn't match")
	}

	token := helpers.GenerateToken(userDB.Id, userDB.Email, userDB.Role, userDB.Fullname)

	return &params.Response{
		Status: http.StatusOK,
		Payload: gin.H{
			"message": "Login successful",
			"token":   token,
		},
	}
}

func (userservice *UserService) CreateUser(request params.CreateUser) *params.Response {
	user := models.User{
		Fullname:  request.Fullname,
		Password:  helpers.HashPassword(request.Password),
		Email:     request.Email,
		Role:      request.Role,
		Phone:     request.Phone,
		Address:   request.Address,
		BirthDate: request.BirthDate,
		Gender:    request.Gender,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userData, err := userservice.userRepo.CreateUser(user)
	if err != nil {
		return helpers.HandleErrorService(http.StatusBadRequest, err.Error())
	}

	return &params.Response{
		Status:  http.StatusCreated,
		Payload: userData,
	}
}
