package utils

import (
	"github.com/triostones/gin-mvc-demo/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Claims struct {
	email string
	jwt.StandardClaims
}

func GenerateAccessToken(id uint, email string) (string, error) {
	jti, _ := uuid.NewRandom()
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"jti":   jti.String(),
		"exp":   time.Now().Add(time.Hour * config.JWT_ACCESS_TOKEN_EXPIRE_TIME).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(config.JWT_SECRET_KEY))
	return signedString, err
}
