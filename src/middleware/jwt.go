package middleware

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/triostones/triostones-backend-gin/src/config"
	"github.com/triostones/triostones-backend-gin/src/dao"
	"github.com/triostones/triostones-backend-gin/src/models"
)

var JWT *jwt.GinJWTMiddleware

func init() {
	JWT, _ = jwt.New(&jwt.GinJWTMiddleware{
		Key: []byte(config.JWT_SECRET_KEY),
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			id := uint(claims["id"].(float64))
			user, err := (&dao.UserDao{}).GetByID(id)
			if err != nil {
				return &models.UserModel{}
			}
			return user
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if u, ok := data.(*models.UserModel); ok && u.Email == config.ADMIN_EMAIL {
				return true
			}
			return false
		},
	})
}
