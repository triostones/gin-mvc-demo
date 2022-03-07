package auth

import "github.com/gin-gonic/gin"

func RegisterAuthGroup(engine *gin.Engine) {
	authGroup := engine.Group("/auth")
	loginResource := LoginResource{path: "/login"}
	{
		loginResource.post(authGroup)
	}
}
