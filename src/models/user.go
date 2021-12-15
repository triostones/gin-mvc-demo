package models

import "github.com/jinzhu/gorm"

type UserModel struct {
	gorm.Model
	Username string `gorm:"size:128;not null"`
	Email    string `gorm:"size:128;not null;unique"`
	Password string `gorm:"size:64;not null"`
}

func (UserModel) TableName() string {
	return "users"
}
