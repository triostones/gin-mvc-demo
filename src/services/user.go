package services

import (
	"github.com/triostones/triostones-backend-gin/src/models"
	"github.com/triostones/triostones-backend-gin/src/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func (s *UserService) Login(email string, password string) (models.UserModel, error) {
	var user models.UserModel
	err := utils.DB.Where(&models.UserModel{Email: email}).First(&user).Error
	if err != nil {
		return user, err
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return models.UserModel{}, nil
	}
	return user, nil
}
