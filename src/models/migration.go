package models

import (
	"github.com/triostones/triostones-backend-gin/src/config"
	"github.com/triostones/triostones-backend-gin/src/utils"
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
