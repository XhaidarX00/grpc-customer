package userservice

import (
	"microservice/helper"
	"microservice/models"
	"microservice/repository"

	"go.uber.org/zap"
)

type UserService interface {
	CheckUserEmail(email string) (*models.User, error)
	ResetUserPassword(input models.LoginRequest) error
	UpdateUser(input models.UpdateUserRequest) error
}

type userService struct {
	Repo repository.Repository
	Log  *zap.Logger
}

// CheckUserEmail implements UserService.
func (u *userService) CheckUserEmail(email string) (*models.User, error) {
	return u.Repo.User.GetByEmail(email)
}

// UpdateData
func (u *userService) UpdateUser(input models.UpdateUserRequest) error {
	return u.Repo.User.Update(input)
}

// ResetUserPassword implements UserService.
func (u *userService) ResetUserPassword(input models.LoginRequest) error {
	resetPassword := models.LoginRequest{
		Email:    input.Email,
		Password: helper.HashPassword(input.Password),
	}
	return u.Repo.User.UpdatePassword(resetPassword)
}

func NewUserService(repo repository.Repository, log *zap.Logger) UserService {
	return &userService{Repo: repo, Log: log}
}
