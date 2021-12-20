package utils

import (
	"github.com/triostones/gin-mvc-demo/src/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB, _ = gorm.Open(sqlite.Open(config.DATABASE_URI), &gorm.Config{})
}
