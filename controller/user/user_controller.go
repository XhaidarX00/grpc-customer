package usercontroller

import (
	"microservice/config"
	"microservice/database"
	"microservice/helper"
	"microservice/models"
	"microservice/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct {
	Service service.Service
	Log     *zap.Logger
	Cacher  database.Cacher
	Config  config.Configuration
}

func NewUserController(service service.Service, log *zap.Logger, cacher database.Cacher, config config.Configuration) *UserController {
	return &UserController{
		Service: service,
		Log:     log,
		Cacher:  cacher,
		Config:  config,
	}
}

func (ctrl *UserController) CheckDetailByEmailUserController(c *gin.Context) {
	request := models.CheckEmailRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		ctrl.Log.Error("Invalid request body", zap.Error(err))
		helper.ResponseError(c, err.Error(), "Invalid request body", http.StatusBadRequest)
		return
	}

	// existedUser, err := ctrl.Service.User.CheckUserEmail(request.Email)
	// if err != nil {
	// 	ctrl.Log.Error("Failed to check user email", zap.Error(err))
	// 	helper.ResponseError(c, err.Error(), "Failed to check user email", http.StatusInternalServerError)
	// 	return
	// }

	// Format respons sebagai map
	now := time.Now()
	response := map[string]interface{}{
		"message": "success get customer",
		"data": map[string]interface{}{
			"id":         1,
			"name":       "Haidar",
			"email":      "haidar@excample.com",
			"role":       "admin",
			"created_at": now,
			"updated_at": now,
		},
	}

	helper.ResponseOK(c, response, "success get customer", http.StatusOK)
}

func (ctrl *UserController) UpdateUserByEmailController(c *gin.Context) {
	// request := models.UpdateUserRequest{}
	// if err := c.ShouldBindJSON(&request); err != nil {
	// 	ctrl.Log.Error("Invalid request body", zap.Error(err))
	// 	helper.ResponseError(c, err.Error(), "Invalid request body", http.StatusBadRequest)
	// 	return
	// }

	// existedUser, err := ctrl.Service.User.CheckUserEmail(request.Email)
	// if err != nil {
	// 	ctrl.Log.Error("Failed to check user email", zap.Error(err))
	// 	helper.ResponseError(c, err.Error(), "user not found", http.StatusInternalServerError)
	// 	return
	// }

	// if request.Name == "" {
	// 	request.Name = existedUser.Name
	// }

	// if request.Email == "" {
	// 	request.Email = existedUser.Email
	// }

	// if request.Password == "" {
	// 	request.Password = existedUser.Password
	// } else {
	// 	request.Password = helper.HashPassword(request.Password)
	// }

	// err = ctrl.Service.User.UpdateUser(request)
	// if err != nil {
	// 	ctrl.Log.Error("Failed to check user email", zap.Error(err))
	// 	helper.ResponseError(c, err.Error(), "Failed to check user email", http.StatusInternalServerError)
	// 	return
	// }

	helper.ResponseOK(c, nil, "success update customer", http.StatusOK)
}
