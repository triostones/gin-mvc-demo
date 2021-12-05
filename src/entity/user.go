package entity

import "github.com/jinzhu/gorm"

type UserEntity struct {
	gorm.Model
	Username string `gorm:"size:128;not null"`
	Email    string `gorm:"size:128;not null;unique"`
	Password string `gorm:"size:64;not null"`
}

func (UserEntity) TableName() string {
	return "users"
}
