package services

import (
	"github.com/triostones/triostones-backend-gin/src/entity"
	"github.com/triostones/triostones-backend-gin/src/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func (s *UserService) Login(email string, password string) (entity.UserEntity, error) {
	var user entity.UserEntity
	err := utils.DB.Where(&entity.UserEntity{Email: email}).First(&user).Error
	if err != nil {
		return user, err
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return entity.UserEntity{}, nil
	}
	return user, nil
}
