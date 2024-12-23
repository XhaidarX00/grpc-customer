package controller

import (
	"microservice/config"
	usercontroller "microservice/controller/user"
	"microservice/database"
	"microservice/service"

	"go.uber.org/zap"
)

type Controller struct {
	User usercontroller.UserController
}

func NewController(service service.Service, logger *zap.Logger, cacher database.Cacher, config config.Configuration) *Controller {
	return &Controller{
		User: *usercontroller.NewUserController(service, logger, cacher, config),
	}
}
