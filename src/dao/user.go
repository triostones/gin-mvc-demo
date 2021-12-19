package dao

import (
	"github.com/triostones/triostones-backend-gin/src/models"
	"github.com/triostones/triostones-backend-gin/src/utils"
)

type UserDao struct{}

func (dao *UserDao) GetByEmail(email string) (*models.UserModel, error) {
	var user models.UserModel
	err := utils.DB.Where(&models.UserModel{Email: email}).First(&user).Error
	return &user, err
}
