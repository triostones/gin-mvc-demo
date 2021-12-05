package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/triostones/triostones-backend-gin/src/config"
	"github.com/google/uuid"
)

type Claims struct {
	email string
	jwt.StandardClaims
}

func GenerateAccessToken(email string) (string, error) {
	jti, _ := uuid.NewRandom()
	claims := Claims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * config.JWT_ACCESS_TOKEN_EXPIRE_TIME).Unix(),
			Id: jti.String(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(config.JWT_SECRET_KEY))
	return signedString, err
}
