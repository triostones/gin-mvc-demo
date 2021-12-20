package resources

import (
	"github.com/gin-gonic/gin"
	"github.com/triostones/gin-mvc-demo/src/resources/auth"
	"github.com/triostones/gin-mvc-demo/src/resources/ping"
)

func RegisterResources(engine *gin.Engine) {
	ping.RegisterPingGroup(engine)
	auth.RegisterAuthGroup(engine)
}
