package repository

import (
	userrepository "microservice/repository/user"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	User userrepository.UserRepository
}

func NewRepository(db *gorm.DB, log *zap.Logger) Repository {
	return Repository{
		User: userrepository.NewUserRepository(db, log),
	}
}
