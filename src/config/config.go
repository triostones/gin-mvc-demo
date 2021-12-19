package config

import "os"

const DATABASE_URI = "./triostones-backend-gin.db"

const CONTEXT_VALIDATOR_KEY = "validator"
const CONTEXT_SERIALIZER_KEY = "serializer"
const JWT_ACCESS_TOKEN_EXPIRE_TIME = 24

var ADMIN_USERNAME = os.Getenv("ADMIN_USERNAME")
var ADMIN_PASSWORD = os.Getenv("ADMIN_PASSWORD")
var ADMIN_EMAIL = os.Getenv("ADMIN_EMAIL")

var JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")

func init() {
	if ADMIN_EMAIL == "" {
		ADMIN_EMAIL = "admin@triostones.com"
	}
	if ADMIN_USERNAME == "" {
		ADMIN_USERNAME = "admin"
	}
	if ADMIN_PASSWORD == "" {
		ADMIN_PASSWORD = "admin@123"
	}
	if JWT_SECRET_KEY == "" {
		JWT_SECRET_KEY = "secret"
	}
}
