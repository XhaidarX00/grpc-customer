package userrepository

import (
	"microservice/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByEmail(email string) (*models.User, error)
	UpdatePassword(resetPasswordInput models.LoginRequest) error
	Update(user models.UpdateUserRequest) error
}

type userRepository struct {
	DB  *gorm.DB
	Log *zap.Logger
}

// GetByEmail implements UserRepository.
func (u *userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := u.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (u *userRepository) Update(user models.UpdateUserRequest) error {
	if err := u.DB.Model(&models.User{}).Where("email = ?", user.Email).Updates(user).Error; err != nil {
		return err
	}

	return nil
}

// UpdatePassword implements UserRepository.
func (u *userRepository) UpdatePassword(resetPasswordInput models.LoginRequest) error {
	return u.DB.Model(&models.User{}).Where("email =?", resetPasswordInput.Email).Update("password", resetPasswordInput.Password).Error
}

func NewUserRepository(db *gorm.DB, log *zap.Logger) UserRepository {
	return &userRepository{DB: db, Log: log}
}
