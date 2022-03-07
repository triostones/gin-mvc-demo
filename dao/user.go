package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/triostones/gin-mvc-demo/models"
	"github.com/triostones/gin-mvc-demo/utils"
)

type UserDao struct{}

func (dao *UserDao) GetByID(id uint) (*models.UserModel, error) {
	var user models.UserModel
	err := utils.DB.Where(&models.UserModel{Model: gorm.Model{ID: id}}).First(&user).Error
	return &user, err
}

func (dao *UserDao) GetByEmail(email string) (*models.UserModel, error) {
	var user models.UserModel
	err := utils.DB.Where(&models.UserModel{Email: email}).First(&user).Error
	return &user, err
}
