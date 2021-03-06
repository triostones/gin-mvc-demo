package auth

import (
	"github.com/triostones/gin-mvc-demo/config"
	"github.com/triostones/gin-mvc-demo/models"
	"github.com/triostones/gin-mvc-demo/serializers/responses"
	"github.com/triostones/gin-mvc-demo/services"
	"github.com/triostones/gin-mvc-demo/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginResource struct {
	path string
}

func (resource *LoginResource) post(rg *gin.RouterGroup) {
	rg.POST(resource.path,
		func(c *gin.Context) {
			var login validators.Login
			if err := c.ShouldBindJSON(&login); err != nil {
				c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
					"status":  http.StatusUnprocessableEntity,
					"success": false,
					"message": err.Error(),
				})
			}
			c.Set(config.CONTEXT_VALIDATOR_KEY, login)
		},
		func(c *gin.Context) {
			login := c.MustGet(config.CONTEXT_VALIDATOR_KEY).(validators.Login)
			if user, err := (&services.UserService{}).Login(login.Email, login.Password); err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"status":  http.StatusInternalServerError,
					"success": false,
					"message": err.Error(),
				})
			} else if user.ID == 0 {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"status":  http.StatusBadRequest,
					"success": false,
					"message": "Invalid email or password",
				})
			} else {
				c.Set(config.CONTEXT_SERIALIZER_KEY, user)
			}
		},
		func(c *gin.Context) {
			user := c.MustGet(config.CONTEXT_SERIALIZER_KEY).(*models.UserModel)
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "",
				"item":    (&responses.LoginResponse{}).Dump(user),
			})
		},
	)
}
