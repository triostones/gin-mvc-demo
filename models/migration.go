package models

import (
	"github.com/triostones/gin-mvc-demo/config"
	"github.com/triostones/gin-mvc-demo/utils"
	"golang.org/x/crypto/bcrypt"
)

func initDatabase() {
	user := UserModel{}
	utils.DB.Where(&UserModel{Email: config.ADMIN_EMAIL}).First(&user)
	if user.ID == 0 {
		passwordHash, _ := bcrypt.GenerateFromPassword(
			[]byte(config.ADMIN_PASSWORD), bcrypt.DefaultCost)
		utils.DB.Create(&UserModel{
			Email:    config.ADMIN_EMAIL,
			Password: string(passwordHash),
			Username: config.ADMIN_USERNAME,
		})
	}
}

func AutoMigrate() {
	utils.DB.AutoMigrate(&UserModel{})
	initDatabase()
}
