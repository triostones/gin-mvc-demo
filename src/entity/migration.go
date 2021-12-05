package entity

import (
	"github.com/triostones/triostones-backend-gin/src/config"
	"github.com/triostones/triostones-backend-gin/src/utils"
	"golang.org/x/crypto/bcrypt"
)

func initDatabase() {
	user := UserEntity{}
	utils.DB.Where(&UserEntity{Email: config.ADMIN_EMAIL}).First(&user)
	if user.ID == 0 {
		passwordHash, _ := bcrypt.GenerateFromPassword(
			[]byte(config.ADMIN_PASSWORD), bcrypt.DefaultCost)
		utils.DB.Create(&UserEntity{
			Email:    config.ADMIN_EMAIL,
			Password: string(passwordHash),
			Username: config.ADMIN_USERNAME,
		})
	}
}

func AutoMigrate() {
	utils.DB.AutoMigrate(&UserEntity{})
	initDatabase()
}
