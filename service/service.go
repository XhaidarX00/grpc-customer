package service

import (
	"microservice/repository"
	userservice "microservice/service/user"

	"go.uber.org/zap"
)

type Service struct {
	User userservice.UserService
}

func NewService(repo repository.Repository, log *zap.Logger) Service {
	return Service{
		User: userservice.NewUserService(repo, log),
	}
}
