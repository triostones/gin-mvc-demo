package ping

import (
	"github.com/gin-gonic/gin"
	"github.com/triostones/triostones-backend-gin/src/middleware"
)

func RegisterPingGroup(engine *gin.Engine) {
	pingGroup := engine.Group("/ping")
	pingResource := PingResource{path: "/"}
	{
		pingResource.get(pingGroup)
		pingGroup.Use(middleware.JWT.MiddlewareFunc())
		pingResource.post(pingGroup)
	}
}
