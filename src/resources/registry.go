package resources

import (
	"github.com/gin-gonic/gin"
	"github.com/triostones/triostones-backend-gin/src/resources/auth"
	"github.com/triostones/triostones-backend-gin/src/resources/ping"
)

func RegisterResources(engine *gin.Engine) {
	ping.RegisterPingGroup(engine)
	auth.RegisterAuthGroup(engine)
}
