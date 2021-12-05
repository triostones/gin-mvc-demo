package config

import "os"

const DATABASE_URI = "./triostones-backend-gin.db"

var ADMIN_USERNAME = os.Getenv("ADMIN_USERNAME")
var ADMIN_PASSWORD = os.Getenv("ADMIN_PASSWORD")
var ADMIN_EMAIL = os.Getenv("ADMIN_EMAIL")

func Initialize() {
	if ADMIN_EMAIL == "" {
		ADMIN_EMAIL = "admin@triostones.com"
	}
	if ADMIN_USERNAME == "" {
		ADMIN_USERNAME = "admin"
	}
	if ADMIN_PASSWORD == "" {
		ADMIN_PASSWORD = "admin@123"
	}
}
